package app

import (
	"github.com/Ressorrrrra/Test-Task/internal/app/data"
	"github.com/Ressorrrrra/Test-Task/internal/pkg/config"
)

type App struct {
	Db  *data.Database
	Cfg *config.Config
}

func New(cfg *config.Config) (app *App, err error) {
	app = &App{Cfg: cfg}

	return
}
