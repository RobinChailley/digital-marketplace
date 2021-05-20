package usecase

import (
	"fmt"
	"marketplace/ads/domain"
	"github.com/go-pg/pg/v10"
)

type DeleteAllMyAdsCmd func (db *pg.DB, user *domain.Account) (error)

func DeleteAllMyAds() DeleteAllMyAdsCmd {

	return func (db *pg.DB, user *domain.Account) (error) {
		ads := domain.Ads{}

		res, err := db.Model(&ads).Where("ads.user_id = ?", user.Id).Delete()

		fmt.Println("res : ", res)
		fmt.Println("err : ", err)

		return nil
	}

}
