package validation

import (
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
)

func ValidateJadwalPegawaiRequest(validator *validator.Validate, log *zerolog.Logger, request *model.JadwalPegawaiRequest) {
	if err := validator.Struct(request); err != nil {
		log.Error().Err(err).Msg("Validation error")
		panic(&exception.BadRequestError{
			Message: "Invalid request",
		})
	}
}
