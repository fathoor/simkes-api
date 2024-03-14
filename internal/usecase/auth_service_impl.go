package usecase

import (
	"github.com/fathoor/simkes-api/internal/exception"
	helper2 "github.com/fathoor/simkes-api/internal/helper"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/repository"
	"github.com/fathoor/simkes-api/internal/validation"
	"time"
)

type authServiceImpl struct {
	repository.AkunRepository
}

func (service *authServiceImpl) Login(request *model.AuthRequest) model.AuthResponse {
	if valid := validation.ValidateAuthRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	akun, err := service.AkunRepository.FindByNIP(request.NIP)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Akun not found",
		})
	}

	if !helper2.DecryptPassword(akun.Password, request.Password) {
		panic(exception.UnauthorizedError{
			Message: "Invalid password",
		})
	}

	token, err := helper2.GenerateJWT(akun.NIP, akun.RoleNama)
	exception.PanicIfError(err)

	response := model.AuthResponse{
		Token:   token,
		Expired: time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05"),
	}

	return response
}

func NewAuthServiceProvider(repository *repository.AkunRepository) AuthService {
	return &authServiceImpl{*repository}
}
