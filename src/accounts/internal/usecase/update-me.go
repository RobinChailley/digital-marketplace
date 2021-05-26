package usecase

import (
	"marketplace/accounts/domain"
	"marketplace/accounts/internal/request"

	"github.com/go-pg/pg/v10"
)

type UpdateMeResponse struct {
	Id       int64        `json:"id"`
	Email    string       `json:"email"`
	Username string       `json:"username"`
	Password string       `json:"-"`
	Balance  float64      `json:"balance"`
	Ads      []domain.Ads `json:"-"`
	Admin    bool         `json:"admin"`
}

type UpdateMeCmd func(db *pg.DB, acc *domain.Account, update *request.UpdateMeRequest) (UpdateMeResponse, error)

func UpdateMe() UpdateMeCmd {

	return func(db *pg.DB, acc *domain.Account, update *request.UpdateMeRequest) (UpdateMeResponse, error) {
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
			return UpdateMeResponse{}, err
		}

		return UpdateMeResponse(*acc), nil
	}
}
