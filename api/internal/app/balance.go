package app

import (
	"errors"

	loyaltymodel "github.com/SashaMelva/anapa_tour/internal/storage/model/loyalty"
)

func (a *App) GetBalanceByAccountId(id int) (*loyaltymodel.BalanceBonus, error) {
	prom, err := a.storage.SelectBalanceByAccount(id)

	if err != nil {
		a.Logger.Error(err)
	}

	return prom, nil
}

func (a *App) ChangeBalance(balance *loyaltymodel.ChangeBalance) error {
	balanceNow, err := a.storage.SelectBalanceByAccount(int(balance.AccountId))
	money := balanceNow.Summ

	if err != nil {
		a.Logger.Error(err)
		return err
	}

	if balance.Plus {
		money += balance.Summ
	}

	if !balance.Plus {
		if balance.Summ > balanceNow.Summ {
			err = errors.New("Ошибка списания балов. Недостаточно баллов")
			a.Logger.Error(err)
			return err
		}
		money -= balance.Summ
	}

	err = a.storage.ChangeBalanceByAccountId(int(balanceNow.AccountId), money)

	if err != nil {
		a.Logger.Error(err)
		return err
	}

	return nil
}

func (a *App) ChangeBalanceByCategory(balance *loyaltymodel.ChangeBalanceForOrganization) error {
	promition, err := a.storage.GetPromotionByOrganizationID(int(balance.OrgCategoryId))

	if err != nil {
		a.Logger.Error(err)
		return err
	}
	balanceNow, err := a.storage.SelectBalanceByAccount(int(balance.AccountId))
	if err != nil {
		a.Logger.Error(err)
		return err
	}

	if promition == nil {
		return errors.New("Акций не найдено!")
	}

	money := balanceNow.Summ + promition[0].Сost

	err = a.storage.ChangeBalanceByAccountId(int(balance.AccountId), money)

	if err != nil {
		a.Logger.Error(err)
		return err
	}

	return nil
}
