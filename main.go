package main

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

var (
	contractAddress     = "0x83642581DB08De4C99726102Fa6ee559cCcE9eD9"
	contractOwnerPriKey = "**"
	rpcClient           = "https://rinkeby.infura.io/v3/5f12b14e30d04834b2a1a870cb6cd1b5"
	chainId             = 4
)

func main() {
	// 连接rpc客户端
	client, err := ethclient.Dial(rpcClient)
	if err != nil {
		log.Fatalf("链接到RPC客户端失败! err: %v", err)
	}
	defer client.Close()

	// 读取私钥
	privateKey, err := crypto.HexToECDSA(contractOwnerPriKey)
	if err != nil {
		log.Fatal(err)
	}
	// 获取到私钥对应的公钥
	publicKey := privateKey.Public()
	// 签名事务
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// 查询nonce与gas的价格
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(int64(chainId)))
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// 开始调用合约
	address := common.HexToAddress(contractAddress)
	act, err := NewCallcontractTransactor(address, client)
	if err != nil {
		log.Fatal(err)
	}
	tx, err := act.Mint(auth, common.HexToAddress("0x130C66f906e03e371D25e447E1fe7d7331631D5a"), new(big.Int).SetInt64(1000000000000000000))
	print("交易哈希为：", tx)
}
