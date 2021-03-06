package main

import (
	"fmt"
	"net/http"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
}

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello Foo")
}

func barHandler(rw http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(rw, "Hello %s!", name)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello world")
	})

	mux.HandleFunc("/bar", barHandler)

	mux.Handle("/foo", &fooHandler{})

	http.ListenAndServe(":3000", mux)
}
