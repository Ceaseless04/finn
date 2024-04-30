package apis

import (
	"context"
	//"fmt"
	"os"
	"log"

	"github.com/joho/godotenv"
	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

// function to get exponential moving average
func GetEMA() {

	// loading env file with checks for when the file is empty
	err := godotenv.Load("../.env");

	if err != nil {
		log.Print(err);
	}

	// new client using the polygon api
	client := polygon.New(os.Getenv("POLYGON_API"));

	// set params
	params := models.GetEMAParams{
		Ticker: "AAPL",
	};
	
	// make request
	res, err := client.GetEMA(context.Background(), &params);
	if err != nil {
		log.Print(err);
	}

	// print result
	log.Print(res);

}