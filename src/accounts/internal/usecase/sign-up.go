package usecase

import (
	"marketplace/accounts/domain"

	"github.com/go-pg/pg/v10"
)

type SignUpCmd func(db *pg.DB, acc *domain.Account) (int64, error)

func SignUp() SignUpCmd {

	return func(db *pg.DB, acc *domain.Account) (int64, error) {
		_, err := db.Model(acc).Insert()

		if err != nil {
			return -1, err
		}
		return acc.Id, nil
	}
}