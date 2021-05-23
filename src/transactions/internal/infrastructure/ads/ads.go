package ads

import (
	"marketplace/transactions/domain"

	"github.com/gin-gonic/gin"
)

type Fetcher interface {
	GetAdsById(c *gin.Context, adsId int64) (domain.Ads, error)
	SetSoldToAds(c *gin.Context, adsId int64) (domain.Ads, error)
}
