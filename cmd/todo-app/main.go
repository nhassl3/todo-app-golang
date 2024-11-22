package main

import (
	"context"
	"errors"
	"github.com/joho/godotenv"
	"github.com/nhassl3/todo-app/pkg/config"
	"github.com/nhassl3/todo-app/pkg/http-server/handlers"
	"github.com/nhassl3/todo-app/pkg/logger/handlers/slogpretty"
	"github.com/nhassl3/todo-app/pkg/repository"
	"github.com/nhassl3/todo-app/pkg/service"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// load dot environment
	if err := godotenv.Load(); err != nil {
		slog.Error("error loading .env variables", slog.String("error", err.Error()))
	}

	// setting configuration of the project
	cfg := config.MustLoad()

	// set up a logger
	log := setupLogger(cfg.Env)
	log.Info("starting todo-app", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	// PostgresDB
	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		log.Error("error connect to database", slog.String("error", err.Error()))
		os.Exit(1)
	}

	// set up the repository, service and handlers
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handlers.NewHandler(services, log)

	// set up http server on localhost:8082
	server := new(Server)
	go func() {
		if err = server.Run(cfg, handler.InitRoutes()); err != nil && !errors.Is(http.ErrServerClosed, err) {
			log.Error("error starting server", slog.String("error", err.Error()))
		}
	}()
	log.Info("Server started", slog.String("host", cfg.HttpServer.Address))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Server is down", slog.String("Status", "ok"))
	if err = server.Shutdown(context.Background()); err != nil {
		log.Error("error shutting down server", slog.String("error", err.Error()))
	}
	if err = db.Close(); err != nil {
		log.Error("error closing database", slog.String("error", err.Error()))
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog(slog.LevelDebug)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			}))
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelInfo,
			}))
	default:
		log = setupPrettySlog()
	}

	return log
}

func setupPrettySlog(level ...slog.Level) *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: level[0],
		},
	}
	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
