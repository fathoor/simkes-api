package validation

import (
	"github.com/fathoor/simkes-api/internal/config"
	"github.com/fathoor/simkes-api/internal/model"
)

func ValidateJabatanRequest(request *model.JabatanRequest) error {
	return config.Validator.Struct(request)
}
