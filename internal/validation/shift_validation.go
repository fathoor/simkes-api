package validation

import (
	"github.com/fathoor/simkes-api/internal/app/validation"
	"github.com/fathoor/simkes-api/internal/model"
)

func ValidateShiftRequest(request *model.ShiftRequest) error {
	return validation.Validator.Struct(request)
}
