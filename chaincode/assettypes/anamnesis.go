package assettypes

import (
	"github.com/hyperledger-labs/cc-tools/assets"
)

// Definição do asset Anamnesis
var Anamnesis = assets.AssetType{
	Tag:         "anamnesis",
	Label:       "Anamnese",
	Description: "Informações da anamnese",
	Props: []assets.AssetProp{
		{
			Tag:      "patient",
			IsKey:    true,
			Required: true,
			Label:    "Paciente",
			DataType: "->patient",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
		{
			Tag:      "timestamp",
			Required: true,
			Label:    "Data e Hora",
			DataType: "datetime",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
	},
}
