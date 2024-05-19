package assettypes

import (
	"github.com/hyperledger-labs/cc-tools/assets"
)

// Definição do asset Patient que herda de User
var Patient = assets.AssetType{
	Tag:         "patient",
	Label:       "Paciente",
	Description: "Informações do paciente",
	Props: append(User.Props,
		assets.AssetProp{
			Tag:      "cpf",
			IsKey:    true,
			Required: true,
			Label:    "CPF",
			DataType: "cpf",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
		assets.AssetProp{
			Tag:      "anamnesis",
			Label:    "Anamnese",
			DataType: "->anamnesis",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
	),
}
