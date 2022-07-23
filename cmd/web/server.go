package main

import (
	"fmt"
	handler "httpPolicy/pkg/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	//routes
	//handles http req
	r.HandleFunc("/http-policy", handler.HandleHttpPolicy).Methods("POST")

	//server config
	fmt.Println("server at PORT 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
