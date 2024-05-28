package assettypes

import (
	"github.com/hyperledger-labs/cc-tools/assets"
)

// Definição do asset Treatment
var Treatment = assets.AssetType{
	Tag:         "treatment",
	Label:       "Tratamento",
	Description: "Informações do tratamento",
	Props: []assets.AssetProp{
		{
			Tag:      "diagnosis",
			IsKey:    true,
			Required: true,
			Label:    "Diagnóstico",
			DataType: "->diagnosis",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
		{
			Tag:      "timestamp",
			IsKey:    true,
			Required: true,
			Label:    "Data e Hora",
			DataType: "datetime",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
		{
			Tag:      "name",
			Required: true,
			Label:    "Nome",
			DataType: "string",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
		{
			Tag:      "obs",
			Required: true,
			Label:    "Observações",
			DataType: "string",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
	},
}
