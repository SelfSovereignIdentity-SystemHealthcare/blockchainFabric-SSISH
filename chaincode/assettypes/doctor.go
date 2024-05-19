package assettypes

import (
	"github.com/hyperledger-labs/cc-tools/assets"
)

// Definição do asset Doctor que herda de User
var Doctor = assets.AssetType{
	Tag:         "doctor",
	Label:       "Médico",
	Description: "Informações do médico",
	Props: append(User.Props,
		assets.AssetProp{
			Tag:      "crm",
			IsKey:    true,
			Required: true,
			Label:    "CRM",
			DataType: "string",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
		assets.AssetProp{
			Tag:      "speciality",
			Required: true,
			Label:    "Especialidade",
			DataType: "string",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
	),
}
