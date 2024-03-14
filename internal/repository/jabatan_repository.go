package repository

import (
	"github.com/fathoor/simkes-api/internal/entity"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type JabatanRepository struct {
	BaseRepository[entity.Jabatan]
	DB *gorm.DB
}

func NewJabatanRepository(i *do.Injector) (*JabatanRepository, error) {
	return &JabatanRepository{
		DB: do.MustInvoke[*gorm.DB](i),
	}, nil
}

func (r *JabatanRepository) FindByJabatan(nama string) (entity.Jabatan, error) {
	var jabatan entity.Jabatan
	err := r.DB.Take(&jabatan, "nama = ?", nama).Error

	return jabatan, err
}
