package app

import (
	"log"
	"net/http"

	"github.com/Ressorrrrra/Test-Task/internal/app/data"
	"github.com/Ressorrrrra/Test-Task/internal/app/data/order"
	"github.com/Ressorrrrra/Test-Task/internal/app/endpoint"
	service "github.com/Ressorrrrra/Test-Task/internal/app/services"
	"github.com/Ressorrrrra/Test-Task/internal/pkg/config"
)

type App struct {
	S   *service.Service
	Cfg *config.Config
	Ep  *endpoint.Endpoint
	Rep *order.Repository
}

func New(cfg *config.Config) (app *App, err error) {
	app = &App{Cfg: cfg}

	database, err := data.New(app.Cfg)
	app.Rep = order.New(database)
	app.S = service.New(app.Rep)
	app.Ep = endpoint.New(app.S)
	if err != nil {
		return nil, err
	}

	return
}

func (app *App) Run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/orders", app.Ep.GetAll)
	mux.HandleFunc("/orders/getById", app.Ep.GetById)
	mux.HandleFunc("/orders/create", app.Ep.Create)
	mux.HandleFunc("/orders/update", app.Ep.Update)
	mux.HandleFunc("/orders/delete", app.Ep.Delete)

	log.Fatal(http.ListenAndServe(":"+app.Cfg.Server.Port, mux))

	return nil
}
