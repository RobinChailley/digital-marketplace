package http

import (
	"marketplace/transactions/domain"
	"marketplace/transactions/internal/request"
	"marketplace/transactions/internal/usecase"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

func AcceptDeclineTransactionHandler(db *pg.DB, cmd usecase.AcceptDeclineTransactionCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("acc").(domain.Account)
		transacId := c.Param("id")

		intTransacId, err := strconv.ParseInt(transacId, 10, 64)

		if err != nil {
			logrus.WithError(err).Error("Bad request. The 'id' field must be an integer.")
			c.Status(http.StatusBadRequest)
			return
		}

		transac, err := cmd(db, c, intTransacId, user.Id)

		if err != nil {
			if err.Error() == "not found" {
				logrus.WithError(err)
				c.Status(http.StatusNotFound)
				return
			} else if strings.Contains(err.Error(), "this transaction is already") {
				logrus.WithError(err)
				c.JSON(http.StatusBadRequest, err.Error())
				return
			}
			logrus.WithError(err).Error("An error has occured.")
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, request.ConvertToResponse([]domain.Transaction{transac}))
	}
}
