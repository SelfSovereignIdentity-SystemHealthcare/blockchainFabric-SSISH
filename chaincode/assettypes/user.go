package assettypes

import (
	"github.com/hyperledger-labs/cc-tools/assets"
)

// Definição do asset User
var User = assets.AssetType{
	Tag:         "user",
	Label:       "Usuário",
	Description: "Informações do usuário",
	Props: []assets.AssetProp{
		{
			Tag:      "email",
			IsKey:    true,
			Required: true,
			Label:    "Email",
			DataType: "string",
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
			Tag:      "role",
			Required: true,
			Label:    "Função",
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
			Tag:      "gender",
			Required: true,
			Label:    "Gênero",
			DataType: "string",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
		{
			Tag:      "birthDate",
			Required: true,
			Label:    "Data de Nascimento",
			DataType: "datetime",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
		{
			Tag:      "phone",
			Required: true,
			Label:    "Telefone",
			DataType: "string",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
		{
			Tag:      "address",
			Required: true,
			Label:    "Endereço",
			DataType: "string",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
	},
}
