// Solidity files have to start with this pragma.
// It will be used by the Solidity compiler to validate its version.
pragma solidity ^0.7.3;

contract Deposit {
    /**
     * A function to deposit value for coinbase.
     */
    function deposit() external payable {
        block.coinbase.transfer(msg.value);
    }
}
