package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"

	"github.com/namsral/flag"
)

func main() {
	var host string
	var port int
	var from string
	var to string
	var pass string
	var body string
	flag.StringVar(&host, "smtp-server", "smtp.gmail.com", "smtp server host")
	flag.IntVar(&port, "smtp-port", 587, "smtp server port")
	flag.StringVar(&pass, "smtp-pass", "smtp.gmail.com", "smtp server host")
	flag.StringVar(&from, "user", "hellojukay@gmail.com", "your email")
	flag.StringVar(&to, "to", "xxx@gmail.com,yyy@gmail.com", "target email")
	flag.StringVar(&body, "body", "hello world", "email body content")
	flag.Parse()
	var sender = Sender{
		Host: host,
		Port: port,
		From: from,
		To:   to,
		Pass: pass,
	}
	if err := sender.Send(); err != nil {
		log.Printf("认证邮件服务器 %s 失败拉，检查帐号密码正确性", fmt.Sprintf("%s:%d", host, port))
		os.Exit(1)
	}
	if host == "" {

	}
}

type Sender struct {
	Host string
	Port int
	User string
	Pass string
	From string
	To   string
	Body []byte
}

func (sender Sender) Send() error {

	msg := "From: " + sender.From + "\n" +
		"To: " + sender.To + "\n" +
		"Subject: Hello there\n\n" +
		string(sender.Body)

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", sender.From, sender.Pass, "smtp.gmail.com"),
		sender.From, strings.Split(sender.To, ","), []byte(msg))
	return err
}
