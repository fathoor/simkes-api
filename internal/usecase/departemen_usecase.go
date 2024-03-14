package usecase

import (
	"github.com/fathoor/simkes-api/internal/entity"
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/repository"
	"github.com/fathoor/simkes-api/internal/validation"
	"github.com/samber/do"
)

type DepartemenUseCase struct {
	DepartemenRepository *repository.DepartemenRepository
}

func NewDepartemenUseCase(i *do.Injector) (*DepartemenUseCase, error) {
	return &DepartemenUseCase{
		DepartemenRepository: do.MustInvoke[*repository.DepartemenRepository](i),
	}, nil
}

func (u *DepartemenUseCase) Create(request *model.DepartemenRequest) model.DepartemenResponse {
	if err := validation.ValidateDepartemenRequest(request); err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	departemen := entity.Departemen{
		Nama: request.Nama,
	}

	if err := u.DepartemenRepository.Insert(&departemen); err != nil {
		exception.PanicIfError(err)
	}

	response := model.DepartemenResponse{
		Nama: departemen.Nama,
	}

	return response
}

func (u *DepartemenUseCase) GetAll() []model.DepartemenResponse {
	departemen, err := u.DepartemenRepository.FindAll()
	exception.PanicIfError(err)

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
	exception.PanicIfError(err)

	response := model.DepartemenResponse{
		Nama: departemen.Nama,
	}

	return response
}

func (u *DepartemenUseCase) Update(d string, request *model.DepartemenRequest) model.DepartemenResponse {
	if err := validation.ValidateDepartemenRequest(request); err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	departemen, err := u.DepartemenRepository.FindByDepartemen(d)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Departemen not found",
		})
	}

	departemen.Nama = request.Nama

	if err := u.DepartemenRepository.Update(&departemen); err != nil {
		exception.PanicIfError(err)
	}

	response := model.DepartemenResponse{
		Nama: departemen.Nama,
	}

	return response
}

func (u *DepartemenUseCase) Delete(d string) {
	departemen, err := u.DepartemenRepository.FindByDepartemen(d)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Departemen not found",
		})
	}

	if err := u.DepartemenRepository.Delete(&departemen); err != nil {
		exception.PanicIfError(err)
	}
}
