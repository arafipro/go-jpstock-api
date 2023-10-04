package main

import (
	"fmt"
	"log"
	"net/http"
	
	_ "github.com/arafipro/go-jpstock-api/pkg/models"
)

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	fmt.Printf("Starting server at port %s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
