package repository

import (
	"github.com/fathoor/simkes-api/internal/entity"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type RoleRepository struct {
	BaseRepository[entity.Role]
	DB *gorm.DB
}

func NewRoleRepository(i *do.Injector) (*RoleRepository, error) {
	return &RoleRepository{
		DB: do.MustInvoke[*gorm.DB](i),
	}, nil
}

func (r *RoleRepository) FindByRole(nama string) (entity.Role, error) {
	var role entity.Role
	err := r.DB.Take(&role, "nama = ?", nama).Error

	return role, err
}
