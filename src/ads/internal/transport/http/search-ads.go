package http

import (
	"marketplace/ads/domain"
	"marketplace/ads/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)


func SearchAdsHandler(db *pg.DB, cmd usecase.SearchAdsCmd) gin.HandlerFunc {
	return func (c *gin.Context) {
		user := c.MustGet("acc").(domain.Account)
		keyword := c.Query("keyword")

		ads, err := cmd(db, &user, keyword)

		if err != nil {
			logrus.WithError(err).Error("Internal server error")
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, ads)
	}
}
