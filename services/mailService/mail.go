package mailservice

import (
	"auth/models"
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"

	"github.com/go-gomail/gomail"
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
	password := "M@her7506"               // companies gmail password

	host := "smtp.gmail.com"
	//host := "smtp.mailtrap.io"
	port := "587"
	msg := buildMail(request)
	auth := smtp.PlainAuth("", username, password, "smtp.gmail.com")
	err := smtp.SendMail(host+":"+port, auth, sender, to, []byte(msg))

	if err != nil {
		log.Fatal(err)
		fmt.Println("mail not sent due to some error on server side!")
	}

	fmt.Println("Mail sent successfully")

}

func Send_Mail(mail string, otp string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "khuntikaran51@gmail.com")
	m.SetHeader("To", mail)
	m.SetHeader("Subject", "Your OTP for verification")
	m.SetBody("text", "Please use"+otp+"as a your OTP")

	host := "smtp.gmail.com"
	port := 587
	username := "khuntikaran51@gmail.com" //company mail address
	password := "M@her7506"

	d := gomail.NewDialer(host, port, username, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	}
}

func buildMail(mail models.Mail) string {
	msg := "Content-Type: text/html;"
	msg += fmt.Sprintf("From: %s\r\n", mail.Sender)
	msg += fmt.Sprintf("To: %s\r\n", mail.To)
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}
