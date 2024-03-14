package validation

import (
	"github.com/fathoor/simkes-api/internal/config"
	"github.com/fathoor/simkes-api/internal/model"
)

func ValidateAkunRequest(request *model.AkunRequest) error {
	return config.Validator.Struct(request)
}

func ValidateAkunUpdateRequest(request *model.AkunUpdateRequest) error {
	return config.Validator.Struct(request)
}
