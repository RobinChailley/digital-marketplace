package usecase

import (
	"github.com/go-pg/pg/v10"
	"marketplace/accounts/domain"
)

type AdminGetAllUsersCmd func(db *pg.DB) ([]domain.Account, error)

func AdminGetAllUsers() AdminGetAllUsersCmd {

	return func(db *pg.DB) ([]domain.Account, error) {

		users := []domain.Account{}

		err := db.Model(&users).Select()

		if err != nil {
			return []domain.Account{}, err
		}

		return users, nil
	}
}
