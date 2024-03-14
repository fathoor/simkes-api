package main

import (
	"fmt"
	"github.com/fathoor/simkes-api/internal/app/provider"
	config2 "github.com/fathoor/simkes-api/internal/config"
	"github.com/fathoor/simkes-api/internal/exception"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var (
		cfg = config2.ProvideConfig()
		app = config2.ProvideApp(cfg)
		db  = config2.ProvideDB(cfg)
		di  = provider.Provider{App: app, DB: db}
	)

	di.Provide()

	err := app.Listen(fmt.Sprintf(":%d", cfg.GetInt("APP_PORT")))
	exception.PanicIfError(err)
}
