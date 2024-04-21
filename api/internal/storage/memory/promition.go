package memory

import loyaltymodel "github.com/SashaMelva/anapa_tour/internal/storage/model/loyalty"

func (s *Storage) CeratePromotion(promition *loyaltymodel.Promotion) (int, error) {
	var pinId int
	query := `insert into promotions(name, description, activ, cost, id_category_organization) values($1, $2, $3, $4, $5) RETURNING id`
	result := s.ConnectionDB.QueryRow(query, promition.Name, promition.Description, promition.Activ, promition.Сost, promition.IdCategoryOrganization) // sql.Result
	err := result.Scan(&pinId)

	if err != nil {
		return 0, err
	}

	return pinId, nil
}

func (s *Storage) DeletePromotionById(id uint32) error {
	query := `delete promotions where id = $1`
	_, err := s.ConnectionDB.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) EditPromotion(prom *loyaltymodel.Promotion) error {
	query := `update promotions set name=$1, description=$2, activ=$3, cost=$4, id_category_organization=$5 where id=$6`
	_, err := s.ConnectionDB.Exec(query, prom.Name, prom.Description, prom.Activ, prom.Сost, prom.IdCategoryOrganization, prom.Id)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) ListAllPromotions() ([]*loyaltymodel.Promotion, error) {
	proms := []*loyaltymodel.Promotion{}
	query := `select id, name, description, activ, cost, id_category_organization from promotions`
	rows, err := s.ConnectionDB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		prom := loyaltymodel.Promotion{}

		if err := rows.Scan(
			&prom.Id,
			&prom.Name,
			&prom.Description,
			&prom.Activ,
			&prom.Сost,
			&prom.IdCategoryOrganization,
		); err != nil {
			return nil, err
		}

		proms = append(proms, &prom)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return proms, nil
}

func (s *Storage) GetPromotionByID(id int) (*loyaltymodel.Promotion, error) {
	prom := loyaltymodel.Promotion{}
	query := `select id, name, description, activ, cost, id_category_organization from promotions where id = $1`
	rows := s.ConnectionDB.QueryRow(query, id)

	if err := rows.Scan(
		&prom.Id,
		&prom.Name,
		&prom.Description,
		&prom.Activ,
		&prom.Сost,
		&prom.IdCategoryOrganization,
	); err != nil {
		return nil, err
	}

	return &prom, nil
}

func (s *Storage) GetPromotionByOrganizationID(orgId int) ([]*loyaltymodel.Promotion, error) {
	proms := []*loyaltymodel.Promotion{}
	query := `select id, name, description, activ, cost, id_category_organization from promotions where id_category_organization = $1`
	rows, err := s.ConnectionDB.Query(query, orgId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		prom := loyaltymodel.Promotion{}

		if err := rows.Scan(
			&prom.Id,
			&prom.Name,
			&prom.Description,
			&prom.Activ,
			&prom.Сost,
			&prom.IdCategoryOrganization,
		); err != nil {
			return nil, err
		}

		proms = append(proms, &prom)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return proms, nil
}
