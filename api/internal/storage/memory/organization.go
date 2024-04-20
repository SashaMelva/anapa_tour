package memory

import (
	"github.com/SashaMelva/anapa_tour/internal/storage/model"
)

func (s *Storage) GetAllOrganizationByAdminId(admiId int) ([]*model.OrganizationFull, error) {
	organizations := []*model.OrganizationFull{}
	query := `select organization.id, organization.name, organization.description, organization.full_name, organization.admin_id, organization.category_id, category_organization.name  from  public.organization 
	left join category_organization On category_organization.id = organization.category_id where admin_id = $1`
	rows, err := s.ConnectionDB.Query(query, admiId) // sql.Result

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		organization := model.OrganizationFull{}

		if err := rows.Scan(
			&organization.Id,
			&organization.Name,
			&organization.Description,
			&organization.FullName,
			&organization.AdminId,
			&organization.CategoryData.Id,
			&organization.CategoryData.Name,
		); err != nil {
			return nil, err
		}

		organizations = append(organizations, &organization)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return organizations, nil
}
