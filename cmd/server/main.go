package main

import (
	"auth/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/signup", handlers.SignupHandler).Methods("POST")
	route.HandleFunc("/otp", otpHandler).Methods("POST")
	http.ListenAndServe(":8000", route)

}
func otpHandler(w http.ResponseWriter, r *http.Request) {

}
