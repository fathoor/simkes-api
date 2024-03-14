package repository

import (
	"github.com/fathoor/simkes-api/internal/entity"
	"github.com/samber/do"
	"gorm.io/gorm"
	"math"
)

type AkunRepository struct {
	BaseRepository[entity.Akun]
	DB *gorm.DB
}

func NewAkunRepository(i *do.Injector) (*AkunRepository, error) {
	return &AkunRepository{
		DB: do.MustInvoke[*gorm.DB](i),
	}, nil
}

func (r *AkunRepository) FindPage(page, size int) ([]entity.Akun, int, error) {
	var akun []entity.Akun
	var total int64

	if err := r.DB.Model(&entity.Akun{}).Count(&total).Error; err != nil {
		return akun, 0, err
	}

	totalPage := int(math.Ceil(float64(total) / float64(size)))

	err := r.DB.Limit(size).Offset((page - 1) * size).Find(&akun).Error

	return akun, totalPage, err
}

func (r *AkunRepository) FindByNIP(nip string) (entity.Akun, error) {
	var akun entity.Akun
	err := r.DB.Take(&akun, "nip = ?", nip).Error

	return akun, err
}
