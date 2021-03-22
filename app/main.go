package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func catchAllHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handler")
	w.Write([]byte("Hello World\n"))
}

func main() {
	log.Println("main")
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.HandlerFunc(catchAllHandler))
	log.Fatal(http.ListenAndServe(":10003", r))
}
