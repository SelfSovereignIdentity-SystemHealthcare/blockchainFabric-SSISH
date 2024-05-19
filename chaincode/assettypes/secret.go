package assettypes

import "github.com/hyperledger-labs/cc-tools/assets"

// Secret is and information available only for org1
// Collections.json configuration is necessary
var Secret = assets.AssetType{
	Tag:         "secret",
	Label:       "Secret",
	Description: "Secret between Org1",

	Readers: []string{"org1MSP", "orgMSP"},
	Props: []assets.AssetProp{
		{
			// Primary Key
			IsKey:    true,
			Tag:      "secretName",
			Label:    "Secret Name",
			DataType: "string",
			Writers:  []string{`org1MSP`, `orgMSP`}, // This means only org1 can create the asset
		},
		{
			// Mandatory Property
			Required: true,
			Tag:      "secret",
			Label:    "Secret",
			DataType: "string",
		},
	},
}
