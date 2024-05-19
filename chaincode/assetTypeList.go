package main

import (
	"github.com/hyperledger-labs/cc-tools-demo/chaincode/assettypes"
	"github.com/hyperledger-labs/cc-tools/assets"
)

var assetTypeList = []assets.AssetType{
	assettypes.Anamnesis,
	assettypes.Exam,
	assettypes.Doctor,
	assettypes.Patient,
	assettypes.Diagnosis,
	assettypes.Treatment,
	assettypes.User,
}
