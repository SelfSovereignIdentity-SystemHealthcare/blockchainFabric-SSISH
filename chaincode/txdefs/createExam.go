package txdefs

import (
	"encoding/json"
	"time"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
	sw "github.com/hyperledger-labs/cc-tools/stubwrapper"
	tx "github.com/hyperledger-labs/cc-tools/transactions"
)

var CreateExam = tx.Transaction{
	Tag:         "createExam",
	Label:       "Create Exam",
	Description: "Create a clinical examination record",
	Method:      "POST",

	Args: []tx.Argument{
		{
			Required:    true,
			Tag:         "patientWalletHolderHash",
			Label:       "Patient Hash",
			Description: "Reference to Patient Asset",
			DataType:    "sha256",
		},
		{
			Required:    true,
			Tag:         "doctorWalletHolderHash",
			Label:       "Doctor Hash",
			Description: "Reference to Doctor Asset",
			DataType:    "sha256",
		},
		{
			Required:    true,
			Tag:         "timestamp",
			Label:       "Timestamp",
			Description: "Timestamp when the examination was performed",
			DataType:    "datetime",
		},
		{
			Required:    true,
			Tag:         "name",
			Label:       "Exam Name",
			Description: "Name of the clinical examination performed",
			DataType:    "string",
		},
		{
			Required:    true,
			Tag:         "urlExamDocument",
			Label:       "Exam Document URL",
			Description: "URL of the exam document",
			DataType:    "string",
		},
		{
			Required:    true,
			Tag:         "diagnosisHash",
			Label:       "Diagnosis Hash",
			Description: "Hash of the diagnosis",
			DataType:    "sha256",
		},
		{
			Required:    true,
			Tag:         "treatmentHash",
			Label:       "Treatment Hash",
			Description: "Hash of the treatment",
			DataType:    "sha256",
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		patientWalletHolderHash, _ := req["patientWalletHolderHash"].(string)
		doctorWalletHolderHash, _ := req["doctorWalletHolderHash"].(string)
		timestamp, _ := req["timestamp"].(time.Time)
		name, _ := req["name"].(string)
		urlExamDocument, _ := req["urlExamDocument"].(string)
		diagnosisHash, _ := req["diagnosisHash"].(string)
		treatmentHash, _ := req["treatmentHash"].(string)

		timestampString := timestamp.Format("2006-01-02T15:04:05.000Z")

		examMap := map[string]interface{}{
			"@assetType":              "exam",
			"patientWalletHolderHash": patientWalletHolderHash,
			"doctorWalletHolderHash":  doctorWalletHolderHash,
			"timestamp":               timestampString,
			"name":                    name,
			"urlExamDocument":         urlExamDocument,
			"diagnosisHash":           diagnosisHash,
			"treatmentHash":           treatmentHash,
		}

		examAsset, err := assets.NewAsset(examMap)
		if err != nil {
			return nil, errors.WrapError(err, "failed to create exam asset")
		}

		res, err := examAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "error saving asset on blockchain")
		}

		examJSON, nerr := json.Marshal(res)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		return examJSON, nil
	},
}
