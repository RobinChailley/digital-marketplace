package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"marketplace/accounts/domain"
	"marketplace/accounts/internal/usecase"
	"net/http"
)

type AddFundsRequest struct {
	Funds int64 `json:"funds"`
}

func AddFundsHandler(db *pg.DB, cmd usecase.AddFundsCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
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
