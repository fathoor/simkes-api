package repository

import (
	"github.com/fathoor/simkes-api/internal/entity"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type DepartemenRepository struct {
	BaseRepository[entity.Departemen]
	DB *gorm.DB
}

func NewDepartemenRepository(i *do.Injector) (*DepartemenRepository, error) {
	return &DepartemenRepository{
		DB: do.MustInvoke[*gorm.DB](i),
	}, nil
}

func (r *DepartemenRepository) FindByDepartemen(nama string) (entity.Departemen, error) {
	var departemen entity.Departemen
	err := r.DB.Take(&departemen, "nama = ?", nama).Error

	return departemen, err
}
