package http

import (
	"net/http"
	"strings"

	"marketplace/ads/domain"
	"marketplace/ads/internal/conf"
	"marketplace/ads/internal/request"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func AuthAdminMiddleware(db *pg.DB, config conf.Configuration) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		tokenSplitted := strings.Split(token, "Bearer ")

		if token == "" {
			c.Status(http.StatusUnauthorized)
			c.Abort()
		}

		if len(tokenSplitted) != 2 {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		claims := &request.Claims{}

		_, err := jwt.ParseWithClaims(tokenSplitted[1], claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Credentials.JwtSecret), nil
		})

		if err != nil {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		var accs []domain.Account
		db.Model(&accs).Where("account.id = ? ", claims.Id).Select()

		if len(accs) < 1 {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		if accs[0].Admin != true {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Set("acc", accs[0])
	}
}
