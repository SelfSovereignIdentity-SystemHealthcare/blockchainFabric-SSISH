package datatypes

import (
	"strings"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
)

var enumUnidadeSensor = assets.DataType{
	AcceptedFormats: []string{"string"},
	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		enumUnidadeSensor, ok := data.(string)
		if !ok {
			return "", nil, errors.NewCCError("A propriedade deve ser uma string", 400)
		}

		// Converte o valor para minúsculas
		var tmpEnumUnidadeSensor = strings.ToLower(enumUnidadeSensor)

		// Verifica se o valor está na lista de valores permitidos
		// Você pode alterar essa lista de acordo com o seu caso de uso
		allowedValues := []string{"bpm", "mmhg", "°c", "%"}
		found := false
		for _, v := range allowedValues {
			if tmpEnumUnidadeSensor == v {
				found = true
				break
			}
		}
		if !found {
			return "", nil, errors.NewCCError("Valor enumUnidadeSensor inválido", 400)
		}

		return enumUnidadeSensor, enumUnidadeSensor, nil
	},
}
