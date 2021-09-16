package main

import (
	"auth/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/signup", handlers.SignupHandler).Methods("POST")
	route.HandleFunc("/otp", handlers.OtpHandler).Methods("POST")
	route.HandleFunc("/refresh", handlers.TokenHandler).Methods("GET")
	http.ListenAndServe(":8000", route)

}
