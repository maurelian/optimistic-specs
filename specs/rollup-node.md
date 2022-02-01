# Rollup Node Specification

<!-- All glossary references in this file. -->
[g-rollup-node]: glossary.md#rollup-node
[g-derivation]: glossary.md#L2-chain-derivation
[g-payload-attr]: glossary.md#payload-attributes
[g-block]: glossary.md#block
[g-exec-engine]: glossary.md#execution-engine
[g-reorg]: glossary.md#re-organization
[g-rollup-driver]: glossary.md#rollup-driver
[g-inception]: glossary.md#L2-chain-inception
[g-receipts]: glossary.md#receipt
[g-deposit-contract]: glossary.md#deposit-contract
[g-deposits]: glossary.md#deposits
[g-deposit-block]: glossary.md#deposit-block
[g-deposited]: glossary.md#deposited-transaction
[g-l1-attr-deposit]: glossary.md#l1-attributes-deposited-transaction
[g-user-deposited]: glossary.md#user-deposited-transaction
[g-l1-attr-predeploy]: glossary.md#l1-attributes-predeployed-contract
[g-depositing-call]: glossary.md#depositing-call
[g-depositing-transaction]: glossary.md#depositing-transaction
[g-mpt]: glossary.md#merkle-patricia-trie

The [rollup node][g-rollup-node] is the component responsible for [deriving the L2 chain][g-derivation] from L1 blocks
(and their associated [receipts][g-receipts]). This process happens in two steps:

1. Read from L1 blocks and associated receipts, in order to generate [payload attributes][g-payload-attr] (essentially
   [a block without output properties][g-block]).
2. Pass the payload attributes to the [execution engine][g-exec-engine], so that the L2 block (including [output block
   properties][g-block]) may be computed.

While this process is conceptually a pure function from the L1 chain to the L2 chain, it is in practice incremental. The
L2 chain is extended whenever new L1 blocks are added to the L1 chain. Similarly, the L2 chain re-organizes whenever the
L1 chain [re-organizes][g-reorg].

The part of the rollup node that derives the L2 chain is called the [rollup driver][g-rollup-driver]. This document is
currently only concerned with the specification of the rollup driver.

## Table of Contents

- [L2 Chain Derivation](#l2-chain-derivation)
  - [From L1 Blocks to Payload Attributes](#from-l1-blocks-to-payload-attributes)
    - [Reading L1 Inputs](#reading-l1-inputs)
    - [Encoding the L1 Attributes Deposited Transaction](#encoding-the-l1-attributes-deposited-transaction)
    - [Encoding User-Deposited Transactions](#encoding-user-deposited-transactions)
    - [Building the Payload Attributes](#building-the-payload-attributes)
  - [From Payload Attributes to L2 Block](#from-payload-attributes-to-l2-block)
    - [Inductive Derivation Step](#inductive-derivation-step)
    - [Engine API Error Handling](#engine-api-error-handling)
    - [Finalization Guarantees](#finalization-guarantees)
  - [From L2 Block to L2 Output Root](#from-l2-block-to-l2-output-root)
  - [Whole L2 Chain Derivation](#whole-l2-chain-derivation)
- [Handling L1 Re-Orgs](#handling-l1-re-orgs)

# L2 Chain Derivation

[l2-chain-derivation]: #l2-chain-derivation

This section specifies how the [rollup driver][g-rollup-driver] derives one L2 [deposit block][g-deposit-block] per
every L1 block. The L2 block carries *[deposited transactions][g-deposited]* of two kinds:

- a single *[L1 attributes deposited transaction][g-l1-attr-deposit]* (always first)
- zero or more *[user-deposited transactions][g-user-deposited]*

------------------------------------------------------------------------------------------------------------------------

## From L1 Blocks to Payload Attributes

### Reading L1 inputs

The rollup reads the following data from each L1 block:

- L1 block attributes
  - block number
  - timestamp
  - basefee
  - *random* (the output of the [`RANDOM` opcode][random])
- L1 log entries emitted for [user deposits][g-deposits], augmented with
  - `blockHeight`: the block-height of the L1 block
  - `transactionIndex`: the transaction-index within the L2 transactions list

[random]: https://eips.ethereum.org/EIPS/eip-4399

> Design note: The extra log entry metadata will be used to ensure that deposited transactions will be unique. Without
> them, two different deposited transaction could have the same exact hash.
>
> We do not use the sender's nonce to ensure uniqueness because this would require an extra L2 EVM state read from the
> [execution engine][g-exec-engine].

The L1 attributes are read from the L1 block header, while deposits are read from the block's [receipts][g-receipts].
Refer to the [**deposit contract specification**][deposit-contract-spec] for details on how deposits are encoded as log
entries.

[deposit-contract-spec]: deposits.md#deposit-contract

### Encoding the L1 Attributes Deposited Transaction

The [L1 attributes deposited transaction][g-l1-attr-deposit] is a call that submits the L1 block attributes (listed
above) to the [L1 attributes predeployed contract][g-l1-attr-predeploy].

To encode the L1 attributes deposited transaction, refer to the following sections of the deposits spec:

- [The Deposited Transaction Type](deposits.md#the-deposited-transaction-type)
- [L1 Attributes Deposited Transaction](deposits.md#l1-attributes-deposited-transaction)

### Encoding User-Deposited Transactions

A [user-deposited-transactions][g-deposited] is an L2 transaction derived from a [user deposit][g-deposits] submitted on
L1 to the [deposit contract][g-deposit-contract]. Refer to the [deposit contract specification][deposit-contract-spec]
for more details.

The user-deposited transaction is derived from the log entry emitted by the [depositing call][g-depositing-call], which
is stored in the [depositing transaction][g-depositing-transaction]'s log receipt.

To encode user-deposited transactions, refer to the following sections of the deposits spec:

- [The Deposited Transaction Type](deposits.md#the-deposited-transaction-type)
- [User-Deposited Transactions](deposits.md#user-deposited-transactions)

### Building the Payload Attributes

[payload attributes]: #building-the-payload-attributes

From the data read from L1 and the encoded transactions, the rollup node constructs the [payload
attributes][g-payload-attr] as an [expanded version][expanded-paylod] of the [`PayloadAttributesV1`] object, which
includes an additional `transactions` field.

The object's properties must be set as follows:

- `timestamp` is set to the timestamp of the L1 block.
- `random` is set to the *random* L1 block attribute
- `suggestedFeeRecipient` is set to the zero-address for deposit-blocks, since there is no sequencer.
- `transactions` is an array of the derived deposits, encoded as per the two preceding sections.

[expanded-payload]: exec-engine.md#extended-payloadattributesv1
[`PayloadAttributesV1`]: https://github.com/ethereum/execution-apis/blob/main/src/engine/specification.md#payloadattributesv1

------------------------------------------------------------------------------------------------------------------------

## From Payload Attributes to L2 Block

Once the [payload attributes] for a given L1 block `B` have been built, and if we have already derived an L2 block from
`B`'s parent block, then we can use the payload attributes to derive a new L2 block.

### Inductive Derivation Step

Let

- `refL2` be the (hash of) the current L2 chain head
- `refL1` be the (hash of) the L1 block from which `refL2` was derived
- `payloadAttributes` be some previously derived [payload attributes] for the L1 block with number `l1Number(refL1) + 1`

Then we can apply the following pseudocode logic to update the state of both the rollup driver and execution engine:

```javascript
// request a new execution payload
forkChoiceState = {
    headBlockHash: refL2,
    safeBlockHash: refL2,
    finalizedBlockHash: l2BlockHashAt(l2Number(refL2) - FINALIZATION_DELAY_BLOCKS)
}
[status, payloadID] = engine_forkchoiceUpdatedV1(forkChoiceState, payloadAttributes)
if (status != "SUCCESS") error()

// retrieve and execute the execution payload
[executionPayload, error] = engine_getPayloadV1(payloadID)
if (error != null) error()
[status, latestValidHash, validationError] = engine_executePayloadV1(executionPayload)
if (status != "VALID" || validationError != null) error()

refL2 = latestValidHash
refL1 = l1HashForNumber(l1Number(refL1) + 1))

// update head to new refL2
forkChoiceState = {
    headBlockHash: refL2,
    safeBlockHash: refL2,
    finalizedBlockHash: l2BlockHashAt(l2Number(headBlockHash) - FINALIZATION_DELAY_BLOCKS)
}
[status, payloadID] = engine_forkchoiceUpdatedV1(refL2, null)
if (status != "SUCCESS") error()
```

The following JSON-RPC methods are part of the [execution engine API][exec-engine]:

> **TODO** fortify the execution engine spec with more information regarding JSON-RPC, notably covering
> information found [here][json-rpc-info-1] and [here][json-rpc-info-2]

[json-rpc-info-1]: https://github.com/ethereum-optimism/optimistic-specs/blob/a3ffa9a8c825d155a0469659b3101db5f41eecc4/specs/rollup-node.md#from-l1-blocks-to-payload-attributes
[json-rpc-info-2]: https://github.com/ethereum-optimism/optimistic-specs/blob/a3ffa9a8c825d155a0469659b3101db5f41eecc4/specs/rollup-node.md#building-the-l2-block-with-the-execution-engine

[exec-engine]: exec-engine.md

- [`engine_forkchoiceUpdatedV1`] — updates the forkchoice (i.e. the chain head) to `headBlockHash` if different, and
  instructs the engine to start building an execution payload given payload attributes the second argument isn't `null`
- [`engine_getPayloadV1`] — retrieves a previously requested execution payload
- [`engine_executePayloadV1`] — executes an execution payload to create a block

[`engine_forkchoiceUpdatedV1`]: exec-engine.md#engine_forkchoiceUpdatedV1
[`engine_getPayloadV1`]: exec-engine.md#engine_executepayloadv1
[`engine_executePayloadV1`]: exec-engine.md#engine_executepayloadv1

The execution payload is an object of type [`ExecutionPayloadV1`].

[`ExecutionPayloadV1`]: https://github.com/ethereum/execution-apis/blob/main/src/engine/specification.md#executionpayloadv1

Within the `forkChoiceState` object, the properties have the following meaning:

- `headBlockHash`: block hash of the last block of the L2 chain, according to the rollup driver.
- `safeBlockHash`: same as `headBlockHash`.
- `finalizedBlockHash`: the hash of the block whose number is `l2Number(headBlockHash) - FINALIZATION_DELAY_BLOCKS` if
  the number of that block is `>= L2_CHAIN_INCEPTION`, 0 otherwise (\*) See the [Finalization Guarantees][finalization]
  section for more details.

(\*) where:

- `FINALIZATION_DELAY_BLOCKS == 50400` (approximately 7 days worth of L1 blocks)
- `L2_CHAIN_INCEPTION` is the [L2 chain inception][g-inception] (the number of the first L1 block for which an L2 block
  was produced).

Finally, the `error()` function signals an error that must be handled by the implementation. Refer to the next section
for more details.

### Engine API Error Handling

[error-handling]: #engine-api-error-handling

All invocations of [`engine_forkchoiceUpdatedV1`], [`engine_getPayloadV1`] and [`engine_executePayloadV1`] by the
rollup driver should not result in errors assuming conformity with the specification. Said otherwise, all errors are
implementation concerns and it is up to them to handle them (e.g. by retrying, or by stopping the chain derivation and
requiring manual user intervention).

The following scenarios are assimilated to errors:

- [`engine_forkchoiceUpdatedV1`] returning a `status` of `"SYNCING"` instead of `"SUCCESS"` whenever passed a
  `headBlockHash` that it retrieved from a previous call to [`engine_executePayloadV1`].
- [`engine_executePayloadV1`] returning a `status` of `"SYNCING"` or `"INVALID"` whenever passed an execution payload
  that was obtained by a previous call to [`engine_getPayloadV1`].

### Finalization Guarantees

[finalization]: #finalization-guarantees

As stated earlier, an L2 block is considered *finalized* after a delay of `FINALIZATION_DELAY_BLOCKS == 50400` L1 blocks
after the L1 block that generated it. This is a duration of approximately 7 days worth of L1 blocks. This is also known
as the "fault proof window", as after this time the block can no longer be challenged by a fault proof.

L1 Ethereum [reaches finality approximately every 12 minutes][l1-finality]. L2 blocks generated from finalized L1 blocks
are "safer" than most recent L2 blocks because they will never disappear from the chain's history because of a re-org.
However, they can still be challenged by a fault proof until the end of the fault proof window.

[l1-finality]: https://www.paradigm.xyz/2021/07/ethereum-reorgs-after-the-merge/

> **TODO** the spec doesn't encode the notion of fault proof yet, revisit this (and include links) when it does

## From L2 Block to L2 Output Root

After processing a block the resulting outputs will need to be synchronized with L1 for trustless execution of
L2-to-L1 messaging, such as withdrawals. To synchronize outputs are merkleized: hashed in a structured form for minimal
proof cost to any piece of data. The merkle-structure is defined with [SSZ], a type system for merkleization and
serialization, used in L1 (beacon-chain). However, we replace `sha256` with `keccak256` to save gas costs in the EVM.

[SSZ]: https://github.com/ethereum/consensus-specs/blob/dev/ssz/simple-serialize.md

```python
class L2Output(Container):
  state_root: Bytes32
  withdrawal_storage_root: Bytes32
  latest_block: ExecutionPayload  # includes block-hash
  history_accumulator_root: Bytes32  # Not functional yet
  extension: Bytes32
```

The `state_root` is the Merkle-Patricia-Trie ([MPT][g-mpt]) root of all execution-layer accounts,
also found in `latest_block.state_root`: this field is frequently used and thus elevated closer to the L2 output root,
reducing the merkle proof depth and thus the cost of usage.

The `withdrawal_storage_root` elevates the Merkle-Patricia-Trie ([MPT][g-mpt]) root of L2 Withdrawal contract storage.
Instead of a MPT proof to the Withdrawal contract account in the account trie,
one can directly access the MPT storage trie root, thus reducing the verification cost of withdrawals on L1.

The `latest_block` is an execution-layer block of L2, represented as the [`ExecutionPayload`][ExecutionPayload] SSZ type
defined in L1. There may be multiple blocks per L2 output root, only the latest is presented.

[ExecutionPayload]: https://github.com/ethereum/consensus-specs/blob/dev/specs/bellatrix/beacon-chain.md#executionpayload

The `history_accumulator_root` is a reserved field, elevating a storage variable of the L2 chain that maintains
the SSZ merkle root of an append-only `List[Bytes32, MAX_ITEM_COUNT]` (`keccak256` SSZ),
where each item is defined as `keccak256(l2_block_hash ++ l2_state_root)`, one per block of the L2 chain.
While reserved, a zeroed `Bytes32` is used instead.
This is a work-in-progress, see [issue 181](https://github.com/ethereum-optimism/optimistic-specs/issues/181).

The `extension` is a zeroed `Bytes32`, to be substituted with a SSZ container to extend merkleized information in future
upgrades. This keeps the static merkle structure forwards-compatible.

## Whole L2 Chain Derivation

The previous two sections presents an inductive process: given that we know the "current" L2 block, well as the next L1
block, then we can derive [payload attributes] for the next L1 block, and from that the next L2 block.

To derive the whole L2 chain from scratch, we simply start with the L2 genesis block as the current L2 block, and the
block at height `L2_CHAIN_INCEPTION + 1` as the next L1 block. Then we iteratively apply the derivation process from the
previous section to each successive L1 block until we have caught up with the L1 head.

> **TODO** specify genesis block

# Handling L1 Re-Orgs

[l1-reorgs]: #handling-L1-re-orgs

The [previous section on L2 chain derivation][l2-chain-derivation] assumes linear progression of the L1 chain. It is
also applicable for batch processing, meaning that any given point in time, the canonical L2 chain is given by
processing the whole L1 chain since the [L2 chain inception][g-inception].

> By itself, the previous section fully specifies the behaviour of the rollup driver. **The current section is
> non-specificative** but shows how L1 re-orgs can be handled in practice.

In practice, the L1 chain is processed incrementally. However, the L1 chain may occasionally [re-organize][g-reorg],
meaning the head of the L1 chain changes to a block that is not the child of the previous head but rather another
descendant of an ancestor of the previous head. In that case, the rollup driver must first search for the common L1
ancestor, and can re-derive the L2 chain from that L1 block and onwards.

The starting point of the re-derivation is a pair `(refL2, nextRefL1)` where `refL2` refers to the L2 block to build
upon and `nextRefL1` refers to the next L1 block to derive from (i.e. if `refL2` is derived from L1 block `refL1`,
`nextRefL1` is the canonicla L1 block at height `l1Number(refL1) + 1`).

In practice, the happy path (no re-org) and the re-org paths are merged. The happy path is simply a special case of the
re-org path where the starting point of the re-derivation is `(currentL2Head, newL1Block)`.

This re-derivation starting point can be found by applying the following algorithm:

1. (Initialization) Set the initial `refL2` to the head block of the L2 execution engine.
2. Set `parentL2` to `refL2`'s parent block and `refL1` to the L1 block that `refL2` was derived from.
3. Fetch `currentL1`, the canonical L1 block at the same height as `refL1`.

- If `currentL1 == refL1`, then `refL2` was built on a canonical L1 block:
  - Find the next L1 block (it may not exist yet) and return `(refL2, nextRefL1)` as the starting point of the
    re-derivation.
    - It is necessary to ensure that no L1 re-org occured during this lookup, i.e. that `nextRefL1.parent == refL1`.
    - If the next L1 block does not exist yet, there is no re-org, and nothing new to derive, and we can abort the
        process.
- Otherwise, if `refL2` is the L2 genesis block, we have re-orged past the genesis block, which is an error that
    requires a re-genesis of the L2 chain to fix (i.e. creating a new genesis
    configuration) (\*)
- Otherwise, if either `currentL1` does not exist, or `currentL1 != refL1`, set `refL2` to `parentL2` and restart this
    algorithm from step 2.
  - Note: if `currentL1` does not exist, it means we are in a re-org to a shorter L1 chain.
  - Note: as an optimization, we can cache `currentL1` and reuse it as the next value of `nextRefL1` to avoid an
        extra lookup.

Note that post-[merge], the depth of re-orgs will be bounded by the [L1 finality delay][l1-finality] (every 2 epochs,
approximately 12 minutes).

(\*) Post-merge, this is only possible for 12 minutes. In practice, we'll pick an already-finalized L1 block as L2
inception point to preclude the possibility of a re-org past genesis, at the cost of a few empty blocks at the start of
the L2 chain.

[merge]: https://ethereum.org/en/eth2/merge/
