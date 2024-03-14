package validation

import (
	"github.com/fathoor/simkes-api/internal/config"
	"github.com/fathoor/simkes-api/internal/model"
)

func ValidateAuthRequest(request *model.AuthRequest) error {
	return config.Validator.Struct(request)
}
