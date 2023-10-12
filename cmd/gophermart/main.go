package main

import (
	"github.com/Tokebay/yandex-diplom/config"
	"github.com/Tokebay/yandex-diplom/internal/api"
	"github.com/Tokebay/yandex-diplom/internal/database"
	"github.com/Tokebay/yandex-diplom/internal/logger"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func main() {
	logger.Initialize("Info")
}

func run() error {
	cfg := config.NewConfig()

	dbInit, err := database.NewPostgreSQL(cfg.DatabaseURI)
	if err != nil {
		logger.Log.Error("Error in NewPostgreSQLStorage", zap.Error(err))
		return err
	}
	router(dbInit)
	return nil
}

func router(db *database.PostgreInit) chi.Router {

	r := chi.NewRouter()
	r.Post("/api/user/register", api.RegisterHandler)
	r.Post("/api/user/login", api.LoginHandler)

	return r
}
