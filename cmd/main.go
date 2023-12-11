package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jeeo/pack-management/internal/api"
	"github.com/jeeo/pack-management/internal/api/handlers"
	"github.com/jeeo/pack-management/internal/application"
	"github.com/jeeo/pack-management/internal/config"
	"github.com/jeeo/pack-management/internal/database"
	"github.com/jeeo/pack-management/internal/repository"
)

func main() {
	config := config.LoadConfig()
	databaseConnection := database.NewPostgresDB(config.DBConfig)
	packRepository := repository.NewPackRepository(databaseConnection)
	orderApplication := application.NewOrderApplication(packRepository)
	packApplication := application.NewPackApplication(packRepository)

	server := api.NewRestServer(config, handlers.NewPackHandler(packApplication), handlers.NewOrderHandler(orderApplication))

	go func() {
		if err := server.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("server crash:", err)
		}

		log.Println("graceful shutdown")
	}()

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	<-signalChannel

	databaseConnection.Close()
	if err := server.Shutdown(); err != nil {
		log.Fatal("error during shutdown the server:", err)
	}

}
