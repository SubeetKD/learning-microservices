package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
    ID          int         `json:"id"`
    Name        string      `json:"name"`
    Description string      `json:"description"`
    Price       float32     `json:"price"`
    SKU         string      `json:"sku"`
    CreatedOn   string      `json:"-"`
    UpdatedOn   string      `json:"-"`
    DeletedOn   string      `json:"-"`
}

type Products []*Product

// directly encode to response writer
func (this *Products)ToJSON(w io.Writer) error {
    encoder := json.NewEncoder(w)
    return encoder.Encode(this)
}

var productList = []*Product{
    &Product{
        ID: 1,
        Name: "First product",
        Description: "First Despcription",
        Price: 0.0,
        SKU: "First sku",
        CreatedOn: time.Now().String(),
        UpdatedOn: time.Now().String(),
    },
    &Product{
        ID: 2,
        Name: "second product",
        Description: "second Despcription",
        Price: 9090909.899,
        SKU: "second sku",
        CreatedOn: time.Now().String(),
        UpdatedOn: time.Now().String(),
    },
}

func GetProduct() Products {
    return productList
}
