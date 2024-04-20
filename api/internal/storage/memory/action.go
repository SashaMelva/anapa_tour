package memory

import loyaltymodel "github.com/SashaMelva/anapa_tour/internal/storage/model/loyalty"

func (s *Storage) CerateAction(action *loyaltymodel.Action) (int, error) {
	var pinId int
	query := `insert into actions(name, description,category_id,date_start, date_end, activ, organization_id) values($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	result := s.ConnectionDB.QueryRow(query, action.Name, action.Description, action.CategoryId, action.DateStart, action.DateEnd, action.Activ, action.OrgId) // sql.Result
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
	query := `update actions set name=$1, description=$2, category_id=$3, date_start=$4, date_end=$5, activ=$6 where id=$7`
	_, err := s.ConnectionDB.Exec(query, action.Name, action.Description, action.CategoryId, action.DateStart, action.DateEnd, action.Activ, action.Id)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) ListAllActions(organizationId int) ([]*loyaltymodel.Action, error) {
	pins := []*loyaltymodel.Action{}
	query := `select id, name, description, category_id, date_start, date_end, activ from actions where organization_id = $1`
	rows, err := s.ConnectionDB.Query(query, organizationId)

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
			&pin.CategoryId,
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

func (s *Storage) ListActiveAction(flag bool, org int) ([]*loyaltymodel.Action, error) {
	pins := []*loyaltymodel.Action{}
	query := `select id, name, description, category_id, date_start, date_end, activ from actions where activ = $1 and organization_id = $2`
	rows, err := s.ConnectionDB.Query(query, flag, org)

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
			&pin.CategoryId,
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

func (s *Storage) ListActiveAndTypeAction(flag bool, typeAction string, orgId int) ([]*loyaltymodel.Action, error) {
	pins := []*loyaltymodel.Action{}
	query := `select id, name, description, category_id, date_start, date_end, activ from actions where activ = $1 and category_id = $2 and organization_id = $3`
	rows, err := s.ConnectionDB.Query(query, flag, typeAction, orgId)

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
			&pin.CategoryId,
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
