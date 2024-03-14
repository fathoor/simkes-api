package validation

import (
	"github.com/fathoor/simkes-api/internal/config"
	"github.com/fathoor/simkes-api/internal/model"
)

func ValidateShiftRequest(request *model.ShiftRequest) error {
	return config.Validator.Struct(request)
}
