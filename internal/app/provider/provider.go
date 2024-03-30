package provider

import (
	"github.com/fathoor/simkes-api/internal/app/route"
	authController "github.com/fathoor/simkes-api/internal/modules/auth/controller"
	authService "github.com/fathoor/simkes-api/internal/modules/auth/service"
	cutiController "github.com/fathoor/simkes-api/internal/modules/cuti/controller"
	cutiRepository "github.com/fathoor/simkes-api/internal/modules/cuti/repository"
	cutiService "github.com/fathoor/simkes-api/internal/modules/cuti/service"
	fileController "github.com/fathoor/simkes-api/internal/modules/file/controller"
	fileService "github.com/fathoor/simkes-api/internal/modules/file/service"
	jadwalPegawaiController "github.com/fathoor/simkes-api/internal/modules/jadwal-pegawai/controller"
	jadwalPegawaiRepository "github.com/fathoor/simkes-api/internal/modules/jadwal-pegawai/repository"
	jadwalPegawaiService "github.com/fathoor/simkes-api/internal/modules/jadwal-pegawai/service"
	kehadiranController "github.com/fathoor/simkes-api/internal/modules/kehadiran/controller"
	kehadiranRepository "github.com/fathoor/simkes-api/internal/modules/kehadiran/repository"
	kehadiranService "github.com/fathoor/simkes-api/internal/modules/kehadiran/service"
	pegawaiController "github.com/fathoor/simkes-api/internal/modules/pegawai/controller"
	pegawaiRepository "github.com/fathoor/simkes-api/internal/modules/pegawai/repository"
	pegawaiService "github.com/fathoor/simkes-api/internal/modules/pegawai/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Provider struct {
	App *fiber.App
	DB  *gorm.DB
}

func (p *Provider) Provide() {
	repositoryAkun := repository.NewAkunRepositoryProvider(p.DB)
	serviceAkun := usecase.NewAkunServiceProvider(&repositoryAkun)
	controllerAkun := controller.NewAkunControllerProvider(&serviceAkun)

	serviceAuth := authService.NewAuthServiceProvider(&repositoryAkun)
	controllerAuth := authController.NewAuthControllerProvider(&serviceAuth)

	repositoryCuti := cutiRepository.NewCutiRepositoryProvider(p.DB)
	serviceCuti := cutiService.NewCutiServiceProvider(&repositoryCuti)
	controllerCuti := cutiController.NewCutiControllerProvider(&serviceCuti)

	serviceFile := fileService.NewFileServiceProvider()
	controllerFile := fileController.NewFileControllerProvider(&serviceFile)

	repositoryJadwalPegawai := jadwalPegawaiRepository.NewJadwalPegawaiRepositoryProvider(p.DB)
	serviceJadwalPegawai := jadwalPegawaiService.NewJadwalPegawaiServiceProvider(&repositoryJadwalPegawai)
	controllerJadwalPegawai := jadwalPegawaiController.NewJadwalPegawaiControllerProvider(&serviceJadwalPegawai)

	repositoryKehadiran := kehadiranRepository.NewKehadiranRepositoryProvider(p.DB)
	serviceKehadiran := kehadiranService.NewKehadiranServiceProvider(&repositoryKehadiran)
	controllerKehadiran := kehadiranController.NewKehadiranControllerProvider(&serviceKehadiran)

	repositoryPegawai := pegawaiRepository.NewPegawaiRepositoryProvider(p.DB)
	servicePegawai := pegawaiService.NewPegawaiServiceProvider(&repositoryPegawai)
	controllerPegawai := pegawaiController.NewPegawaiControllerProvider(&servicePegawai)

	router := route.Route{
		App:                     p.App,
		AkunController:          controllerAkun,
		AuthController:          controllerAuth,
		CutiController:          controllerCuti,
		FileController:          controllerFile,
		JadwalPegawaiController: controllerJadwalPegawai,
		KehadiranController:     controllerKehadiran,
		PegawaiController:       controllerPegawai,
	}

	router.Setup()
}
