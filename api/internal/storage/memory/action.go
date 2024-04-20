package memory

import loyaltymodel "github.com/SashaMelva/anapa_tour/internal/storage/model/loyalty"

func (s *Storage) CerateAction(action *loyaltymodel.Action) (int, error) {
	var pinId int
	query := `insert into actions(name, description,type,date_start, date_end, activ) values($1, $2, $3, $4, $5, $6) RETURNING id`
	result := s.ConnectionDB.QueryRow(query, action.Name, action.Description, action.Type, action.DateStart, action.DateEnd, action.Activ) // sql.Result
	err := result.Scan(&pinId)

	if err != nil {
		return 0, err
	}

	return pinId, nil
}

func (s *Storage) ChangeActiveActionById(id uint32, activ bool) error {
	query := `update actions set activ=$1 where id=$2`
	_, err := s.ConnectionDB.Exec(query, activ, id)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) DeleteActionById(id uint32) error {
	query := `delete actions events where id = $1`
	_, err := s.ConnectionDB.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) EditAction(action *loyaltymodel.Action) error {
	query := `update actions set name=$1, description=$2, type=$3, date_start=$4, date_end=$5, activ=$6 where id=$7`
	_, err := s.ConnectionDB.Exec(query, action.Name, action.Description, action.Type, action.DateStart, action.DateEnd, action.Activ, action.Id)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) ListAllActions() ([]*loyaltymodel.Action, error) {
	pins := []*loyaltymodel.Action{}
	query := `select id, name, description, type, date_start, date_end, activ from actions`
	rows, err := s.ConnectionDB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		pin := loyaltymodel.Action{}

		if err := rows.Scan(
			&pin.Id,
			&pin.Name,
			&pin.Description,
			&pin.Type,
			&pin.DateStart,
			&pin.DateEnd,
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

func (s *Storage) ListActiveAction(flag bool) ([]*loyaltymodel.Action, error) {
	pins := []*loyaltymodel.Action{}
	query := `select id, name, description, type, date_start, date_end, activ from actions where activ = $1`
	rows, err := s.ConnectionDB.Query(query, flag)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		pin := loyaltymodel.Action{}

		if err := rows.Scan(
			&pin.Id,
			&pin.Name,
			&pin.Description,
			&pin.Type,
			&pin.DateStart,
			&pin.DateEnd,
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

func (s *Storage) ListActiveAndTypeAction(flag bool, typeAction string) ([]*loyaltymodel.Action, error) {
	pins := []*loyaltymodel.Action{}
	query := `select id, name, description, type, date_start, date_end, activ from actions where activ = $1 and type = $2`
	rows, err := s.ConnectionDB.Query(query, flag, typeAction)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		pin := loyaltymodel.Action{}

		if err := rows.Scan(
			&pin.Id,
			&pin.Name,
			&pin.Description,
			&pin.Type,
			&pin.DateStart,
			&pin.DateEnd,
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
