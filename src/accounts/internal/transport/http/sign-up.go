package http

import (
	"marketplace/accounts/domain"
	"marketplace/accounts/internal/usecase"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type SignUpRequest struct {
	Email    string         `json:"email"`
	Password string         `json:"password"`
	Username string					`json:"username"`
}

type SignUpResponse struct {
	Id 			int64 					`json:"id"`
}

func SignUpHandler(db *pg.DB, cmd usecase.SignUpCmd) gin.HandlerFunc {
	return func (c *gin.Context) {
		signUpReq := &SignUpRequest{}
		err := c.BindJSON(signUpReq)
		if err != nil {
			logrus.WithError(err).Error("Bad request. Datas are not well formated.")
			c.Status(http.StatusBadRequest)
			return
		}

		hashedPassword, err:= bcrypt.GenerateFromPassword([]byte(signUpReq.Password), bcrypt.MinCost)

		if err != nil {
			logrus.WithError(err).Error()
			c.Status(http.StatusInternalServerError)
			return
		}

		accountID, err := cmd(db, &domain.Account{
			Email: signUpReq.Email,
			Username: signUpReq.Username,
			Password: string(hashedPassword),
		})

		if err != nil {
			if strings.Contains(err.Error(), "ERROR #23505 duplicate key value violates unique constraint") == true {
				c.Status(http.StatusConflict)
				return
			}
			logrus.WithError(err).Error()
			c.Status(http.StatusInternalServerError)
			return
		}

		if accountID == -1 {
			logrus.WithError(err).Error()
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusCreated, SignUpResponse{Id: accountID})
	}
}