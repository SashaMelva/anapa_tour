package app

import (
	"errors"

	marcketmodel "github.com/SashaMelva/anapa_tour/internal/storage/model/marcket"
)

func (a *App) GetAiiProduction() ([]*marcketmodel.Production, error) {
	prod, err := a.storage.GetAllProductions()

	if err != nil {
		return nil, err
	}

	return prod, nil
}

func (a *App) ByProductions(byData marcketmodel.ByProduction) error {

	balanceNow, err := a.storage.SelectBalanceByAccount(int(byData.IdAccount))
	if err != nil {
		a.Logger.Error(err)
		return err
	}

	if balanceNow.Summ < byData.Price {
		return errors.New("Недостаточно баллов")
	}

	money := balanceNow.Summ - byData.Price

	err = a.storage.ChangeBalanceByAccountId(int(byData.IdAccount), money)

	if err != nil {
		a.Logger.Error(err)
		return err
	}

	err = a.storage.ByProductions(byData.ContScklad-1, byData.IdProd)

	if err != nil {
		return err
	}

	return nil
}
