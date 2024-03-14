package route

import (
	authController "github.com/fathoor/simkes-api/internal/controller"
	middleware2 "github.com/fathoor/simkes-api/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

type Route struct {
	App                     *fiber.App
	AkunController          authController.AkunController
	AuthController          authController.AuthController
	CutiController          authController.CutiController
	DepartemenController    authController.DepartemenController
	FileController          authController.FileController
	JabatanController       authController.JabatanController
	JadwalPegawaiController authController.JadwalPegawaiController
	KehadiranController     authController.KehadiranController
	PegawaiController       authController.PegawaiController
	RoleController          authController.RoleController
	ShiftController         authController.ShiftController
}

func (r *Route) Setup() {
	akun := r.App.Group("/v1/akun", middleware2.Authenticate("Public"))
	auth := r.App.Group("/v1/auth")
	cuti := r.App.Group("/v1/cuti", middleware2.Authenticate("Pegawai"))
	departemen := r.App.Group("/v1/departemen", middleware2.Authenticate("Admin"))
	file := r.App.Group("/v1/file", middleware2.Authenticate("Public"))
	jabatan := r.App.Group("/v1/jabatan", middleware2.Authenticate("Admin"))
	jadwalPegawai := r.App.Group("/v1/jadwal-pegawai", middleware2.Authenticate("Admin"))
	kehadiran := r.App.Group("/v1/kehadiran", middleware2.Authenticate("Admin"))
	pegawai := r.App.Group("/v1/pegawai", middleware2.Authenticate("Pegawai"))
	role := r.App.Group("/v1/role", middleware2.Authenticate("Admin"))
	shift := r.App.Group("/v1/shift", middleware2.Authenticate("Admin"))

	akun.Post("/", r.AkunController.Create, middleware2.Authenticate("Admin"))
	akun.Get("/", r.AkunController.Get, middleware2.Authenticate("Pegawai"))
	akun.Get("/:nip", r.AkunController.GetByNIP, middleware2.Authenticate("Pegawai"))
	akun.Put("/:nip", r.AkunController.Update, middleware2.Authenticate("Pegawai"), middleware2.AuthorizeNIP())
	akun.Delete("/:nip", r.AkunController.Delete, middleware2.Authenticate("Admin"))

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
	pegawai.Put("/:nip", r.PegawaiController.Update, middleware2.AuthorizeNIP())
	pegawai.Delete("/:nip", r.PegawaiController.Delete, middleware2.Authenticate("Admin"))

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
