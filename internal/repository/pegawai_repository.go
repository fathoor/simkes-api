package repository

import (
	"github.com/fathoor/simkes-api/internal/entity"
	"github.com/samber/do"
	"gorm.io/gorm"
	"math"
)

type PegawaiRepository struct {
	BaseRepository[entity.Pegawai]
	DB *gorm.DB
}

func NewPegawaiRepository(i *do.Injector) (*PegawaiRepository, error) {
	return &PegawaiRepository{
		DB: do.MustInvoke[*gorm.DB](i),
	}, nil
}

func (r *PegawaiRepository) FindPage(page, size int) ([]entity.Pegawai, int, error) {
	var pegawai []entity.Pegawai
	var total int64

	if err := r.DB.Model(&entity.Pegawai{}).Count(&total).Error; err != nil {
		return pegawai, 0, err
	}

	totalPage := int(math.Ceil(float64(total) / float64(size)))

	err := r.DB.Limit(size).Offset((page - 1) * size).Find(&pegawai).Error

	return pegawai, totalPage, err
}

func (r *PegawaiRepository) FindByNIP(nip string) (entity.Pegawai, error) {
	var pegawai entity.Pegawai
	err := r.DB.Take(&pegawai, "nip = ?", nip).Error

	return pegawai, err
}
