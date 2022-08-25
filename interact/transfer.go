package interact

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

// TransferFrom
//
//	@Description: 这里仅实现了从私钥存储的账户向其他账户进行转账。因为需要私钥进行签名交易。
//		作为一个后端程序，拿到用户的私钥是相当不容易的。但是，也没必要拿到。
//	@Author: tianyuan01528@foxmail.com
//	@param fromAddress
//	@param toAddress
//	@param amount
func TransferFrom(userAuth, toAddress string, amount int64) {
	// 连接rpc客户端
	client, err := ethclient.Dial(rpcClient)
	if err != nil {
		log.Fatalf("链接到RPC客户端失败! err: %v", err)
	}
	defer client.Close()

	// 读取私钥
	privateKey, err := crypto.HexToECDSA(contractOwnerPriKey)
	if err != nil {
		log.Fatalf("解析私钥失败! err: %v", err)
	}
	// 获取到私钥对应的公钥
	publicKey := privateKey.Public()
	// 签名事务
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	pubAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// 查询nonce与gas的价格
	nonce, err := client.PendingNonceAt(context.Background(), pubAddress)
	if err != nil {
		log.Fatalf("从RPC客户端查询nonce失败! err: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("从RPC客户端查询gas失败! err: %v", err)
	}

	//  获取交易签名信息
	// 	Kick！ 这里正在寻找远程签名的方法。目前的解决方案最好还是使用前端与用户做交互，
	//	将交易信息传递给用户进行签名，之后返回被签名的信息并广播。
	// 	这里名没有给出实现哈哈哈哈哈哈。具体的实现参照  https://goethereumbook.org/zh/transfer-tokens/
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(int64(chainId)))
	if err != nil {
		log.Fatalf("创建Transactor失败! err: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(contractAddress)
	act, err := NewCallcontractTransactor(address, client)
	if err != nil {
		log.Fatalf("合约调用异常! err: %v", err)
	}
	decimal := new(big.Int).SetInt64(1000000000000000000)
	value := new(big.Int).SetInt64(amount)
	tx, err := act.Transfer(auth, common.HexToAddress(toAddress), new(big.Int).Mul(decimal, value))
	println("转账给:", toAddress, "总量:", value, " 成功! 交易哈希为：", tx)
}
