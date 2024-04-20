package app

import (
	"context"
	"fmt"

	loyaltymodel "github.com/SashaMelva/anapa_tour/internal/storage/model/loyalty"
)

func (a *App) CeratePromotion(promotion *loyaltymodel.Promotion) (int, error) {
	id, err := a.storage.CeratePromotion(promotion)

	if err != nil {
		a.Logger.Error(err)
	} else {
		a.Logger.Info(fmt.Sprintf("Create action whith id = %v", id))
	}

	return id, nil
}

func (a *App) DeletePromotionByID(id uint32) error {
	err := a.storage.DeletePromotionById(id)

	if err != nil {
		a.Logger.Error(err)
	} else {
		a.Logger.Info(fmt.Sprintf("Delet Action whith id = %v", id))
	}

	return err
}

func (a *App) EditPromotion(ctx context.Context, prompt *loyaltymodel.Promotion) error {
	err := a.storage.EditPromotion(prompt)

	if err != nil {
		a.Logger.Error(err)
	} else {
		a.Logger.Info(fmt.Sprintf("Update event whith id = %v", prompt.Id))
	}

	return err
}

func (a *App) GetPromotionByID(id int) (*loyaltymodel.Promotion, error) {
	prom, err := a.storage.GetPromotionByID(id)

	if err != nil {
		a.Logger.Error(err)
	}

	return prom, nil
}

func (a *App) GetPromotionByOrganizationID(org int) ([]*loyaltymodel.Promotion, error) {
	pins, err := a.storage.GetPromotionByOrganizationID(org)

	if err != nil {
		a.Logger.Error(err)
	}

	return pins, nil
}
