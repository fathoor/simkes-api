package usecase

import (
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/helper"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/repository"
	"github.com/fathoor/simkes-api/internal/validation"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"github.com/samber/do"
	"time"
)

type AuthUseCase struct {
	AkunRepository *repository.AkunRepository
	Log            *zerolog.Logger
	Validator      *validator.Validate
}

func NewAuthUseCase(i *do.Injector) (*AuthUseCase, error) {
	return &AuthUseCase{
		AkunRepository: do.MustInvoke[*repository.AkunRepository](i),
		Log:            do.MustInvoke[*zerolog.Logger](i),
		Validator:      do.MustInvoke[*validator.Validate](i),
	}, nil
}

func (u *AuthUseCase) Login(request *model.AuthRequest) model.AuthResponse {
	validation.ValidateAuthRequest(u.Validator, u.Log, request)

	akun, err := u.AkunRepository.FindByNIP(request.NIP)
	if err != nil {
		u.Log.Info().Str("nip", request.NIP).Msg("Akun not found")
		panic(exception.NotFoundError{
			Message: "Akun not found",
		})
	}

	if !helper.DecryptPassword(akun.Password, request.Password) {
		u.Log.Info().Str("nip", request.NIP).Msg("Invalid password")
		panic(exception.UnauthorizedError{
			Message: "Invalid password",
		})
	}

	token, err := helper.GenerateJWT(akun.NIP, akun.RoleNama)
	if err != nil {
		u.Log.Error().Err(err).Msg("Failed to generate JWT")
		panic(exception.InternalServerError{
			Message: "Failed to generate JWT",
		})
	}

	response := model.AuthResponse{
		Token:   token,
		Expired: time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05"),
	}

	return response
}
