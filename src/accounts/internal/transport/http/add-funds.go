package http

import (
	"marketplace/accounts/domain"
	"github.com/go-pg/pg/v10"
	"marketplace/accounts/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type AddFundsRequest struct {
	Funds int64  		`json:"funds"`
}

func AddFundsHandler(db *pg.DB, cmd usecase.AddFundsCmd) gin.HandlerFunc {
	return func (c *gin.Context) {
		addFundsRequest := &AddFundsRequest{}
		err := c.BindJSON(addFundsRequest)

		if err != nil {
			logrus.WithError(err).Error("Bad request. Data are not well formated.")
			c.Status(http.StatusBadRequest)
			return
		}

		user := c.MustGet("acc").(domain.Account)

		balance, err := cmd(db, &user, addFundsRequest.Funds)

		if err != nil {
			logrus.WithError(err).Error()
			c.Status(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusCreated, balance)
	}
}