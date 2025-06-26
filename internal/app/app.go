package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"golang.org/x/sync/errgroup"
	"net"
	"net/http"
	"ngMarketplace/config"
	"ngMarketplace/internal/category"
	"ngMarketplace/internal/transport/http/router"
	"ngMarketplace/pkg/logger"
	"ngMarketplace/pkg/postgres"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type App struct {
	cfg        *config.Config
	wg         sync.WaitGroup
	router     *gin.Engine
	httpServer *http.Server
	logger     logger.Logger
	pg         *postgres.Postgres
}

// New collects everything needed to start the app
func New(cfg *config.Config) (*App, error) {
	var a = &App{
		wg: sync.WaitGroup{},
	}

	l := logger.New(logger.WithLevel(cfg.Log.Level), logger.WithIsJSON(true), logger.WithAddSource(true))
	l.Debug("logger initialized")

	l.Debug("Configurations: %v", cfg)

	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Sprintf("app - Run - postgres.New: %v", err))
	}
	l.Debug("PostgreSQL initialized")

	//runner := async.NewBackgroundRunner(&a.wg)

	categoryRepo := category.NewRepository(*pg)
	categoryUseCase := category.NewUseCase(*categoryRepo)
	categoryHandler := category.NewHandler(categoryUseCase, l)

	router := router.NewRouter()
	categoryHandler.Register(router)

	a.cfg = cfg
	a.router = router
	a.logger = l
	a.pg = pg

	return a, nil
}

// Run start the application either by http, grpc or etc.
func (a *App) Run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	grp, ctx := errgroup.WithContext(ctx)

	grp.Go(func() error {
		return a.startHTTP(ctx)
	})

	err := grp.Wait()
	switch {
	case err == nil || errors.Is(err, context.Canceled):
		a.logger.Info("Everything is good. Server shutting down...")
	case errors.Is(err, http.ErrServerClosed):
		a.logger.Info("server shutdown")
	default:
		a.logger.Error(fmt.Sprintf("app.Run: %v", err))
		return err
	}

	a.logger.Info("waiting for background tasks to finish...")
	a.wg.Wait()

	if a.pg != nil {
		a.pg.Close()
		a.logger.Info("Postgres connection closed")
	}

	return nil
}

func (a *App) startHTTP(ctx context.Context) error {
	const op = "startHTTP"

	a.logger.Info("HTTP Server initializing")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", "0.0.0.0", a.cfg.HTTP.Port))
	if err != nil {
		a.logger.Fatal("%s: failed to create listener: ", op, err)
		return err
	}

	c := cors.New(cors.Options{
		AllowedOrigins:     a.cfg.HTTP.CORS.AllowedOrigins,
		AllowedMethods:     a.cfg.HTTP.CORS.AllowedMethods,
		AllowedHeaders:     a.cfg.HTTP.CORS.AllowedHeaders,
		ExposedHeaders:     a.cfg.HTTP.CORS.ExposedHeaders,
		AllowCredentials:   a.cfg.HTTP.CORS.AllowCredentials,
		OptionsPassthrough: a.cfg.HTTP.CORS.OptionsPassthrough,
		Debug:              a.cfg.HTTP.CORS.Debug,
		Logger:             a.logger,
	})

	handler := c.Handler(a.router)

	a.httpServer = &http.Server{
		Handler:      handler,
		WriteTimeout: a.cfg.HTTP.WriteTimeout,
		ReadTimeout:  a.cfg.HTTP.ReadTimeout,
	}

	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		a.logger.Info("Shutting down HTTP server...")
		err = a.httpServer.Shutdown(shutdownCtx)
		if err != nil {
			a.logger.Error("failed to shutdown server: %v", err)
		}
	}()

	if err = a.httpServer.Serve(listener); err != nil {
		switch {
		case err == nil:
			a.logger.Info("server exited with no error")
			return nil
		case errors.Is(err, http.ErrServerClosed):
			a.logger.Info("server shutdown")
		default:
			a.logger.Fatal("failed to start server")
		}
	}

	return err
}
