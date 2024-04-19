package memory

import autenficationmodel "github.com/SashaMelva/anapa_tour/api/internal/storage/model/autenfication"

func (s *Storage) CheckUniqueLogin(account *autenficationmodel.Account) (int, error) {
	var accountId int
	query := `select id from users where login = $1 RETURNING id`
	result := s.ConnectionDB.QueryRow(query, account.Login) // sql.Result
	err := result.Scan(&accountId)

	if err != nil {
		return 0, err
	}

	return accountId, nil
}

func (s *Storage) RegisterAccount(account *autenficationmodel.Account) (int, error) {
	var eventId int
	query := `insert into accounts(login, password) values($1, $2) RETURNING id`
	result := s.ConnectionDB.QueryRow(query, account.Login, account.Password, account.Role) // sql.Result
	err := result.Scan(&eventId)

	if err != nil {
		return 0, err
	}

	return eventId, nil
}
