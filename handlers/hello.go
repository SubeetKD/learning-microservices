package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
    l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
    return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, w *http.Request) {
    h.l.Println("hello world")
    d, err := ioutil.ReadAll(w.Body)
    if err != nil {
        http.Error(rw, "Oops", http.StatusBadRequest)
        return
    }
    fmt.Fprintf(rw, "Hello %s", d)
}