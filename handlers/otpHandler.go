package handlers

import (
	otpservice "auth/services/OTPService"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type OTP struct {
	Otp string `json:"otp"`
}

func OtpHandler(w http.ResponseWriter, r *http.Request) {
	Otp := OTP{}
	otp := json.NewDecoder(r.Body).Decode(&Otp)
	otpjson, err := json.Marshal(otp)
	if err != nil {
		log.Fatal(err)
	}
	key := r.Header.Get(http.StatusText(http.StatusOK))
	fmt.Println("this is the key we got from client in our header", key)
	fmt.Println(otpjson)
	//uid := w.Header().Get(http.StatusText(http.StatusOK))
	//uuid := w.Header().Get(key)
	//fmt.Println("this is a unique string in the client request header", uuid)
	fmt.Println(Otp.Otp)
	v := otpservice.Compare(Otp.Otp, key)
	val := strconv.FormatBool(v)
	w.Write([]byte(val))

}
