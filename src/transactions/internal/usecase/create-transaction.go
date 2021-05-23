package usecase

import (
	"errors"
	"marketplace/transactions/domain"
	"marketplace/transactions/internal/infrastructure/accounts"
	"marketplace/transactions/internal/infrastructure/ads"
	"marketplace/transactions/internal/request"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

type CreateTransactionCmd func (db *pg.DB, c *gin.Context, createTransaction *request.CreateTransactionRequest, user domain.Account) (domain.Transaction, error)

func CreateTransaction(adsFetcher ads.Fetcher, accountsFetcher accounts.Fetcher) CreateTransactionCmd {

	return func (db *pg.DB, c *gin.Context, createTransaction *request.CreateTransactionRequest, user domain.Account) (domain.Transaction, error) {

		ads, err := adsFetcher.GetAdsById(c, createTransaction.AdsId)
	
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				return domain.Transaction{}, errors.New("Not found")
			}
			return domain.Transaction{}, err
		}


		seller, err := accountsFetcher.GetUserById(c, ads.UserId)

		if err != nil {
			logrus.WithError(err).Error("Can not use the accountsFetcher to GetUserById")
			return domain.Transaction{}, err
		}

		if seller.Id == user.Id {
			return domain.Transaction{}, errors.New("You can not make an offer for an add that you created.")
		}

		transaction := domain.Transaction{
			Buyer: &user,
			BuyerId: user.Id,
			Seller: &seller,
			SellerId: seller.Id,
			Ads: &ads,
			AdsId: ads.Id,
			Messages: []domain.Message{},
			Bid: createTransaction.Bid,
			Status: "PROPOSITION",
		}

		_, err = db.Model(&transaction).Insert()

		if err != nil {
			return domain.Transaction{}, nil
		}

		return transaction, nil
	}

}
