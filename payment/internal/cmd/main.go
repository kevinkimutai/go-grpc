package main

import (
	"fmt"

	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/kevinkimutai/go-grpc/payment/internal/adapters/db"
	"github.com/kevinkimutai/go-grpc/payment/internal/adapters/grpc"
	"github.com/kevinkimutai/go-grpc/payment/internal/application/core/api"
	log "github.com/sirupsen/logrus"
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

	//Convert Port to int
	portInt, err := strconv.Atoi(PORT)

	if err != nil {
		fmt.Println("Error during conversion PORT to int")
		return
	}

	//Connect To DB
	dbAdapter, err := db.NewAdapter(DBURL)

	if err != nil {
		log.Fatal("Error connecting to DB.", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, portInt)
	grpcAdapter.Run()

}
