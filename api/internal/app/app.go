package app

import (
	memorystorage "github.com/SashaMelva/anapa_tour/api/internal/storage/memory"
	"go.uber.org/zap"
)

type App struct {
	storage *memorystorage.Storage
	Logger  *zap.SugaredLogger
}

func New(logger *zap.SugaredLogger, storage *memorystorage.Storage) *App {
	return &App{
		storage: storage,
		Logger:  logger,
	}
}


func (a *)