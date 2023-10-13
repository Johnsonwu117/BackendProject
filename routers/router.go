package routers

import (
	dbtest "backendProject/db"
	"backendProject/resolver"
	"github.com/gin-gonic/gin"
	"strconv"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		input := &dbtest.InputCoin{}

		input.Source = c.Query("source")
		input.Target = c.Query("target")
		amountStr := c.Query("amount")
		input.Amount, _ = strconv.Atoi(amountStr)

		output := resolver.ChangeCoin(input)
		c.JSON(200, gin.H{
			"success": true,
			"data":    output,
		})
	})

	return router
}
