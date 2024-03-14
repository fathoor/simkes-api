package router

import (
	"github.com/fathoor/simkes-api/internal/controller"
	"github.com/fathoor/simkes-api/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
)

type Router struct {
	App                     *fiber.App
	AkunController          *controller.AkunController
	AuthController          *controller.AuthController
	CutiController          *controller.CutiController
	DepartemenController    *controller.DepartemenController
	FileController          *controller.FileController
	JabatanController       *controller.JabatanController
	JadwalPegawaiController *controller.JadwalPegawaiController
	KehadiranController     *controller.KehadiranController
	PegawaiController       *controller.PegawaiController
	RoleController          *controller.RoleController
	ShiftController         *controller.ShiftController
}

func NewRouter(i *do.Injector) (*Router, error) {
	return &Router{
		App:                     do.MustInvoke[*fiber.App](i),
		AkunController:          do.MustInvoke[*controller.AkunController](i),
		AuthController:          do.MustInvoke[*controller.AuthController](i),
		CutiController:          do.MustInvoke[*controller.CutiController](i),
		DepartemenController:    do.MustInvoke[*controller.DepartemenController](i),
		FileController:          do.MustInvoke[*controller.FileController](i),
		JabatanController:       do.MustInvoke[*controller.JabatanController](i),
		JadwalPegawaiController: do.MustInvoke[*controller.JadwalPegawaiController](i),
		KehadiranController:     do.MustInvoke[*controller.KehadiranController](i),
		PegawaiController:       do.MustInvoke[*controller.PegawaiController](i),
		RoleController:          do.MustInvoke[*controller.RoleController](i),
		ShiftController:         do.MustInvoke[*controller.ShiftController](i),
	}, nil
}

func (r *Router) Route() {
	akun := r.App.Group("/v1/akun", middleware.Authenticate("Public"))
	auth := r.App.Group("/v1/auth")
	cuti := r.App.Group("/v1/cuti", middleware.Authenticate("Pegawai"))
	departemen := r.App.Group("/v1/departemen", middleware.Authenticate("Admin"))
	file := r.App.Group("/v1/file", middleware.Authenticate("Public"))
	jabatan := r.App.Group("/v1/jabatan", middleware.Authenticate("Admin"))
	jadwalPegawai := r.App.Group("/v1/jadwal-pegawai", middleware.Authenticate("Admin"))
	kehadiran := r.App.Group("/v1/kehadiran", middleware.Authenticate("Admin"))
	pegawai := r.App.Group("/v1/pegawai", middleware.Authenticate("Pegawai"))
	role := r.App.Group("/v1/role", middleware.Authenticate("Admin"))
	shift := r.App.Group("/v1/shift", middleware.Authenticate("Admin"))

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

	departemen.Post("/", r.DepartemenController.Create)
	departemen.Get("/", r.DepartemenController.Get)
	departemen.Get("/:departemen", r.DepartemenController.GetByNama)
	departemen.Put("/:departemen", r.DepartemenController.Update)
	departemen.Delete("/:departemen", r.DepartemenController.Delete)

	file.Post("/", r.FileController.Upload)
	file.Get("/:filetype/:filename/download", r.FileController.Download)
	file.Get("/:filetype/:filename", r.FileController.View)
	file.Delete("/:filetype/:filename", r.FileController.Delete)

	jabatan.Post("/", r.JabatanController.Create)
	jabatan.Get("/", r.JabatanController.Get)
	jabatan.Get("/:jabatan", r.JabatanController.GetByNama)
	jabatan.Put("/:jabatan", r.JabatanController.Update)
	jabatan.Delete("/:jabatan", r.JabatanController.Delete)

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

	role.Post("/", r.RoleController.Create)
	role.Get("/", r.RoleController.Get)
	role.Get("/:role", r.RoleController.GetByNama)
	role.Put("/:role", r.RoleController.Update)
	role.Delete("/:role", r.RoleController.Delete)

	shift.Post("/", r.ShiftController.Create)
	shift.Get("/", r.ShiftController.Get)
	shift.Get("/:shift", r.ShiftController.GetByNama)
	shift.Put("/:shift", r.ShiftController.Update)
}
