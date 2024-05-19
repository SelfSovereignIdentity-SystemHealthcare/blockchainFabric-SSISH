package assettypes

import (
	"github.com/hyperledger-labs/cc-tools/errors"
	"time"
)

// Função de validação de data
func dateValidation(date interface{}) error {
	dateStr, ok := date.(string)
	if !ok {
		return errors.NewCCError("o valor deve ser uma string", 400)
	}

	// Verifica se a string de data está no formato RFC3339 válido
	_, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return errors.NewCCError("o valor não é uma data RFC3339 válida", 400)
	}

	return nil
}
