package validation

import (
	"github.com/fathoor/simkes-api/internal/config"
	"github.com/fathoor/simkes-api/internal/model"
)

func ValidateDepartemenRequest(request *model.DepartemenRequest) error {
	return config.Validator.Struct(request)
}
