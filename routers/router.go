package routers

import (
	dbtest "backendProject/db"
	"backendProject/resolver"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func Router() *gin.Engine {
	router := gin.New()

	router.GET("/swap/coin", handleRequest)

	return router
}

func handleRequest(c *gin.Context) {
	input := &dbtest.InputCoin{}

	input.Source = c.Query("source")
	input.Target = c.Query("target")
	amountStr := c.Query("amount")

	if isValidSource(input.Source) && isValidTarget(input.Target) {
		input.Amount, _ = strconv.ParseFloat(amountStr, 64)

		output := resolver.ChangeCoin(input)
		amountWithComma := formatAmount(output.Amount)

		c.JSON(http.StatusOK, gin.H{
			"msg":    "success",
			"amount": amountWithComma,
		})
	} else {
		var errorMsg string
		if !isValidSource(input.Source) {
			errorMsg = "source coin not supported"
		} else {
			errorMsg = "target coin not supported"
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": errorMsg,
		})
	}
}

func isValidSource(source string) bool {
	return source == "TWD" || source == "USD" || source == "JPY"
}

func isValidTarget(target string) bool {
	return target == "TWD" || target == "USD" || target == "JPY"
}

func formatAmount(amount float64) string {
	amountFormatted := strconv.FormatFloat(amount, 'f', 2, 64)
	return addCommas(amountFormatted)
}

func addCommas(s string) string {

	parts := strings.Split(s, ".")
	integerPart := parts[0]
	decimalPart := ""
	if len(parts) > 1 {
		decimalPart = "." + parts[1]
	}

	n := len(integerPart)
	if n <= 3 {
		return s
	}
	return addCommas(integerPart[:n-3]) + "," + integerPart[n-3:] + decimalPart
}
