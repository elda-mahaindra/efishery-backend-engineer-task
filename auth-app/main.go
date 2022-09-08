package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"auth-app/api"
	"auth-app/store"
	"auth-app/token"
	"auth-app/util"
)

const (
	USER_DATA_FILE = "./data/user_data"
)

func setupApiServer(config util.Config, tokenManager token.Manager) (*api.ApiServer, error) {
	userStore := store.NewInMemoryUserStore(USER_DATA_FILE)

	err := userStore.PopulateDataFromFile()
	if err != nil {
		return nil, fmt.Errorf("failed to populate data from file: %v", err.Error())
	}

	apiServer := api.NewApiServer(config, userStore, tokenManager)

	return apiServer, nil
}

func main() {
	// make sure we can get the file and line number that cause the app to crash
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// load environment variables from .env file
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config file: %v", err.Error())
	}

	tokenManager, err := token.NewJWTManager(config.Secret)
	if err != nil {
		log.Fatalf("Failed to create token manager: %v", err.Error())
	}

	apiServer, err := setupApiServer(config, tokenManager)
	if err != nil {
		log.Fatalf("Failed to create api server: %v", err.Error())
	}

	// serve the api server
	go func() {
		err = apiServer.Start(config)
		if err != nil {
			log.Fatal("Failed to start api server:", err)
		}
	}()

	// wait for ctrl + c to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// block until a signal is received
	<-ch

	log.Println("Stopping the server...")
	log.Println("End of Program...")
}
