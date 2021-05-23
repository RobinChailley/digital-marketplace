package usecase

import (
	"errors"
	"marketplace/ads/domain"

	"github.com/go-pg/pg/v10"
)

type GetAdsByIdCmd func (db *pg.DB, adsId int64) (domain.Ads, error)

func GetAdsById() GetAdsByIdCmd {
	return func (db *pg.DB, adsId int64) (domain.Ads, error) {

		var ads domain.Ads

		err := db.Model(&ads).
			Where("ads.Id = ?", adsId).
			Select()

		if ads == (domain.Ads{}) {
			return domain.Ads{}, errors.New("Not found")
		}

		if err != nil {
			return domain.Ads{}, err
		}

		return ads, nil
	}
}
