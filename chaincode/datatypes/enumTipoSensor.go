package datatypes

import (
	"strings"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
)

var enumTipoSensor = assets.DataType{
	AcceptedFormats: []string{"string"},
	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		enumTipoSensor, ok := data.(string)
		if !ok {
			return "", nil, errors.NewCCError("propriedade deve ser uma string", 400)
		}

		// Converte o valor para minúsculas
		enumTipoSensor = strings.ToLower(enumTipoSensor)

		// Verifica se o valor está na lista de valores permitidos
		// Você pode alterar essa lista de acordo com o seu caso de uso
		allowedValues := []string{"frequencia_cardiaca", "pressao_arterial", "temperatura", "oxigenacao_do_sangue"}
		found := false
		for _, v := range allowedValues {
			if enumTipoSensor == v {
				found = true
				break
			}
		}
		if !found {
			return "", nil, errors.NewCCError("valor enumTipoSensor inválido", 400)
		}

		return enumTipoSensor, enumTipoSensor, nil
	},
}
