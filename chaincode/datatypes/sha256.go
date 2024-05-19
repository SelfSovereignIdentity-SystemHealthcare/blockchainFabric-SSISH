package datatypes

import (
	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
)

// Definição do Datatype SHA256
var Sha256 = assets.DataType{
	AcceptedFormats: []string{"string"},
	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		hash, ok := data.(string)
		if !ok {
			return "", nil, errors.NewCCError("a propriedade deve ser uma string", 400)
		}

		// ...

		return hash, hash, nil
	},
}
