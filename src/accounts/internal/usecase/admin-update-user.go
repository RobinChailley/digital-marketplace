package usecase

import (
	"errors"
	"github.com/go-pg/pg/v10"
	"marketplace/accounts/domain"
	"marketplace/accounts/internal/request"
)

type AdminUpdateUserCmd func(db *pg.DB, updateReq request.UpdateUserRequest, userId int64) (domain.Account, error)

func AdminUpdateUser() AdminUpdateUserCmd {

	return func(db *pg.DB, updateReq request.UpdateUserRequest, userId int64) (domain.Account, error) {
		user := domain.Account{Id: userId}

		err := db.Model(&user).WherePK().Select()

		if err != nil {
			return domain.Account{}, errors.New("not found")
		}

		if updateReq.Email != "" {
			user.Email = updateReq.Email
		}

		if updateReq.Username != "" {
			user.Username = updateReq.Username
		}

		if updateReq.Password != "" {
			user.Password = updateReq.Password
		}

		if updateReq.Balance != 0 {
			user.Balance = updateReq.Balance
		}

		if updateReq.Admin != false {
			user.Admin = updateReq.Admin
		}

		_, err = db.Model(&user).WherePK().Update()

		if err != nil {
			return domain.Account{}, err
		}

		return user, nil
	}
}
