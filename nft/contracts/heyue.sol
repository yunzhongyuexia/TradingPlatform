// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;//版本号和迁移脚本一致

// 引入OpenZeppelin的ERC721标准合约
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract MintNFT is ERC721, Ownable {
    // 定义NFT的结构体，包含id，名称和价值
    struct NFT {
        uint256 id;
        string name;
        uint256 value;
    }

    // 存储NFT的信息
    mapping(uint256 => NFT) public nfts;
    uint256 public nextNFTId = 1;

    // 更改合约拥有者
    function transferOwnership(address newOwner) public onlyOwner {
        _transferOwnership(newOwner);
    }

    // 重写ERC721的ownerOf函数，返回NFT的额外信息
    function ownerOf(uint256 tokenId) public view override returns (address owner, NFT memory nft) {
        owner = _ownerOf(tokenId);
        nft = nfts[tokenId];
    }

    // 铸造新的NFT并分配给指定的地址
    function mintNFT(address to, string memory name, uint256 value) public onlyOwner {
        uint256 tokenId = nextNFTId++;
        _nft = nfts[tokenId] = NFT(tokenId, name, value);
        _mint(to, tokenId);
    }

    // 执行合成的函数，将多个NFT合成一个新的NFT
    function synthesisNFT(uint256[] memory tokenIds, string memory name, uint256 value) public {
        require(tokenIds.length > 1, "At least two NFTs are required for synthesis");
        require(_isApprovedOrOwner(_msgSender(), tokenIds[0]), "Not owner nor approved");

        // 合成逻辑，这里简单地取第一个NFT的id作为新NFT的id
        uint256 newNFTId = tokenIds[0];
        _nft = nfts[newNFTId] = NFT(newNFTId, name, value);

        // 将用于合成的原始NFT设置为不存在
        for (uint256 i = 0; i < tokenIds.length; i++) {
            _burn(tokenIds[i]);
        }

        // 将新合成的NFT分配给调用者
        _mint(_msgSender(), newNFTId);
    }

    // 空投NFT的函数
    function airdropNFT(address[] memory _recipients, uint256[] memory _values) public onlyOwner {
        require(_recipients.length == _values.length, "Arrays length must match");

        for (uint256 i = 0; i < _recipients.length; i++) {
            // 铸造新的NFT并分配给指定的地址
            _mint(_recipients[i], nextNFTId++);
            // 为新铸造的NFT设置属性
            nfts[nextNFTId - 1] = NFT(nextNFTId - 1, "Unique NFT Name", _values[i]);
        }
    }
}