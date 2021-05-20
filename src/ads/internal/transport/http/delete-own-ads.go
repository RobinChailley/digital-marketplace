package http

import (
	"marketplace/ads/domain"
	"marketplace/ads/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)


func DeleteOwnAdsHandler(db *pg.DB, cmd usecase.DeleteOwnAdsCmd) gin.HandlerFunc {
	return func (c *gin.Context) {
		user := c.MustGet("acc").(domain.Account)
		adsId := c.Param("id")

		intAdsId, err := strconv.Atoi(adsId)

		if err != nil {
			logrus.WithError(err).Error("Ads'id must be a integer.")
			c.Status(http.StatusBadRequest)
			return
		}

		err = cmd(db, &user, intAdsId)

		if err != nil {
			if err.Error() == "Not found" {
				logrus.WithError(err)
				c.Status(http.StatusNotFound)
			} else if err.Error() == "Unauthorized" {
				logrus.WithError(err)
				c.Status(http.StatusUnauthorized)
			} else {
				logrus.WithError(err).Error("An error has occured.")
				c.Status(http.StatusInternalServerError)
			}
			return
		}

		c.Status(http.StatusOK)
	}
}
