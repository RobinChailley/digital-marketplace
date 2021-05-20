package usecase

import (
	"marketplace/accounts/domain"
	"github.com/go-pg/pg/v10"
)

type AddFundsCmd func (db *pg.DB, acc *domain.Account, fundsToAdd int64) (int64, error)

func AddFunds() AddFundsCmd {

	return func (db *pg.DB, acc *domain.Account, fundsToAdd int64 ) (int64, error) {
		acc.Balance += fundsToAdd

		_, err := db.Model(acc).WherePK().Update()

		if err != nil {
			return -1, err
		}

		return acc.Balance, nil
	}
}