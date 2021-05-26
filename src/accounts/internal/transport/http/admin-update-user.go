package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"marketplace/accounts/internal/request"
	"marketplace/accounts/internal/usecase"
	"net/http"
	"strconv"
)

func AdminUpdateUserHandler(db *pg.DB, cmd usecase.AdminUpdateUserCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		intId, err := strconv.ParseInt(id, 10, 64)

		if err != nil {
			logrus.WithError(err).Error("Bad request : The 'id' field must be an integer")
			c.Status(http.StatusBadRequest)
			return
		}

		updateReq := request.UpdateUserRequest{Email: "", Username: "", Password: "", Balance: 0, Admin: false}
		err = c.BindJSON(&updateReq)

		if err != nil {
			logrus.WithError(err).Error("Bad request : The body is not well formated")
			c.Status(http.StatusBadRequest)
			return
		}

		newUser, err := cmd(db, updateReq, intId)

		if err != nil {
			if err.Error() == "not found" {
				logrus.WithError(err).Error("not found")
				c.Status(http.StatusNotFound)
				return
			}
			logrus.WithError(err).Error("internal server error")
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, newUser)
	}
}
