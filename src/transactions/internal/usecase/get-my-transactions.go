package usecase

import (
	"marketplace/transactions/domain"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

// type MyTransacsResponse struct {
// 	Id 				int64				`json:"id"`
// 	Buyer 		int64				`json:"id"`
// 	BuyerId 	int64				`json:"id"`
// 	Seller 		int64				`json:"id"`
// 	SellerId	int64				`json:"id"`
// 	Ads 			int64				`json:"id"`
// 	AdsId 		int64				`json:"id"`
// 	Messages 	int64				`json:"id"`
// 	Bid			 	float64			`json:"bid"`
// 	Status	 	string			`json:"status"`
// }

type GetMyTransactionsCmd func (db *pg.DB, userId int64) ([]domain.Transaction, error)

func GetMyTransactions() GetMyTransactionsCmd {

	return func (db *pg.DB, userId int64) ([]domain.Transaction, error) {
		var transacs []domain.Transaction

		err := db.Model(&transacs).
			Column("transaction.*").
			Relation("Buyer").
			Relation("Ads").
			Relation("Seller").
			Relation("Messages").
			Where("transaction.seller_id = ?", userId).
			WhereOr("transaction.buyer_id = ?", userId).
			Select()
	
		if err != nil {
			logrus.WithError(err)
			return []domain.Transaction{}, nil 
		}

		return transacs, nil
	}

}
