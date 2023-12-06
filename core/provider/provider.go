package provider

import (
	akunController "github.com/fathoor/simkes-api/module/akun/akun/controller"
	akunRepository "github.com/fathoor/simkes-api/module/akun/akun/repository"
	akunService "github.com/fathoor/simkes-api/module/akun/akun/service"
	roleController "github.com/fathoor/simkes-api/module/akun/role/controller"
	roleRepository "github.com/fathoor/simkes-api/module/akun/role/repository"
	roleService "github.com/fathoor/simkes-api/module/akun/role/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ProvideModule(app *fiber.App, db *gorm.DB) {
	ProvideRole(app, db)
	ProvideAkun(app, db)
}

func ProvideRole(app *fiber.App, db *gorm.DB) {
	repository := roleRepository.ProvideRoleRepository(db)
	service := roleService.ProvideRoleService(&repository)
	controller := roleController.ProvideRoleController(&service)

	controller.Route(app)
}

func ProvideAkun(app *fiber.App, db *gorm.DB) {
	repository := akunRepository.ProvideAkunRepository(db)
	service := akunService.ProvideAkunService(&repository)
	controller := akunController.ProvideAkunController(&service)

	controller.Route(app)
}