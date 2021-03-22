package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n"))
}

func main() {
	log.Println("main")
	r := mux.NewRouter()
	r.HandleFunc("*", Handler)
	log.Fatal(http.ListenAndServe(":10003", r))
}
