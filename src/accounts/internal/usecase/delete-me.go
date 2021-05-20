package usecase

import (
	"marketplace/accounts/domain"
	"marketplace/accounts/internal/infrastructure/ads"
	"github.com/go-pg/pg/v10"
	"github.com/gin-gonic/gin"
)

type DeleteMeCmd func (db *pg.DB, c *gin.Context, acc *domain.Account) error

func DeleteMe(adsFetcher ads.Fetcher) DeleteMeCmd {

	return func (db *pg.DB, c *gin.Context, acc *domain.Account) (error) {

		err := adsFetcher.DeleteAllMyAds(c)

		_, err = db.Model(acc).Where("id = ?", acc.Id).Delete()


		if err != nil {
			return err
		}

		return nil
	}
}