package usecase

import (
	"github.com/fathoor/simkes-api/internal/entity"
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/helper"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/repository"
	"github.com/fathoor/simkes-api/internal/validation"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"github.com/samber/do"
)

type AkunUseCase struct {
	AkunRepository *repository.AkunRepository
	Log            *zerolog.Logger
	Validator      *validator.Validate
}

func NewAkunUseCase(i *do.Injector) (*AkunUseCase, error) {
	return &AkunUseCase{
		AkunRepository: do.MustInvoke[*repository.AkunRepository](i),
		Log:            do.MustInvoke[*zerolog.Logger](i),
		Validator:      do.MustInvoke[*validator.Validate](i),
	}, nil
}

func (u *AkunUseCase) Create(request *model.AkunRequest) model.AkunResponse {
	validation.ValidateAkunRequest(u.Validator, u.Log, request)

	encrypted, err := helper.EncryptPassword(request.Password)
	if err != nil {
		u.Log.Error().Err(err).Msg("Failed to encrypt password")
		panic(exception.InternalServerError{
			Message: "Failed to encrypt password",
		})
	}

	akun := entity.Akun{
		NIP:      request.NIP,
		Email:    request.Email,
		Password: string(encrypted),
		RoleNama: request.RoleNama,
	}

	if err := u.AkunRepository.Insert(&akun); err != nil {
		u.Log.Error().Err(err).Msg("Failed to create akun")
		panic(exception.InternalServerError{
			Message: "Failed to create akun",
		})
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
	if err != nil {
		u.Log.Error().Err(err).Msg("Failed to get akun")
		panic(exception.InternalServerError{
			Message: "Failed to get akun",
		})
	}

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
	if err != nil {
		u.Log.Error().Err(err).Msg("Failed to get akun")
		panic(exception.InternalServerError{
			Message: "Failed to get akun",
		})
	}

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
		u.Log.Info().Str("nip", nip).Msg("Akun not found")
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
	validation.ValidateAkunRequest(u.Validator, u.Log, request)

	akun, err := u.AkunRepository.FindByNIP(nip)
	if err != nil {
		u.Log.Info().Str("nip", nip).Msg("Akun not found")
		panic(exception.NotFoundError{
			Message: "Akun not found",
		})
	}

	if request.Password != "" {
		encrypted, err := helper.EncryptPassword(request.Password)
		if err != nil {
			u.Log.Error().Err(err).Msg("Failed to encrypt password")
			panic(exception.InternalServerError{
				Message: "Failed to encrypt password",
			})
		}

		akun.Password = string(encrypted)
	}

	akun.Email = request.Email

	response := model.AkunResponse{
		NIP:      akun.NIP,
		Email:    akun.Email,
		RoleNama: akun.RoleNama,
	}

	if err := u.AkunRepository.Update(&akun); err != nil {
		if err != nil {
			u.Log.Error().Err(err).Msg("Failed to update akun")
			panic(exception.InternalServerError{
				Message: "Failed to update akun",
			})
		}
	}

	return response
}

func (u *AkunUseCase) UpdateAdmin(nip string, request *model.AkunRequest) model.AkunResponse {
	validation.ValidateAkunRequest(u.Validator, u.Log, request)

	akun, err := u.AkunRepository.FindByNIP(nip)
	if err != nil {
		u.Log.Info().Str("nip", nip).Msg("Akun not found")
		panic(exception.NotFoundError{
			Message: "Akun not found",
		})
	}

	if request.Password != "" {
		encrypted, err := helper.EncryptPassword(request.Password)
		if err != nil {
			u.Log.Error().Err(err).Msg("Failed to encrypt password")
			panic(exception.InternalServerError{
				Message: "Failed to encrypt password",
			})
		}

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
		if err != nil {
			u.Log.Error().Err(err).Msg("Failed to update akun")
			panic(exception.InternalServerError{
				Message: "Failed to update akun",
			})
		}
	}

	return response
}

func (u *AkunUseCase) Delete(nip string) {
	akun, err := u.AkunRepository.FindByNIP(nip)
	if err != nil {
		u.Log.Info().Str("nip", nip).Msg("Akun not found")
		panic(exception.NotFoundError{
			Message: "Akun not found",
		})
	}

	if err := u.AkunRepository.Delete(&akun); err != nil {
		u.Log.Error().Err(err).Msg("Failed to delete akun")
		panic(exception.InternalServerError{
			Message: "Failed to delete akun",
		})
	}
}
