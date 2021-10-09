package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/SubeetKD/learning-microservices/handlers"
)

// basic http handler
func main() {

	logger := log.New(os.Stdout, "product-api", log.LstdFlags)

	// initialize handlers
	hello_handler := handlers.NewHello(logger)
	goodbye_handler := handlers.NewGoodbye(logger)
    product_handler := handlers.NewProduct(logger)

	// create mux
	serve_mux := http.NewServeMux()

	// set handlers
	serve_mux.Handle("/shit", hello_handler)
	serve_mux.Handle("/", product_handler)
	serve_mux.Handle("/goodbye", goodbye_handler)

	// set you custom server
	server := &http.Server{
		Addr:         ":9090",
		Handler:      serve_mux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

    go func() {
		// start the custom server
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
            os.Exit(1)
		}
	}()

    sigChan := make(chan os.Signal)
    signal.Notify(sigChan, os.Interrupt)
    signal.Notify(sigChan, os.Kill)

    sig := <- sigChan
    logger.Println("Recieved termination, graceful shutdown", sig)

    timeout_context, _ := context.WithTimeout(context.Background(), 30 * time.Second)

	server.Shutdown(timeout_context)

	// listen to the server with the given handlers
	// this uses the default the server
	// http.ListenAndServe(":9090", serve_mux)
}
