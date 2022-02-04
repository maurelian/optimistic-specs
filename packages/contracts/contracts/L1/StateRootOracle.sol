// SPDX-License-Identifier: MIT
pragma solidity >0.5.0 <0.9.0;

// This should be an UPGRADABLE CONTRACT.
// We will eventually need to add support for concurrent
// state root proposals, but until we have a full dispute
// game spec that we are happy with, it is better to KEEP IT SIMPLE.
contract StateRootOracle {
    uint256 public submissionFrequency;
    uint256 public l2BlockTime;
    mapping(uint256 => bytes32) public stateRoots;
    uint256 public latestBlockTimestamp;
    uint256 public startingBlockTimestamp;
    uint256 public historicalTotalBlocks;

    /**********
     * Events *
     **********/

    event StateRootAppended(uint256 indexed _timestamp, bytes32 _stateRoot);

    event StateRootDeleted(uint256 indexed _timestamp, bytes32 _stateRoot);

    constructor(
        uint256 _submissionFrequency,
        uint256 _l2BlockTime,
        bytes32 _genesisRoot,
        uint256 _historicalTotalBlocks
    ) {
        require(_submissionFrequency > 0, "submission frequency must be positive");
        require(_l2BlockTime > 0, "l2BlockTime must be positive");

        submissionFrequency = _submissionFrequency;
        l2BlockTime = _l2BlockTime;
        stateRoots[block.timestamp] = _genesisRoot;
        latestBlockTimestamp = block.timestamp;
        startingBlockTimestamp = block.timestamp;
        historicalTotalBlocks = _historicalTotalBlocks;
    }

    function appendStateRoot(bytes32 _stateRoot, uint256 _timestamp) external {
        require(
            block.timestamp > _timestamp,
            "Cannot append state roots from the future. Come on bruh."
        );
        require(_stateRoot != bytes32(0), "Cannot submit empty state root");
        require(
            _timestamp == latestBlockTimestamp + submissionFrequency,
            "Must submit state root for every 25 minutes"
        );
        stateRoots[_timestamp] = _stateRoot;
        latestBlockTimestamp = _timestamp;
        emit StateRootAppended(_timestamp, _stateRoot);
    }

    function getTotalSubmittedStateRoots() public view returns (uint256) {
        return (latestBlockTimestamp - startingBlockTimestamp) / submissionFrequency;
    }

    function deleteStateBatch(uint256 _timestamp) external {
        uint256 _nextTimestamp = _timestamp + submissionFrequency;
        require(stateRoots[_nextTimestamp] == bytes32(0), "Must delete tip state root.");
        bytes32 _oldStateRoot = stateRoots[_timestamp];
        stateRoots[_timestamp] = bytes32(0);
        emit StateRootDeleted(_timestamp, _oldStateRoot);
        return;
    }

    function verifyStateCommitment()
        external
        pure
        returns (
            // bytes32 _element,
            // bytes32 _stateRoot,
            // bytes memory _proof
            bool _verified
        )
    {
        // TODO -- Merkle proof
        // TODO -- Decide if we want to put the withdrawals closer to the root
        // of the
        // commitment.
        return false;
    }

    function insideFraudProofWindow(uint256 _timestamp) external view returns (bool _inside) {
        require(stateRoots[_timestamp] != bytes32(0), "State root not submitted yet.");
        // if stateRoots[_blockNumber].timestamp is 1 week old) {
        //     return true
        // }
        return false;
    }

    function nextTimestamp() external view returns (uint256) {
        return latestBlockTimestamp + submissionFrequency;
    }

    function computeL2BlockNumber(uint256 _timestamp) external view returns (uint256) {
        require(_timestamp >= startingBlockTimestamp, "timestamp before start");
        return historicalTotalBlocks + (_timestamp - startingBlockTimestamp) / l2BlockTime;
    }
}
