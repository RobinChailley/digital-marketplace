package http

import (
	"net/http"
	"strings"

	"marketplace/transactions/domain"
	"marketplace/transactions/internal/conf"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func AuthMiddleware(db *pg.DB, config conf.Configuration) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		tokenSplitted := strings.Split(token, "Bearer ")

		if token == "" {
			c.Status(http.StatusUnauthorized)
			c.Abort()
		}

		claims := &usecase.Claims{}

		_, err := jwt.ParseWithClaims(tokenSplitted[1], claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Credentials.JwtSecret), nil
		})

		if err != nil {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		var accs []domain.Account;
		db.Model(&accs).Where("account.id = ? ", claims.Id).Select()

		if len(accs) < 1 {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Set("acc", accs[0])
	}
}