package usecase

import (
	"github.com/fathoor/simkes-api/internal/model"
)

type RoleService interface {
	Create(request *model.RoleRequest) model.RoleResponse
	GetAll() []model.RoleResponse
	GetByRole(r string) model.RoleResponse
	Update(r string, request *model.RoleRequest) model.RoleResponse
	Delete(r string)
}
