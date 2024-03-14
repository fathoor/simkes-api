package usecase

import (
	"fmt"
	"github.com/fathoor/simkes-api/internal/entity"
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/repository"
	"github.com/fathoor/simkes-api/internal/validation"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/samber/do"
	"time"
)

type KehadiranUseCase struct {
	KehadiranRepository *repository.KehadiranRepository
	ShiftRepository     *repository.ShiftRepository
	Log                 *zerolog.Logger
	Validator           *validator.Validate
}

func NewKehadiranUseCase(i *do.Injector) (*KehadiranUseCase, error) {
	return &KehadiranUseCase{
		KehadiranRepository: do.MustInvoke[*repository.KehadiranRepository](i),
		ShiftRepository:     do.MustInvoke[*repository.ShiftRepository](i),
		Log:                 do.MustInvoke[*zerolog.Logger](i),
		Validator:           do.MustInvoke[*validator.Validate](i),
	}, nil
}

func (u *KehadiranUseCase) CheckIn(request *model.KehadiranRequest) model.KehadiranResponse {
	if valid := validation.ValidateKehadiranRequest(u.Validator, u.Log, request); valid != nil {
		u.Log.Error().Err(valid).Msg("Validation error")
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	shift, err := u.ShiftRepository.FindByNama(request.ShiftNama)
	if err != nil {
		u.Log.Info().Str("shift", request.ShiftNama).Msg("Shift not found")
		panic(exception.NotFoundError{
			Message: "Shift not found",
		})
	}

	tanggal, err := time.Parse("2006-01-02", request.Tanggal)
	if err != nil {
		u.Log.Info().Str("tanggal", request.Tanggal).Msg("Invalid date")
		panic(exception.BadRequestError{
			Message: "Invalid date",
		})
	}

	jamMasuk := time.Now()
	date := time.Now().Format("2006-01-02")

	shiftMasuk := shift.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05")
	masuk, err := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("%s %s", date, shiftMasuk), time.FixedZone("WIB", 7*60*60))
	if err != nil {
		u.Log.Info().Str("jam_masuk", shiftMasuk).Msg("Invalid time")
		panic(exception.BadRequestError{
			Message: "Invalid time",
		})
	}

	var keterangan string
	if jamMasuk.After(masuk) {
		keterangan = "Terlambat"
	} else {
		keterangan = "Hadir"
	}

	kehadiran := entity.Kehadiran{
		ID:         uuid.New(),
		NIP:        request.NIP,
		Tanggal:    tanggal,
		ShiftNama:  request.ShiftNama,
		JamMasuk:   jamMasuk,
		Keterangan: keterangan,
	}

	if err := u.KehadiranRepository.Insert(&kehadiran); err != nil {
		u.Log.Error().Err(err).Msg("Failed to insert kehadiran")
		panic(exception.InternalServerError{
			Message: "Failed to insert kehadiran",
		})
	}

	response := model.KehadiranResponse{
		ID:      kehadiran.ID.String(),
		NIP:     kehadiran.NIP,
		Tanggal: kehadiran.Tanggal.Format("2006-01-02"),
		Shift: model.KehadiranShiftResponse{
			Nama:      kehadiran.ShiftNama,
			JamMasuk:  shift.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
			JamKeluar: shift.JamKeluar.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
		},
		JamMasuk:   kehadiran.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
		Keterangan: kehadiran.Keterangan,
	}

	return response
}

func (u *KehadiranUseCase) CheckOut(request *model.KehadiranRequest) model.KehadiranResponse {
	if valid := validation.ValidateKehadiranRequest(u.Validator, u.Log, request); valid != nil {
		u.Log.Error().Err(valid).Msg("Validation error")
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	shift, err := u.ShiftRepository.FindByNama(request.ShiftNama)
	if err != nil {
		u.Log.Info().Str("shift", request.ShiftNama).Msg("Shift not found")
		panic(exception.NotFoundError{
			Message: "Shift not found",
		})
	}

	kehadiran, err := u.KehadiranRepository.FindLatestByNIP(request.NIP)
	if err != nil {
		u.Log.Info().Str("nip", request.NIP).Msg("Kehadiran not found")
		panic(exception.NotFoundError{
			Message: "Kehadiran not found",
		})
	}

	jamKeluar := time.Now()

	kehadiran.JamKeluar = jamKeluar

	if err := u.KehadiranRepository.Update(&kehadiran); err != nil {
		u.Log.Error().Err(err).Msg("Failed to update kehadiran")
		panic(exception.InternalServerError{
			Message: "Failed to update kehadiran",
		})
	}

	response := model.KehadiranResponse{
		ID:      kehadiran.ID.String(),
		NIP:     kehadiran.NIP,
		Tanggal: kehadiran.Tanggal.Format("2006-01-02"),
		Shift: model.KehadiranShiftResponse{
			Nama:      kehadiran.ShiftNama,
			JamMasuk:  shift.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
			JamKeluar: shift.JamKeluar.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
		},
		JamMasuk:   kehadiran.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
		JamKeluar:  kehadiran.JamKeluar.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
		Keterangan: kehadiran.Keterangan,
	}

	return response
}

func (u *KehadiranUseCase) GetAll() []model.KehadiranResponse {
	kehadiran, err := u.KehadiranRepository.FindAll()
	if err != nil {
		u.Log.Error().Err(err).Msg("Failed to get kehadiran")
		panic(exception.InternalServerError{
			Message: "Failed to get kehadiran",
		})
	}

	response := make([]model.KehadiranResponse, len(kehadiran))
	for i, kehadiran := range kehadiran {
		response[i] = model.KehadiranResponse{
			ID:      kehadiran.ID.String(),
			NIP:     kehadiran.NIP,
			Tanggal: kehadiran.Tanggal.Format("2006-01-02"),
			Shift: model.KehadiranShiftResponse{
				Nama:      kehadiran.ShiftNama,
				JamMasuk:  kehadiran.Shift.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
				JamKeluar: kehadiran.Shift.JamKeluar.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
			},
			JamMasuk:   kehadiran.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
			JamKeluar:  kehadiran.JamKeluar.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
			Keterangan: kehadiran.Keterangan,
		}
	}

	return response
}

func (u *KehadiranUseCase) GetByNIP(nip string) []model.KehadiranResponse {
	kehadiran, err := u.KehadiranRepository.FindByNIP(nip)
	if err != nil {
		u.Log.Error().Err(err).Str("nip", nip).Msg("Failed to get kehadiran")
		panic(exception.InternalServerError{
			Message: "Failed to get kehadiran",
		})
	}

	response := make([]model.KehadiranResponse, len(kehadiran))
	for i, kehadiran := range kehadiran {
		response[i] = model.KehadiranResponse{
			ID:      kehadiran.ID.String(),
			NIP:     kehadiran.NIP,
			Tanggal: kehadiran.Tanggal.Format("2006-01-02"),
			Shift: model.KehadiranShiftResponse{
				Nama:      kehadiran.ShiftNama,
				JamMasuk:  kehadiran.Shift.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
				JamKeluar: kehadiran.Shift.JamKeluar.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
			},
			JamMasuk:   kehadiran.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
			JamKeluar:  kehadiran.JamKeluar.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
			Keterangan: kehadiran.Keterangan,
		}
	}

	return response
}

func (u *KehadiranUseCase) GetByID(id string) model.KehadiranResponse {
	kehadiranID, err := uuid.Parse(id)
	if err != nil {
		u.Log.Info().Str("id", id).Msg("Invalid ID")
		panic(exception.BadRequestError{
			Message: "Invalid ID",
		})
	}

	kehadiran, err := u.KehadiranRepository.FindByID(kehadiranID)
	if err != nil {
		u.Log.Info().Str("id", id).Msg("Kehadiran not found")
		panic(exception.NotFoundError{
			Message: "Kehadiran not found",
		})
	}

	response := model.KehadiranResponse{
		ID:      kehadiran.ID.String(),
		NIP:     kehadiran.NIP,
		Tanggal: kehadiran.Tanggal.Format("2006-01-02"),
		Shift: model.KehadiranShiftResponse{
			Nama:      kehadiran.ShiftNama,
			JamMasuk:  kehadiran.Shift.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
			JamKeluar: kehadiran.Shift.JamKeluar.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
		},
		JamMasuk:   kehadiran.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
		JamKeluar:  kehadiran.JamKeluar.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
		Keterangan: kehadiran.Keterangan,
	}

	return response
}

func (u *KehadiranUseCase) Update(id string, request *model.KehadiranUpdateRequest) model.KehadiranResponse {
	if valid := validation.ValidateKehadiranUpdateRequest(u.Validator, u.Log, request); valid != nil {
		u.Log.Error().Err(valid).Msg("Validation error")
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	kehadiranID, err := uuid.Parse(id)
	if err != nil {
		u.Log.Info().Str("id", id).Msg("Invalid ID")
		panic(exception.BadRequestError{
			Message: "Invalid ID",
		})
	}

	kehadiran, err := u.KehadiranRepository.FindByID(kehadiranID)
	if err != nil {
		u.Log.Info().Str("id", id).Msg("Kehadiran not found")
		panic(exception.NotFoundError{
			Message: "Kehadiran not found",
		})
	}

	tanggal, err := time.Parse("2006-01-02", request.Tanggal)
	if err != nil {
		u.Log.Info().Str("tanggal", request.Tanggal).Msg("Invalid date")
		panic(exception.BadRequestError{
			Message: "Invalid date",
		})
	}

	jamMasuk, err := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("1970-01-01 %s", request.JamMasuk), time.FixedZone("WIB", 7*60*60))
	if err != nil {
		u.Log.Info().Str("jam_masuk", request.JamMasuk).Msg("Invalid time")
		panic(exception.BadRequestError{
			Message: "Invalid time",
		})
	}

	jamKeluar, err := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("1970-01-01 %s", request.JamKeluar), time.FixedZone("WIB", 7*60*60))
	if err != nil {
		u.Log.Info().Str("jam_keluar", request.JamKeluar).Msg("Invalid time")
		panic(exception.BadRequestError{
			Message: "Invalid time",
		})
	}

	kehadiran.Tanggal = tanggal
	kehadiran.ShiftNama = request.ShiftNama
	kehadiran.JamMasuk = jamMasuk
	kehadiran.JamKeluar = jamKeluar
	kehadiran.Keterangan = request.Keterangan

	if err := u.KehadiranRepository.Update(&kehadiran); err != nil {
		u.Log.Error().Err(err).Msg("Failed to update kehadiran")
		panic(exception.InternalServerError{
			Message: "Failed to update kehadiran",
		})
	}

	response := model.KehadiranResponse{
		ID:      kehadiran.ID.String(),
		NIP:     kehadiran.NIP,
		Tanggal: kehadiran.Tanggal.Format("2006-01-02"),
		Shift: model.KehadiranShiftResponse{
			Nama:      kehadiran.ShiftNama,
			JamMasuk:  kehadiran.Shift.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
			JamKeluar: kehadiran.Shift.JamKeluar.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
		},
		JamMasuk:   kehadiran.JamMasuk.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
		JamKeluar:  kehadiran.JamKeluar.In(time.FixedZone("WIB", 7*60*60)).Format("15:04:05"),
		Keterangan: kehadiran.Keterangan,
	}

	return response

}

func (u *KehadiranUseCase) Delete(id string) {
	kehadiranID, err := uuid.Parse(id)
	if err != nil {
		u.Log.Info().Str("id", id).Msg("Invalid ID")
		panic(exception.BadRequestError{
			Message: "Invalid ID",
		})
	}

	kehadiran, err := u.KehadiranRepository.FindByID(kehadiranID)
	if err != nil {
		u.Log.Info().Str("id", id).Msg("Kehadiran not found")
		panic(exception.NotFoundError{
			Message: "Kehadiran not found",
		})
	}

	if err := u.KehadiranRepository.Delete(&kehadiran); err != nil {
		u.Log.Error().Err(err).Msg("Failed to delete kehadiran")
		panic(exception.InternalServerError{
			Message: "Failed to delete kehadiran",
		})
	}
}
