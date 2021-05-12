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
	var user string
	var to string
	var pass string
	var body string
	flag.StringVar(&host, "smtp-server", "smtp.gmail.com", "smtp server host")
	flag.IntVar(&port, "smtp-port", 587, "smtp server port")
	flag.StringVar(&pass, "smtp-pass", "smtp.gmail.com", "smtp server host")
	flag.StringVar(&user, "user", "hellojukay@gmail.com", "your email")
	flag.StringVar(&to, "to", "xxx@gmail.com,yyy@gmail.com", "target email")
	flag.StringVar(&body, "body", "hello world", "email body content")
	flag.Parse()
	var sender = Sender{
		Host: host,
		Port: port,
		User: user,
		To:   to,
		Pass: pass,
	}
	if err := sender.Send(); err != nil {
		log.Printf("认证邮件服务器 %s 失败拉，检查帐号密码正确性,%s", fmt.Sprintf("%s:%d", host, port), err)
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
	To   string
	Body []byte
}

func (sender Sender) Send() error {

	msg := "From: " + sender.User + "\n" +
		"To: " + sender.To + "\n" +
		"Subject: Hello there\n\n" +
		string(sender.Body)

	log.Print(msg)
	auth := smtp.PlainAuth("", sender.User, sender.Pass, sender.Host)
	err := smtp.SendMail(fmt.Sprintf("%s:%d", sender.Host, sender.Port),
		auth,
		sender.User, strings.Split(sender.To, ","), []byte(msg))
	return err
}
