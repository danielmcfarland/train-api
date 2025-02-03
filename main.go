package main

import (
	"fmt"
	"github.com/danielmcfarland/train-api/middleware"
	"github.com/danielmcfarland/train-api/routes"
	"log"
	"net/http"
	"os"
)

var webPort = "8080"

func init() {
	fmt.Println("Application Initialisation")

	if len(os.Getenv("WEB_PORT")) != 0 {
		webPort = os.Getenv("WEB_PORT")
	}
}

func main() {
	router := http.NewServeMux()

	routes.LoadRoutes(router)

	stack := middleware.CreateStack(middleware.Logging)

	server := http.Server{
		Addr:    ":" + webPort,
		Handler: stack(router),
	}

	log.Println(fmt.Sprintf("Server listening on port: %v", webPort))

	server.ListenAndServe()
}
