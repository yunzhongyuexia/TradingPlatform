// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

// 引入OpenZeppelin的ERC20和Ownable标准合约
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

    // 执行空投的函数
//    function airdropNFT(address[] memory recipients) public onlyOwner {
//        require(recipients.length <= 255, "Too many recipients in one transaction"); // 防止一次性空投过多地址
//
//        for (uint256 i = 0; i < recipients.length; i++) {
//            // 假设NFT的tokenId是递增的
//            uint256 tokenId = 1 + i + airdropAmount * (block.timestamp / 1 days); // 简单的示例，实际中tokenId应由业务逻辑决定
//            token.mint(recipients[i], tokenId);
//        }
//    }
}
