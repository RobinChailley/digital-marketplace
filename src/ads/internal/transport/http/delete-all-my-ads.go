package http

import (
	"github.com/go-pg/pg/v10"
	"marketplace/ads/internal/usecase"
	"marketplace/ads/domain"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)


func DeleteAllMyAdsHandler(db *pg.DB, cmd usecase.DeleteAllMyAdsCmd) gin.HandlerFunc {
	return func (c *gin.Context) {
		user := c.MustGet("acc").(domain.Account)

		err := cmd(db, &user)

		if err != nil {
			logrus.WithError(err).Error("An error has occured.")
			c.Status(http.StatusInternalServerError)
			return
		}

		c.Status(http.StatusOK)
	}
}
