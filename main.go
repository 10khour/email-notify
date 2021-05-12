package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
	"strings"

	"github.com/namsral/flag"
)

var sender Sender

func init() {
	var (
		host      string
		port      int
		user      string
		to        string
		pass      string
		subject   string
		emailFile string
	)
	flag.StringVar(&host, "smtp-server", "smtp.gmail.com", "smtp server host")
	flag.IntVar(&port, "smtp-port", 587, "smtp server port")
	flag.StringVar(&pass, "smtp-pass", "smtp.gmail.com", "smtp server host")
	flag.StringVar(&user, "user", "hellojukay@gmail.com", "your email")
	flag.StringVar(&to, "to", "xxx@gmail.com,yyy@gmail.com", "target email")
	flag.StringVar(&emailFile, "path", "mail.txt", "email body path")
	flag.StringVar(&subject, "subject", "Hello there", "email subject")
	flag.Parse()
	buffer, err := ioutil.ReadFile(emailFile)
	if err != nil {
		log.Printf("无法读取邮件文件%s %s", emailFile, err)
		os.Exit(1)
	}
	sender = Sender{
		Host:    host,
		Port:    port,
		User:    user,
		To:      to,
		Pass:    pass,
		Body:    os.ExpandEnv(string(buffer)),
		Subject: subject,
	}
}
func main() {
	if err := sender.Send(); err != nil {
		log.Printf("认证邮件服务器 %s 失败，检查帐号密码正确性,%s", fmt.Sprintf("%s:%d", sender.Host, sender.Port), err)
		os.Exit(1)
	}
	log.Printf("成功发送邮件给 %s", sender.To)
}

type Sender struct {
	Host    string
	Port    int
	User    string
	Pass    string
	To      string
	Subject string
	Body    string
}

func (sender Sender) Send() error {
	msg := "From: " + sender.User + "\n" +
		"To: " + sender.To + "\n" +
		"MIME-version: 1.0;\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\n" +
		"Subject: " + sender.Subject + "\n\n" +
		string(sender.Body)

	log.Printf("认证邮件服务器 smtp:%s:%d", sender.Host, sender.Port)

	// 认证这里有坑，参考: http://being23.github.io/2015/09/17/%E4%BD%BF%E7%94%A8golang%E5%8F%91%E9%80%81%E9%82%AE%E4%BB%B6/
	auth := LoginAuth(strings.Split(sender.User, "@")[0], sender.Pass)
	err := smtp.SendMail(fmt.Sprintf("%s:%d", sender.Host, sender.Port),
		auth,
		sender.User, strings.Split(sender.To, ","), []byte(msg))
	return err
}

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unkown fromServer")
		}
	}
	return nil, nil
}
