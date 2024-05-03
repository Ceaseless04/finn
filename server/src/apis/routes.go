package apis;

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/polygon-io/client-go/rest/models"
);

type Ticker struct {
	Name string		`json:"name"`
	Limit int		`json:"limit"`
	EMA models.SingleIndicatorValues	`json:"EMA"`
}

func getTicker() gin.HandlerFunc {
	AAPL := Ticker{
		Name: "AAPL",
		Limit: 100,
		EMA: Calc_EMA("AAPL", 100),
	}

	fn := func(c *gin.Context) {
		c.JSON(http.StatusOK, AAPL);
	}

	return gin.HandlerFunc(fn);
}

func Display() {
	router := gin.Default();
	router.GET("/apis/EMA", getTicker());
	router.Run("localhost:8080");
}
