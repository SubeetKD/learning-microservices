package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// refactoring the repo
type Hello struct {
    logger *log.Logger
}

func NewHello(logger *log.Logger) *Hello {
    return &Hello{logger}
}

func (this *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
    this.logger.Println("Hello world")
    d, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(rw, "some thing went wrong", http.StatusBadRequest)
        return
    }
    fmt.Fprintf(rw, "Hello %s bitch\n", d)
}
