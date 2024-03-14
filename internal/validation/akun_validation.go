package validation

import (
	"github.com/fathoor/simkes-api/internal/app/validation"
	"github.com/fathoor/simkes-api/internal/model"
)

func ValidateAkunRequest(request *model.AkunRequest) error {
	return validation.Validator.Struct(request)
}

func ValidateAkunUpdateRequest(request *model.AkunUpdateRequest) error {
	return validation.Validator.Struct(request)
}
