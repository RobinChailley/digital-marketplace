package http

import (
	"marketplace/accounts/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

type UpdateBalanceRequest struct {
	Balance 		float64  				`json:"balance"`
}

func UpdateBalanceByIdHandler(db *pg.DB, cmd usecase.UpdateBalanceByIdCmd) gin.HandlerFunc {
	return func (c *gin.Context) {
		userId := c.Param("id")
		intUserId, err := strconv.ParseInt(userId, 10, 64)

		if err != nil {
			logrus.Error("Bad request : The user id must be an integer")
			c.Status(http.StatusBadRequest)
			return
		}

		updateBalanceRequest := &UpdateBalanceRequest{}
		err = c.BindJSON(&updateBalanceRequest)

		if err != nil {
			logrus.Error("Bad request : The json body is not well formated.")
			c.Status(http.StatusBadRequest)
			return
		}


		err = cmd(db, intUserId, updateBalanceRequest.Balance)

		if err != nil {
			if err.Error() == "Balance is too low." {
				logrus.WithError(err).Error("Balance is too low.")
				c.Status(http.StatusUnauthorized)
				return
			} else if err.Error() == "not found" {
				logrus.WithError(err).Error("not found")
				c.Status(http.StatusNotFound)
				return
			}
			logrus.WithError(err).Error(err)
			c.Status(http.StatusInternalServerError)
		}

		c.Status(http.StatusOK)
	}
}