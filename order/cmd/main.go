package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/kevinkimutai/go-grpc/order/adapters/db"
	"github.com/kevinkimutai/go-grpc/order/adapters/grpc"
	"github.com/kevinkimutai/go-grpc/order/application/core/api"
	paymentAdaptor "github.com/kevinkimutai/go-grpc/payment/internal/adapters/payment"
)

func main() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//GET ENV VARIABLES
	PORT := os.Getenv("PORT")
	DBURL := os.Getenv("DB_URL")
	PAYMENT_URL := os.Getenv("PAYMENT_URL")

	//Convert Port to int
	portInt, err := strconv.Atoi(PORT)

	if err != nil {
		fmt.Println("Error during conversion PORT to int")
		return
	}

	//Connect To DB
	dbAdapter, err := db.NewAdapter(DBURL)

	//connect To Payment Service
	paymentAdapter, err := paymentAdapter.NewAdapter(PAYMENT_URL)

	if err != nil {
		log.Fatal("Error connecting to DB.", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, portInt)
	grpcAdapter.Run()

}
