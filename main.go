package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/hello", helloHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("frontend")))

	http.Handle("/", r)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World! %#v", r)
}
