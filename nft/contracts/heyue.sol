// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

// MintToken 就是编译后生成的json文件的文件名
contract MintToken is ERC20 {
    address public minter;

    event MinterChanged(address indexed from, address to);

    constructor() ERC20("AToken", "ATK") {
        minter = msg.sender;
    }

    function mintTokens(address to, uint256 amount) external {
        require(msg.sender == minter, "Only minter can mint tokens");
        _mint(to, amount);
    }

    function changeMinter(address newMinter) external {
        require(msg.sender == minter, "Only current minter can change minter");
        emit MinterChanged(minter, newMinter);
        minter = newMinter;
    }
}

