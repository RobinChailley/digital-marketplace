package usecase

import (
	"errors"
	"github.com/go-pg/pg/v10"
	"marketplace/accounts/domain"
)

type AdminGetUserCmd func(db *pg.DB, userId int64) (domain.Account, error)

func AdminGetUser() AdminGetUserCmd {

	return func(db *pg.DB, userId int64) (domain.Account, error) {
		user := domain.Account{Id: userId}
		err := db.Model(&user).WherePK().Select()
		if err != nil {
			return domain.Account{}, err
		}

		if user.Email == "" {
			return domain.Account{}, errors.New("not found")
		}
		return user, nil
	}
}
