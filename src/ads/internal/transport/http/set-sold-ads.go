package http

import (
	"marketplace/ads/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)


func SetSoldAdsHandler(db *pg.DB, cmd usecase.SetSoldAdsCmd) gin.HandlerFunc {
	return func (c *gin.Context) {
		id := c.Param("id")
		intId, err := strconv.ParseInt(id, 10, 64)

		if err != nil {
			logrus.WithError(err).Error("Bad request : The id must be an integer")
			c.Status(http.StatusBadRequest)
			return			
		}

		ads, err := cmd(db, intId)

		if err != nil {
			logrus.WithError(err).Error("Internal server error")
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, ads)
	}
}
