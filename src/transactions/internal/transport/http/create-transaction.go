package http

import (
	"marketplace/transactions/domain"
	"marketplace/transactions/internal/request"
	"marketplace/transactions/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

func CreateTransactionHandler(db *pg.DB, cmd usecase.CreateTransactionCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		createTransaction := &request.CreateTransactionRequest{}
		err := c.BindJSON(createTransaction)
		user := c.MustGet("acc").(domain.Account)

		if err != nil {
			logrus.WithError(err).Error("Bad request. Data are not well formated.")
			c.Status(http.StatusBadRequest)
			return
		}

		transac, err := cmd(db, c, createTransaction, user)

		if err != nil {
			if err.Error() == "Not found" {
				logrus.WithError(err).Error("The ads is not found.")
				c.Status(http.StatusNotFound)
				return
			} else if err.Error() == "You can not make an offer for an add that you created." {
				logrus.WithError(err).Error("You can not make an offer for an add that you created.")
				c.Status(http.StatusBadRequest)
				return
			} else if err.Error() == "too expensive" {
				logrus.WithError(err).Error("You can not make this transaction because you don't have enough money on your account")
				c.Status(http.StatusUnauthorized)
				return
			}
			logrus.WithError(err).Error("Bad request. Data are not well formated.")
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusCreated, request.ConvertToResponse([]domain.Transaction{transac}))
	}
}
