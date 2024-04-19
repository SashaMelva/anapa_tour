package memory

import (
	loyaltymodel "github.com/SashaMelva/anapa_tour/internal/storage/model/loyalty"
)

func (s *Storage) CeratePin(pin *loyaltymodel.Pin) (int, error) {
	var pinId int
	query := `insert into pins(coordinat, name, description, activ) values($1, $2, $3, $4) RETURNING id`
	result := s.ConnectionDB.QueryRow(query, pin.Coordinat, pin.Name, pin.Description, pin.Activ) // sql.Result
	err := result.Scan(&pinId)

	if err != nil {
		return 0, err
	}

	return pinId, nil
}

func (s *Storage) ChangeActivePointById(id uint32, activ bool) error {
	query := `update pins set activ=$1 where id=$2`
	_, err := s.ConnectionDB.Exec(query, activ, id)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) DeletePointById(id uint32) error {
	query := `delete pins events where id = $1`
	_, err := s.ConnectionDB.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) EditPin(pin *loyaltymodel.Pin) error {
	query := `update pins set name=$1, description=$2, coordinat=$3, activ=$4 where id=$5`
	_, err := s.ConnectionDB.Exec(query, pin.Name, pin.Description, pin.Coordinat, pin.Activ, pin.Id)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) ListAllPins() ([]*loyaltymodel.Pin, error) {
	pins := []*loyaltymodel.Pin{}
	query := `select id, name, description, coordinat, activ from pins`
	rows, err := s.ConnectionDB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		pin := loyaltymodel.Pin{}

		if err := rows.Scan(
			&pin.Id,
			&pin.Name,
			&pin.Description,
			&pin.Coordinat,
			&pin.Activ,
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

func (s *Storage) ListActivePins() ([]*loyaltymodel.Pin, error) {
	pins := []*loyaltymodel.Pin{}
	query := `select id, name, description, coordinat, activ from pins where activ = true`
	rows, err := s.ConnectionDB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		pin := loyaltymodel.Pin{}

		if err := rows.Scan(
			&pin.Id,
			&pin.Name,
			&pin.Description,
			&pin.Coordinat,
			&pin.Activ,
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
