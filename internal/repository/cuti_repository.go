package repository

import (
	"github.com/fathoor/simkes-api/internal/entity"
	"github.com/google/uuid"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type CutiRepository struct {
	BaseRepository[entity.Cuti]
	DB *gorm.DB
}

func NewCutiRepository(i *do.Injector) (*CutiRepository, error) {
	return &CutiRepository{
		DB: do.MustInvoke[*gorm.DB](i),
	}, nil
}

func (r *CutiRepository) FindByNIP(nip string) ([]entity.Cuti, error) {
	var cuti []entity.Cuti
	err := r.DB.Preload("Pegawai").Find(&cuti, "nip = ?", nip).Error

	return cuti, err
}

func (r *CutiRepository) FindByID(id uuid.UUID) (entity.Cuti, error) {
	var cuti entity.Cuti
	err := r.DB.Preload("Pegawai").Take(&cuti, "id = ?", id).Error

	return cuti, err
}
