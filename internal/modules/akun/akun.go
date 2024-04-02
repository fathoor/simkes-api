package akun

import (
	"github.com/fathoor/simkes-api/internal/app/config"
	"github.com/fathoor/simkes-api/internal/modules/akun/internal/controller"
	"github.com/fathoor/simkes-api/internal/modules/akun/internal/repository/postgres"
	"github.com/fathoor/simkes-api/internal/modules/akun/internal/router"
	"github.com/fathoor/simkes-api/internal/modules/akun/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func ProvideAkun(app *fiber.App, cfg *config.Config, db *sqlx.DB, validator *config.Validator) {
	akunRepository := postgres.NewAkunRepository(db)
	akunUseCase := usecase.NewAkunUseCase(&akunRepository, cfg)
	akunController := controller.NewAkunController(akunUseCase, validator)

	alamatRepository := postgres.NewAlamatRepository(db)
	alamatUseCase := usecase.NewAlamatUseCase(&alamatRepository)
	alamatController := controller.NewAlamatController(alamatUseCase, validator)

	router.Route(
		app,
		akunController,
		alamatController,
	)
}
