package service

import (
	"fmt"
	"github.com/fathoor/simkes-api/internal/app/exception"
	"github.com/fathoor/simkes-api/internal/shift/entity"
	"github.com/fathoor/simkes-api/internal/shift/model"
	"github.com/fathoor/simkes-api/internal/shift/repository"
	"github.com/fathoor/simkes-api/internal/shift/validation"
	"time"
)

type shiftServiceImpl struct {
	repository.ShiftRepository
}

func (service *shiftServiceImpl) Create(request *model.ShiftRequest) model.ShiftResponse {
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

	if err := service.ShiftRepository.Insert(&shift); err != nil {
		exception.PanicIfError(err)
	}

	response := model.ShiftResponse{
		Nama:      shift.Nama,
		JamMasuk:  shift.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
		JamKeluar: shift.JamKeluar.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
	}

	return response
}

func (service *shiftServiceImpl) GetAll() []model.ShiftResponse {
	shift, err := service.ShiftRepository.FindAll()
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

func (service *shiftServiceImpl) GetByNama(nama string) model.ShiftResponse {
	shift, err := service.ShiftRepository.FindByNama(nama)
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

func (service *shiftServiceImpl) Update(nama string, request *model.ShiftRequest) model.ShiftResponse {
	if valid := validation.ValidateShiftRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	shift, err := service.ShiftRepository.FindByNama(nama)
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

	if err := service.ShiftRepository.Update(&shift); err != nil {
		exception.PanicIfError(err)
	}

	response := model.ShiftResponse{
		Nama:      shift.Nama,
		JamMasuk:  shift.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
		JamKeluar: shift.JamKeluar.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
	}

	return response
}

func (service *shiftServiceImpl) Delete(nama string) {
	shift, err := service.ShiftRepository.FindByNama(nama)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Shift not found",
		})
	}

	if err := service.ShiftRepository.Delete(&shift); err != nil {
		exception.PanicIfError(err)
	}
}

func NewShiftServiceProvider(repository *repository.ShiftRepository) ShiftService {
	return &shiftServiceImpl{*repository}
}
