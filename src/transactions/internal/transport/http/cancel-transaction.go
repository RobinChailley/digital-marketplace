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

func CancelTransactionHandler(db *pg.DB, cmd usecase.CancelTransactionCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("acc").(domain.Account)
		transacId := c.Param("id")
		intTransacId, err := strconv.ParseInt(transacId, 10, 64)

		if err != nil {
			logrus.WithError(err).Error("Bad request. The 'id' field must be an integer.")
			c.Status(http.StatusBadRequest)
			return
		}

		transac, err := cmd(db, intTransacId, user.Id)

		if err != nil {
			if err.Error() == "not found" {
				logrus.WithError(err)
				c.Status(http.StatusNotFound)
				return
			} else if err.Error() == "unauthorized" {
				logrus.WithError(err).Error("You can not cancel this transaction because you are not the seller of it.")
				c.Status(http.StatusUnauthorized)
				return
			}
			logrus.WithError(err).Error("An error has occured.")
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, request.ConvertToResponse([]domain.Transaction{transac}))
	}
}
