package usecase

import (
	"errors"
	"fmt"
	"marketplace/transactions/domain"

	"github.com/go-pg/pg/v10"
)

type CancelTransactionCmd func(db *pg.DB, transacId int64, userId int64) (domain.Transaction, error)

func CancelTransaction() CancelTransactionCmd {

	return func(db *pg.DB, transacId int64, userId int64) (domain.Transaction, error) {
		transac := domain.Transaction{Id: transacId}
		err := db.Model(&transac).
			Column("transaction.*").
			Relation("Buyer").
			Relation("Ads").
			Relation("Seller").
			Relation("Messages").
			WherePK().
			Select()

		if err != nil {
			return domain.Transaction{}, errors.New("not found")
		}

		if transac.SellerId != userId {
			return domain.Transaction{}, errors.New("unauthorized")
		}

		if transac.Status == "PROPOSITION" {
			transac.Status = "CANCELLED"
		} else if transac.Status == "CANCELLED" {
			transac.Status = "PROPOSITION"
		} else {
			return domain.Transaction{}, errors.New(fmt.Sprintf("Can not cancel a transaction when his status is : %s", transac.Status))
		}

		_, err = db.Model(&transac).WherePK().Update()

		if err != nil {
			return domain.Transaction{}, errors.New("can not update")
		}

		return transac, nil
	}

}
