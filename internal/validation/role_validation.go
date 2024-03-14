package validation

import (
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
)

func ValidateRoleRequest(validator *validator.Validate, log *zerolog.Logger, request *model.RoleRequest) {
	if request.Nama == "Admin" {
		log.Info().Str("nama", request.Nama).Msg("Forbidden role")
		panic(exception.ForbiddenError{
			Message: "You are not allowed to modify this role!",
		})
	}

	if err := validator.Struct(request); err != nil {
		log.Error().Err(err).Msg("Validation error")
		panic(&exception.BadRequestError{
			Message: "Invalid request",
		})
	}
}
