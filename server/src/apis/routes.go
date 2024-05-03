package apis;

import (
	"net/http"
	"github.com/gin-gonic/gin"
);

type Ticker struct {
	name string;
	limit int;
	EMA []float64;
}


func createTicker() gin.HandlerFunc {
	var AAPL Ticker;

	AAPL.name = "AAPL";
	AAPL.limit = 100;
	AAPL.EMA = Calc_EMA(AAPL.name, AAPL.limit);

	fn := func(c *gin.Context) {
		c.JSON(http.StatusOK, AAPL);
	}

	return gin.HandlerFunc(fn);
}

func Display() {
	router := gin.Default();
	router.GET("/apis/EMA", createTicker());
	router.Run("localhost:8080");
}
