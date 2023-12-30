package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	greater()
	r := mux.NewRouter()
	r.HandleFunc("/", serverBase)
	log.Fatal(http.ListenAndServe(":8000", r))

}

func greater() {
	fmt.Println("Using Gorilla mod")
}

func serverBase(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>USing Gorrilaa mod in webpage </h1>"))

}
