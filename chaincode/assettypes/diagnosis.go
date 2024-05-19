package assettypes

import (
	"github.com/hyperledger-labs/cc-tools/assets"
)

// Definição do asset Diagnosis
var Diagnosis = assets.AssetType{
	Tag:         "diagnosis",
	Label:       "Diagnóstico",
	Description: "Informações do diagnóstico",
	Props: []assets.AssetProp{
		{
			Tag:      "exam",
			IsKey:    true,
			Required: true,
			Label:    "Exame",
			DataType: "->exam",
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
