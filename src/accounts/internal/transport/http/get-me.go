package http

import (
	"marketplace/accounts/domain"
	"marketplace/accounts/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

type GetMeResponse struct {
	Id 			int64  			`json:"id"`
	Email 		string 			`json:"email"`
	Username 	string 			`json:"username"`
	Password 	string 			`json:"-"`
	Balance 	float64  			`json:"balance"`
	Ads			[]domain.Ads   	`json:"ads"`
}

func GetMeHandler(db *pg.DB, cmd usecase.GetMeCmd) gin.HandlerFunc {
	return func (c *gin.Context) {
		user := c.MustGet("acc").(domain.Account)


		fullUser, err := cmd(db, c, &user)

		if err != nil {

			logrus.WithError(err).Error("An error has occured")
			c.Status(http.StatusInternalServerError)
			return
		}

		getMeResponse := GetMeResponse(*fullUser)

		c.JSON(http.StatusOK, getMeResponse)
	}
}