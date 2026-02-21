package main

import (
	"log"

	"github.com/NikolayNam/collabsphere-go/cmd/app"
	"github.com/NikolayNam/collabsphere-go/cmd/httpserver"

	"github.com/NikolayNam/collabsphere-go/internal/config"
)

func main() {
	// 1) config (env + secrets + TZ)
	conf := config.New()

	// 2) build app (router + huma + module registration)
	application := app.New(conf)

	// 3) run http server (timeouts + graceful shutdown)
	if err := httpserver.Run(application.Router, conf.APP.Address,
		httpserver.Options{
			ReadTimeout:       conf.APP.TimeoutRead,
			WriteTimeout:      conf.APP.TimeoutWrite,
			IdleTimeout:       conf.APP.TimeoutIdle,
			ReadHeaderTimeout: 5, // seconds; либо тоже вынеси в конфиг
			ShutdownTimeout:   5, // seconds; либо тоже вынеси в конфиг
		}); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
