package usecase

import (
	"errors"
	"marketplace/ads/domain"
	"github.com/go-pg/pg/v10"
)

type DeleteOwnAdsCmd func (db *pg.DB, user *domain.Account, adsId int) (error)

func DeleteOwnAds() DeleteOwnAdsCmd {

	return func (db *pg.DB, user *domain.Account, adsId int) (error) {

		ads := &domain.Ads{Id: int64(adsId)}

		err := db.Model(ads).WherePK().Select()

		if *ads == (domain.Ads{Id: int64(adsId)}) {
			return errors.New("Not found")
		}

		if err != nil {
			return err
		}

		if ads.UserId != user.Id {
			return errors.New("Unauthorized")
		}

		_, err = db.Model(ads).WherePK().Delete()

		if err != nil {
			return err
		}

		return nil
	}

}
