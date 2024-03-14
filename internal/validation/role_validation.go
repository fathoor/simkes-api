package validation

import (
	"github.com/fathoor/simkes-api/internal/config"
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/model"
)

func ValidateRoleRequest(request *model.RoleRequest) error {
	if request.Nama == "Admin" {
		panic(exception.ForbiddenError{
			Message: "You are not allowed to modify this role!",
		})
	}

	return config.Validator.Struct(request)
}
