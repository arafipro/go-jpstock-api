package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arafipro/go-jpstock/pkg/config"
	"github.com/arafipro/go-jpstock/pkg/models"
)

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	fmt.Printf("Starting server at port %s\n", server.Addr)
	config.Connect()
	models.CreateMarketTable()
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
