package memory

import loyaltymodel "github.com/SashaMelva/anapa_tour/internal/storage/model/loyalty"

func (s *Storage) CreateBalanceBonus(accauntId int) error {
	var balanseId int
	query := `insert into balance_bonus(account_id, summ) values($1, $2) RETURNING id`
	result := s.ConnectionDB.QueryRow(query, accauntId, 0) // sql.Result
	err := result.Scan(&balanseId)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) ChangeBalanceByAccountId(balanceId int, summ int) error {
	query := `update balance_bonus set summ=$1 where account_id=$2`
	_, err := s.ConnectionDB.Exec(query, summ, balanceId)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) SelectBalanceByAccount(accountId int) (*loyaltymodel.BalanceBonus, error) {
	var bonus loyaltymodel.BalanceBonus
	query := `select * from balance_bonus  where account_id=$1`
	rows := s.ConnectionDB.QueryRow(query, accountId)

	if err := rows.Scan(
		&bonus.Id,
		&bonus.AccountId,
		&bonus.Summ,
	); err != nil {
		return nil, err
	}

	return &bonus, nil
}
