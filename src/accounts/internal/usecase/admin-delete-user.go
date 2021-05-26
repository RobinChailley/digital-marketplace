package usecase

import (
	"errors"
	"fmt"
	"github.com/go-pg/pg/v10"
	"marketplace/accounts/domain"
)

type AdminDeleteUserCmd func(db *pg.DB, userId int64) error

func AdminDeleteUser() AdminDeleteUserCmd {

	return func(db *pg.DB, userId int64) error {
		user := domain.Account{
			Id: userId,
		}

		err := db.Model(&user).WherePK().Select()

		if err != nil {
			return errors.New("not found")
		}

		fmt.Println("user : ", user)

		if user.Email == "" {
			return errors.New("not found")
		}

		_, err = db.Model(&user).WherePK().Delete()

		if err != nil {
			return err
		}
		return nil
	}
}
