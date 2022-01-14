//SPDX-License-Identifier: MIT
pragma solidity 0.8.11;

/**
 * @title DepositFeed
 */
contract DepositFeed {
    /**
     * Deposits which create a new contract must set the recipient to address(0).
     */
    error NonZeroCreationTarget();

    // Constant for address aliasing
    uint160 private constant OFFSET = uint160(0x1111000000000000000000000000000000001111);

    /**
     * Event with the parameters required to derive transactions on L2.
     */
    event TransactionDeposited(
        address indexed from,
        address indexed to,
        uint256 mint,
        uint256 value,
        uint256 gasLimit,
        bool isCreation,
        bytes data
    );

    /**
     * Accepts deposits of ETH and data, and emits a TransactionDeposited event for use in deriving
     * deposit transactions.
     * @param _to The L2 destination address.
     * @param _value The ETH value to send in the deposit transaction.
     * @param _gasLimit The L2 gasLimit.
     * @param _isCreation Whether or not the transaction should be contract creation.
     * @param _data The input data.
     */
    function depositTransaction(
        address _to,
        uint256 _value,
        uint256 _gasLimit,
        bool _isCreation,
        bytes memory _data
    ) external payable {
        if (_isCreation && _to != address(0)) {
            revert NonZeroCreationTarget();
        }

        address from = msg.sender;
        // Transform the from-address to its alias if the caller is a contract.
        if (msg.sender != tx.origin) {
            unchecked {
                from = address(uint160(msg.sender) + OFFSET);
            }
        }

        emit TransactionDeposited(from, _to, msg.value, _value, _gasLimit, _isCreation, _data);
    }
}
