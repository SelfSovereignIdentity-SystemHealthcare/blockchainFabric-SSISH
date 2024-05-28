package assettypes

import (
	"github.com/hyperledger-labs/cc-tools/assets"
)

// Definição do asset VerifiableCredential
var VerifiableCredential = assets.AssetType{
	Tag:         "verifiableCredential",
	Label:       "Credencial Verificável",
	Description: "Uma credencial que pode ser verificada",
	Props: []assets.AssetProp{
		{
			Tag:      "issuerHash",
			IsKey:    true,
			Required: true,
			Label:    "Hash do Emissor",
			DataType: "sha256",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
		{
			Tag:      "receiverHash",
			Required: true,
			Label:    "Hash de conexão com o Receptor",
			DataType: "sha256",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
		{
			Tag:      "credentialType",
			Required: true,
			Label:    "Tipo de Credencial",
			DataType: "string",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
		{
			Tag:      "credentialData",
			IsKey:    true,
			Required: true,
			Label:    "Dados da Credencial",
			DataType: "string",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
		{
			Tag:      "status",
			Required: true,
			Label:    "Status",
			DataType: "string",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
	},
}
