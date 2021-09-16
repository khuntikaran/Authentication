package smsservice

import (
	"fmt"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func MessageContructor(otp string) string {
	sms := fmt.Sprintf("Please use %v as you OTP", otp)
	return sms
}

func SendSms(ph string, otp string) {
	var AccountId = "AC1509f6b44774fc8e0b40b5d50a3730be"
	var AccountToken = "58647f05defc53aa16741d0273c4172e"
	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username: AccountId,
		Password: AccountToken,
	})
	params := openapi.CreateMessageParams{}
	params.SetFrom("+12019493769")
	params.SetTo(ph)
	params.SetBody(otp)
	resp, err := client.ApiV2010.CreateMessage(&params)
	if err != nil {
		fmt.Println("the error occures while sending OTP", err)
	}
	fmt.Println(resp.Body)
}
