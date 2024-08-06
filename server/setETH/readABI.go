package setETH

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"io/ioutil"
	"log"
	"strings"
)

// ABI文件地址
var abiFileAddress = "D:/project/TradingPlatform/server/setETH/MintToken.abi"

// ReadABIFromFile 从指定路径读取ABI文件并解析为ABI结构
func ReadABIFromFile() {

	// 读取ABI文件
	abiFile, err := ioutil.ReadFile(abiFileAddress)
	if err != nil {
		log.Fatal("failed to read ABI file: %w", err)
	}

	// 解析ABI
	contractABI, err := abi.JSON(strings.NewReader(string(abiFile)))
	if err != nil {
		log.Fatal("failed to parse ABI: %w", err)
	}

	//打印方法
	for name, method := range contractABI.Methods {
		fmt.Printf("Method: %s,Inputs: %v, Outputs: %v\n", name, method.Inputs, method.Outputs)
	}
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
