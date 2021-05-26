package http

import (
	"marketplace/transactions/domain"
	"marketplace/transactions/internal/request"
	"marketplace/transactions/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)


func PostMessageOnTransactionHandler(db *pg.DB, cmd usecase.PostMessageOnTransactionCmd) gin.HandlerFunc {
	return func (c *gin.Context) {
		messageRequest := &request.MessageRequest{}
		transacId := c.Param("id")
		intTransacId, err := strconv.ParseInt(transacId, 10, 64)

		if err != nil {
			logrus.WithError(err).Error("Bad request. The 'id' field must be an integer")
			c.Status(http.StatusBadRequest)
			return
		}

		err = c.BindJSON(messageRequest)

		if err != nil {
			logrus.WithError(err).Error("Bad request. Data are not well formated.")
			c.Status(http.StatusBadRequest)
			return
		}

		user := c.MustGet("acc").(domain.Account)
		messageRequest.TransactionId = intTransacId


		transac, err := cmd(db, messageRequest, user)

		if err != nil {
			if err.Error() == "Not found" {
				logrus.WithError(err).Error("The ads is not found.")
				c.Status(http.StatusNotFound)
				return
			} else if err.Error() == "You can not make an offer for an add that you created." {
				logrus.WithError(err).Error("You can not make an offer for an add that you created.")
				c.Status(http.StatusBadRequest)
				return
			} else if err.Error() == "unauthorized : this transaction is not yours" {
				logrus.WithError(err).Error("unauthorized : this transaction is not yours")
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
