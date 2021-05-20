package http

import (
	"marketplace/accounts/domain"
	"marketplace/accounts/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

func DeleteMeHandler(db *pg.DB, cmd usecase.DeleteMeCmd) gin.HandlerFunc {
	return func (c *gin.Context) {
		user := c.MustGet("acc").(domain.Account)

		err := cmd(db, c, &user)

		if err != nil {
			logrus.WithError(err).Error()
			c.Status(http.StatusInternalServerError)
			return
		}

		c.Status(http.StatusOK)
	}
}