package handlers

import (
	"auth/models"
	otpservice "auth/services/OTPService"
	jwtservice "auth/services/jwtService"
	smsservice "auth/services/smsService"

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
		fmt.Println("something went wrong", err)
	}
	phone := user.Phone
	mail := user.Email
	fmt.Println("this is signuphandler " + phone + mail)
	otp, uuid := otpservice.StoreData(phone)
	fmt.Println(otp)
	sms := smsservice.MessageContructor(otp.String())
	//mailservice.Send_Mail(mail, otp.String())
	smsservice.SendSms(phone, sms)

	//mailservice.SendMail(mail, otp.String())
	t, err := jwtservice.CreateTokenPair(user.Name)
	if err != nil {
		fmt.Println(err)
	}
	j, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
	}
	/*	tokens := map[string]string{
		"accesstoken":t.AccessToken,
		"refreshtoken":t.RefreshToken,
	}*/
	fmt.Println("this is a token: ", t.RefreshToken)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add(http.StatusText(http.StatusOK), uuid.String())
	//w.Header().Add("Tokens", t.AccessToken)
	//w.Header().Add("RefreshToken", t.RefreshToken)
	w.Write(j)
	w.WriteHeader(http.StatusOK)
	w.Write(userJson)
}
