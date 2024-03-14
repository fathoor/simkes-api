package provider

import (
	authController "github.com/fathoor/simkes-api/internal/controller"
	akunRepository "github.com/fathoor/simkes-api/internal/repository"
	"github.com/fathoor/simkes-api/internal/route"
	authService "github.com/fathoor/simkes-api/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Provider struct {
	App *fiber.App
	DB  *gorm.DB
}

func (p *Provider) Provide() {
	repositoryAkun := akunRepository.NewAkunRepositoryProvider(p.DB)
	serviceAkun := authService.NewAkunServiceProvider(&repositoryAkun)
	controllerAkun := authController.NewAkunControllerProvider(&serviceAkun)

	serviceAuth := authService.NewAuthServiceProvider(&repositoryAkun)
	controllerAuth := authController.NewAuthControllerProvider(&serviceAuth)

	repositoryCuti := akunRepository.NewCutiRepositoryProvider(p.DB)
	serviceCuti := authService.NewCutiServiceProvider(&repositoryCuti)
	controllerCuti := authController.NewCutiControllerProvider(&serviceCuti)

	repositoryDepartemen := akunRepository.NewDepartemenRepositoryProvider(p.DB)
	serviceDepartemen := authService.NewDepartemenServiceProvider(&repositoryDepartemen)
	controllerDepartemen := authController.NewDepartemenControllerProvider(&serviceDepartemen)

	serviceFile := authService.NewFileServiceProvider()
	controllerFile := authController.NewFileControllerProvider(&serviceFile)

	repositoryJabatan := akunRepository.NewJabatanRepositoryProvider(p.DB)
	serviceJabatan := authService.NewJabatanServiceProvider(&repositoryJabatan)
	controllerJabatan := authController.NewJabatanControllerProvider(&serviceJabatan)

	repositoryJadwalPegawai := akunRepository.NewJadwalPegawaiRepositoryProvider(p.DB)
	serviceJadwalPegawai := authService.NewJadwalPegawaiServiceProvider(&repositoryJadwalPegawai)
	controllerJadwalPegawai := authController.NewJadwalPegawaiControllerProvider(&serviceJadwalPegawai)

	repositoryShift := akunRepository.NewShiftRepositoryProvider(p.DB)
	serviceShift := authService.NewShiftServiceProvider(&repositoryShift)
	controllerShift := authController.NewShiftControllerProvider(&serviceShift)

	repositoryKehadiran := akunRepository.NewKehadiranRepositoryProvider(p.DB)
	serviceKehadiran := authService.NewKehadiranServiceProvider(&repositoryKehadiran, &repositoryShift)
	controllerKehadiran := authController.NewKehadiranControllerProvider(&serviceKehadiran)

	repositoryPegawai := akunRepository.NewPegawaiRepositoryProvider(p.DB)
	servicePegawai := authService.NewPegawaiServiceProvider(&repositoryPegawai)
	controllerPegawai := authController.NewPegawaiControllerProvider(&servicePegawai)

	repositoryRole := akunRepository.NewRoleRepositoryProvider(p.DB)
	serviceRole := authService.NewRoleServiceProvider(&repositoryRole)
	controllerRole := authController.NewRoleControllerProvider(&serviceRole)

	router := route.Route{
		App:                     p.App,
		AkunController:          controllerAkun,
		AuthController:          controllerAuth,
		CutiController:          controllerCuti,
		DepartemenController:    controllerDepartemen,
		FileController:          controllerFile,
		JabatanController:       controllerJabatan,
		JadwalPegawaiController: controllerJadwalPegawai,
		KehadiranController:     controllerKehadiran,
		PegawaiController:       controllerPegawai,
		RoleController:          controllerRole,
		ShiftController:         controllerShift,
	}

	router.Setup()
}
