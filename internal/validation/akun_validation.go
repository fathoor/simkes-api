package validation

import (
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
)

func ValidateAkunRequest(validator *validator.Validate, log *zerolog.Logger, request *model.AkunRequest) {
	if err := validator.Struct(request); err != nil {
		log.Info().Err(err).Msg("Validation error")
		panic(&exception.BadRequestError{
			Message: "Invalid request",
		})
	}
}

func ValidateAkunUpdateRequest(validator *validator.Validate, log *zerolog.Logger, request *model.AkunUpdateRequest) {
	if err := validator.Struct(request); err != nil {
		log.Info().Err(err).Msg("Validation error")
		panic(&exception.BadRequestError{
			Message: "Invalid request",
		})
	}
}
