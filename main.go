package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net/smtp"
	"strconv"
)

func main() {
	sendMail()
}

func generateRandNum(phone string) *big.Int {

	val, err := strconv.Atoi(phone)

	if err != nil {
		log.Fatal(err)
	}
	p, err := rand.Int(rand.Reader, big.NewInt(int64(val)))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)
	return p
}

type Mail struct {
	Sender  string
	To      []string
	Subject string
	Body    string
}

func sendMail() {
	// In production we will store all our confidential data like sender,username,password,addr into a separate .env file.
	g := generateRandNum("9553550050")
	sender := "noreply@gmail.com"
	to := []string{"user50@gmail.com"}
	subject := "Your OTP for verification..."
	body := " Please use" + g.String() + "  as a your OTP"

	request := Mail{
		Sender:  sender,
		To:      to,
		Subject: subject,
		Body:    body,
	}
	username := "noreply@gmail.com" //company mail address
	password := "750505"            // companies gmail password

	addr := "smtp.gmail.com:587"
	//host := "smtp.mailtrap.io"

	msg := buildMail(request)
	auth := smtp.CRAMMD5Auth(username, password)
	err := smtp.SendMail(addr, auth, sender, to, []byte(msg))

	if err != nil {
		log.Fatal(err)
		fmt.Println("mail not sent due to some error on server side!")
	}

	fmt.Println("Mail sent successfully")

}

func buildMail(mail Mail) string {
	msg := "Content-Type: text/html;"
	msg += fmt.Sprintf("From: %s\r\n", mail.Sender)
	msg += fmt.Sprintf("To: %s\r\n", mail.To)
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}
