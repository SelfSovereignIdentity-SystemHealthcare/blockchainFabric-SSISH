package main

import (
	"github.com/hyperledger-labs/cc-tools-demo/chaincode/txdefs"
	tx "github.com/hyperledger-labs/cc-tools/transactions"
)

var txList = []tx.Transaction{
	tx.CreateAsset,
	tx.UpdateAsset,

	txdefs.CreateExam,
}