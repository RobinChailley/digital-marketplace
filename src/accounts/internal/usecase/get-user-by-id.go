package usecase

import (
	"fmt"
	"marketplace/accounts/domain"
	"marketplace/accounts/internal/infrastructure/ads"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

type GetUserByIdCmd func (db *pg.DB, c *gin.Context, userId int64) (*domain.Account, error)

func GetUserById(adsFetcher ads.Fetcher) GetUserByIdCmd {

	return func (db *pg.DB, c *gin.Context, userId int64) (*domain.Account, error) {

		acc := domain.Account{Id: userId}
		err := db.Model(&acc).WherePK().Select()

		fmt.Println(" acc: ", acc)

		if acc.Password == "" && acc.Email == "" && acc.Id == 0 {
			logrus.Error("acc.Password == \"\" && acc.Email == \"\" && acc.Id == 0")
			return nil, nil
		}

		if err != nil {
			logrus.WithError(err).Error("db.Model(&acc).WherePK().Select()")
			return nil, err
		}

		var adss []domain.Ads

		adss, err = adsFetcher.GetMyAds(c)

		if err != nil {
			logrus.WithError(err).Error("adss, err = adsFetcher.GetMyAds(c)")
			return nil, err
		}

		acc.Ads = adss

		return &acc, nil
	}
}