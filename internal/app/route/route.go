package route

import (
	"github.com/fathoor/simkes-api/internal/app/middleware"
	authController "github.com/fathoor/simkes-api/internal/modules/auth/controller"
	cutiController "github.com/fathoor/simkes-api/internal/modules/cuti/controller"
	fileController "github.com/fathoor/simkes-api/internal/modules/file/controller"
	jadwalPegawaiController "github.com/fathoor/simkes-api/internal/modules/jadwal-pegawai/controller"
	kehadiranController "github.com/fathoor/simkes-api/internal/modules/kehadiran/controller"
	pegawaiController "github.com/fathoor/simkes-api/internal/modules/pegawai/controller"
	"github.com/gofiber/fiber/v2"
)

type Route struct {
	App                     *fiber.App
	AkunController          controller.AkunController
	AuthController          authController.AuthController
	CutiController          cutiController.CutiController
	FileController          fileController.FileController
	JadwalPegawaiController jadwalPegawaiController.JadwalPegawaiController
	KehadiranController     kehadiranController.KehadiranController
	PegawaiController       pegawaiController.PegawaiController
}

func (r *Route) Setup() {
	akun := r.App.Group("/v1/akun", middleware.Authenticate("Public"))
	auth := r.App.Group("/v1/auth")
	cuti := r.App.Group("/v1/cuti", middleware.Authenticate("Pegawai"))
	file := r.App.Group("/v1/file", middleware.Authenticate("Public"))
	jadwalPegawai := r.App.Group("/v1/jadwal-pegawai", middleware.Authenticate("Admin"))
	kehadiran := r.App.Group("/v1/kehadiran", middleware.Authenticate("Admin"))
	pegawai := r.App.Group("/v1/pegawai", middleware.Authenticate("Pegawai"))

	akun.Post("/", r.AkunController.Create, middleware.Authenticate("Admin"))
	akun.Get("/", r.AkunController.Get, middleware.Authenticate("Pegawai"))
	akun.Get("/:nip", r.AkunController.GetByNIP, middleware.Authenticate("Pegawai"))
	akun.Put("/:nip", r.AkunController.Update, middleware.Authenticate("Pegawai"), middleware.AuthorizeNIP())
	akun.Delete("/:nip", r.AkunController.Delete, middleware.Authenticate("Admin"))

	auth.Post("/login", r.AuthController.Login)

	cuti.Post("/", r.CutiController.Create)
	cuti.Get("/", r.CutiController.Get)
	cuti.Get("/:id", r.CutiController.GetByID)
	cuti.Put("/:id", r.CutiController.Update)
	cuti.Delete("/:id", r.CutiController.Delete)

	file.Post("/", r.FileController.Upload)
	file.Get("/:filetype/:filename/download", r.FileController.Download)
	file.Get("/:filetype/:filename", r.FileController.View)
	file.Delete("/:filetype/:filename", r.FileController.Delete)

	jadwalPegawai.Post("/", r.JadwalPegawaiController.Create)
	jadwalPegawai.Get("/", r.JadwalPegawaiController.Get)
	jadwalPegawai.Get("/:tahun/:bulan/:hari/:nip", r.JadwalPegawaiController.GetByPK)
	jadwalPegawai.Put("/:tahun/:bulan/:hari/:nip", r.JadwalPegawaiController.Update)
	jadwalPegawai.Delete("/:tahun/:bulan/:hari/:nip", r.JadwalPegawaiController.Delete)

	kehadiran.Post("/checkin", r.KehadiranController.CheckIn)
	kehadiran.Post("/checkout", r.KehadiranController.CheckOut)
	kehadiran.Get("/", r.KehadiranController.Get)
	kehadiran.Get("/:id", r.KehadiranController.GetByID)
	kehadiran.Put("/:id", r.KehadiranController.Update)
	kehadiran.Delete("/:id", r.KehadiranController.Delete)

	pegawai.Post("/", r.PegawaiController.Create)
	pegawai.Get("/", r.PegawaiController.Get)
	pegawai.Get("/:nip", r.PegawaiController.GetByNIP)
	pegawai.Put("/:nip", r.PegawaiController.Update, middleware.AuthorizeNIP())
	pegawai.Delete("/:nip", r.PegawaiController.Delete, middleware.Authenticate("Admin"))
}
