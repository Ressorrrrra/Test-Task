package main

import (
	"log"

	"github.com/Ressorrrrra/Test-Task/internal/app"
	"github.com/Ressorrrrra/Test-Task/internal/pkg/config"
)

func main() {
	cfg, err := config.ConfigureFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	app, err := app.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	app.Run()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(app.Cfg.Db.Port)
}
