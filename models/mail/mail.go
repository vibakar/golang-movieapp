package mail

import (
	"gopkg.in/gomail.v2"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"crypto/rand"
	"strconv"
)

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
var fromEmail = beego.AppConfig.String("fromEmail")
var configEmail = beego.AppConfig.String("configEmail")
var configEmailPwd = beego.AppConfig.String("configEmailPwd")

func SendMail(email string, code float64){
	m := gomail.NewMessage()
	m.SetHeader("From", fromEmail)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Email verification!")
	m.SetBody("text/html", `<p>Hai!!</p>
                            <p>Kindly use this verification code to verify your email <strong>`+strconv.FormatFloat(code, 'f', 0, 64)+`</strong></p>`)

	d := gomail.NewDialer("smtp.gmail.com", 587, configEmail, configEmailPwd)

	if err := d.DialAndSend(m); err != nil {
		beego.Error("-------------------------- Email not sent to the user --------------------------", email, err)
	} else {
		fmt.Println("Mail sent successfully to the user", email, code)
	}
}

func GenerateRandNo(max int) float64 {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	randNo := string(b)
	s, err := strconv.ParseFloat(randNo, 64)
	if err == nil {
		beego.Info("Random no generated successfully")
		return s
	} else {
		beego.Error("Failed to generate the random number", err)
		return 789232
	}
}