package service_eth

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"strings"
)

var (
	NFTClient *ethclient.Client
	JSONABI   abi.ABI
)

const (
	// ClientURL holesky测试网
	holeskyURL = "https://ethereum-holesky.core.chainstack.com/5cb26e90d4925a7f24a1c7b51c2d8263"

	//测试链ID
	chainID = 17000

	// ABI文件地址
	abiFileAddress = "D:/project/TradingPlatform/server/config/MintToken.abi"

	// Account1 私钥
	Account1 = "3b35e382c432bfdfbcafcdf5e428fff8998ccce50204ce0336025e5a01bd63e6"

	// Account1Address Account1地址
	Account1Address = "0x44cc3545d27fc25a4fe827e76ade6f74d4ac7a02"

	// Account2 私钥
	Account2 = "5e1d6f1a64cc1d3ad1c155fcdb3165a0a844d797a6e99f770910fb5bc81134b3"

	// Account2Address Account2地址
	Account2Address = "0x8ff0bf277ea5ae4e3047c178923d973e6eac27f1"

	// 你的NFT智能合约地址
	nftContractAddress = "0x2A650BEBBC791B6Fefbe6F437916D7662Da15877"
)

func ConnectNFTClient() {
	// 连接到以太坊节点
	client, err := ethclient.Dial(holeskyURL)
	if err != nil {
		log.Fatal(err)
	}

	NFTClient = client

	// 从文件系统中读取ABI
	abiBytes, err := ioutil.ReadFile(abiFileAddress)
	if err != nil {
		log.Fatalf("Failed to read ABI file: %v", err)
	}

	// 解析ABI
	contractABI, err := abi.JSON(strings.NewReader(string(abiBytes)))
	if err != nil {
		log.Fatalf("failed to parse ABI: %v", err)
	}
	JSONABI = contractABI

	//打印api
	for _, method := range contractABI.Methods {
		fmt.Printf("ABI method: %s\n", method.String())
	}

	// 将ABI序列化为[]byte
	//abiJSON, err := json.Marshal(contractABI)
	//if err != nil {
	//	log.Fatalf("failed to marshal ABI to JSON: %v", err)
	//}
}

//Method: transfer,Inputs: [{to address false} {value uint256 false}], Outputs: [{ bool false}]
//Method: changeMinter,Inputs: [{newMinter address false}], Outputs: []
//Method: symbol,Inputs: [], Outputs: [{ string false}]
//Method: approve,Inputs: [{spender address false} {value uint256 false}], Outputs: [{ bool false}]
//Method: balanceOf,Inputs: [{account address false}], Outputs: [{ uint256 false}]
//Method: decimals,Inputs: [], Outputs: [{ uint8 false}]
//Method: minter,Inputs: [], Outputs: [{ address false}]
//Method: name,Inputs: [], Outputs: [{ string false}]
//Method: totalSupply,Inputs: [], Outputs: [{ uint256 false}]
//Method: transferFrom,Inputs: [{from address false} {to address false} {value uint256 false}], Outputs: [{ bool false}]
//Method: allowance,Inputs: [{owner address false} {spender address false}], Outputs: [{ uint256 false}]
//Method: mintTokens,Inputs: [{to address false} {amount uint256 false}], Outputs: []

//transfer: 允许账户发送指定数量的代币到另一个账户。需要接收方地址和代币数量作为输入。
//
//changeMinter: 更改合约的minter角色，通常用于多签钱包或权限管理。输入参数为新minter的地址。
//
//symbol: 返回代币的符号，例如"ETH"、"DAI"等。
//
//approve: 允许账户授权第三方（spender）使用指定数量的代币。这通常用于允许代币在不转移所有权的情况下被使用或转移。
//
//balanceOf: 查询指定账户的代币余额。
//
//decimals: 返回代币的小数位数，通常ERC-20代币的小数位数为18。
//
//minter: 查询当前的minter地址。
//
//name: 返回代币的名称，例如"Ethereum"、"Maker"等。
//
//totalSupply: 返回代币的总供应量。
//
//transferFrom: 从授权的账户转移代币到另一个账户。这通常用于实现代币的代理交易。
//
//allowance: 查询某个账户授权给另一个账户使用的代币数量。
//
//mintTokens: 铸造新代币并分配给指定账户。这个方法的存在表明这个代币合约可能是一个可增发的代币合约
