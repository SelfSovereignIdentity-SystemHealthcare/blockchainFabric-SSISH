package datatypes

import (
	"github.com/hyperledger-labs/cc-tools/assets"
)

// CustomDataTypes contain the user-defined primary data types
var CustomDataTypes = map[string]assets.DataType{
	"cpf":               cpf,
	"sha256":            Sha256,
	"enumTipoSensor":    enumTipoSensor,
	"enumUnidadeSensor": enumUnidadeSensor,
}
