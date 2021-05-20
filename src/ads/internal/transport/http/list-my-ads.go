package http

import (
	"marketplace/ads/domain"
	"marketplace/ads/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)


func ListMyAdsHandler(db *pg.DB, cmd usecase.ListUserAdsCmd) gin.HandlerFunc {
	return func (c *gin.Context) {
		user := c.MustGet("acc").(domain.Account)

		ads, err := cmd(db, user.Id)

		if err != nil {
			logrus.WithError(err).Error("Internal server error")
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, ads)
	}
}
