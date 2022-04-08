/* Imports: External */
import { Wallet, providers, Contract } from 'ethers'
import { bool, cleanEnv, num, str } from 'envalid'
import dotenv from 'dotenv'

/* Imports: Internal */
const portalArtifact = require('../../../contracts/artifacts/contracts/L1/OptimismPortal.sol/OptimismPortal.json')

dotenv.config()

const procEnv = cleanEnv(process.env, {
  L1_URL: str({default: 'http://localhost:8545'}),
  L1_POLLING_INTERVAL: num({default: 10}),

  L2_URL: str({default: 'http://localhost:9545'}),
  L2_POLLING_INTERVAL: num({default: 1}),

  OPTIMISM_PORTAL_ADDRESS: str(),

  PRIVATE_KEY: str({
    default:
      'ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80',
  }),

  MOCHA_TIMEOUT: num({
    default: 120_000,
  }),
  MOCHA_BAIL: bool({
    default: false,
  }),
})

/// Helper class for instantiating a test environment with a funded account
export class OptimismEnv {
  // The wallets
  l1Wallet: Wallet
  l2Wallet: Wallet

  // The providers
  l1Provider: providers.JsonRpcProvider
  l2Provider: providers.JsonRpcProvider

  // Contracts
  optimismPortal: Contract

  constructor() {
    const l1Provider = new providers.JsonRpcProvider(procEnv.L1_URL)
    l1Provider.pollingInterval = procEnv.L1_POLLING_INTERVAL

    const l2Provider = new providers.JsonRpcProvider(procEnv.L2_URL)
    l2Provider.pollingInterval = procEnv.L2_POLLING_INTERVAL

    const l1Wallet = new Wallet(procEnv.PRIVATE_KEY, l1Provider)
    const l2Wallet = new Wallet(procEnv.PRIVATE_KEY, l2Provider)

    if (!procEnv.OPTIMISM_PORTAL_ADDRESS) {
      throw new Error('Must define an OptimismPortal address.')
    }

    this.l1Wallet = l1Wallet
    this.l2Wallet = l2Wallet
    this.l1Provider = l1Provider
    this.l2Provider = l2Provider
    this.optimismPortal = new Contract(
      procEnv.OPTIMISM_PORTAL_ADDRESS,
      portalArtifact.abi,
    )
  }
}

const env = new OptimismEnv()
export default env