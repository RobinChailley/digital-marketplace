package usecase

import (
	"errors"
	"marketplace/ads/domain"
	"marketplace/ads/internal/request"
	"github.com/go-pg/pg/v10"
)

type UpdateMyAdsCmd func (db *pg.DB, updateAds *request.UpdateAdsRequest, adsId int64, userId int64) (*domain.Ads, error)

func UpdateMyAds() UpdateMyAdsCmd {

	return func (db *pg.DB, updateAds *request.UpdateAdsRequest, adsId int64, userId int64) (*domain.Ads, error) {
		ads := &domain.Ads{Id: adsId}

		err := db.Model(ads).WherePK().Select()

		if err != nil {
			return nil, errors.New("Not found")
		}

		if ads == (&domain.Ads{Id: adsId}) {
			return nil, errors.New("Not found")
		}

		if ads.UserId != userId {
			return nil, errors.New("Unauthorized")
		}

		if updateAds.Title != "" {
			ads.Title = updateAds.Title
		}

		if updateAds.Description != "" {
			ads.Description = updateAds.Description
		}

		if updateAds.Price != 0 {
			ads.Price = updateAds.Price
		}

		if updateAds.Picture != "" {
			ads.Picture = updateAds.Picture
		}

		_, err = db.Model(ads).WherePK().Update()

		if err != nil {
			return nil, err
		}

		return ads, nil
	}
}