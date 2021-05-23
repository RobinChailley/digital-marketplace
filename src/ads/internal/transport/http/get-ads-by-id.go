package http

import (
	"marketplace/ads/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)


func GetAdsByIdHandler(db *pg.DB, cmd usecase.GetAdsByIdCmd) gin.HandlerFunc {
	return func (c *gin.Context) {
		adsId := c.Param("id")
		intAdsId, err := strconv.ParseInt(adsId, 10, 64)

		if err != nil {
			logrus.WithError(err).Error("Bad request : The id param must be an integer")
			c.Status(http.StatusBadRequest)
			return
		}		

		ads, err := cmd(db, intAdsId)

		if err != nil {
			if err.Error() == "Not found" {
				logrus.WithError(err).Error("Not found")
				c.Status(http.StatusNotFound)
				return
			}
			logrus.WithError(err).Error("Internal server error")
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, ads)
	}
}
