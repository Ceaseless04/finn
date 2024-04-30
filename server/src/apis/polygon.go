package apis;

import (
	"context"
	"fmt"
	"os"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

// function to get exponential moving average
func getEMA() {
	
	// new client using the polygon api
	client := polygon.New(os.Getenv("POLYGON_API"));

	// set params
	params := &models.GetLastTradeParams{
		Ticker: "AAPL",
	};

	// make request
	res, err := client.GetLastTrade(context.Background(), params);
	if err != nil {
		fmt.Println(err);
	}

	// print result
	fmt.Println(res);

	fmt.Println(client);
}