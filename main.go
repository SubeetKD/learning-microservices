package main

import (
	"log"
	"net/http"
	"os"

	"github.com/SubeetKD/learning-microservices/handlers"
	"github.com/prometheus/common/log"
)

func main() {
    l := log.NewLogger(os.Stdout, "product-api", log.LstdFlags)
    hh := handlers.NewHello(l)
    http.HandleFunc()
    http.ListenAndServe(":9090", nil)
}
