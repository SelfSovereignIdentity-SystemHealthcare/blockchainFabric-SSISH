package assettypes

import (
	"github.com/hyperledger-labs/cc-tools/assets"
)

// Definição do asset Wallet
var Wallet = assets.AssetType{
	Tag:         "wallet",
	Label:       "Carteira",
	Description: "Uma carteira digital para armazenar credenciais",
	Props: []assets.AssetProp{
		{
			Tag:      "holderHash",
			IsKey:    true,
			Required: true,
			Label:    "Hash do Portador",
			DataType: "sha256",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
		{
			Tag:      "email",
			Required: true,
			Label:    "E-mail",
			DataType: "string",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
		{
			Tag:      "credentials",
			Label:    "Credenciais",
			DataType: "[]->verifiableCredential",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
		{
			Tag:      "verifications",
			Label:    "Verificações",
			DataType: "[]->credentialVerification",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
	},
}
