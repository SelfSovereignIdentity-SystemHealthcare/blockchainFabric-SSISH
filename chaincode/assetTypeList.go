package main

import (
	"github.com/hyperledger-labs/cc-tools-demo/chaincode/assettypes"
	"github.com/hyperledger-labs/cc-tools/assets"
)

var assetTypeList = []assets.AssetType{
	assettypes.Exam,
	assettypes.Diagnosis,
	assettypes.Treatment,
	assettypes.DID,
	assettypes.SSISHEvent,
	assettypes.CredentialVerification,
	assettypes.VerifiableCredential,
	assettypes.Wallet,
}
