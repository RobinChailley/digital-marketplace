package usecase

import (
	"marketplace/accounts/domain"
	"marketplace/accounts/internal/infrastructure/ads"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

type GetUserCmd func (db *pg.DB, c *gin.Context, email string) (*domain.Account, error)

func GetUser(adsFetcher ads.Fetcher) GetUserCmd {

	return func (db *pg.DB, c *gin.Context, email string) (*domain.Account, error) {

		var acc domain.Account;
		err := db.Model(&acc).Where("account.email = ? ", email).Select()

		if acc.Password == "" && acc.Email == "" && acc.Balance == 0 && acc.Id == 0 {
			return nil, nil
		}

		if err != nil {
			return nil, err
		}

		var adss []domain.Ads

		adss, err = adsFetcher.GetMyAds(c)

		if err != nil {
			return nil, err
		}

		acc.Ads = adss

		acc.Balance = -1

		return &acc, nil
	}
}