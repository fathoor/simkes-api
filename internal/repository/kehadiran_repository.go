package repository

import (
	"github.com/fathoor/simkes-api/internal/entity"
	"github.com/google/uuid"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type KehadiranRepository struct {
	BaseRepository[entity.Kehadiran]
	DB *gorm.DB
}

func NewKehadiranRepository(i *do.Injector) (*KehadiranRepository, error) {
	return &KehadiranRepository{
		DB: do.MustInvoke[*gorm.DB](i),
	}, nil
}

func (r *KehadiranRepository) FindByNIP(nip string) ([]entity.Kehadiran, error) {
	var kehadiran []entity.Kehadiran
	err := r.DB.Preload("Shift").Find(&kehadiran, "nip = ?", nip).Error

	return kehadiran, err
}

func (r *KehadiranRepository) FindByID(id uuid.UUID) (entity.Kehadiran, error) {
	var kehadiran entity.Kehadiran
	err := r.DB.Preload("Shift").Take(&kehadiran, "id = ?", id).Error

	return kehadiran, err
}

func (r *KehadiranRepository) FindLatestByNIP(nip string) (entity.Kehadiran, error) {
	var kehadiran entity.Kehadiran
	err := r.DB.Preload("Shift").Order("tanggal desc").Take(&kehadiran, "nip = ?", nip).Error

	return kehadiran, err
}
