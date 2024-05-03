package apis;

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	
	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)



// loading env file to create new client with Polygon API
var env = godotenv.Load("../.env");
var client = polygon.New(os.Getenv("POLYGON_API"));

// function to get exponential moving averages from ticker
func Calc_EMA(name string, limit int) []{

	if env != nil {
		log.Fatal("Error loading env file");
	}

	// set params
	params := models.GetEMAParams{
		Ticker: name,
	}.WithLimit(limit);
	
	// make request
	res, err := client.GetEMA(context.Background(), params);
	if err != nil {
		log.Print(err);
	}

	// print result
	fmt.Println(res);

	// gettting values array to store Value in ema_value
	values := res.Results.Values;
	ema_values := []float64{};

	for i, v := range values {
		ema_values[i] = v.Value;
	}

	return ema_values; // return array
}
