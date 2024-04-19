package app

import (
	"errors"

	"github.com/SashaMelva/anapa_tour/internal/storage/memory"
	autenficationmodel "github.com/SashaMelva/anapa_tour/internal/storage/model/autenfication"
	"go.uber.org/zap"
)

type App struct {
	storage *memory.Storage
	Logger  *zap.SugaredLogger
}

func New(logger *zap.SugaredLogger, storage *memory.Storage) *App {
	return &App{
		storage: storage,
		Logger:  logger,
	}
}

func (a *App) RegisterAccount(user *autenficationmodel.Account) (int, error) {
	if user.Login == "" {
		return 0, errors.New("Поле логин не валидно")
	}

	createdAccountId, err := a.storage.RegisterAccount(user)

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
