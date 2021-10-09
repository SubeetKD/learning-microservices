package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
    logger *log.Logger
}

func NewGoodbye(logger *log.Logger) *Goodbye {
    return &Goodbye{logger: logger}
}

func (this *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
    log.Println("Goodbye")
}
