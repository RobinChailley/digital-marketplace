package accounts

import (
	"marketplace/transactions/domain"

	"github.com/gin-gonic/gin"
)

type Fetcher interface {
	GetUserById(c *gin.Context, userId int64) (domain.Account, error)
	UpdateUserBalanceById(c *gin.Context, userId int64, deltaBalance float64) (error)
}
