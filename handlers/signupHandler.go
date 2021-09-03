package handlers

import (
	"auth/models"
	otpservice "auth/services/OTPService"
	mailservice "auth/services/mailService"
	"encoding/json"
	"fmt"
	"net/http"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(err.Error()))
	}
	userJson, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	phone := user.Phone
	mail := user.Email
	otp, uuid := otpservice.StoreData(phone)
	go mailservice.SendMail(mail, otp.String())

	w.Header().Set("Content-Type", "application/json")
	w.Header().Add(user.Name, uuid.String())
	w.WriteHeader(http.StatusOK)
	w.Write(userJson)
}
