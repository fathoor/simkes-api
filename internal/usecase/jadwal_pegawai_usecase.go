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

type JadwalPegawaiUseCase struct {
	JadwalPegawaiRepository *repository.JadwalPegawaiRepository
	Log                     *zerolog.Logger
	Validator               *validator.Validate
}

func NewJadwalPegawaiUseCase(i *do.Injector) (*JadwalPegawaiUseCase, error) {
	return &JadwalPegawaiUseCase{
		JadwalPegawaiRepository: do.MustInvoke[*repository.JadwalPegawaiRepository](i),
		Log:                     do.MustInvoke[*zerolog.Logger](i),
		Validator:               do.MustInvoke[*validator.Validate](i),
	}, nil
}

func (u *JadwalPegawaiUseCase) Create(request *model.JadwalPegawaiRequest) model.JadwalPegawaiResponse {
	validation.ValidateJadwalPegawaiRequest(u.Validator, u.Log, request)

	jadwalPegawai := entity.JadwalPegawai{
		NIP:       request.NIP,
		Tahun:     request.Tahun,
		Bulan:     request.Bulan,
		Hari:      request.Hari,
		ShiftNama: request.ShiftNama,
	}

	if err := u.JadwalPegawaiRepository.Insert(&jadwalPegawai); err != nil {
		u.Log.Error().Err(err).Msg("Failed to insert Jadwal Pegawai")
		panic(exception.InternalServerError{
			Message: "Failed to insert Jadwal Pegawai",
		})
	}

	response := model.JadwalPegawaiResponse{
		NIP:       jadwalPegawai.NIP,
		Tahun:     jadwalPegawai.Tahun,
		Bulan:     jadwalPegawai.Bulan,
		Hari:      jadwalPegawai.Hari,
		ShiftNama: jadwalPegawai.ShiftNama,
	}

	return response
}

func (u *JadwalPegawaiUseCase) GetAll() []model.JadwalPegawaiResponse {
	jadwalPegawai, err := u.JadwalPegawaiRepository.FindAll()
	if err != nil {
		u.Log.Error().Err(err).Msg("Failed to get Jadwal Pegawai")
		panic(exception.InternalServerError{
			Message: "Failed to get Jadwal Pegawai",
		})
	}

	response := make([]model.JadwalPegawaiResponse, len(jadwalPegawai))
	for i, jadwalPegawai := range jadwalPegawai {
		response[i] = model.JadwalPegawaiResponse{
			NIP:       jadwalPegawai.NIP,
			Tahun:     jadwalPegawai.Tahun,
			Bulan:     jadwalPegawai.Bulan,
			Hari:      jadwalPegawai.Hari,
			ShiftNama: jadwalPegawai.ShiftNama,
		}
	}

	return response
}

func (u *JadwalPegawaiUseCase) GetByNIP(nip string) []model.JadwalPegawaiResponse {
	jadwalPegawai, err := u.JadwalPegawaiRepository.FindByNIP(nip)
	if err != nil {
		u.Log.Info().Str("nip", nip).Msg("Jadwal Pegawai not found")
		panic(exception.NotFoundError{
			Message: "Jadwal Pegawai not found",
		})
	}

	response := make([]model.JadwalPegawaiResponse, len(jadwalPegawai))
	for i, jadwalPegawai := range jadwalPegawai {
		response[i] = model.JadwalPegawaiResponse{
			NIP:       jadwalPegawai.NIP,
			Tahun:     jadwalPegawai.Tahun,
			Bulan:     jadwalPegawai.Bulan,
			Hari:      jadwalPegawai.Hari,
			ShiftNama: jadwalPegawai.ShiftNama,
		}
	}

	return response
}

func (u *JadwalPegawaiUseCase) GetByTahunBulan(tahun, bulan int16) []model.JadwalPegawaiResponse {
	jadwalPegawai, err := u.JadwalPegawaiRepository.FindByTahunBulan(tahun, bulan)
	if err != nil {
		u.Log.Info().Int16("tahun", tahun).Int16("bulan", bulan).Msg("Jadwal Pegawai not found")
		panic(exception.NotFoundError{
			Message: "Jadwal Pegawai not found",
		})
	}

	response := make([]model.JadwalPegawaiResponse, len(jadwalPegawai))
	for i, jadwalPegawai := range jadwalPegawai {
		response[i] = model.JadwalPegawaiResponse{
			NIP:       jadwalPegawai.NIP,
			Tahun:     jadwalPegawai.Tahun,
			Bulan:     jadwalPegawai.Bulan,
			Hari:      jadwalPegawai.Hari,
			ShiftNama: jadwalPegawai.ShiftNama,
		}
	}

	return response
}

func (u *JadwalPegawaiUseCase) GetByPK(nip string, tahun, bulan, hari int16) model.JadwalPegawaiResponse {
	jadwalPegawai, err := u.JadwalPegawaiRepository.FindByPK(nip, tahun, bulan, hari)
	if err != nil {
		u.Log.Info().Str("nip", nip).Int16("tahun", tahun).Int16("bulan", bulan).Int16("hari", hari).Msg("Jadwal Pegawai not found")
		panic(exception.NotFoundError{
			Message: "Jadwal Pegawai not found",
		})
	}

	response := model.JadwalPegawaiResponse{
		NIP:       jadwalPegawai.NIP,
		Tahun:     jadwalPegawai.Tahun,
		Bulan:     jadwalPegawai.Bulan,
		Hari:      jadwalPegawai.Hari,
		ShiftNama: jadwalPegawai.ShiftNama,
	}

	return response
}

func (u *JadwalPegawaiUseCase) Update(nip string, tahun, bulan, hari int16, request *model.JadwalPegawaiRequest) model.JadwalPegawaiResponse {
	validation.ValidateJadwalPegawaiRequest(u.Validator, u.Log, request)

	jadwalPegawai, err := u.JadwalPegawaiRepository.FindByPK(nip, tahun, bulan, hari)
	if err != nil {
		u.Log.Info().Str("nip", nip).Int16("tahun", tahun).Int16("bulan", bulan).Int16("hari", hari).Msg("Jadwal Pegawai not found")
		panic(exception.NotFoundError{
			Message: "Jadwal Pegawai not found",
		})
	}

	jadwalPegawai.ShiftNama = request.ShiftNama

	if err := u.JadwalPegawaiRepository.Update(&jadwalPegawai); err != nil {
		u.Log.Error().Err(err).Msg("Failed to update Jadwal Pegawai")
		panic(exception.InternalServerError{
			Message: "Failed to update Jadwal Pegawai",
		})
	}

	response := model.JadwalPegawaiResponse{
		NIP:       jadwalPegawai.NIP,
		Tahun:     jadwalPegawai.Tahun,
		Bulan:     jadwalPegawai.Bulan,
		Hari:      jadwalPegawai.Hari,
		ShiftNama: jadwalPegawai.ShiftNama,
	}

	return response
}

func (u *JadwalPegawaiUseCase) Delete(nip string, tahun, bulan, hari int16) {
	jadwalPegawai, err := u.JadwalPegawaiRepository.FindByPK(nip, tahun, bulan, hari)
	if err != nil {
		u.Log.Info().Str("nip", nip).Int16("tahun", tahun).Int16("bulan", bulan).Int16("hari", hari).Msg("Jadwal Pegawai not found")
		panic(exception.NotFoundError{
			Message: "Jadwal Pegawai not found",
		})
	}

	if err := u.JadwalPegawaiRepository.Delete(&jadwalPegawai); err != nil {
		u.Log.Error().Err(err).Msg("Failed to delete Jadwal Pegawai")
		panic(exception.InternalServerError{
			Message: "Failed to delete Jadwal Pegawai",
		})
	}
}
