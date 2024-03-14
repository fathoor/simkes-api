package usecase

import (
	"github.com/fathoor/simkes-api/internal/entity"
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/repository"
	"github.com/fathoor/simkes-api/internal/validation"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"github.com/samber/do"
)

type DepartemenUseCase struct {
	DepartemenRepository *repository.DepartemenRepository
	Log                  *zerolog.Logger
	Validator            *validator.Validate
}

func NewDepartemenUseCase(i *do.Injector) (*DepartemenUseCase, error) {
	return &DepartemenUseCase{
		DepartemenRepository: do.MustInvoke[*repository.DepartemenRepository](i),
		Log:                  do.MustInvoke[*zerolog.Logger](i),
		Validator:            do.MustInvoke[*validator.Validate](i),
	}, nil
}

func (u *DepartemenUseCase) Create(request *model.DepartemenRequest) model.DepartemenResponse {
	validation.ValidateDepartemenRequest(u.Validator, u.Log, request)

	departemen := entity.Departemen{
		Nama: request.Nama,
	}

	if err := u.DepartemenRepository.Insert(&departemen); err != nil {
		u.Log.Error().Err(err).Msg("Failed to insert departemen")
		panic(exception.InternalServerError{
			Message: "Failed to insert departemen",
		})
	}

	response := model.DepartemenResponse{
		Nama: departemen.Nama,
	}

	return response
}

func (u *DepartemenUseCase) GetAll() []model.DepartemenResponse {
	departemen, err := u.DepartemenRepository.FindAll()
	if err != nil {
		u.Log.Error().Err(err).Msg("Failed to get departemen")
		panic(exception.InternalServerError{
			Message: "Failed to get departemen",
		})
	}

	response := make([]model.DepartemenResponse, len(departemen))
	for i, departemen := range departemen {
		response[i] = model.DepartemenResponse{
			Nama: departemen.Nama,
		}
	}

	return response
}

func (u *DepartemenUseCase) GetByDepartemen(d string) model.DepartemenResponse {
	departemen, err := u.DepartemenRepository.FindByDepartemen(d)
	if err != nil {
		u.Log.Info().Str("departemen", d).Msg("Departemen not found")
		panic(exception.NotFoundError{
			Message: "Departemen not found",
		})
	}

	response := model.DepartemenResponse{
		Nama: departemen.Nama,
	}

	return response
}

func (u *DepartemenUseCase) Update(d string, request *model.DepartemenRequest) model.DepartemenResponse {
	validation.ValidateDepartemenRequest(u.Validator, u.Log, request)

	departemen, err := u.DepartemenRepository.FindByDepartemen(d)
	if err != nil {
		u.Log.Info().Str("departemen", d).Msg("Departemen not found")
		panic(exception.NotFoundError{
			Message: "Departemen not found",
		})
	}

	departemen.Nama = request.Nama

	if err := u.DepartemenRepository.Update(&departemen); err != nil {
		u.Log.Error().Err(err).Msg("Failed to update departemen")
		panic(exception.InternalServerError{
			Message: "Failed to update departemen",
		})
	}

	response := model.DepartemenResponse{
		Nama: departemen.Nama,
	}

	return response
}

func (u *DepartemenUseCase) Delete(d string) {
	departemen, err := u.DepartemenRepository.FindByDepartemen(d)
	if err != nil {
		u.Log.Info().Str("departemen", d).Msg("Departemen not found")
		panic(exception.NotFoundError{
			Message: "Departemen not found",
		})
	}

	if err := u.DepartemenRepository.Delete(&departemen); err != nil {
		u.Log.Error().Err(err).Msg("Failed to delete departemen")
		panic(exception.InternalServerError{
			Message: "Failed to delete departemen",
		})
	}
}
