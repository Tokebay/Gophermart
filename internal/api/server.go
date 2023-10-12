package api

import (
	"net/http"

	"github.com/Tokebay/yandex-diplom/config"
	"github.com/Tokebay/yandex-diplom/internal/database"
	"github.com/Tokebay/yandex-diplom/internal/logger"
	service "github.com/Tokebay/yandex-diplom/internal/services"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type Server struct {
	config   *config.Config
	database *database.PostgreInit
	router   *chi.Mux
	user     *service.UserHandler
}

func NewServer(cfg *config.Config, db *database.PostgreInit) *Server {
	s := &Server{
		config:   cfg,
		database: db,
		router:   chi.NewRouter(),
	}

	s.setupRoutes()
	return s
}

func (s *Server) setupRoutes() {
	// Настройка маршрутов и их обработчиков.
	s.router.Post("/api/user/register", s.user.RegisterHandler)
	s.router.Post("/api/user/login", s.user.LoginHandler)

}

func (s *Server) Start() error {
	logger.Log.Info("Server is starting at address " + s.config.RunAddress)
	err := http.ListenAndServe(s.config.RunAddress, s.router)
	if err != nil {
		logger.Log.Error("Error starting server", zap.Error(err))
		return err
	}
	return nil
}
