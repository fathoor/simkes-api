package usecase

import (
	"github.com/fathoor/simkes-api/internal/entity"
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/repository"
	"github.com/fathoor/simkes-api/internal/validation"
	"github.com/samber/do"
	"time"
)

type PegawaiUseCase struct {
	PegawaiRepository *repository.PegawaiRepository
}

func NewPegawaiUseCase(i *do.Injector) (*PegawaiUseCase, error) {
	return &PegawaiUseCase{
		PegawaiRepository: do.MustInvoke[*repository.PegawaiRepository](i),
	}, nil
}

func (u *PegawaiUseCase) Create(request *model.PegawaiRequest) model.PegawaiResponse {
	if valid := validation.ValidatePegawaiRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	tanggalLahir, err := time.Parse("2006-01-02", request.TanggalLahir)
	exception.PanicIfError(err)

	tanggalMasuk, err := time.Parse("2006-01-02", request.TanggalMasuk)
	exception.PanicIfError(err)

	pegawai := entity.Pegawai{
		NIP:            request.NIP,
		NIK:            request.NIK,
		Nama:           request.Nama,
		JenisKelamin:   request.JenisKelamin,
		JabatanNama:    request.JabatanNama,
		DepartemenNama: request.DepartemenNama,
		StatusKerja:    request.StatusKerja,
		Pendidikan:     request.Pendidikan,
		TempatLahir:    request.TempatLahir,
		TanggalLahir:   tanggalLahir,
		Alamat:         request.Alamat,
		AlamatLat:      request.AlamatLat,
		AlamatLon:      request.AlamatLon,
		Telepon:        request.Telepon,
		TanggalMasuk:   tanggalMasuk,
		Foto:           request.Foto,
	}

	if err := u.PegawaiRepository.Insert(&pegawai); err != nil {
		exception.PanicIfError(err)
	}

	response := model.PegawaiResponse{
		NIP:            pegawai.NIP,
		NIK:            pegawai.NIK,
		Nama:           pegawai.Nama,
		JenisKelamin:   pegawai.JenisKelamin,
		JabatanNama:    pegawai.JabatanNama,
		DepartemenNama: pegawai.DepartemenNama,
		StatusKerja:    pegawai.StatusKerja,
		Pendidikan:     pegawai.Pendidikan,
		TempatLahir:    pegawai.TempatLahir,
		TanggalLahir:   pegawai.TanggalLahir.Format("2006-01-02"),
		Alamat:         pegawai.Alamat,
		AlamatLat:      pegawai.AlamatLat,
		AlamatLon:      pegawai.AlamatLon,
		Telepon:        pegawai.Telepon,
		TanggalMasuk:   pegawai.TanggalMasuk.Format("2006-01-02"),
		Foto:           pegawai.Foto,
	}

	return response
}

func (u *PegawaiUseCase) GetAll() []model.PegawaiResponse {
	pegawai, err := u.PegawaiRepository.FindAll()
	exception.PanicIfError(err)

	response := make([]model.PegawaiResponse, len(pegawai))
	for i, pegawai := range pegawai {
		response[i] = model.PegawaiResponse{
			NIP:            pegawai.NIP,
			NIK:            pegawai.NIK,
			Nama:           pegawai.Nama,
			JenisKelamin:   pegawai.JenisKelamin,
			JabatanNama:    pegawai.JabatanNama,
			DepartemenNama: pegawai.DepartemenNama,
			StatusKerja:    pegawai.StatusKerja,
			Pendidikan:     pegawai.Pendidikan,
			TempatLahir:    pegawai.TempatLahir,
			TanggalLahir:   pegawai.TanggalLahir.Format("2006-01-02"),
			Alamat:         pegawai.Alamat,
			AlamatLat:      pegawai.AlamatLat,
			AlamatLon:      pegawai.AlamatLon,
			Telepon:        pegawai.Telepon,
			TanggalMasuk:   pegawai.TanggalMasuk.Format("2006-01-02"),
			Foto:           pegawai.Foto,
		}
	}

	return response
}

func (u *PegawaiUseCase) GetPage(page, size int) model.PegawaiPageResponse {
	pegawai, total, err := u.PegawaiRepository.FindPage(page, size)
	exception.PanicIfError(err)

	response := make([]model.PegawaiResponse, len(pegawai))
	for i, pegawai := range pegawai {
		response[i] = model.PegawaiResponse{
			NIP:            pegawai.NIP,
			NIK:            pegawai.NIK,
			Nama:           pegawai.Nama,
			JenisKelamin:   pegawai.JenisKelamin,
			JabatanNama:    pegawai.JabatanNama,
			DepartemenNama: pegawai.DepartemenNama,
			StatusKerja:    pegawai.StatusKerja,
			Pendidikan:     pegawai.Pendidikan,
			TempatLahir:    pegawai.TempatLahir,
			TanggalLahir:   pegawai.TanggalLahir.Format("2006-01-02"),
			Alamat:         pegawai.Alamat,
			AlamatLat:      pegawai.AlamatLat,
			AlamatLon:      pegawai.AlamatLon,
			Telepon:        pegawai.Telepon,
			TanggalMasuk:   pegawai.TanggalMasuk.Format("2006-01-02"),
			Foto:           pegawai.Foto,
		}
	}

	pagedResponse := model.PegawaiPageResponse{
		Pegawai: response,
		Page:    page,
		Size:    size,
		Total:   total,
	}

	return pagedResponse
}

func (u *PegawaiUseCase) GetByNIP(nip string) model.PegawaiResponse {
	pegawai, err := u.PegawaiRepository.FindByNIP(nip)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Pegawai not found",
		})
	}

	response := model.PegawaiResponse{
		NIP:            pegawai.NIP,
		NIK:            pegawai.NIK,
		Nama:           pegawai.Nama,
		JenisKelamin:   pegawai.JenisKelamin,
		JabatanNama:    pegawai.JabatanNama,
		DepartemenNama: pegawai.DepartemenNama,
		StatusKerja:    pegawai.StatusKerja,
		Pendidikan:     pegawai.Pendidikan,
		TempatLahir:    pegawai.TempatLahir,
		TanggalLahir:   pegawai.TanggalLahir.Format("2006-01-02"),
		Alamat:         pegawai.Alamat,
		AlamatLat:      pegawai.AlamatLat,
		AlamatLon:      pegawai.AlamatLon,
		Telepon:        pegawai.Telepon,
		TanggalMasuk:   pegawai.TanggalMasuk.Format("2006-01-02"),
		Foto:           pegawai.Foto,
	}

	return response
}

func (u *PegawaiUseCase) Update(nip string, request *model.PegawaiRequest) model.PegawaiResponse {
	if valid := validation.ValidatePegawaiRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	pegawai, err := u.PegawaiRepository.FindByNIP(nip)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Pegawai not found",
		})
	}

	tanggalLahir, err := time.Parse("2006-01-02", request.TanggalLahir)
	exception.PanicIfError(err)

	tanggalMasuk, err := time.Parse("2006-01-02", request.TanggalMasuk)
	exception.PanicIfError(err)

	pegawai.NIK = request.NIK
	pegawai.Nama = request.Nama
	pegawai.JenisKelamin = request.JenisKelamin
	pegawai.JabatanNama = request.JabatanNama
	pegawai.DepartemenNama = request.DepartemenNama
	pegawai.StatusKerja = request.StatusKerja
	pegawai.Pendidikan = request.Pendidikan
	pegawai.TempatLahir = request.TempatLahir
	pegawai.TanggalLahir = tanggalLahir
	pegawai.Alamat = request.Alamat
	pegawai.AlamatLat = request.AlamatLat
	pegawai.AlamatLon = request.AlamatLon
	pegawai.Telepon = request.Telepon
	pegawai.TanggalMasuk = tanggalMasuk
	pegawai.Foto = request.Foto

	if err := u.PegawaiRepository.Update(&pegawai); err != nil {
		exception.PanicIfError(err)
	}

	response := model.PegawaiResponse{
		NIP:            pegawai.NIP,
		NIK:            pegawai.NIK,
		Nama:           pegawai.Nama,
		JenisKelamin:   pegawai.JenisKelamin,
		JabatanNama:    pegawai.JabatanNama,
		DepartemenNama: pegawai.DepartemenNama,
		StatusKerja:    pegawai.StatusKerja,
		Pendidikan:     pegawai.Pendidikan,
		TempatLahir:    pegawai.TempatLahir,
		TanggalLahir:   pegawai.TanggalLahir.Format("2006-01-02"),
		Alamat:         pegawai.Alamat,
		AlamatLat:      pegawai.AlamatLat,
		AlamatLon:      pegawai.AlamatLon,
		Telepon:        pegawai.Telepon,
		TanggalMasuk:   pegawai.TanggalMasuk.Format("2006-01-02"),
		Foto:           pegawai.Foto,
	}

	return response
}

func (u *PegawaiUseCase) Delete(nip string) {
	pegawai, err := u.PegawaiRepository.FindByNIP(nip)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Pegawai not found",
		})
	}

	if err := u.PegawaiRepository.Delete(&pegawai); err != nil {
		exception.PanicIfError(err)
	}
}
