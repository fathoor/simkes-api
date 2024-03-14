package validation

import (
	"github.com/fathoor/simkes-api/internal/config"
	"github.com/fathoor/simkes-api/internal/model"
)

func ValidatePegawaiRequest(request *model.PegawaiRequest) error {
	return config.Validator.Struct(request)
}
