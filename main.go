package main

import (
	"bytes"
	"html/template"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

type Data struct {
	Subject               string
	Name                  string
	Email                 string
	EmailVerificationLink string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	data := Data{
		Subject:               "Verification Email",
		Name:                  "Argana Phangquestian",
		Email:                 "arganaphangquestian@gmail.com",
		EmailVerificationLink: "https://facebook.com",
	}

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("MAIL_EMAIL"))
	m.SetHeader("To", data.Email)
	m.SetHeader("Subject", data.Subject)
	render, err := renderTemplate(data)
	if err != nil {
		log.Fatal(err)
	}
	m.SetBody("text/html", *render)

	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("MAIL_EMAIL"), os.Getenv("MAIL_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func renderTemplate(data Data) (*string, error) {
	t := template.Must(template.ParseFiles("template.html"))
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		return nil, err
	}
	result := tpl.String()
	return &result, nil
}
