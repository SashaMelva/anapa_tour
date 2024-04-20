package app

import (
	"context"
	"fmt"

	loyaltymodel "github.com/SashaMelva/anapa_tour/internal/storage/model/loyalty"
)

func (a *App) CerateAction(action *loyaltymodel.Action) (int, error) {
	id, err := a.storage.CerateAction(action)

	if err != nil {
		a.Logger.Error(err)
	} else {
		a.Logger.Info(fmt.Sprintf("Create action whith id = %v", id))
	}

	return id, nil
}

func (a *App) DeleteActionByID(pinId uint32) error {
	err := a.storage.DeleteActionById(pinId)

	if err != nil {
		a.Logger.Error(err)
	} else {
		a.Logger.Info(fmt.Sprintf("Delet Action whith id = %v", pinId))
	}

	return err
}

func (a *App) EditAction(ctx context.Context, action *loyaltymodel.Action) error {
	err := a.storage.EditAction(action)

	if err != nil {
		a.Logger.Error(err)
	} else {
		a.Logger.Info(fmt.Sprintf("Update event whith id = %v", action.Id))
	}

	return err
}

func (a *App) GetActiveAction(activ bool) ([]*loyaltymodel.Action, error) {
	pins, err := a.storage.ListActiveAction(activ)

	if err != nil {
		a.Logger.Error(err)
	}

	return pins, nil
}
func (a *App) GetActiveAndTypeAction(activ bool, typeAction string) ([]*loyaltymodel.Action, error) {
	pins, err := a.storage.ListActiveAndTypeAction(activ, typeAction)

	if err != nil {
		a.Logger.Error(err)
	}

	return pins, nil
}

func (a *App) GetAllAction() ([]*loyaltymodel.Action, error) {
	pins, err := a.storage.ListAllActions()

	if err != nil {
		a.Logger.Error(err)
	}

	return pins, nil
}
