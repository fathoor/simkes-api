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

type JabatanUseCase struct {
	JabatanRepository *repository.JabatanRepository
	Log               *zerolog.Logger
	Validator         *validator.Validate
}

func NewJabatanUseCase(i *do.Injector) (*JabatanUseCase, error) {
	return &JabatanUseCase{
		JabatanRepository: do.MustInvoke[*repository.JabatanRepository](i),
		Log:               do.MustInvoke[*zerolog.Logger](i),
		Validator:         do.MustInvoke[*validator.Validate](i),
	}, nil
}

func (u *JabatanUseCase) Create(request *model.JabatanRequest) model.JabatanResponse {
	validation.ValidateJabatanRequest(u.Validator, u.Log, request)

	jabatan := entity.Jabatan{
		Nama:      request.Nama,
		Jenjang:   request.Jenjang,
		GajiPokok: request.GajiPokok,
		Tunjangan: request.Tunjangan,
	}

	if err := u.JabatanRepository.Insert(&jabatan); err != nil {
		u.Log.Error().Err(err).Msg("Failed to insert jabatan")
		panic(exception.InternalServerError{
			Message: "Failed to insert jabatan",
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

func (u *JabatanUseCase) GetAll() []model.JabatanResponse {
	jabatan, err := u.JabatanRepository.FindAll()
	if err != nil {
		u.Log.Error().Err(err).Msg("Failed to get jabatan")
		panic(exception.InternalServerError{
			Message: "Failed to get jabatan",
		})
	}

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
		u.Log.Info().Str("jabatan", j).Msg("Jabatan not found")
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
	validation.ValidateJabatanRequest(u.Validator, u.Log, request)

	jabatan, err := u.JabatanRepository.FindByJabatan(j)
	if err != nil {
		u.Log.Info().Str("jabatan", j).Msg("Jabatan not found")
		panic(exception.NotFoundError{
			Message: "Jabatan not found",
		})
	}

	jabatan.Nama = request.Nama
	jabatan.Jenjang = request.Jenjang
	jabatan.GajiPokok = request.GajiPokok
	jabatan.Tunjangan = request.Tunjangan

	if err := u.JabatanRepository.Update(&jabatan); err != nil {
		u.Log.Error().Err(err).Msg("Failed to update jabatan")
		panic(exception.InternalServerError{
			Message: "Failed to update jabatan",
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

func (u *JabatanUseCase) Delete(j string) {
	jabatan, err := u.JabatanRepository.FindByJabatan(j)
	if err != nil {
		u.Log.Info().Str("jabatan", j).Msg("Jabatan not found")
		panic(exception.NotFoundError{
			Message: "Jabatan not found",
		})
	}

	if err := u.JabatanRepository.Delete(&jabatan); err != nil {
		u.Log.Error().Err(err).Msg("Failed to delete jabatan")
		panic(exception.InternalServerError{
			Message: "Failed to delete jabatan",
		})
	}
}
