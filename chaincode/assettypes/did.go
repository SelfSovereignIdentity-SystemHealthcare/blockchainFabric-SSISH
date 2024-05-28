package assettypes

import (
	"github.com/hyperledger-labs/cc-tools/assets"
)

// Definição do asset DID
var DID = assets.AssetType{
	Tag:         "did",
	Label:       "Identidade Digital Descentralizada",
	Description: "Identidade digital vinculada a uma carteira",
	Props: []assets.AssetProp{
		{
			Tag:      "walletHash",
			IsKey:    true,
			Required: true,
			Label:    "Hash da Carteira",
			DataType: "sha256",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
		{
			Tag:      "userController",
			Required: true,
			Label:    "Controlador do Usuário",
			DataType: "sha256",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
		{
			Tag:      "publicKey",
			Required: true,
			Label:    "Chave Pública",
			DataType: "string",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
		{
			Tag:      "authenticationMethods",
			Required: true,
			Label:    "Métodos de Autenticação",
			DataType: "string",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
		{
			Tag:      "services",
			Required: true,
			Label:    "Serviços",
			DataType: "string",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
		{
			Tag:      "timestamp",
			Required: true,
			Label:    "Data e Hora",
			DataType: "datetime",
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
