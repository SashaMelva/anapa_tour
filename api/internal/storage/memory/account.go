package memory

import autenficationmodel "github.com/SashaMelva/anapa_tour/internal/storage/model/autenfication"

func (s *Storage) CheckUniqueLogin(login string) (int, error) {
	var accountId int
	query := `select id from users where login = $1 RETURNING id`
	result := s.ConnectionDB.QueryRow(query, login) // sql.Result
	err := result.Scan(&accountId)

	if err != nil {
		return 0, err
	}

	return accountId, nil
}

func (s *Storage) SaveJwtToken(token string, idAccount uint32) error {
	var accountId int
	query := `insert into accounts(idAccount, token) values($1, $2) where idAccount= $2 RETURNING id`
	result := s.ConnectionDB.QueryRow(query, token, idAccount) // sql.Result
	err := result.Scan(&accountId)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetAccountByLogin(login string) (autenficationmodel.Account, error) {
	var account autenficationmodel.Account
	query := `select id, login, password from users where login = $1`
	result := s.ConnectionDB.QueryRow(query, login) // sql.Result
	err := result.Scan(&account)

	if err != nil {
		return autenficationmodel.Account{}, err
	}

	return account, nil
}

func (s *Storage) RegisterAccount(account *autenficationmodel.Account) (int, error) {
	var accountId int
	query := `insert into accounts(login, password) values($1, $2) RETURNING id`
	result := s.ConnectionDB.QueryRow(query, account.Login, account.Password, account.Role) // sql.Result
	err := result.Scan(&accountId)

	if err != nil {
		return 0, err
	}

	return accountId, nil
}
