package service

import (
	"github.com/fathoor/simkes-api/internal/app/exception"
	"github.com/fathoor/simkes-api/internal/kehadiran/entity"
	"github.com/fathoor/simkes-api/internal/kehadiran/model"
	"github.com/fathoor/simkes-api/internal/kehadiran/repository"
	"github.com/fathoor/simkes-api/internal/kehadiran/validation"
	"github.com/google/uuid"
	"time"
)

type kehadiranServiceImpl struct {
	repository.KehadiranRepository
}

func (service *kehadiranServiceImpl) CheckIn(request *model.KehadiranRequest) model.KehadiranResponse {
	if valid := validation.ValidateKehadiranRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	tanggal, err := time.Parse("2006-01-02", request.Tanggal)
	exception.PanicIfError(err)

	kehadiran := entity.Kehadiran{
		ID:        uuid.New(),
		NIP:       request.NIP,
		Tanggal:   tanggal,
		ShiftNama: request.ShiftNama,
		JamMasuk:  time.Now(),
	}

	if err := service.KehadiranRepository.Insert(&kehadiran); err != nil {
		exception.PanicIfError(err)
	}

	response := model.KehadiranResponse{
		ID:      kehadiran.ID.String(),
		NIP:     kehadiran.NIP,
		Tanggal: kehadiran.Tanggal.Format("2006-01-02"),
		Shift: model.KehadiranShiftResponse{
			Nama:      kehadiran.Shift.Nama,
			JamMasuk:  kehadiran.Shift.JamMasuk.Format("15:04:05"),
			JamKeluar: kehadiran.Shift.JamKeluar.Format("15:04:05"),
		},
		JamMasuk:   kehadiran.JamMasuk.Format("15:04:05"),
		JamKeluar:  kehadiran.JamKeluar.Format("15:04:05"),
		Keterangan: kehadiran.Keterangan,
	}

	return response
}

func (service *kehadiranServiceImpl) CheckOut(request *model.KehadiranRequest) model.KehadiranResponse {
	//TODO implement me
	panic("implement me")
}

func (service *kehadiranServiceImpl) GetAll() []model.KehadiranResponse {
	kehadiran, err := service.KehadiranRepository.FindAll()
	exception.PanicIfError(err)

	var response []model.KehadiranResponse
	for i, kehadiran := range kehadiran {
		response[i] = model.KehadiranResponse{
			ID:      kehadiran.ID.String(),
			NIP:     kehadiran.NIP,
			Tanggal: kehadiran.Tanggal.Format("2006-01-02"),
			Shift: model.KehadiranShiftResponse{
				Nama:      kehadiran.Shift.Nama,
				JamMasuk:  kehadiran.Shift.JamMasuk.Format("15:04:05"),
				JamKeluar: kehadiran.Shift.JamKeluar.Format("15:04:05"),
			},
			JamMasuk:   kehadiran.JamMasuk.Format("15:04:05"),
			JamKeluar:  kehadiran.JamKeluar.Format("15:04:05"),
			Keterangan: kehadiran.Keterangan,
		}
	}

	return response
}

func (service *kehadiranServiceImpl) GetByNIP(nip string) []model.KehadiranResponse {
	kehadiran, err := service.KehadiranRepository.FindByNIP(nip)
	exception.PanicIfError(err)

	var response []model.KehadiranResponse
	for i, kehadiran := range kehadiran {
		response[i] = model.KehadiranResponse{
			ID:      kehadiran.ID.String(),
			NIP:     kehadiran.NIP,
			Tanggal: kehadiran.Tanggal.Format("2006-01-02"),
			Shift: model.KehadiranShiftResponse{
				Nama:      kehadiran.Shift.Nama,
				JamMasuk:  kehadiran.Shift.JamMasuk.Format("15:04:05"),
				JamKeluar: kehadiran.Shift.JamKeluar.Format("15:04:05"),
			},
			JamMasuk:   kehadiran.JamMasuk.Format("15:04:05"),
			JamKeluar:  kehadiran.JamKeluar.Format("15:04:05"),
			Keterangan: kehadiran.Keterangan,
		}
	}

	return response
}

func (service *kehadiranServiceImpl) GetByID(id string) model.KehadiranResponse {
	kehadiranID, err := uuid.Parse(id)
	exception.PanicIfError(err)

	kehadiran, err := service.KehadiranRepository.FindByID(kehadiranID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Kehadiran not found",
		})
	}

	response := model.KehadiranResponse{
		ID:      kehadiran.ID.String(),
		NIP:     kehadiran.NIP,
		Tanggal: kehadiran.Tanggal.Format("2006-01-02"),
		Shift: model.KehadiranShiftResponse{
			Nama:      kehadiran.Shift.Nama,
			JamMasuk:  kehadiran.Shift.JamMasuk.Format("15:04:05"),
			JamKeluar: kehadiran.Shift.JamKeluar.Format("15:04:05"),
		},
		JamMasuk:   kehadiran.JamMasuk.Format("15:04:05"),
		JamKeluar:  kehadiran.JamKeluar.Format("15:04:05"),
		Keterangan: kehadiran.Keterangan,
	}

	return response
}

func (service *kehadiranServiceImpl) Update(id string, request *model.KehadiranUpdateRequest) model.KehadiranResponse {
	if valid := validation.ValidateKehadiranUpdateRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	kehadiranID, err := uuid.Parse(id)
	exception.PanicIfError(err)

	kehadiran, err := service.KehadiranRepository.FindByID(kehadiranID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Kehadiran not found",
		})
	}

	tanggal, err := time.Parse("2006-01-02", request.Tanggal)
	exception.PanicIfError(err)

	jamMasuk, err := time.Parse("15:04:05", request.JamMasuk)
	exception.PanicIfError(err)

	jamKeluar, err := time.Parse("15:04:05", request.JamKeluar)
	exception.PanicIfError(err)

	kehadiran.Tanggal = tanggal
	kehadiran.ShiftNama = request.ShiftNama
	kehadiran.JamMasuk = jamMasuk
	kehadiran.JamKeluar = jamKeluar
	kehadiran.Keterangan = request.Keterangan

	if err := service.KehadiranRepository.Update(&kehadiran); err != nil {
		exception.PanicIfError(err)
	}

	response := model.KehadiranResponse{
		ID:      kehadiran.ID.String(),
		NIP:     kehadiran.NIP,
		Tanggal: kehadiran.Tanggal.Format("2006-01-02"),
		Shift: model.KehadiranShiftResponse{
			Nama:      kehadiran.Shift.Nama,
			JamMasuk:  kehadiran.Shift.JamMasuk.Format("15:04:05"),
			JamKeluar: kehadiran.Shift.JamKeluar.Format("15:04:05"),
		},
		JamMasuk:   kehadiran.JamMasuk.Format("15:04:05"),
		JamKeluar:  kehadiran.JamKeluar.Format("15:04:05"),
		Keterangan: kehadiran.Keterangan,
	}

	return response

}

func (service *kehadiranServiceImpl) Delete(id string) {
	kehadiranID, err := uuid.Parse(id)
	exception.PanicIfError(err)

	kehadiran, err := service.KehadiranRepository.FindByID(kehadiranID)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Kehadiran not found",
		})
	}

	if err := service.KehadiranRepository.Delete(&kehadiran); err != nil {
		exception.PanicIfError(err)
	}
}

func NewKehadiranServiceProvider(repository *repository.KehadiranRepository) KehadiranService {
	return &kehadiranServiceImpl{*repository}
}
