package utils

import (
	"crypto/rand"
	"crypto/tls"
	"encoding/hex"
	"fmt"

	"gopkg.in/gomail.v2"
)

func RandomToken() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func SendEmail(email string, link string) {

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "example@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", email)

	// Set E-Mail subject
	m.SetHeader("Subject", "Reset password")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/html", "<a href=\""+link+"\">Reset password</a>")

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "example@gmail.com", "bjcw klus iiht wllh")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return
}
