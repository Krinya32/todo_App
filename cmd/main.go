package main

import (
	"github.com/krinya32/todoApp"
	handlers2 "github.com/krinya32/todoApp/pkg/handlers"
	"log"
)

func main() {
	handlers := new(handlers2.Handler)

	srv := new(todoApp.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal("Error occured while running http server: %s", err.Error())
	}
}
