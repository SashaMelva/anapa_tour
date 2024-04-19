package memory

import (
	"fmt"

	autenficationmodel "github.com/SashaMelva/anapa_tour/internal/storage/model/autenfication"
)

func (s *Storage) CheckUniqueLogin(login string) (int, error) {
	var accountId int
	query := `select id from accounts where login = $1`
	result := s.ConnectionDB.QueryRow(query, login)

	err := result.Scan(&accountId)
	fmt.Println(accountId)

	if err != nil {
		return 0, nil
	}

	return accountId, nil
}

func (s *Storage) CheckJwtToken(idAccount uint32) (bool, error) {
	var accountId int
	query := `select id from tokens where accounId=$1`
	result := s.ConnectionDB.QueryRow(query, idAccount) // sql.Result
	err := result.Scan(&accountId)

	if err != nil {
		return true, nil
	}

	return false, nil
}

func (s *Storage) UpdateJwtToken(token string, idAccount uint32) error {
	var accountId int
	query := `UPDATE set accounId = $1, token =$2 where accounId = $1`
	result := s.ConnectionDB.QueryRow(query, idAccount, token) // sql.Result
	err := result.Scan(&accountId)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) SaveJwtToken(token string, idAccount uint32) error {
	var accountId int
	query := `insert into tokens(accounId, token) values($1, $2)`
	result := s.ConnectionDB.QueryRow(query, idAccount, token) // sql.Result
	err := result.Scan(&accountId)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetAccountByLogin(login string) (autenficationmodel.Account, error) {
	var account autenficationmodel.Account
	query := `select id, login, password from accounts where login = $1`
	result := s.ConnectionDB.QueryRow(query, login) // sql.Result
	err := result.Scan(&account.Id, &account.Login, &account.Password)

	if err != nil {
		return autenficationmodel.Account{}, err
	}

	return account, nil
}

func (s *Storage) RegisterAccount(account *autenficationmodel.Account) (int, error) {
	var accountId int
	query := `insert into accounts(login, password, role) values($1, $2, $3) RETURNING id`
	result := s.ConnectionDB.QueryRow(query, account.Login, account.Password, account.Role) // sql.Result
	err := result.Scan(&accountId)

	if err != nil {
		return 0, err
	}

	return accountId, nil
}
