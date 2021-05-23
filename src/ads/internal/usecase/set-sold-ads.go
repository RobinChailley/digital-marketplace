package usecase

import (
	"errors"
	"marketplace/ads/domain"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

type SetSoldAdsCmd func (db *pg.DB, adsId int64) (domain.Ads, error)

func SetSoldAds() SetSoldAdsCmd {
	return func (db *pg.DB, adsId int64) (domain.Ads, error) {
		var ads domain.Ads
		err := db.Model(&ads).
			Where("ads.Id = ?", adsId).
			Select()

		if err != nil {
			return domain.Ads{}, err
		}

		ads.Sold = true

		_, err = db.Model(&ads).WherePK().Update()

		if err != nil {
			logrus.Error("Can not save the updated ads.")
			return domain.Ads{}, errors.New("can not save the updated ads")
		}

		return ads, nil
	}
}
