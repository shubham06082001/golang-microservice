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

func newHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) serveHTTP(rw http.ResponseWriter, r *http.Request) {
	log.Println("Hello")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s", d)
}
