package usecase

import (
	"marketplace/ads/domain"
	"marketplace/ads/internal/request"

	"github.com/go-pg/pg/v10"
)

type CreateAdsCmd func (db *pg.DB, createAds *request.UpdateAdsRequest, userId int64) (*domain.Ads, error)

func CreateAds() CreateAdsCmd {

	return func (db *pg.DB, createAds *request.UpdateAdsRequest, userId int64) (*domain.Ads, error) {

		ads := &domain.Ads{}

		ads.Title = createAds.Title
		ads.Description = createAds.Description
		ads.Price = createAds.Price
		ads.UserId = userId
		ads.Picture = createAds.Picture

		_, err := db.Model(ads).Insert()

		if err != nil {
			return nil, err
		}

		return ads, nil
	}

}
