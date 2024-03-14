package config

import (
	"github.com/fathoor/simkes-api/internal/controller"
	"github.com/fathoor/simkes-api/internal/repository"
	"github.com/fathoor/simkes-api/internal/router"
	"github.com/fathoor/simkes-api/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func NewContainer() *do.Injector {
	injector := do.New()

	// Config
	do.Provide[*zerolog.Logger](injector, func(i *do.Injector) (*zerolog.Logger, error) {
		return NewLogger(), nil
	})
	do.Provide[*Config](injector, NewConfig)
	cfg := do.MustInvoke[*Config](injector)
	log := do.MustInvoke[*zerolog.Logger](injector)

	// App
	do.Provide[*fiber.App](injector, func(i *do.Injector) (*fiber.App, error) {
		return NewApp(cfg), nil
	})
	do.Provide[*gorm.DB](injector, func(i *do.Injector) (*gorm.DB, error) {
		return NewDB(cfg, log), nil
	})
	do.Provide[*validator.Validate](injector, func(i *do.Injector) (*validator.Validate, error) {
		return NewValidator(), nil
	})

	// Repository
	do.Provide[*repository.AkunRepository](injector, repository.NewAkunRepository)
	do.Provide[*repository.CutiRepository](injector, repository.NewCutiRepository)
	do.Provide[*repository.DepartemenRepository](injector, repository.NewDepartemenRepository)
	do.Provide[*repository.JabatanRepository](injector, repository.NewJabatanRepository)
	do.Provide[*repository.JadwalPegawaiRepository](injector, repository.NewJadwalPegawaiRepository)
	do.Provide[*repository.KehadiranRepository](injector, repository.NewKehadiranRepository)
	do.Provide[*repository.PegawaiRepository](injector, repository.NewPegawaiRepository)
	do.Provide[*repository.RoleRepository](injector, repository.NewRoleRepository)
	do.Provide[*repository.ShiftRepository](injector, repository.NewShiftRepository)

	// UseCase
	do.Provide[*usecase.AkunUseCase](injector, usecase.NewAkunUseCase)
	do.Provide[*usecase.AuthUseCase](injector, usecase.NewAuthUseCase)
	do.Provide[*usecase.CutiUseCase](injector, usecase.NewCutiUseCase)
	do.Provide[*usecase.DepartemenUseCase](injector, usecase.NewDepartemenUseCase)
	do.Provide[*usecase.FileUseCase](injector, usecase.NewFileUseCase)
	do.Provide[*usecase.JabatanUseCase](injector, usecase.NewJabatanUseCase)
	do.Provide[*usecase.JadwalPegawaiUseCase](injector, usecase.NewJadwalPegawaiUseCase)
	do.Provide[*usecase.KehadiranUseCase](injector, usecase.NewKehadiranUseCase)
	do.Provide[*usecase.PegawaiUseCase](injector, usecase.NewPegawaiUseCase)
	do.Provide[*usecase.RoleUseCase](injector, usecase.NewRoleUseCase)
	do.Provide[*usecase.ShiftUseCase](injector, usecase.NewShiftUseCase)

	// Controller
	do.Provide[*controller.AkunController](injector, controller.NewAkunController)
	do.Provide[*controller.AuthController](injector, controller.NewAuthController)
	do.Provide[*controller.CutiController](injector, controller.NewCutiController)
	do.Provide[*controller.DepartemenController](injector, controller.NewDepartemenController)
	do.Provide[*controller.FileController](injector, controller.NewFileController)
	do.Provide[*controller.JabatanController](injector, controller.NewJabatanController)
	do.Provide[*controller.JadwalPegawaiController](injector, controller.NewJadwalPegawaiController)
	do.Provide[*controller.KehadiranController](injector, controller.NewKehadiranController)
	do.Provide[*controller.PegawaiController](injector, controller.NewPegawaiController)
	do.Provide[*controller.RoleController](injector, controller.NewRoleController)
	do.Provide[*controller.ShiftController](injector, controller.NewShiftController)

	// Router
	do.Provide[*router.Router](injector, router.NewRouter)

	return injector
}
