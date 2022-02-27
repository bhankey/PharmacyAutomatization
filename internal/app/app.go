package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bhankey/BD_lab/backend/internal/app/container"
	configinternal "github.com/bhankey/BD_lab/backend/internal/config"
	httphandler "github.com/bhankey/BD_lab/backend/internal/delivery/http"
	"github.com/bhankey/BD_lab/backend/internal/delivery/http/swaggerhandler"
	"github.com/bhankey/BD_lab/backend/pkg/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type App struct {
	server    *http.Server
	container *container.Container
	logger    logger.Logger
}

const shutDownTimeoutSeconds = 10

func NewApp(configPath string) (*App, error) {
	config, err := configinternal.GetConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to init app because of config error: %w", err)
	}

	log, err := logger.GetLogger(config.Logger.Path, config.Logger.Level, true)
	if err != nil {
		return nil, fmt.Errorf("failed to init logger error: %w", err)
	}

	log.Info("try to init data source resource")
	// TODO remove dataSource struct
	dataSources, err := newDataSource(config)
	if err != nil {
		return nil, err
	}

	dependencies := container.NewContainer(log, dataSources.db, dataSources.db, dataSources.redisClient)

	baseHandler := httphandler.NewHandler(log)

	swaggerHandler := swaggerhandler.NewSwaggerHandler(baseHandler)

	// TODO move to different package or function
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
	}))

	router.Use(func(handler http.Handler) http.Handler {
		return httphandler.LoggingMiddleware(log)(handler)
	})

	router.Mount("/docs", swaggerHandler.Router)

	server := &http.Server{
		Addr:    ":" + config.Server.Port,
		Handler: router,
	}

	return &App{
		logger:    log,
		server:    server,
		container: dependencies,
	}, nil
}

func (a *App) Start() {
	a.logger.Info("staring server on port: " + a.server.Addr)
	go func() {
		if err := a.server.ListenAndServe(); err != nil {
			a.logger.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	a.logger.Info("received signal to shutdown server")
	ctx, cancel := context.WithTimeout(context.Background(), shutDownTimeoutSeconds*time.Second)
	defer cancel()
	if err := a.server.Shutdown(ctx); err != nil {
		a.logger.Error(err)
	}

	<-ctx.Done()

	a.container.CloseAllConnections()

	a.logger.Info("server was shutdown")
}
