package usecase

import (
	"marketplace/ads/domain"
	"github.com/go-pg/pg/v10"
)

type SearchAdsCmd func (db *pg.DB, user *domain.Account, keyword string) ([]domain.Ads, error)

func SearchAds() SearchAdsCmd {
	return func (db *pg.DB, user *domain.Account, keyword string) ([]domain.Ads, error) {

		var adsArray []domain.Ads

		err := db.Model(&adsArray).
			Where("ads.Title LIKE ?", "%" + keyword + "%").
			WhereOr("ads.Description LIKE ?", "%" + keyword + "%").
			Select()

		if err != nil {
			return []domain.Ads{}, err
		}

		return adsArray, nil
	}
}
