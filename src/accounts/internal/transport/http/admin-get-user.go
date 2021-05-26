package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"marketplace/accounts/internal/usecase"
	"net/http"
	"strconv"
)

func AdminGetUserHandler(db *pg.DB, cmd usecase.AdminGetUserCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		intId, err := strconv.ParseInt(id, 10, 64)

		if err != nil {
			logrus.WithError(err).Error("Bad request. The 'id' field must be an integer")
			c.Status(http.StatusBadRequest)
			return
		}

		user, err := cmd(db, intId)

		if err != nil {
			if err.Error() == "not found" {
				logrus.WithError(err)
				c.Status(http.StatusNotFound)
				return
			}
			logrus.WithError(err)
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
