package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/daudfauzy98/Pengenalan-Microservice/auth-service/handler"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/validate-admin", http.HandlerFunc(handler.ValidateAdmin))

	fmt.Println("Auth service listen on 8001")
	log.Panic(http.ListenAndServe(":8001", router))
}
