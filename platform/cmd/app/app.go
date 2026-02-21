package app

import (
	"github.com/NikolayNam/collabsphere-go/internal/users/transport/http"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/NikolayNam/collabsphere-go/internal/config"
	appLogger "github.com/NikolayNam/collabsphere-go/internal/logger"
	mw "github.com/NikolayNam/collabsphere-go/internal/middleware"
	"github.com/NikolayNam/collabsphere-go/internal/system"
)

type App struct {
	Router chi.Router
	API    huma.API
}

func New(conf *config.Config) *App {
	router := chi.NewRouter()

	// ЕДИНЫЙ базовый логгер на всё приложение
	log := appLogger.New()

	// middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestSize(1 << 20))

	// инжектим request-scoped logger в context
	router.Use(mw.LoggerContext(log, "X-Organization-ID"))

	// единый access log (убери chi/middleware.Logger)
	router.Use(mw.AccessLog())

	// huma config
	cfg := huma.DefaultConfig(conf.APP.Title, conf.APP.Version)
	cfg.CreateHooks = nil // убирает $schema в ответах

	api := humachi.New(router, cfg)

	// регистрация модулей
	system.Register(api)
	http.Register(api)

	return &App{
		Router: router,
		API:    api,
	}
}
