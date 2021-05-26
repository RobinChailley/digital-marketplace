package usecase

import (
	"errors"
	"fmt"
	"marketplace/transactions/domain"
	"marketplace/transactions/internal/infrastructure/accounts"
	"marketplace/transactions/internal/infrastructure/ads"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

type AcceptDeclineTransactionCmd func(db *pg.DB, c *gin.Context, transacId int64, userId int64) (domain.Transaction, error)

func AcceptDeclineTransaction(acceptOrDecline string, adsFetcher ads.Fetcher, accountsFetcher accounts.Fetcher) AcceptDeclineTransactionCmd {

	return func(db *pg.DB, c *gin.Context, transacId int64, userId int64) (domain.Transaction, error) {
		var transac domain.Transaction

		err := db.Model(&transac).
			Column("transaction.*").
			Relation("Buyer").
			Relation("Ads").
			Relation("Seller").
			Relation("Messages").
			Where("transaction.seller_id = ?", userId).
			Where("transaction.id = ?", transacId).
			Select()

		if transac.Status != "PROPOSITION" {
			return domain.Transaction{}, errors.New(fmt.Sprintf("this transaction is already '%s'", transac.Status))
		}

		if transac.Id == 0 {
			return domain.Transaction{}, errors.New("not found")
		}

		if err != nil {
			logrus.WithError(err)
			return domain.Transaction{}, err
		}

		if acceptOrDecline == "accept" {
			err = accountsFetcher.UpdateUserBalanceById(c, transac.BuyerId, -transac.Bid)
			if err != nil {
				logrus.Error("The buyer does not have enough money to buy your product.")
				return domain.Transaction{}, errors.New("the buyer does not have enough money to buy your product")
			}
		}

		if acceptOrDecline == "accept" {
			transac.Status = "ACCEPTED"
		} else {
			transac.Status = "DECLINED"
		}

		_, err = db.Model(&transac).WherePK().Update()

		if err != nil {
			logrus.WithError(err)
			return domain.Transaction{}, err
		}

		if acceptOrDecline != "accept" {
			return transac, nil
		}

		err = accountsFetcher.UpdateUserBalanceById(c, transac.SellerId, transac.Bid)

		if err != nil {
			logrus.Error("Can not add money on the seller accounts")
			return domain.Transaction{}, errors.New("can not add money on the seller accounts")
		}

		soldAds, err := adsFetcher.SetSoldToAds(c, transac.AdsId)

		if err != nil {
			logrus.Error(fmt.Sprintf("Can not set this ads sold (id: %d)", transac.AdsId))
			return transac, errors.New(fmt.Sprintf("can not set this ads sold (id: %d)", transac.AdsId))
		}

		transac.Ads = &soldAds

		return transac, nil
	}

}
