package database

import (
	"database/sql"

	"github.com/Tokebay/yandex-diplom/internal/logger"
	"github.com/Tokebay/yandex-diplom/internal/models"
	goose "github.com/pressly/goose/v3"
	"go.uber.org/zap"
)

type PostgreInit struct {
	db *sql.DB
}

func NewPostgreSQL(dsn string) (*PostgreInit, error) {
	// Выполнить миграции
	db, err := goose.OpenDBWithDriver("pgx", dsn)
	if err != nil {
		logger.Log.Error("Error open conn", zap.Error(err))
		return nil, err
	}
	err = goose.Up(db, "./migration")
	if err != nil {
		logger.Log.Error("Error goose UP", zap.Error(err))
		return nil, err
	}

	// Вернуть созданный объект PostgreSQLStorage
	return &PostgreInit{db: db}, nil
}

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByLogin(login string) (*models.User, error)
}
