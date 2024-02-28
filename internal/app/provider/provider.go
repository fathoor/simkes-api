package provider

import (
	akunController "github.com/fathoor/simkes-api/internal/akun/controller"
	akunRepository "github.com/fathoor/simkes-api/internal/akun/repository"
	akunService "github.com/fathoor/simkes-api/internal/akun/service"
	"github.com/fathoor/simkes-api/internal/app/route"
	authController "github.com/fathoor/simkes-api/internal/auth/controller"
	authService "github.com/fathoor/simkes-api/internal/auth/service"
	departemenController "github.com/fathoor/simkes-api/internal/departemen/controller"
	departemenRepository "github.com/fathoor/simkes-api/internal/departemen/repository"
	departemenService "github.com/fathoor/simkes-api/internal/departemen/service"
	fileController "github.com/fathoor/simkes-api/internal/file/controller"
	fileService "github.com/fathoor/simkes-api/internal/file/service"
	jabatanController "github.com/fathoor/simkes-api/internal/jabatan/controller"
	jabatanRepository "github.com/fathoor/simkes-api/internal/jabatan/repository"
	jabatanService "github.com/fathoor/simkes-api/internal/jabatan/service"
	jadwalPegawaiController "github.com/fathoor/simkes-api/internal/jadwal-pegawai/controller"
	jadwalPegawaiRepository "github.com/fathoor/simkes-api/internal/jadwal-pegawai/repository"
	jadwalPegawaiService "github.com/fathoor/simkes-api/internal/jadwal-pegawai/service"
	pegawaiController "github.com/fathoor/simkes-api/internal/pegawai/controller"
	pegawaiRepository "github.com/fathoor/simkes-api/internal/pegawai/repository"
	pegawaiService "github.com/fathoor/simkes-api/internal/pegawai/service"
	roleController "github.com/fathoor/simkes-api/internal/role/controller"
	roleRepository "github.com/fathoor/simkes-api/internal/role/repository"
	roleService "github.com/fathoor/simkes-api/internal/role/service"
	shiftController "github.com/fathoor/simkes-api/internal/shift/controller"
	shiftRepository "github.com/fathoor/simkes-api/internal/shift/repository"
	shiftService "github.com/fathoor/simkes-api/internal/shift/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Provider struct {
	App *fiber.App
	DB  *gorm.DB
}

func (p *Provider) Provide() {
	repositoryAkun := akunRepository.NewAkunRepositoryProvider(p.DB)
	serviceAkun := akunService.NewAkunServiceProvider(&repositoryAkun)
	controllerAkun := akunController.NewAkunControllerProvider(&serviceAkun)

	serviceAuth := authService.NewAuthServiceProvider(&repositoryAkun)
	controllerAuth := authController.NewAuthControllerProvider(&serviceAuth)

	repositoryDepartemen := departemenRepository.NewDepartemenRepositoryProvider(p.DB)
	serviceDepartemen := departemenService.NewDepartemenServiceProvider(&repositoryDepartemen)
	controllerDepartemen := departemenController.NewDepartemenControllerProvider(&serviceDepartemen)

	serviceFile := fileService.NewFileServiceProvider()
	controllerFile := fileController.NewFileControllerProvider(&serviceFile)

	repositoryJabatan := jabatanRepository.NewJabatanRepositoryProvider(p.DB)
	serviceJabatan := jabatanService.NewJabatanServiceProvider(&repositoryJabatan)
	controllerJabatan := jabatanController.NewJabatanControllerProvider(&serviceJabatan)

	repositoryJadwalPegawai := jadwalPegawaiRepository.NewJadwalPegawaiRepositoryProvider(p.DB)
	serviceJadwalPegawai := jadwalPegawaiService.NewJadwalPegawaiServiceProvider(&repositoryJadwalPegawai)
	controllerJadwalPegawai := jadwalPegawaiController.NewJadwalPegawaiControllerProvider(&serviceJadwalPegawai)

	repositoryPegawai := pegawaiRepository.NewPegawaiRepositoryProvider(p.DB)
	servicePegawai := pegawaiService.NewPegawaiServiceProvider(&repositoryPegawai)
	controllerPegawai := pegawaiController.NewPegawaiControllerProvider(&servicePegawai)

	repositoryRole := roleRepository.NewRoleRepositoryProvider(p.DB)
	serviceRole := roleService.NewRoleServiceProvider(&repositoryRole)
	controllerRole := roleController.NewRoleControllerProvider(&serviceRole)

	repositoryShift := shiftRepository.NewShiftRepositoryProvider(p.DB)
	serviceShift := shiftService.NewShiftServiceProvider(&repositoryShift)
	controllerShift := shiftController.NewShiftControllerProvider(&serviceShift)

	router := route.Route{
		App:                     p.App,
		AkunController:          controllerAkun,
		AuthController:          controllerAuth,
		DepartemenController:    controllerDepartemen,
		FileController:          controllerFile,
		JabatanController:       controllerJabatan,
		JadwalPegawaiController: controllerJadwalPegawai,
		PegawaiController:       controllerPegawai,
		RoleController:          controllerRole,
		ShiftController:         controllerShift,
	}

	router.Setup()
}
