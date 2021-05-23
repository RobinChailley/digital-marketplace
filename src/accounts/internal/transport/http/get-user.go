package http

import (
	"marketplace/accounts/domain"
	"marketplace/accounts/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

type GetUserResponse struct {
	Id 				int64  					`json:"id"`
	Email 		string 					`json:"email"`
	Username 	string 					`json:"username"`
	Password 	string 					`json:"-"`
	Balance 	float64  					`json:"-"`
	Ads				[]domain.Ads   	`json:"ads"`
}

func GetUserHandler(db *pg.DB, cmd usecase.GetUserCmd) gin.HandlerFunc {
	return func (c *gin.Context) {
		userEmail := c.Param("email")

		if userEmail == "" {
			logrus.Error("The field 'email' is required.")
			c.Status(http.StatusBadRequest)
			return
		}

		user, err := cmd(db, c, userEmail)

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