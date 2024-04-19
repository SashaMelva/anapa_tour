package app

import "context"

func (a *App) GetAllHotels(ctx context.Context) error {

	_, err := a.storage.GetAllEvents()

	if err != nil {
		return nil, err
	}

	return nil
}
