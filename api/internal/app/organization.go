package app

import (
	"github.com/SashaMelva/anapa_tour/internal/storage/model"
)

func (a *App) GetOrganizationByAdminId(id int) ([]*model.OrganizationFull, error) {
	pins, err := a.storage.GetAllOrganizationByAdminId(id)

	if err != nil {
		a.Logger.Error(err)
	}

	return pins, nil
}

func (a *App) GetOrganizationByCategory(id int) ([]*model.OrganizationFull, error) {
	pins, err := a.storage.GetAllOrganizationByCategory(id)

	if err != nil {
		a.Logger.Error(err)
	}

	return pins, nil
}
