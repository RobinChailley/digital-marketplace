package http

import (
	"marketplace/ads/domain"
	"marketplace/ads/internal/request"
	"marketplace/ads/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

func UpdateMyAdsHandler(db *pg.DB, cmd usecase.UpdateMyAdsCmd) gin.HandlerFunc {
	return func (c *gin.Context) {
		user := c.MustGet("acc").(domain.Account)
		adsId := c.Param("id")

		updateAdsRequest := &request.UpdateAdsRequest{}
		err := c.BindJSON(updateAdsRequest)

		if err != nil {
			logrus.WithError(err).Error("Bad request. Data are not well formated.")
			c.Status(http.StatusBadRequest)
			return
		}

		intAdsId, err := strconv.ParseInt(adsId, 10, 64)

		if err != nil {
			logrus.WithError(err).Error("Bad request. The 'id' field must be an integer.")
			c.Status(http.StatusBadRequest)
			return
		}

		updatedAds, err := cmd(db, updateAdsRequest, intAdsId, user.Id)

		if err != nil {
			if err.Error() == "Not found" {
				logrus.WithError(err)
				c.Status(http.StatusNotFound)
			} else if err.Error() == "Unauthorized" {
				logrus.WithError(err)
				c.Status(http.StatusUnauthorized)
			} else {
				logrus.WithError(err).Error("Bad request. Data are not well formated.")
				c.Status(http.StatusInternalServerError)
			}
			return
		}

		c.JSON(http.StatusOK, updatedAds)

	}
}