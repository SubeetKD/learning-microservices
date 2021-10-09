package handlers

import (
	"log"
	"net/http"

	"github.com/SubeetKD/learning-microservices/data"
)

type Product struct {
    logger *log.Logger
}

func NewProduct(logger *log.Logger) *Product {
    return &Product{logger: logger}
}

func (this *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        this.getProducts(rw, r)
        return
    }

    rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (this *Product) getProducts(rw http.ResponseWriter, r *http.Request) {
    list := data.GetProduct()
    this.logger.Printf("got list of product %v", list)
    err := list.ToJSON(rw)
    if err != nil {
        this.logger.Println("Error converting data to json")
        http.Error(rw, "Unable to parse data", http.StatusBadGateway)
        return
    }
}
