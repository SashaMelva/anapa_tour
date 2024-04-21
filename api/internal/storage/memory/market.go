package memory

import (
	marcketmodel "github.com/SashaMelva/anapa_tour/internal/storage/model/marcket"
)

func (s *Storage) GetAllProductions() ([]*marcketmodel.Production, error) {
	pins := []*marcketmodel.Production{}
	query := `select id, name, description, price, count from productions`
	rows, err := s.ConnectionDB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		pin := marcketmodel.Production{}

		if err := rows.Scan(
			&pin.Id,
			&pin.Name,
			&pin.Description,
			&pin.Price,
			&pin.Count,
		); err != nil {
			return nil, err
		}

		pins = append(pins, &pin)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return pins, nil
}

func (s *Storage) ByProductions(counts int, id int) error {
	query := `update productions set count=$1 where id=$2`
	_, err := s.ConnectionDB.Exec(query, counts, id)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) CreateTicket(time string, idacc int, id int) error {
	query := `insert into ticket_market(date, account_id, product_id) values($1, $2, $3) `
	_, err := s.ConnectionDB.Exec(query, time, idacc, id)

	if err != nil {
		return err
	}

	return nil
}
