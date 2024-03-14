package usecase

import (
	"fmt"
	"github.com/fathoor/simkes-api/internal/entity"
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/repository"
	"github.com/fathoor/simkes-api/internal/validation"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"github.com/samber/do"
	"time"
)

type ShiftUseCase struct {
	ShiftRepository *repository.ShiftRepository
	Log             *zerolog.Logger
	Validator       *validator.Validate
}

func NewShiftUseCase(i *do.Injector) (*ShiftUseCase, error) {
	return &ShiftUseCase{
		ShiftRepository: do.MustInvoke[*repository.ShiftRepository](i),
		Log:             do.MustInvoke[*zerolog.Logger](i),
		Validator:       do.MustInvoke[*validator.Validate](i),
	}, nil
}

func (u *ShiftUseCase) Create(request *model.ShiftRequest) model.ShiftResponse {
	validation.ValidateShiftRequest(u.Validator, u.Log, request)

	jamMasuk, err := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("1970-01-01 %s", request.JamMasuk), time.FixedZone("WIB", 7*60*60))
	if err != nil {
		u.Log.Info().Str("jam_masuk", request.JamMasuk).Msg("Invalid jam masuk")
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	jamKeluar, err := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("1970-01-01 %s", request.JamKeluar), time.FixedZone("WIB", 7*60*60))
	if err != nil {
		u.Log.Info().Str("jam_keluar", request.JamKeluar).Msg("Invalid jam keluar")
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	shift := entity.Shift{
		Nama:      request.Nama,
		JamMasuk:  jamMasuk,
		JamKeluar: jamKeluar,
	}

	if err := u.ShiftRepository.Insert(&shift); err != nil {
		u.Log.Error().Err(err).Msg("Failed to insert shift")
		panic(exception.InternalServerError{
			Message: "Failed to insert shift",
		})
	}

	response := model.ShiftResponse{
		Nama:      shift.Nama,
		JamMasuk:  shift.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
		JamKeluar: shift.JamKeluar.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
	}

	return response
}

func (u *ShiftUseCase) GetAll() []model.ShiftResponse {
	shift, err := u.ShiftRepository.FindAll()
	if err != nil {
		u.Log.Error().Err(err).Msg("Failed to get shift")
		panic(exception.InternalServerError{
			Message: "Failed to get shift",
		})
	}

	response := make([]model.ShiftResponse, len(shift))
	for i, shift := range shift {
		response[i] = model.ShiftResponse{
			Nama:      shift.Nama,
			JamMasuk:  shift.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
			JamKeluar: shift.JamKeluar.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
		}
	}

	return response
}

func (u *ShiftUseCase) GetByNama(nama string) model.ShiftResponse {
	shift, err := u.ShiftRepository.FindByNama(nama)
	if err != nil {
		u.Log.Info().Str("nama", nama).Msg("Shift not found")
		panic(exception.NotFoundError{
			Message: "Shift not found",
		})
	}

	response := model.ShiftResponse{
		Nama:      shift.Nama,
		JamMasuk:  shift.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
		JamKeluar: shift.JamKeluar.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
	}

	return response
}

func (u *ShiftUseCase) Update(nama string, request *model.ShiftRequest) model.ShiftResponse {
	validation.ValidateShiftRequest(u.Validator, u.Log, request)

	shift, err := u.ShiftRepository.FindByNama(nama)
	if err != nil {
		u.Log.Info().Str("nama", nama).Msg("Shift not found")
		panic(exception.NotFoundError{
			Message: "Shift not found",
		})
	}

	jamMasuk, err := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("1970-01-01 %s", request.JamMasuk), time.FixedZone("WIB", 7*60*60))
	if err != nil {
		u.Log.Info().Str("jam_masuk", request.JamMasuk).Msg("Invalid jam masuk")
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	jamKeluar, err := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("1970-01-01 %s", request.JamKeluar), time.FixedZone("WIB", 7*60*60))
	if err != nil {
		u.Log.Info().Str("jam_keluar", request.JamKeluar).Msg("Invalid jam keluar")
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	shift.Nama = request.Nama
	shift.JamMasuk = jamMasuk
	shift.JamKeluar = jamKeluar

	if err := u.ShiftRepository.Update(&shift); err != nil {
		u.Log.Error().Err(err).Msg("Failed to update shift")
		panic(exception.InternalServerError{
			Message: "Failed to update shift",
		})
	}

	response := model.ShiftResponse{
		Nama:      shift.Nama,
		JamMasuk:  shift.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
		JamKeluar: shift.JamKeluar.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
	}

	return response
}

func (u *ShiftUseCase) Delete(nama string) {
	shift, err := u.ShiftRepository.FindByNama(nama)
	if err != nil {
		u.Log.Info().Str("nama", nama).Msg("Shift not found")
		panic(exception.NotFoundError{
			Message: "Shift not found",
		})
	}

	if err := u.ShiftRepository.Delete(&shift); err != nil {
		u.Log.Error().Err(err).Msg("Failed to delete shift")
		panic(exception.InternalServerError{
			Message: "Failed to delete shift",
		})
	}
}
