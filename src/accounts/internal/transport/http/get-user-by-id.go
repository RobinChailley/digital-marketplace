package http

import (
	"marketplace/accounts/domain"
	"marketplace/accounts/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

type GetUserByIdResponse struct {
	Id       int64        `json:"id"`
	Email    string       `json:"email"`
	Username string       `json:"username"`
	Password string       `json:"-"`
	Balance  float64      `json:"-"`
	Ads      []domain.Ads `json:"ads"`
	Admin    bool         `json:"admin"`
}

func GetUserByIdHandler(db *pg.DB, cmd usecase.GetUserByIdCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("id")
		intUserId, err := strconv.ParseInt(userId, 10, 64)

		if err != nil {
			logrus.WithError(err).Error("Bad request : The id param must be an integer")
			c.Status(http.StatusBadRequest)
			return
		}

		user, err := cmd(db, c, intUserId)

		if user == nil {
			logrus.Error("This account can not be found.")
			c.Status(http.StatusNotFound)
			return
		}

		if err != nil {
			logrus.WithError(err).Error("An error has occured.")
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, GetUserResponse(*user))
	}
}
