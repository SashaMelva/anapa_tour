package app

import (
	"errors"

	memorystorage "github.com/SashaMelva/anapa_tour/api/internal/storage/memory"
	autenficationmodel "github.com/SashaMelva/anapa_tour/internal/storage/model/autenfication"
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

func (a *App) RegisterUser(user autenficationmodel.Account) (uint32, error) {
	if user.Login == "" {
		return 0, errors.New("Поле логин не валидно")
	}

	createdAccountId, err := a.storage.CreateUser(user)

	if err != nil {
		return 0, err
	}

	return createdAccountId, err
}

func (a *App) CheckUniqueLogin(login string) error {
	id, err := a.storage.CheckUniqueLogin(login)

	if err != nil {
		return err
	}
	if id > 0 {
		return errors.New("Пользователь с данным логином уже существует")
	}

	return nil
}
