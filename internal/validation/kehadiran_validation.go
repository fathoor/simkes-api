package validation

import (
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"time"
)

func ValidateKehadiranRequest(validator *validator.Validate, log *zerolog.Logger, request *model.KehadiranRequest) error {
	if _, err := time.Parse("2006-01-02", request.Tanggal); err != nil {
		log.Info().Str("tanggal", request.Tanggal).Msg("Invalid date format")
		return &exception.BadRequestError{
			Message: "Invalid date format",
		}
	}

	return validator.Struct(request)
}

func ValidateKehadiranUpdateRequest(validator *validator.Validate, log *zerolog.Logger, request *model.KehadiranUpdateRequest) error {
	if _, err := time.Parse("2006-01-02", request.Tanggal); err != nil {
		log.Info().Str("tanggal", request.Tanggal).Msg("Invalid date format")
		return &exception.BadRequestError{
			Message: "Invalid date format",
		}
	}

	if request.JamMasuk != "" {
		if _, err := time.Parse("15:04:05", request.JamMasuk); err != nil {
			log.Info().Str("jam_masuk", request.JamMasuk).Msg("Invalid time format")
			return &exception.BadRequestError{
				Message: "Invalid time format",
			}
		}
	}

	if request.JamKeluar != "" {
		if _, err := time.Parse("15:04:05", request.JamKeluar); err != nil {
			log.Info().Str("jam_keluar", request.JamKeluar).Msg("Invalid time format")
			return &exception.BadRequestError{
				Message: "Invalid time format",
			}
		}
	}

	return validator.Struct(request)
}
