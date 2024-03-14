package usecase

import (
	"github.com/fathoor/simkes-api/internal/entity"
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/repository"
	"github.com/fathoor/simkes-api/internal/validation"
	"github.com/samber/do"
)

type JabatanUseCase struct {
	JabatanRepository *repository.JabatanRepository
}

func NewJabatanUseCase(i *do.Injector) (*JabatanUseCase, error) {
	return &JabatanUseCase{
		JabatanRepository: do.MustInvoke[*repository.JabatanRepository](i),
	}, nil
}

func (u *JabatanUseCase) Create(request *model.JabatanRequest) model.JabatanResponse {
	if err := validation.ValidateJabatanRequest(request); err != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	jabatan := entity.Jabatan{
		Nama:      request.Nama,
		Jenjang:   request.Jenjang,
		GajiPokok: request.GajiPokok,
		Tunjangan: request.Tunjangan,
	}

	if err := u.JabatanRepository.Insert(&jabatan); err != nil {
		exception.PanicIfError(err)
	}

	response := model.JabatanResponse{
		Nama:      jabatan.Nama,
		Jenjang:   jabatan.Jenjang,
		GajiPokok: jabatan.GajiPokok,
		Tunjangan: jabatan.Tunjangan,
	}

	return response
}

func (u *JabatanUseCase) GetAll() []model.JabatanResponse {
	jabatan, err := u.JabatanRepository.FindAll()
	exception.PanicIfError(err)

	response := make([]model.JabatanResponse, len(jabatan))
	for i, jabatan := range jabatan {
		response[i] = model.JabatanResponse{
			Nama:      jabatan.Nama,
			Jenjang:   jabatan.Jenjang,
			GajiPokok: jabatan.GajiPokok,
			Tunjangan: jabatan.Tunjangan,
		}
	}

	return response
}

func (u *JabatanUseCase) GetByJabatan(j string) model.JabatanResponse {
	jabatan, err := u.JabatanRepository.FindByJabatan(j)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Jabatan not found",
		})
	}

	response := model.JabatanResponse{
		Nama:      jabatan.Nama,
		Jenjang:   jabatan.Jenjang,
		GajiPokok: jabatan.GajiPokok,
		Tunjangan: jabatan.Tunjangan,
	}

	return response
}

func (u *JabatanUseCase) Update(j string, request *model.JabatanRequest) model.JabatanResponse {
	if valid := validation.ValidateJabatanRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	jabatan, err := u.JabatanRepository.FindByJabatan(j)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Jabatan not found",
		})
	}

	jabatan.Nama = request.Nama
	jabatan.Jenjang = request.Jenjang
	jabatan.GajiPokok = request.GajiPokok
	jabatan.Tunjangan = request.Tunjangan

	if err := u.JabatanRepository.Update(&jabatan); err != nil {
		exception.PanicIfError(err)
	}

	response := model.JabatanResponse{
		Nama:      jabatan.Nama,
		Jenjang:   jabatan.Jenjang,
		GajiPokok: jabatan.GajiPokok,
		Tunjangan: jabatan.Tunjangan,
	}

	return response
}

func (u *JabatanUseCase) Delete(j string) {
	jabatan, err := u.JabatanRepository.FindByJabatan(j)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Jabatan not found",
		})
	}

	if err := u.JabatanRepository.Delete(&jabatan); err != nil {
		exception.PanicIfError(err)
	}
}
