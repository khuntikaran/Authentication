package mailservice

import (
	"auth/models"
	"fmt"
	"log"
	"net/smtp"
)

func SendMail(mailid string, otp string) {
	// In production we will store all our confidential data like sender,username,password,addr into a separate .env file.
	//g := generateRandNum("9553550050")
	sender := "khuntikaran51@gmail.com"
	to := []string{mailid}
	subject := "Your OTP for verification..."
	body := " Please use" + otp + "  as a your OTP"

	request := models.Mail{
		Sender:  sender,
		To:      to,
		Subject: subject,
		Body:    body,
	}
	username := "khuntikaran51@gmail.com" //company mail address
	password := "maher7505"               // companies gmail password

	addr := "smtp.gmail.com:587"
	//host := "smtp.mailtrap.io"

	msg := buildMail(request)
	auth := smtp.PlainAuth("", username, password, "smtp.gmail.com")
	err := smtp.SendMail(addr, auth, sender, to, []byte(msg))

	if err != nil {
		log.Fatal(err)
		fmt.Println("mail not sent due to some error on server side!")
	}

	fmt.Println("Mail sent successfully")

}

func buildMail(mail models.Mail) string {
	msg := "Content-Type: text/html;"
	msg += fmt.Sprintf("From: %s\r\n", mail.Sender)
	msg += fmt.Sprintf("To: %s\r\n", mail.To)
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}
