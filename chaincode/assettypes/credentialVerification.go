package assettypes

import (
	"github.com/hyperledger-labs/cc-tools/assets"
)

// Definição do asset CredentialVerification
var CredentialVerification = assets.AssetType{
	Tag:         "credentialVerification",
	Label:       "Verificação de Credencial",
	Description: "Resultado da verificação de uma credencial",
	Props: []assets.AssetProp{
		{
			Tag:      "verifierHash",
			IsKey:    true,
			Required: true,
			Label:    "Hash do Verificador",
			DataType: "sha256",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
		{
			Tag:      "verification",
			Required: true,
			Label:    "Verificação",
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
			Tag:      "result",
			Required: true,
			Label:    "Resultado",
			DataType: "string",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
	},
}
