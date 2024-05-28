package assettypes

import (
	"github.com/hyperledger-labs/cc-tools/assets"
)

// Definição do asset Exam
var Exam = assets.AssetType{
	Tag:         "exam",
	Label:       "Exame",
	Description: "Informações do exame",
	Props: []assets.AssetProp{
		{
			Tag:      "patientWalletHolderHash",
			IsKey:    true,
			Required: true,
			Label:    "Hash do Paciente",
			DataType: "sha256",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
		{
			Tag:      "doctorWalletHolderHash",
			Required: true,
			IsKey:    true,
			Label:    "Hash do Médico",
			DataType: "sha256",
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
			Tag:      "urlExamDocument",
			Required: true,
			Label:    "URL do Documento do Exame",
			DataType: "string",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
		{
			Tag:      "diagnosisHash",
			Label:    "Hash do Diagnóstico",
			DataType: "sha256",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
		{
			Tag:      "treatmentHash",
			Label:    "Hash do Tratamento",
			DataType: "sha256",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
	},
}
