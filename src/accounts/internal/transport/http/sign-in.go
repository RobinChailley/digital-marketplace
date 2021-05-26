package http

import (
	"marketplace/accounts/domain"
	"marketplace/accounts/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

type SignInRequest struct {
	Password string         `json:"password"`
	Email 	 string			`json:"email"`
}

type SignInResponse struct {
	JwtToken string   		`json:"jwttoken"`
}

func SignInHandler(db *pg.DB, cmd usecase.SignInCmd) gin.HandlerFunc {
	return func (c *gin.Context) {
		signInReq := &SignInRequest{}
		err := c.BindJSON(signInReq)
		if err != nil {
			logrus.WithError(err).Error("Bad request. Datas are not well formated.")
			c.Status(http.StatusBadRequest)
			return
		}

		tokenString, err := cmd(db, &domain.Account{
			Email: signInReq.Email,
			Password: signInReq.Password,
		})

		if err != nil {
			logrus.WithError(err).Error()
			c.Status(http.StatusUnauthorized)
			return
		}

		if tokenString == "" {
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, SignInResponse{JwtToken: tokenString})
	}
}