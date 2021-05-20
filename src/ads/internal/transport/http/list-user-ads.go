package http

import (
	"marketplace/ads/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)


func ListUserAdsHandler(db *pg.DB, cmd usecase.ListUserAdsCmd) gin.HandlerFunc {
	return func (c *gin.Context) {
		userIdStr := c.Param("id")
		userIdInt, err := strconv.ParseInt(userIdStr, 10, 64)

		if err != nil {
			logrus.WithError(err).Error("Bad request. Id must be an integer.")
			c.Status(http.StatusBadRequest)
			return
		}

		ads, err := cmd(db, userIdInt)

		if err != nil {
			logrus.WithError(err).Error("Internal server error")
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, ads)
	}
}
