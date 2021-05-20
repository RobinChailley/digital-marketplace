package accounts

import (
	"github.com/gin-gonic/gin"
)

type Fetcher interface {
	Test(c *gin.Context) (error)
}
