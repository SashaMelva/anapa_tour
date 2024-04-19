package memory

import (
	loyaltymodel "github.com/SashaMelva/anapa_tour/sapi/internal/storage/model/loyalty"
)

func (s *Storage) CeratePin(pin *loyaltymodel.Pin) (int, error) {
	var pinId int
	query := `insert into accounts(idAccount, token) values($1, $2) where idAccount= $2 RETURNING id`
	result := s.ConnectionDB.QueryRow(query, pin, idAccount) // sql.Result
	err := result.Scan(&pinId)

	if err != nil {
		return 0, err
	}

	return pinId, nil
}

func (s *Storage) DeletePointById(id uint32) error {
	query := `delete from pins where id = $1`
	_, err := s.ConnectionDB.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) EditPin(pin *loyaltymodel.Pin) error {
	query := `update events set title=$1, description=$2, date_time_start=$3, date_time_end=$4 where id=$5`
	_, err := s.ConnectionDB.Exec(query, pin.Title, event.Description, event.DateTimeStart, event.DateTimeEnd, event.ID)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) ListAllPins() ([]*loyaltymodel.Pin, error) {
	pins := []*loyaltymodel.Pin{}
	query := `select id, title, date_time_start, date_time_end, description from events where date_time_start >= $1::timestamp and date_time_end < $2::timestamp`
	rows, err := s.ConnectionDB.Query(query, storage.Date(dateStart), storage.Date(dateEnd))

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		pin := loyaltymodel.Pin{}

		if err := rows.Scan(
			&pin.ID,
			&pin.Title,
			&pin.Description,
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
