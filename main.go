package main

import (
	"log"
	"net/http"
	"os"

	"github.com/webmachinedev/core"
)

func main() {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "80"
	}
	http.HandleFunc("/", core.Handler)
	log.Default().Println("listening on port "+port)
	http.ListenAndServe(":"+port, nil)
}
