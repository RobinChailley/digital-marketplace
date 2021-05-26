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

func GetMyTransactionsHandler(db *pg.DB, cmd usecase.GetMyTransactionsCmd) gin.HandlerFunc {
	return func (c *gin.Context) {
		user := c.MustGet("acc").(domain.Account)
		transacs, err := cmd(db, user.Id)

		if err != nil {
			logrus.WithError(err).Error("An error has occured.")
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, request.ConvertToResponse(transacs))
	}
}
