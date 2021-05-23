package usecase

import (
	"errors"
	"fmt"
	"marketplace/accounts/domain"

	"github.com/go-pg/pg/v10"
)

type UpdateBalanceByIdCmd func(db *pg.DB, userId int64, deltaBalance float64) (error)

func UpdateBalanceById() UpdateBalanceByIdCmd {

	return func(db *pg.DB, userId int64, deltaBalance float64) (error) {

		acc := domain.Account{Id: userId}

		err := db.Model(&acc).WherePK().Select()

		if err != nil {
			fmt.Println("error 1 : ", err)
			return err
		}

		if acc.Email == "" {
			fmt.Println("error 2 : ", err)
			return errors.New("not found")
		}

		if acc.Balance + deltaBalance < 0 {
			fmt.Println("error 3 : ", err)
			return errors.New("Balance is too low.")
		}

		acc.Balance += deltaBalance

		_, err = db.Model(&acc).WherePK().Update()

		if err != nil {
			fmt.Println("error 4 : ", err)
			return err
		}

		return nil
	}
}