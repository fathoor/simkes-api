package validation

import (
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"time"
)

func ValidateCutiCreateRequest(validator *validator.Validate, log *zerolog.Logger, request *model.CutiCreateRequest) {
	if _, err := time.Parse("2006-01-02", request.TanggalMulai); err != nil {
		log.Info().Str("tanggal_mulai", request.TanggalMulai).Msg("Invalid date format")
		panic(exception.BadRequestError{
			Message: "Invalid date format",
		})
	}

	if _, err := time.Parse("2006-01-02", request.TanggalSelesai); err != nil {
		log.Info().Str("tanggal_selesai", request.TanggalSelesai).Msg("Invalid date format")
		panic(exception.BadRequestError{
			Message: "Invalid date format",
		})
	}

	if err := validator.Struct(request); err != nil {
		log.Error().Err(err).Msg("Validation error")
		panic(exception.BadRequestError{
			Message: "Invalid request",
		})
	}
}

func ValidateCutiUpdateRequest(validator *validator.Validate, log *zerolog.Logger, request *model.CutiUpdateRequest) {
	if _, err := time.Parse("2006-01-02", request.TanggalMulai); err != nil {
		log.Info().Str("tanggal_mulai", request.TanggalMulai).Msg("Invalid date format")
		panic(exception.BadRequestError{
			Message: "Invalid date format",
		})
	}

	if _, err := time.Parse("2006-01-02", request.TanggalSelesai); err != nil {
		log.Info().Str("tanggal_selesai", request.TanggalSelesai).Msg("Invalid date format")
		panic(exception.BadRequestError{
			Message: "Invalid date format",
		})
	}

	if err := validator.Struct(request); err != nil {
		log.Error().Err(err).Msg("Validation error")
		panic(exception.BadRequestError{
			Message: "Invalid request",
		})
	}
}
