package assettypes

import (
	"github.com/hyperledger-labs/cc-tools/assets"
)

// Definição do asset SSISHEvent
var SSISHEvent = assets.AssetType{
	Tag:         "ssishEvent",
	Label:       "Evento SSISH",
	Description: "Evento relacionado ao sistema SSISH",
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
			Tag:      "eventType",
			IsKey:    true,
			Required: true,
			Label:    "Tipo de Evento",
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
			Tag:      "eventDetails",
			Required: true,
			Label:    "Detalhes do Evento",
			DataType: "string",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
	},
}
