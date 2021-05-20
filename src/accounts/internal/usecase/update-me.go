package usecase

import (
	"marketplace/accounts/domain"
	"github.com/go-pg/pg/v10"
	"marketplace/accounts/internal/request"
)

type UpdateMeCmd func (db *pg.DB, acc *domain.Account, update *request.UpdateMeRequest) (*domain.Account, error)

func UpdateMe() UpdateMeCmd {

	return func (db *pg.DB, acc *domain.Account, update *request.UpdateMeRequest) (*domain.Account, error) {
		if update.Email != "" {
			acc.Email = update.Email
		}

		if update.Password != "" {
			acc.Password = update.Password
		}

		if update.Username != "" {
			acc.Username = update.Username
		}

		_, err := db.Model(acc).WherePK().Update()

		if err != nil {
			return nil, err
		}


		acc.Password = "-/-"  // TODO : How to do it cleanly?

		return acc, nil
	}
}