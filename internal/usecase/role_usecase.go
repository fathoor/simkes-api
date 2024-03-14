package usecase

import (
	"github.com/fathoor/simkes-api/internal/entity"
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/repository"
	"github.com/fathoor/simkes-api/internal/validation"
	"github.com/samber/do"
)

type RoleUseCase struct {
	RoleRepository *repository.RoleRepository
}

func NewRoleUseCase(i *do.Injector) (*RoleUseCase, error) {
	return &RoleUseCase{
		RoleRepository: do.MustInvoke[*repository.RoleRepository](i),
	}, nil
}

func (u *RoleUseCase) Create(request *model.RoleRequest) model.RoleResponse {
	if valid := validation.ValidateRoleRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	role := entity.Role{
		Nama: request.Nama,
	}

	if err := u.RoleRepository.Insert(&role); err != nil {
		exception.PanicIfError(err)
	}

	response := model.RoleResponse{
		Nama: role.Nama,
	}

	return response
}

func (u *RoleUseCase) GetAll() []model.RoleResponse {
	roles, err := u.RoleRepository.FindAll()
	exception.PanicIfError(err)

	response := make([]model.RoleResponse, len(roles))
	for i, role := range roles {
		response[i] = model.RoleResponse{
			Nama: role.Nama,
		}
	}

	return response
}

func (u *RoleUseCase) GetByRole(r string) model.RoleResponse {
	role, err := u.RoleRepository.FindByRole(r)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Role not found",
		})
	}

	response := model.RoleResponse{
		Nama: role.Nama,
	}

	return response
}

func (u *RoleUseCase) Update(r string, request *model.RoleRequest) model.RoleResponse {
	if r == "Admin" {
		panic(exception.ForbiddenError{
			Message: "Role Admin is forbidden to be updated",
		})
	}

	if valid := validation.ValidateRoleRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	role, err := u.RoleRepository.FindByRole(r)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Role not found",
		})
	}

	role.Nama = request.Nama

	if err := u.RoleRepository.Update(&role); err != nil {
		exception.PanicIfError(err)
	}

	response := model.RoleResponse{
		Nama: role.Nama,
	}

	return response
}

func (u *RoleUseCase) Delete(r string) {
	if r == "Admin" {
		panic(exception.ForbiddenError{
			Message: "Role Admin is forbidden to be deleted",
		})
	}

	role, err := u.RoleRepository.FindByRole(r)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Role not found",
		})
	}

	if err := u.RoleRepository.Delete(&role); err != nil {
		exception.PanicIfError(err)
	}
}
