package validation

import (
	"github.com/fathoor/simkes-api/internal/config"
	"github.com/fathoor/simkes-api/internal/model"
)

func ValidateJadwalPegawaiRequest(request *model.JadwalPegawaiRequest) error {
	return config.Validator.Struct(request)
}
