package usecase

import (
	"fmt"
	"github.com/fathoor/simkes-api/internal/entity"
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/repository"
	"github.com/fathoor/simkes-api/internal/validation"
	"github.com/samber/do"
	"time"
)

type ShiftUseCase struct {
	repository.ShiftRepository
}

func NewShiftUseCase(i *do.Injector) (*ShiftUseCase, error) {
	return &ShiftUseCase{
		ShiftRepository: do.MustInvoke[repository.ShiftRepository](i),
	}, nil
}

func (u *ShiftUseCase) Create(request *model.ShiftRequest) model.ShiftResponse {
	if valid := validation.ValidateShiftRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	jamMasuk, err := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("1970-01-01 %s", request.JamMasuk), time.FixedZone("WIB", 7*60*60))
	exception.PanicIfError(err)

	jamKeluar, err := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("1970-01-01 %s", request.JamKeluar), time.FixedZone("WIB", 7*60*60))
	exception.PanicIfError(err)

	shift := entity.Shift{
		Nama:      request.Nama,
		JamMasuk:  jamMasuk,
		JamKeluar: jamKeluar,
	}

	if err := u.ShiftRepository.Insert(&shift); err != nil {
		exception.PanicIfError(err)
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
	exception.PanicIfError(err)

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
	if valid := validation.ValidateShiftRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	shift, err := u.ShiftRepository.FindByNama(nama)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Shift not found",
		})
	}

	jamMasuk, err := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("1970-01-01 %s", request.JamMasuk), time.FixedZone("WIB", 7*60*60))
	exception.PanicIfError(err)

	jamKeluar, err := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("1970-01-01 %s", request.JamKeluar), time.FixedZone("WIB", 7*60*60))
	exception.PanicIfError(err)

	shift.Nama = request.Nama
	shift.JamMasuk = jamMasuk
	shift.JamKeluar = jamKeluar

	if err := u.ShiftRepository.Update(&shift); err != nil {
		exception.PanicIfError(err)
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
		panic(exception.NotFoundError{
			Message: "Shift not found",
		})
	}

	if err := u.ShiftRepository.Delete(&shift); err != nil {
		exception.PanicIfError(err)
	}
}
