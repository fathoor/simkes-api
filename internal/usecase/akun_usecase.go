package usecase

import (
	"github.com/fathoor/simkes-api/internal/entity"
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/helper"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/repository"
	"github.com/fathoor/simkes-api/internal/validation"
	"github.com/samber/do"
)

type AkunUseCase struct {
	AkunRepository *repository.AkunRepository
}

func NewAkunUseCase(i *do.Injector) (*AkunUseCase, error) {
	return &AkunUseCase{
		AkunRepository: do.MustInvoke[*repository.AkunRepository](i),
	}, nil
}

func (u *AkunUseCase) Create(request *model.AkunRequest) model.AkunResponse {
	if valid := validation.ValidateAkunRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	encrypted, err := helper.EncryptPassword(request.Password)
	exception.PanicIfError(err)

	akun := entity.Akun{
		NIP:      request.NIP,
		Email:    request.Email,
		Password: string(encrypted),
		RoleNama: request.RoleNama,
	}

	if err := u.AkunRepository.Insert(&akun); err != nil {
		exception.PanicIfError(err)
	}

	response := model.AkunResponse{
		NIP:      akun.NIP,
		Email:    akun.Email,
		RoleNama: akun.RoleNama,
	}

	return response
}

func (u *AkunUseCase) GetAll() []model.AkunResponse {
	akun, err := u.AkunRepository.FindAll()
	exception.PanicIfError(err)

	response := make([]model.AkunResponse, len(akun))
	for i, akun := range akun {
		response[i] = model.AkunResponse{
			NIP:      akun.NIP,
			Email:    akun.Email,
			RoleNama: akun.RoleNama,
		}
	}

	return response
}

func (u *AkunUseCase) GetPage(page, size int) model.AkunPageResponse {
	akun, total, err := u.AkunRepository.FindPage(page, size)
	exception.PanicIfError(err)

	response := make([]model.AkunResponse, len(akun))
	for i, akun := range akun {
		response[i] = model.AkunResponse{
			NIP:      akun.NIP,
			Email:    akun.Email,
			RoleNama: akun.RoleNama,
		}
	}

	pagedResponse := model.AkunPageResponse{
		Akun:  response,
		Page:  page,
		Size:  size,
		Total: total,
	}

	return pagedResponse
}

func (u *AkunUseCase) GetByNIP(nip string) model.AkunResponse {
	akun, err := u.AkunRepository.FindByNIP(nip)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Akun not found",
		})
	}

	response := model.AkunResponse{
		NIP:      akun.NIP,
		Email:    akun.Email,
		RoleNama: akun.RoleNama,
	}

	return response
}

func (u *AkunUseCase) Update(nip string, request *model.AkunRequest) model.AkunResponse {
	if valid := validation.ValidateAkunRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	akun, err := u.AkunRepository.FindByNIP(nip)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Akun not found",
		})
	}

	if request.Password != "" {
		encrypted, err := helper.EncryptPassword(request.Password)
		exception.PanicIfError(err)

		akun.Password = string(encrypted)
	}

	akun.Email = request.Email

	response := model.AkunResponse{
		NIP:      akun.NIP,
		Email:    akun.Email,
		RoleNama: akun.RoleNama,
	}

	if err := u.AkunRepository.Update(&akun); err != nil {
		exception.PanicIfError(err)
	}

	return response
}

func (u *AkunUseCase) UpdateAdmin(nip string, request *model.AkunRequest) model.AkunResponse {
	if valid := validation.ValidateAkunRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	akun, err := u.AkunRepository.FindByNIP(nip)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Akun not found",
		})
	}

	if request.Password != "" {
		encrypted, err := helper.EncryptPassword(request.Password)
		exception.PanicIfError(err)

		akun.Password = string(encrypted)
	}

	akun.NIP = request.NIP
	akun.Email = request.Email
	akun.RoleNama = request.RoleNama

	response := model.AkunResponse{
		NIP:      akun.NIP,
		Email:    akun.Email,
		RoleNama: akun.RoleNama,
	}

	if err := u.AkunRepository.Update(&akun); err != nil {
		exception.PanicIfError(err)
	}

	return response
}

func (u *AkunUseCase) Delete(nip string) {
	akun, err := u.AkunRepository.FindByNIP(nip)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Akun not found",
		})
	}

	if err := u.AkunRepository.Delete(&akun); err != nil {
		exception.PanicIfError(err)
	}
}
