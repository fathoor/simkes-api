package usecase

import (
	"github.com/fathoor/simkes-api/internal/entity"
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/repository"
	"github.com/fathoor/simkes-api/internal/validation"
	"github.com/samber/do"
)

type JadwalPegawaiUseCase struct {
	JadwalPegawaiRepository *repository.JadwalPegawaiRepository
}

func NewJadwalPegawaiUseCase(i *do.Injector) (*JadwalPegawaiUseCase, error) {
	return &JadwalPegawaiUseCase{
		JadwalPegawaiRepository: do.MustInvoke[*repository.JadwalPegawaiRepository](i),
	}, nil
}

func (u *JadwalPegawaiUseCase) Create(request *model.JadwalPegawaiRequest) model.JadwalPegawaiResponse {
	if valid := validation.ValidateJadwalPegawaiRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	jadwalPegawai := entity.JadwalPegawai{
		NIP:       request.NIP,
		Tahun:     request.Tahun,
		Bulan:     request.Bulan,
		Hari:      request.Hari,
		ShiftNama: request.ShiftNama,
	}

	if err := u.JadwalPegawaiRepository.Insert(&jadwalPegawai); err != nil {
		exception.PanicIfError(err)
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
	exception.PanicIfError(err)

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
	exception.PanicIfError(err)

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
	exception.PanicIfError(err)

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
	if valid := validation.ValidateJadwalPegawaiRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	jadwalPegawai, err := u.JadwalPegawaiRepository.FindByPK(nip, tahun, bulan, hari)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Jadwal Pegawai not found",
		})
	}

	jadwalPegawai.ShiftNama = request.ShiftNama

	if err := u.JadwalPegawaiRepository.Update(&jadwalPegawai); err != nil {
		exception.PanicIfError(err)
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
		panic(exception.NotFoundError{
			Message: "Jadwal Pegawai not found",
		})
	}

	if err := u.JadwalPegawaiRepository.Delete(&jadwalPegawai); err != nil {
		exception.PanicIfError(err)
	}
}
