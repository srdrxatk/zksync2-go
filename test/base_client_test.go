package test

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/clients"
	"github.com/zksync-sdk/zksync2-go/contracts/erc20"
	zkTypes "github.com/zksync-sdk/zksync2-go/types"
	"github.com/zksync-sdk/zksync2-go/utils"
	"math/big"
	"testing"
)

func TestIntegrationBaseClient_Dial(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")
	assert.NotNil(t, client, "clients.Dial should return a non-nil client")
}

func TestIntegrationBaseClient_DialContext(t *testing.T) {
	client, err := clients.DialContext(context.Background(), ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.DialContext should not return an error")
	assert.NotNil(t, client, "clients.DialContext should return a non-nil client")
}

func TestIntegrationBaseClient_ChainID(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	chainID, err := client.ChainID(context.Background())

	assert.NoError(t, err, "ChainID should not return an error")
	assert.NotNil(t, chainID, "ChainID should return a non-nil chainID")
}

func TestIntegrationBaseClient_BlockByHash(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	blockTmp, err := client.BlockByNumber(context.Background(), nil)
	assert.NoError(t, err, "BlockByNumber should not return an error")

	block, err := client.BlockByHash(context.Background(), blockTmp.Hash)

	assert.NoError(t, err, "BlockByHash should not return an error")
	assert.NotNil(t, block, "BlockByHash should return a non-nil block")
}

func TestIntegrationBaseClient_BlockByNumber(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	block, err := client.BlockByNumber(context.Background(), nil)

	assert.NoError(t, err, "BlockByNumber should not return an error")
	assert.NotNil(t, block, "BlockByNumber should return a non-nil block")
}

func TestIntegrationBaseClient_BlockNumber(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	blockNumber, err := client.BlockNumber(context.Background())

	assert.NoError(t, err, "BlockNumber should not return an error")
	assert.NotNil(t, blockNumber, "BlockNumber should return a non-nil block number")
}

func TestIntegrationBaseClient_PeerCount(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	peerCount, err := client.PeerCount(context.Background())

	assert.NoError(t, err, "PeerCount should not return an error")
	assert.NotNil(t, peerCount, "PeerCount should return a non-nil peer count")
}

func TestIntegrationBaseClient_HeaderByHash(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	block, err := client.BlockByNumber(context.Background(), nil)
	assert.NoError(t, err, "BlockByNumber should not return an error")

	header, err := client.HeaderByHash(context.Background(), block.Hash)

	assert.NoError(t, err, "HeaderByHash should not return an error")
	assert.NotNil(t, header, "HeaderByHash should return a non-nil header")
}

func TestIntegrationBaseClient_HeaderByNumber(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	header, err := client.HeaderByNumber(context.Background(), nil)

	assert.NoError(t, err, "HeaderByNumber should not return an error")
	assert.NotNil(t, header, "HeaderByNumber should return a non-nil header")
}

func TestIntegrationBaseClient_TransactionByHash(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	block, err := client.BlockByNumber(context.Background(), nil)
	assert.NoError(t, err, "BlockByNumber should not return an error")

	tx, isPending, err := client.TransactionByHash(context.Background(), block.Transactions[0].Hash)

	assert.NoError(t, err, "TransactionByHash should not return an error")
	assert.NotNil(t, tx, "TransactionByHash should return a non-nil transaction")
	assert.NotNil(t, isPending, "TransactionByHash should return a non-nil pending status")
}

func TestIntegrationBaseClient_TransactionSender(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	block, err := client.BlockByNumber(context.Background(), nil)
	assert.NoError(t, err, "BlockByNumber should not return an error")

	tx, _, err := client.TransactionByHash(context.Background(), block.Transactions[0].Hash)
	assert.NoError(t, err, "TransactionByHash should not return an error")

	sender, err := client.TransactionSender(context.Background(), tx, block.Hash, 0)

	assert.NoError(t, err, "TransactionSender should not return an error")
	assert.NotNil(t, sender, "TransactionSender should return a non-nil sender")
}

func TestIntegrationBaseClient_TransactionCount(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	block, err := client.BlockByNumber(context.Background(), nil)
	assert.NoError(t, err, "BlockByNumber should not return an error")

	txCount, err := client.TransactionCount(context.Background(), block.Hash)

	assert.NoError(t, err, "HeaderByNumber should not return an error")
	assert.NotNil(t, txCount, "HeaderByNumber should return a non-nil transaction count")
}

func TestIntegrationBaseClient_TransactionInBlock(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	block, err := client.BlockByNumber(context.Background(), nil)
	assert.NoError(t, err, "BlockByNumber should not return an error")

	tx, err := client.TransactionInBlock(context.Background(), block.Hash, 0)

	assert.NoError(t, err, "TransactionInBlock should not return an error")
	assert.NotNil(t, tx, "TransactionInBlock should return a non-nil transaction")
}

func TestIntegrationBaseClient_TransactionReceipt(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	block, err := client.BlockByNumber(context.Background(), nil)
	assert.NoError(t, err, "BlockByNumber should not return an error")

	txReceipt, err := client.TransactionReceipt(context.Background(), block.Transactions[0].Hash)

	assert.NoError(t, err, "TransactionReceipt should not return an error")
	assert.NotNil(t, txReceipt, "TransactionReceipt should return a non-nil transaction receipt")
}

func TestIntegrationBaseClient_SyncProgress(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	sync, err := client.SyncProgress(context.Background())

	assert.NoError(t, err, "SyncProgress should not return an error")
	assert.Nil(t, sync, "SyncProgress should return a nil sync progress - nodes are always synced")
}

//func TestIntegrationBaseClient_SubscribeNewHead(t *testing.T) {
//	client, err := clients.Dial(ZkSyncEraProvider)
//	defer client.Close()
//	assert.NoError(t, err, "clients.Dial should not return an error")
//
//	headers := make(chan *types.Header)
//	sub, err := client.SubscribeNewHead(context.Background(), headers)
//	if err != nil {
//		log.Panic(err)
//	}
//	defer sub.Unsubscribe()
//
//	for {
//		select {
//		case errSub := <-sub.Err():
//			log.Println(errSub)
//			break
//		case header := <-headers:
//			fmt.Println("New block header:", header.Number.String())
//		}
//	}
//}

func TestIntegrationBaseClient_NetworkID(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	networkID, err := client.NetworkID(context.Background())

	assert.NoError(t, err, "NetworkID should not return an error")
	assert.NotNil(t, networkID, "NetworkID should return a non-nil network ID")
}

func TestIntegrationBaseClient_BalanceAt(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	balance, err := client.BalanceAt(context.Background(), Address, nil)

	assert.NoError(t, err, "BalanceAt should not return an error")
	assert.NotNil(t, balance, "BalanceAt should return a non-nil balance")
}

func TestIntegrationBaseClient_StorageAt(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	storage, err := client.StorageAt(context.Background(), L2Dai, common.HexToHash("0"), nil)

	assert.NoError(t, err, "StorageAt should not return an error")
	assert.NotNil(t, storage, "StorageAt should return a non-nil storage slot")
}

func TestIntegrationBaseClient_CodeAt(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	code, err := client.CodeAt(context.Background(), L2Dai, nil)

	assert.NoError(t, err, "CodeAt should not return an error")
	assert.NotNil(t, code, "CodeAt should return a non-nil bytecode")
}

func TestIntegrationBaseClient_NonceAt(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	nonce, err := client.NonceAt(context.Background(), Address, nil)

	assert.NoError(t, err, "NonceAt should not return an error")
	assert.NotNil(t, nonce, "NonceAt should return a non-nil nonce")
}

func TestIntegrationBaseClient_FilterLogs(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   big.NewInt(1000),
		Addresses: []common.Address{L2Dai},
	})

	assert.NoError(t, err, "FilterLogs should not return an error")
	assert.GreaterOrEqual(t, len(logs), 0, "FilterLogs should return some logs")
}

func TestIntegrationBaseClient_FilterLogsL2(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	logs, err := client.FilterLogsL2(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   big.NewInt(1000),
		Addresses: []common.Address{L2Dai},
	})

	assert.NoError(t, err, "FilterLogsL2 should not return an error")
	assert.GreaterOrEqual(t, len(logs), 0, "FilterLogsL2 should return some logs")

}

//func TestIntegrationBaseClient_SubscribeFilterLogs(t *testing.T) {
//	token := readToken()
//
//	client, err := clients.Dial(ZkSyncEraProvider)
//	defer client.Close()
//	assert.NoError(t, err, "clients.Dial should not return an error")
//
//	filterLogs := make(chan zkTypes.Log)
//	sub, err := client.SubscribeFilterLogs(context.Background(), ethereum.FilterQuery{
//		FromBlock: big.NewInt(0),
//		ToBlock:   big.NewInt(1000),
//		Addresses: []common.Address{common.HexToAddress(token.L2Address)},
//	}, filterLogs)
//	if err != nil {
//		log.Panic(err)
//	}
//	defer sub.Unsubscribe()
//
//	for {
//		select {
//		case err := <-sub.Err():
//			log.Fatal(err)
//		case l := <-filterLogs:
//			fmt.Println("Address: ", l.Address)
//			fmt.Println("Data", l.Data)
//		}
//	}
//}

//func TestIntegrationBaseClient_SubscribeFilterLogsL2(t *testing.T) {
//	token := readToken()
//
//	client, err := clients.Dial(ZkSyncEraProvider)
//	defer client.Close()
//	assert.NoError(t, err, "clients.Dial should not return an error")
//
//	filterLogs := make(chan zkTypes.Log)
//	sub, err := client.SubscribeFilterLogsL2(context.Background(), ethereum.FilterQuery{
//		FromBlock: big.NewInt(0),
//		ToBlock:   big.NewInt(1000),
//		Addresses: []common.Address{common.HexToAddress(token.L2Address)},
//	}, filterLogs)
//	if err != nil {
//		log.Panic(err)
//	}
//	defer sub.Unsubscribe()
//
//	for {
//		select {
//		case err := <-sub.Err():
//			log.Fatal(err)
//		case l := <-filterLogs:
//			fmt.Println("Address: ", l.Address)
//			fmt.Println("Data", l.Data)
//		}
//	}
//}

func TestIntegrationBaseClient_PendingBalanceAt(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	balance, err := client.PendingBalanceAt(context.Background(), Address)

	assert.NoError(t, err, "PendingBalanceAt should not return an error")
	assert.NotNil(t, balance, "PendingBalanceAt should return a non-nil balance")
}

func TestIntegrationBaseClient_PendingStorageAt(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	storage, err := client.PendingStorageAt(context.Background(), L2Dai, common.HexToHash("0"))

	assert.NoError(t, err, "PendingStorageAt should not return an error")
	assert.NotNil(t, storage, "PendingStorageAt should return a non-nil storage slot")
}

func TestIntegrationBaseClient_PendingCodeAt(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	code, err := client.PendingCodeAt(context.Background(), L2Dai)

	assert.NoError(t, err, "PendingCodeAt should not return an error")
	assert.NotNil(t, code, "PendingCodeAt should return a non-nil bytecode")
}

func TestIntegrationBaseClient_PendingNonceAt(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	nonce, err := client.PendingNonceAt(context.Background(), Address)

	assert.NoError(t, err, "PendingNonceAt should not return an error")
	assert.NotNil(t, nonce, "PendingNonceAt should return a non-nil nonce")
}

// BUG in zkSync node. When there are no pending transaction,
// the return value should be 0x0 instead of a null. The response
// cannot be parsed in hex number.

//func TestIntegrationBaseClient_PendingTransactionCount(t *testing.T) {
//	client, err := clients.Dial(ZkSyncEraProvider)
//	defer client.Close()
//	assert.NoError(t, err, "clients.Dial should not return an error")
//
//	transactionCount, err := client.PendingTransactionCount(context.Background())
//
//	assert.NoError(t, err, "PendingTransactionCount should not return an error")
//	assert.GreaterOrEqual(t, transactionCount, uint64(0), "PendingTransactionCount should return a non-negative number")
//}

func TestIntegrationBaseClient_CallContract(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	tokenAbi, err := erc20.IERC20MetaData.GetAbi()
	assert.NoError(t, err, "bind.GetAbi should not return an error")

	symbolCalldata, err := tokenAbi.Pack("symbol")
	assert.NoError(t, err, "abi.Pack should not return an error")

	result, err := client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &L2Dai,
		Data: symbolCalldata,
	}, nil)
	assert.NoError(t, err, "CallContract should not return an error")

	unpack, err := tokenAbi.Unpack("symbol", result)
	assert.NoError(t, err, "abi.Unpack should not return an error")

	symbol := *abi.ConvertType(unpack[0], new(string)).(*string)
	assert.Equal(t, "DAI", symbol, "Symbols should be the same")
}

func TestIntegrationBaseClient_CallContractL2(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	tokenAbi, err := erc20.IERC20MetaData.GetAbi()
	assert.NoError(t, err, "bind.GetAbi should not return an error")

	symbolCalldata, err := tokenAbi.Pack("symbol")
	assert.NoError(t, err, "abi.Pack should not return an error")

	result, err := client.CallContractL2(context.Background(), zkTypes.CallMsg{
		CallMsg: ethereum.CallMsg{
			To:   &L2Dai,
			Data: symbolCalldata,
		},
	}, nil)
	assert.NoError(t, err, "CallContractL2 should not return an error")

	unpack, err := tokenAbi.Unpack("symbol", result)
	assert.NoError(t, err, "abi.Unpack should not return an error")

	symbol := *abi.ConvertType(unpack[0], new(string)).(*string)
	assert.Equal(t, "DAI", symbol, "Symbols should be the same")
}

func TestIntegrationBaseClient_CallContractAtHash(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	tokenAbi, err := erc20.IERC20MetaData.GetAbi()
	assert.NoError(t, err, "bind.GetAbi should not return an error")

	symbolCalldata, err := tokenAbi.Pack("symbol")
	assert.NoError(t, err, "abi.Pack should not return an error")

	block, err := client.BlockByNumber(context.Background(), nil)
	assert.NoError(t, err, "BlockByNumber should not return an error")

	result, err := client.CallContractAtHash(context.Background(), ethereum.CallMsg{
		To:   &L2Dai,
		Data: symbolCalldata,
	}, block.Hash)
	assert.NoError(t, err, "CallContractAtHash should not return an error")

	unpack, err := tokenAbi.Unpack("symbol", result)
	assert.NoError(t, err, "abi.Unpack should not return an error")

	symbol := *abi.ConvertType(unpack[0], new(string)).(*string)
	assert.Equal(t, "DAI", symbol, "Symbols should be the same")
}

func TestIntegrationBaseClient_CallContractAtHashL2(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	tokenAbi, err := erc20.IERC20MetaData.GetAbi()
	assert.NoError(t, err, "bind.GetAbi should not return an error")

	symbolCalldata, err := tokenAbi.Pack("symbol")
	assert.NoError(t, err, "abi.Pack should not return an error")

	block, err := client.BlockByNumber(context.Background(), nil)
	assert.NoError(t, err, "BlockByNumber should not return an error")

	result, err := client.CallContractAtHashL2(context.Background(), zkTypes.CallMsg{
		CallMsg: ethereum.CallMsg{
			To:   &L2Dai,
			Data: symbolCalldata,
		},
	}, block.Hash)
	assert.NoError(t, err, "CallContractAtHashL2 should not return an error")

	unpack, err := tokenAbi.Unpack("symbol", result)
	assert.NoError(t, err, "abi.Unpack should not return an error")

	symbol := *abi.ConvertType(unpack[0], new(string)).(*string)
	assert.Equal(t, "DAI", symbol, "Symbols should be the same")
}

func TestIntegrationBaseClient_PendingCallContract(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	tokenAbi, err := erc20.IERC20MetaData.GetAbi()
	assert.NoError(t, err, "bind.GetAbi should not return an error")

	symbolCalldata, err := tokenAbi.Pack("symbol")
	assert.NoError(t, err, "abi.Pack should not return an error")

	result, err := client.PendingCallContract(context.Background(), ethereum.CallMsg{
		To:   &L2Dai,
		Data: symbolCalldata,
	})
	assert.NoError(t, err, "PendingCallContract should not return an error")

	unpack, err := tokenAbi.Unpack("symbol", result)
	assert.NoError(t, err, "abi.Unpack should not return an error")

	symbol := *abi.ConvertType(unpack[0], new(string)).(*string)
	assert.Equal(t, "DAI", symbol, "Symbols should be the same")
}

func TestIntegrationBaseClient_PendingCallContractL2(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	tokenAbi, err := erc20.IERC20MetaData.GetAbi()
	assert.NoError(t, err, "bind.GetAbi should not return an error")

	symbolCalldata, err := tokenAbi.Pack("symbol")
	assert.NoError(t, err, "abi.Pack should not return an error")

	result, err := client.PendingCallContractL2(context.Background(), zkTypes.CallMsg{
		CallMsg: ethereum.CallMsg{
			To:   &L2Dai,
			Data: symbolCalldata,
		},
	})
	assert.NoError(t, err, "PendingCallContractL2 should not return an error")

	unpack, err := tokenAbi.Unpack("symbol", result)
	assert.NoError(t, err, "abi.Unpack should not return an error")

	symbol := *abi.ConvertType(unpack[0], new(string)).(*string)
	assert.Equal(t, "DAI", symbol, "Symbols should be the same")
}

func TestIntegrationBaseClient_SuggestGasPrice(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	gasPrice, err := client.SuggestGasPrice(context.Background())

	assert.NoError(t, err, "SuggestGasPrice should not return an error")
	assert.True(t, gasPrice.Cmp(big.NewInt(0)) > 0, "SuggestGasPrice should return a positive number")
}

func TestIntegrationBaseClient_SuggestGasTipCap(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	tip, err := client.SuggestGasTipCap(context.Background())

	assert.NoError(t, err, "SuggestGasTipCap should not return an error")
	assert.True(t, tip.Cmp(big.NewInt(0)) > 0, "SuggestGasTipCap should return a positive number")
}

func TestIntegrationBaseClient_EstimateGas(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	tokenAbi, err := erc20.IERC20MetaData.GetAbi()
	assert.NoError(t, err, "bind.GetAbi should not return an error")

	approveTokenCalldata, err := tokenAbi.Pack("approve", Receiver, big.NewInt(1))
	assert.NoError(t, err, "abi.Pack should not return an error")

	gas, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: Address,
		To:   &L2Dai,
		Data: approveTokenCalldata,
	})

	assert.NoError(t, err, "EstimateGas should not return an error")
	assert.Greater(t, gas, uint64(0), "EstimateGas should return a positive number")
}

func TestIntegrationBaseClient_EstimateGasL2(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	tokenAbi, err := erc20.IERC20MetaData.GetAbi()
	assert.NoError(t, err, "bind.GetAbi should not return an error")

	approveTokenCalldata, err := tokenAbi.Pack("approve", Receiver, big.NewInt(1))
	assert.NoError(t, err, "abi.Pack should not return an error")

	gas, err := client.EstimateGasL2(context.Background(), zkTypes.CallMsg{
		CallMsg: ethereum.CallMsg{
			From: Address,
			To:   &L2Dai,
			Data: approveTokenCalldata,
		},
	})

	assert.NoError(t, err, "EstimateGas should not return an error")
	assert.Greater(t, gas, uint64(0), "EstimateGas should return a positive number")
}

func TestIntegrationBaseClient_SendTransaction(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	privateKey, err := crypto.HexToECDSA(PrivateKey)
	assert.NoError(t, err, "crypto.HexToECDSA should not return an error")

	tokenAbi, err := erc20.IERC20MetaData.GetAbi()
	assert.NoError(t, err, "bind.GetAbi should not return an error")

	approveTokenCalldata, err := tokenAbi.Pack("approve", Receiver, big.NewInt(1))
	assert.NoError(t, err, "abi.Pack should not return an error")

	chainID, err := client.ChainID(context.Background())
	assert.NoError(t, err, "ChainID should not return an error")

	nonce, err := client.NonceAt(context.Background(), Address, nil)
	assert.NoError(t, err, "NonceAt should not return an error")

	gas, err := client.EstimateGasL2(context.Background(), zkTypes.CallMsg{
		CallMsg: ethereum.CallMsg{
			From: Address,
			To:   &L2Dai,
			Data: approveTokenCalldata,
		},
	})
	assert.NoError(t, err, "EstimateGas should not return an error")

	gasPrice, err := client.SuggestGasPrice(context.Background())
	assert.NoError(t, err, "SuggestGasPrice should not return an error")

	transaction := types.NewTx(
		&types.DynamicFeeTx{
			To:        &L2Dai,
			Nonce:     nonce,
			Gas:       gas,
			GasFeeCap: gasPrice,
			Data:      approveTokenCalldata,
		})

	signedTx, err := types.SignTx(transaction, types.NewLondonSigner(chainID), privateKey)
	assert.NoError(t, err, "types.SignTx should not return an error")

	err = client.SendTransaction(context.Background(), signedTx)
	assert.NoError(t, err, "SendTransaction should not return an error")

	txReceipt, err := client.WaitMined(context.Background(), signedTx.Hash())
	assert.NoError(t, err, "client.WaitMined should not return an error")

	assert.NotNil(t, txReceipt.BlockHash, "Transaction should be mined")
}

func TestIntegrationBaseClient_SendRawTransaction(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	w, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, nil)
	assert.NoError(t, err, "NewWallet should not return an error")

	tokenAbi, err := erc20.IERC20MetaData.GetAbi()
	assert.NoError(t, err, "bind.GetAbi should not return an error")

	approveTokenCalldata, err := tokenAbi.Pack("approve", Receiver, big.NewInt(1))
	assert.NoError(t, err, "abi.Pack should not return an error")

	tx, err := w.PopulateTransaction(context.Background(), accounts.Transaction{
		To:   &L2Dai,
		Data: approveTokenCalldata,
	})
	assert.NoError(t, err, "PopulateTransaction should not return an error")

	signedTx, err := w.SignTransaction(tx)
	assert.NoError(t, err, "SignTransaction should not return an error")

	txHash, err := client.SendRawTransaction(context.Background(), signedTx)
	assert.NoError(t, err, "SendRawTransaction should not return an error")

	txReceipt, err := client.WaitMined(context.Background(), txHash)
	assert.NoError(t, err, "client.WaitMined should not return an error")

	assert.NotNil(t, txReceipt.BlockHash, "Transaction should be mined")
}

func TestIntegrationBaseClient_WaitMined(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	w, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, nil)
	assert.NoError(t, err, "NewWallet should not return an error")

	tx, err := w.Transfer(nil, accounts.TransferTransaction{
		To:     Receiver,
		Amount: big.NewInt(7_000_000_000),
		Token:  utils.EthAddress,
	})
	assert.NoError(t, err, "Transfer should not return an error")

	txReceipt, err := client.WaitMined(context.Background(), tx.Hash())
	assert.NoError(t, err, "client.WaitMined should not return an error")

	assert.NotNil(t, txReceipt.BlockHash, "Transaction should be mined")
}

func TestIntegrationBaseClient_WaitFinalized(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	w, err := accounts.NewWallet(common.Hex2Bytes(PrivateKey), &client, nil)
	assert.NoError(t, err, "NewWallet should not return an error")

	tx, err := w.Transfer(nil, accounts.TransferTransaction{
		To:     Receiver,
		Amount: big.NewInt(7_000_000_000),
		Token:  utils.EthAddress,
	})
	assert.NoError(t, err, "Transfer should not return an error")

	txReceipt, err := client.WaitFinalized(context.Background(), tx.Hash())
	assert.NoError(t, err, "client.WaitMined should not return an error")

	assert.NotNil(t, txReceipt.BlockHash, "Transaction should be mined")
}

func TestIntegrationBaseClient_MainContractAddress(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	mainContract, err := client.MainContractAddress(context.Background())

	assert.NoError(t, err, "MainContractAddress should not return an error")
	assert.NotNil(t, mainContract, "MainContractAddress should return a non-nil address")
}

func TestIntegrationBaseClient_TestnetPaymaster(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	_, err = client.TestnetPaymaster(context.Background())

	assert.NoError(t, err, "TestnetPaymaster should not return an error")
}

func TestIntegrationBaseClient_BridgeContracts(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	bridgeContracts, err := client.BridgeContracts(context.Background())

	assert.NoError(t, err, "BridgeContracts should not return an error")
	assert.NotNil(t, bridgeContracts, "BridgeContracts should return a non-nil bridge contracts")
}

func TestIntegrationBaseClient_ContractAccountInfo(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	accountInfo, err := client.ContractAccountInfo(context.Background(), L2Dai)

	assert.NoError(t, err, "ContractAccountInfo should not return an error")
	assert.NotNil(t, accountInfo, "ContractAccountInfo should return a non-nil account information")
}

func TestIntegrationBaseClient_L1ChainID(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	l1ChainID, err := client.L1ChainID(context.Background())

	assert.NoError(t, err, "L1ChainID should not return an error")
	assert.NotNil(t, l1ChainID, "L1ChainID should return a non-nil chain ID")
}

func TestIntegrationBaseClient_L1BatchNumber(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	l1BatchNumber, err := client.L1BatchNumber(context.Background())

	assert.NoError(t, err, "L1BatchNumber should not return an error")
	assert.True(t, l1BatchNumber.Cmp(big.NewInt(0)) > 0, "L1BatchNumber should return a positive number")
}

func TestIntegrationBaseClient_L1BatchBlockRange(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	l1BatchNumber, err := client.L1BatchNumber(context.Background())
	assert.NoError(t, err, "L1BatchNumber should not return an error")

	l1ChainID, err := client.L1BatchBlockRange(context.Background(), l1BatchNumber)

	assert.NoError(t, err, "L1BatchBlockRange should not return an error")
	assert.NotNil(t, l1ChainID, "L1BatchBlockRange should return a non-nil block range")
}

func TestIntegrationBaseClient_L1BatchDetails(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	l1BatchNumber, err := client.L1BatchNumber(context.Background())
	assert.NoError(t, err, "L1BatchNumber should not return an error")

	batchDetails, err := client.L1BatchDetails(context.Background(), l1BatchNumber)

	assert.NoError(t, err, "L1BatchDetails should not return an error")
	assert.NotNil(t, batchDetails, "L1BatchDetails should return a non-nil batch details")
}

func TestIntegrationBaseClient_BlockDetails(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	blockNumber, err := client.BlockNumber(context.Background())
	assert.NoError(t, err, "BlockNumber should not return an error")

	blockDetails, err := client.BlockDetails(context.Background(), uint32(blockNumber))

	assert.NoError(t, err, "BlockDetails should not return an error")
	assert.NotNil(t, blockDetails, "BlockDetails should return a non-nil block details")
}

func TestIntegrationBaseClient_TransactionDetails(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	txTmp, _, err := client.TransactionByHash(context.Background(), L2DepositTx)
	assert.NoError(t, err, "TransactionByHash should not return an error")

	tx, err := client.TransactionInBlock(context.Background(), *txTmp.BlockHash, 0)
	assert.NoError(t, err, "TransactionInBlock should not return an error")

	txDetails, err := client.TransactionDetails(context.Background(), tx.Hash)

	assert.NoError(t, err, "TransactionDetails should not return an error")
	assert.NotNil(t, txDetails, "TransactionDetails should return a non-nil block details")
}

func TestIntegrationBaseClient_L2TransactionFromPriorityOp(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	ethClient, err := ethclient.Dial(EthereumProvider)
	assert.NoError(t, err, "ethclient.Dial should not return an error")
	defer ethClient.Close()

	l1Receipt, err := ethClient.TransactionReceipt(context.Background(), L1DepositTx)
	assert.NoError(t, err, "ethclient.TransactionReceipt should not return an error")

	l2Tx, err := client.L2TransactionFromPriorityOp(context.Background(), l1Receipt)

	assert.NoError(t, err, "ethclient.TransactionReceipt should not return an error")
	assert.NotNil(t, l2Tx, "EstimateFee should return a non-nil transaction")
}

func TestIntegrationBaseClient_L2TokenAddress(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	l2Address, err := client.L2TokenAddress(context.Background(), L1Dai)

	assert.NoError(t, err, "L2TokenAddress should not return an error")
	assert.Equal(t, L2Dai, l2Address, "L2 token addresses should be the same")
}

func TestIntegrationBaseClient_L1TokenAddress(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	l1Address, err := client.L1TokenAddress(context.Background(), L2Dai)

	assert.NoError(t, err, "L1TokenAddress should not return an error")
	assert.Equal(t, L1Dai, l1Address, "L1 token addresses should be the same")
}

func TestIntegrationBaseClient_AllAccountBalances(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	balances, err := client.AllAccountBalances(context.Background(), Address)

	assert.NoError(t, err, "AllAccountBalances should not return an error")
	assert.Len(t, balances, 2, "Should have ETH and DAI balance")
}

func TestIntegrationBaseClient_EstimateFee(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	fee, err := client.EstimateFee(context.Background(), zkTypes.CallMsg{
		CallMsg: ethereum.CallMsg{
			From:  Address,
			To:    &Receiver,
			Value: big.NewInt(7_000_000_000),
		},
	})

	assert.NoError(t, err, "EstimateFee should not return an error")
	assert.NotNil(t, fee, "EstimateFee should return a non-nil fee")
}

func TestIntegrationBaseClient_EstimateGasL1(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	gas, err := client.EstimateGasL1(context.Background(), zkTypes.CallMsg{
		CallMsg: ethereum.CallMsg{
			From:  Address,
			To:    &Receiver,
			Value: big.NewInt(7_000_000_000),
		},
		Meta: &zkTypes.Eip712Meta{
			GasPerPubdata: utils.NewBig(utils.RequiredL1ToL2GasPerPubdataLimit.Int64()),
		},
	})

	assert.NoError(t, err, "EstimateGasL1 should not return an error")
	assert.NotNil(t, gas, "EstimateGasL1 should return a non-nil gas")
}

func TestIntegrationBaseClient_EstimateGasTransfer(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	gas, err := client.EstimateGasTransfer(context.Background(), clients.TransferCallMsg{
		From:   Address,
		To:     Receiver,
		Amount: big.NewInt(7_000_000_000),
	})

	assert.NoError(t, err, "EstimateGasTransfer should not return an error")
	assert.Greater(t, gas, uint64(0), "EstimateGasTransfer should return a positive number")
}

func TestIntegrationBaseClient_EstimateGasWithdraw(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	gas, err := client.EstimateGasWithdraw(context.Background(), clients.WithdrawalCallMsg{
		From:   Address,
		To:     Receiver,
		Amount: big.NewInt(7_000_000_000),
	})

	assert.NoError(t, err, "EstimateGasWithdraw should not return an error")
	assert.Greater(t, gas, uint64(0), "EstimateGasWithdraw should return a positive number")
}

func TestIntegrationBaseClient_EstimateL1ToL2Execute(t *testing.T) {
	client, err := clients.Dial(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.Dial should not return an error")

	mainContractAddress, err := client.MainContractAddress(context.Background())
	assert.NoError(t, err, "MainContractAddress should not return an error")

	gas, err := client.EstimateL1ToL2Execute(context.Background(), zkTypes.CallMsg{
		CallMsg: ethereum.CallMsg{
			From:  Address,
			To:    &mainContractAddress,
			Value: big.NewInt(7_000_000_000),
		},
	})

	assert.NoError(t, err, "EstimateL1ToL2Execute should not return an error")
	assert.Greater(t, gas, uint64(0), "EstimateL1ToL2Execute should return a positive number")
}

//func TestIntegrationBaseClient_Proof(t *testing.T) {
//	client, err := clients.Dial(ZkSyncEraProvider)
//	defer client.Close()
//	assert.NoError(t, err, "clients.Dial should not return an error")
//
//	baseClient, ok := client.(*clients.BaseClient)
//	assert.True(t, ok, "Casting should not return error")
//
//	addressPadded := common.LeftPadBytes(Address.Bytes(), 32)
//
//	// Convert the slot number to a hex string and pad it to 32 bytes
//	slotBytes := common.Hex2Bytes("0x00") // slot with index 0
//	slotPadded := common.LeftPadBytes(slotBytes, 32)
//
//	// Concatenate the padded address and slot number
//	concatenated := append(addressPadded, slotPadded...)
//	storageKey := crypto.Keccak256(concatenated)
//
//	l1BatchNumber, err := client.L1BatchNumber(context.Background())
//	assert.NoError(t, err, "L1BatchNumber should not return an error")
//
//	storageProof, err := baseClient.Proof(
//		context.Background(),
//		utils.NonceHolderAddress,
//		[]common.Hash{common.BytesToHash(storageKey)},
//		l1BatchNumber)
//
//	assert.NoError(t, err, "Proof should not return an error")
//	assert.NotNil(t, storageProof, "Proof should return a non-nil value")
//}
