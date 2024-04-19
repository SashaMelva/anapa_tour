package app

import (
	"context"
	"fmt"

	loyaltymodel "github.com/SashaMelva/anapa_tour/internal/storage/model/loyalty"
)

func (a *App) CeratePin(pin *loyaltymodel.Pin) (int, error) {
	id, err := a.storage.CeratePin(pin)

	if err != nil {
		a.Logger.Error(err)
	} else {
		a.Logger.Info(fmt.Sprintf("Create event whith id = %v", id))
	}

	return id, nil
}

func (a *App) DeletePinByID(pinId uint32) error {
	err := a.storage.DeletePointById(pinId)

	if err != nil {
		a.Logger.Error(err)
	} else {
		a.Logger.Info(fmt.Sprintf("Delet event whith id = %v", pinId))
	}

	return err
}

func (a *App) GetAllPins(ctx context.Context) ([]*loyaltymodel.Pin, error) {
	pins, err := a.storage.ListAllPins()

	if err != nil {
		a.Logger.Error(err)
	}

	return pins, nil
}

func (a *App) EditPin(ctx context.Context, pin *loyaltymodel.Pin) error {
	err := a.storage.EditPin(pin)

	if err != nil {
		a.Logger.Error(err)
	} else {
		a.Logger.Info(fmt.Sprintf("Update event whith id = %v", pin.Id))
	}

	return err
}
