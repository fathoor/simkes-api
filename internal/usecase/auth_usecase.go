package usecase

import (
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/helper"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/repository"
	"github.com/fathoor/simkes-api/internal/validation"
	"github.com/samber/do"
	"time"
)

type AuthUseCase struct {
	AkunRepository *repository.AkunRepository
}

func NewAuthUseCase(i *do.Injector) (*AuthUseCase, error) {
	return &AuthUseCase{
		AkunRepository: do.MustInvoke[*repository.AkunRepository](i),
	}, nil
}

func (u *AuthUseCase) Login(request *model.AuthRequest) model.AuthResponse {
	if valid := validation.ValidateAuthRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	akun, err := u.AkunRepository.FindByNIP(request.NIP)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Akun not found",
		})
	}

	if !helper.DecryptPassword(akun.Password, request.Password) {
		panic(exception.UnauthorizedError{
			Message: "Invalid password",
		})
	}

	token, err := helper.GenerateJWT(akun.NIP, akun.RoleNama)
	exception.PanicIfError(err)

	response := model.AuthResponse{
		Token:   token,
		Expired: time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05"),
	}

	return response
}
