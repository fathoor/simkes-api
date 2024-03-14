package usecase

import (
	"github.com/fathoor/simkes-api/internal/entity"
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/repository"
	"github.com/fathoor/simkes-api/internal/validation"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"github.com/samber/do"
)

type RoleUseCase struct {
	RoleRepository *repository.RoleRepository
	Log            *zerolog.Logger
	Validator      *validator.Validate
}

func NewRoleUseCase(i *do.Injector) (*RoleUseCase, error) {
	return &RoleUseCase{
		RoleRepository: do.MustInvoke[*repository.RoleRepository](i),
		Log:            do.MustInvoke[*zerolog.Logger](i),
		Validator:      do.MustInvoke[*validator.Validate](i),
	}, nil
}

func (u *RoleUseCase) Create(request *model.RoleRequest) model.RoleResponse {
	validation.ValidateRoleRequest(u.Validator, u.Log, request)

	role := entity.Role{
		Nama: request.Nama,
	}

	if err := u.RoleRepository.Insert(&role); err != nil {
		u.Log.Error().Err(err).Msg("Failed to insert role")
		panic(exception.InternalServerError{
			Message: "Failed to insert role",
		})
	}

	response := model.RoleResponse{
		Nama: role.Nama,
	}

	return response
}

func (u *RoleUseCase) GetAll() []model.RoleResponse {
	roles, err := u.RoleRepository.FindAll()
	if err != nil {
		u.Log.Error().Err(err).Msg("Failed to get roles")
		panic(exception.InternalServerError{
			Message: "Failed to get roles",
		})
	}

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
		u.Log.Info().Str("role", r).Msg("Role not found")
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
		u.Log.Info().Str("role", r).Msg("Forbidden role")
		panic(exception.ForbiddenError{
			Message: "Role Admin is forbidden to be updated",
		})
	}

	validation.ValidateRoleRequest(u.Validator, u.Log, request)

	role, err := u.RoleRepository.FindByRole(r)
	if err != nil {
		u.Log.Info().Str("role", r).Msg("Role not found")
		panic(exception.NotFoundError{
			Message: "Role not found",
		})
	}

	role.Nama = request.Nama

	if err := u.RoleRepository.Update(&role); err != nil {
		u.Log.Error().Err(err).Msg("Failed to update role")
		panic(exception.InternalServerError{
			Message: "Failed to update role",
		})
	}

	response := model.RoleResponse{
		Nama: role.Nama,
	}

	return response
}

func (u *RoleUseCase) Delete(r string) {
	if r == "Admin" {
		u.Log.Info().Str("role", r).Msg("Forbidden role")
		panic(exception.ForbiddenError{
			Message: "Role Admin is forbidden to be deleted",
		})
	}

	role, err := u.RoleRepository.FindByRole(r)
	if err != nil {
		u.Log.Info().Str("role", r).Msg("Role not found")
		panic(exception.NotFoundError{
			Message: "Role not found",
		})
	}

	if err := u.RoleRepository.Delete(&role); err != nil {
		u.Log.Error().Err(err).Msg("Failed to delete role")
		panic(exception.InternalServerError{
			Message: "Failed to delete role",
		})
	}
}
