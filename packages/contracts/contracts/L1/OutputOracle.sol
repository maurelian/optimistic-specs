// SPDX-License-Identifier: MIT
pragma solidity >0.5.0 <0.9.0;

// This should be an UPGRADABLE CONTRACT.
// We will eventually need to add support for concurrent
// state root proposals, but until we have a full dispute
// game spec that we are happy with, it is better to KEEP IT SIMPLE.
contract OutputOracle {
    // Todo: change naming from StateRoot to Output

    // Constants
    // Submit a state root every 100 L1 blocks
    uint256 public submissionFrequency = 100;

    // This struct can be stored as a hash to optimize storage costs.
    // However, for testing purposes this is fine. Plus this is not
    // a very hot code path so I prefer at least starting out with a
    // readable version.
    struct StateRootSubmission {
        bytes32 stateRoot;
        uint256 timestamp;
    }
    mapping(uint256=>StateRootSubmission) public stateRoots;
    uint256 public latestBlockNumber;
    uint256 public startingBlockNumber;

    /**********
     * Events *
     **********/

    event StateRootAppended(
        uint256 indexed _blockNumber,
        bytes32 _stateRoot
    );

    event StateRootDeleted(uint256 indexed _blockNumber, bytes32 _stateRoot);


    constructor(
        bytes32 _genesisRoot
    ) {
        stateRoots[block.number] = StateRootSubmission({
            stateRoot: _genesisRoot,
            timestamp: block.timestamp
        });
        latestBlockNumber = block.number;
        startingBlockNumber = block.number;
    }

    function appendStateRoot(bytes32 _stateRoot, uint256 _blockNumber) external {
        require(block.number > _blockNumber, "Cannot append state roots from the future. Come on bruh.");
        require(_stateRoot != bytes32(0), "Cannot submit empty state root");
        require(_blockNumber == latestBlockNumber + submissionFrequency, "Must submit state root for every 100 blocks");
        stateRoots[_blockNumber] = StateRootSubmission({
            stateRoot: _stateRoot,
            timestamp: block.timestamp
        });
        emit StateRootAppended(_blockNumber, _stateRoot);
    }

    function getTotalSubmittedStateRoots() public view returns(uint256) {
        return (latestBlockNumber - startingBlockNumber) / submissionFrequency;
    }

    function deleteStateBatch(uint256 _blockNumber) external {
        uint256 _nextBlockNum = _blockNumber + submissionFrequency;
        require(stateRoots[_nextBlockNum].stateRoot == bytes32(0), "Must delete tip state root.");
        require(stateRoots[_nextBlockNum].timestamp == uint256(0), "Must delete tip state root.");
        bytes32 _oldStateRoot = stateRoots[_blockNumber].stateRoot;
        stateRoots[_blockNumber] = StateRootSubmission({
            stateRoot: bytes32(0),
            timestamp: uint256(0)
        });
        emit StateRootDeleted(_blockNumber, _oldStateRoot);
        return;
    }

    function verifyStateCommitment(
        bytes32 _element,
        bytes32 _stateRoot,
        bytes memory _proof
    ) external view returns (bool _verified) {
        // TODO -- Merkle proof
        // TODO -- Decide if we want to put the withdrawals closer to the root of the
        // commitment.
        return false;
    }

    function insideFraudProofWindow(uint256 _blockNumber)
        external
        view
        returns (bool _inside) {
            require(stateRoots[_blockNumber].timestamp != uint256(0), "State root not submitted yet.");
            // if stateRoots[_blockNumber].timestamp is 1 week old) {
            //     return true
            // }
            return false;
        }
}hh