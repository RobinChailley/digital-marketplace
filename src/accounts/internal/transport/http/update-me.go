package http

import (
	"github.com/go-pg/pg/v10"
	"marketplace/accounts/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"marketplace/accounts/internal/request"
	"marketplace/accounts/domain"
	"net/http"
)



func UpdateMeHandler(db *pg.DB, cmd usecase.UpdateMeCmd) gin.HandlerFunc {
	return func (c *gin.Context) {
		updateMeRequest := &request.UpdateMeRequest{}
		err := c.BindJSON(updateMeRequest)

		if err != nil {
			logrus.WithError(err).Error("Bad request. Datas are not well formated.")
			c.Status(http.StatusBadRequest)
			return
		}

		user := c.MustGet("acc").(domain.Account)


		updateMeResponse, err := cmd(db, &user, updateMeRequest)

		if err != nil {
			logrus.WithError(err).Error()
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusCreated, updateMeResponse)
	}
}