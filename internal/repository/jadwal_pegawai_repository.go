package repository

import (
	"github.com/fathoor/simkes-api/internal/entity"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type JadwalPegawaiRepository struct {
	BaseRepository[entity.JadwalPegawai]
	DB *gorm.DB
}

func NewJadwalPegawaiRepository(i *do.Injector) (*JadwalPegawaiRepository, error) {
	return &JadwalPegawaiRepository{
		DB: do.MustInvoke[*gorm.DB](i),
	}, nil
}

func (r *JadwalPegawaiRepository) FindByNIP(nip string) ([]entity.JadwalPegawai, error) {
	var jadwalPegawai []entity.JadwalPegawai
	err := r.DB.Find(&jadwalPegawai, "nip = ?", nip).Error

	return jadwalPegawai, err
}

func (r *JadwalPegawaiRepository) FindByTahunBulan(tahun, bulan int16) ([]entity.JadwalPegawai, error) {
	var jadwalPegawai []entity.JadwalPegawai
	err := r.DB.Find(&jadwalPegawai, "tahun = ? AND bulan = ?", tahun, bulan).Error

	return jadwalPegawai, err
}

func (r *JadwalPegawaiRepository) FindByPK(nip string, tahun, bulan, hari int16) (entity.JadwalPegawai, error) {
	var jadwalPegawai entity.JadwalPegawai
	err := r.DB.Take(&jadwalPegawai, "nip = ? AND tahun = ? AND bulan = ? AND hari = ?", nip, tahun, bulan, hari).Error

	return jadwalPegawai, err
}
