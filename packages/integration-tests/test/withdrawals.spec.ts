/* Imports: External */
import { BigNumber, constants, Contract, ContractReceipt, utils, Wallet } from 'ethers'
import { awaitCondition } from '@eth-optimism/core-utils'
import * as rlp from 'rlp'
import { Block } from '@ethersproject/abstract-provider'

/* Imports: Internal */
import { WITHDRAWER_ADDR } from './shared/constants'
import env from './shared/env'
import { expect } from './shared/setup'

const withdrawerArtifact = require('../../contracts/artifacts/contracts/L2/Withdrawer.sol/Withdrawer.json')
const l2OOracleArtifact = require('../../contracts/artifacts/contracts/L1/L2OutputOracle.sol/L2OutputOracle.json')
const counterArtifact = require('../artifacts/Counter.sol/Counter.json')
const multiDepositorArtifact = require('../artifacts/MultiDepositor.sol/MultiDepositor.json')

describe('Withdrawals', () => {
  let portal: Contract
  let withdrawer: Contract

  let recipient: Wallet

  before(async () => {
    portal = env.optimismPortal

    withdrawer = new Contract(
      WITHDRAWER_ADDR,
      withdrawerArtifact.abi,
    )
  })

  describe.skip('simple withdrawals', () => {
    let nonce: BigNumber
    let burnBlock: Block
    let withdrawalHash: string

    before(async function () {
      this.timeout(60_000)
      recipient = Wallet.createRandom().connect(env.l2Provider)
      withdrawer = withdrawer.connect(recipient)

      let tx = await portal.connect(env.l1Wallet).depositTransaction(
        recipient.address,
        utils.parseEther('1.337'),
        '3000000',
        false,
        [],
        {
          value: utils.parseEther('1.337'),
        },
      )
      await tx.wait()

      await awaitCondition(async () => {
        const bal = await recipient.getBalance()
        return bal.eq(tx.value)
      })

      tx = await env.l1Wallet.sendTransaction({
        to: recipient.address,
        value: utils.parseEther('2')
      })
      await tx.wait()
    })

    it('should create a withdrawal on L2', async () => {
      nonce = await withdrawer.nonce()
      const tx = await withdrawer.initiateWithdrawal(
        recipient.address,
        '3000000',
        [],
        {
          value: utils.parseEther('1'),
        },
      )
      const receipt: ContractReceipt = await tx.wait()
      expect(receipt.events!.length).to.eq(1)
      expect(receipt.events![0].args).to.deep.eq([
        nonce,
        recipient.address,
        recipient.address,
        BigNumber.from(utils.parseEther('1')),
        BigNumber.from('3000000'),
        '0x',
      ])

      burnBlock = await env.l2Provider.getBlock(receipt.blockHash)
      withdrawalHash = utils.keccak256(
        utils.defaultAbiCoder.encode(
          ['uint256', 'address', 'address', 'uint256', 'uint256', 'bytes'],
          [
            utils.hexZeroPad(nonce.toHexString(), 32),
            recipient.address,
            recipient.address,
            utils.hexZeroPad(utils.parseEther('1').toHexString(), 32),
            utils.hexZeroPad(BigNumber.from('3000000').toHexString(), 32),
            '0x',
          ],
        ),
      )
    })

    it('should verify the withdrawal on L1', async () => {
      recipient = recipient.connect(env.l1Provider)
      portal = portal.connect(recipient)
      const oracle = new Contract(
        await portal.L2_ORACLE(),
        l2OOracleArtifact.abi
      ).connect(recipient);

      // await new Promise((resolve) => setTimeout(resolve, 30000));

      let targetOutputTimestamp: number;
      if (
        (await oracle.getL2Output(burnBlock.timestamp)) ==
        "0x" + "00".repeat(32)
      ) {
        await awaitCondition(
          async () => {
            const nextTimestamp = await oracle.nextTimestamp();
            console.log("Burnblock timestamp is:", burnBlock.timestamp);
            console.log(
              "nextTimestamp to append is:",
              nextTimestamp.toNumber()
            );
            console.log(
              "difference:",
              burnBlock.timestamp - nextTimestamp.toNumber()
            );

            const bal = await recipient.getBalance();

            // update this number on each iteration. When the condition is met, this will be the
            // timestamp we want to prove against.
            targetOutputTimestamp = nextTimestamp;
            return nextTimestamp.toNumber() > burnBlock.timestamp;
          },
          2000,
          100
        );
      }

      const storageSlot = "00".repeat(31) + "01"; // i.e the second variable declared in the contract
      const proof = await env.l2Provider.send("eth_getProof", [
        WITHDRAWER_ADDR,
        [utils.keccak256(withdrawalHash + storageSlot)],
        BigNumber.from(targetOutputTimestamp)
          .toHexString()
          .replace("0x0", "0x"),
      ]);

      const { targetStateRoot, targetHash } = await env.l2Provider.send(
        "eth_getBlockByNumber",
        [
          BigNumber.from(targetOutputTimestamp)
            .toHexString()
            .replace("0x0", "0x"),
          false,
        ]
      );

      const tx = await portal.finalizeWithdrawalTransaction(
        nonce,
        recipient.address,
        recipient.address,
        utils.parseEther('1'),
        '3000000',
        '0x',
        targetOutputTimestamp,
        {
          version: constants.HashZero,
          stateRoot: targetStateRoot,
          withdrawerStorageRoot: proof.storageHash,
          latestBlockhash: targetHash,
        },
        rlp.encode(proof.storageProof[0].proof),
        {
          gasLimit: 3_000_000
        }
      )
      const receipt = await tx.wait()
      console.log(receipt)
    }).timeout(120_000)
  })
})
