package usecase

import (
	"marketplace/ads/domain"
	"github.com/go-pg/pg/v10"
)

type ListUserAdsCmd func (db *pg.DB, userId int64) ([]domain.Ads, error)

func ListUserAds() ListUserAdsCmd {
	return func (db *pg.DB, userId int64) ([]domain.Ads, error) {

		var adsArray []domain.Ads

		err := db.Model(&adsArray).
			Where("ads.User_Id = ?", userId).
			Select()

		if err != nil {
			return []domain.Ads{}, err
		}

		return adsArray, nil
	}
}
