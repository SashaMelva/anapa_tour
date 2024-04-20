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

type RequestAuth struct {
	Token     string `json:"Token"`
	Role      string `json:"role"`
	IdAccount uint32 `json:"id_account"`
}

func (a *App) LoginAccout(account *autenficationmodel.Account) (RequestAuth, error) {
	accountData, err := a.storage.GetAccountByLogin(account.Login)
	req := RequestAuth{}

	if err != nil {
		return req, err
	}
	if accountData.Password == "" {
		return req, errors.New("Пользователя с таким email не существует")
	}
	if accountData.Password != account.Password {
		return req, errors.New("Неверный пароль")
	}

	token, err := pkg.GenerateToken(account.Id, account.Role, a.JwtKey)

	if err != nil {
		log.Fatal(err)
		return req, errors.New("Ошибка генерации JWT Токена")
	}

	a.Logger.Info(token)
	flag, err := a.storage.CheckJwtToken(accountData.Id)
	a.Logger.Info(flag, accountData.Id)
	if flag {
		err = a.storage.SaveJwtToken(token, accountData.Id)
	} else {
		err = a.storage.UpdateJwtToken(token, accountData.Id)
	}
	if err != nil {
		return req, err
	}

	return RequestAuth{
		Token:     token,
		Role:      accountData.Role,
		IdAccount: accountData.Id}, nil
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
