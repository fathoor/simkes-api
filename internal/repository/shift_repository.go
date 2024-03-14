package repository

import (
	"github.com/fathoor/simkes-api/internal/entity"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type ShiftRepository struct {
	BaseRepository[entity.Shift]
	DB *gorm.DB
}

func NewShiftRepository(i *do.Injector) (*ShiftRepository, error) {
	return &ShiftRepository{
		DB: do.MustInvoke[*gorm.DB](i),
	}, nil
}

func (r *ShiftRepository) FindByNama(nama string) (entity.Shift, error) {
	var shift entity.Shift
	err := r.DB.Take(&shift, "nama = ?", nama).Error

	return shift, err
}
