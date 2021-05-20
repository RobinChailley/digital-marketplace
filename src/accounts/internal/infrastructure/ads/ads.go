package ads

import (
	"marketplace/accounts/domain"
	"github.com/gin-gonic/gin"
)

type Fetcher interface {
	GetMyAds(c *gin.Context) ([]domain.Ads, error)
	DeleteAllMyAds(c *gin.Context) (error)
}
