package app

import (
	"errors"
	"log"

	"github.com/SashaMelva/anapa_tour/internal/storage/memory"
	autenficationmodel "github.com/SashaMelva/anapa_tour/internal/storage/model/autenfication"
	"github.com/SashaMelva/anapa_tour/server/pkg"
	"go.uber.org/zap"
)

type App struct {
	storage *memory.Storage
	Logger  *zap.SugaredLogger
	JwtKey  string
}

func New(logger *zap.SugaredLogger, storage *memory.Storage, jwtKey string) *App {
	return &App{
		storage: storage,
		Logger:  logger,
		JwtKey:  jwtKey,
	}
}

func (a *App) LoginAccout(account *autenficationmodel.Account) (string, error) {
	accountData, err := a.storage.GetAccountByLogin(account.Login)

	if err != nil {
		return "", err
	}
	if accountData.Password == "" {
		return "", errors.New("Пользователя с таким email не существует")
	}
	if accountData.Password != account.Password {
		return "", errors.New("Неверный пароль")
	}

	token, err := pkg.GenerateToken(account.Id, account.Role, a.JwtKey)

	if err != nil {
		log.Fatal(err)
		return "", errors.New("Ошибка генерации JWT Токена")
	}

	a.Logger.Info(token)
	flag, err := a.storage.CheckJwtToken(account.Id)
	a.Logger.Info(flag)
	if flag {
		err = a.storage.SaveJwtToken(token, account.Id)
	} else {
		err = a.storage.UpdateJwtToken(token, account.Id)
	}
	if err != nil {
		return "", err
	}

	return token, nil
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
