package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"marketplace/accounts/internal/usecase"
	"net/http"
)

func AdminGetAllUsersHandler(db *pg.DB, cmd usecase.AdminGetAllUsersCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := cmd(db)

		if err != nil {
			logrus.WithError(err)
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, users)
	}
}
