package usecase

import (
	"marketplace/accounts/domain"
	"marketplace/accounts/internal/infrastructure/ads"
	"github.com/go-pg/pg/v10"
	"github.com/gin-gonic/gin"
)

type GetMeCmd func (db *pg.DB, c *gin.Context, user *domain.Account) (*domain.Account, error)

func GetMe(adsFetcher ads.Fetcher) GetMeCmd {

	return func (db *pg.DB, c *gin.Context, user *domain.Account) (*domain.Account, error) {

		var adss []domain.Ads

		adss, err := adsFetcher.GetMyAds(c)

		if err != nil {
			return nil, err
		}

		user.Ads = adss

		return user, nil
	}
}