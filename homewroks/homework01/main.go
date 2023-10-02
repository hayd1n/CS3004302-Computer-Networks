package main

import (
	"CRT-HAO/CS3004302/homework01/smtp"
	"flag"
	"fmt"
	"os"
)

var (
	host  string
	from  string
	to    string
	title string
	body  string
)

func init() {
	flag.StringVar(&host, "host", "", "smtp server address")
	flag.StringVar(&from, "from", "", "sender email address")
	flag.StringVar(&to, "to", "", "receiver email address")
	flag.StringVar(&title, "title", "", "email title")
	flag.StringVar(&body, "body", "", "email content")
}

func main() {
	flag.Parse()

	if host == "" {
		fmt.Println("Error: host not set")
		os.Exit(1)
	}

	if from == "" || to == "" {
		fmt.Println("Error: sender or receiver email not set")
		os.Exit(1)
	}

	smtp := smtp.NewSmtp()
	smtp.Host = host

	// connect to smtp server
	if err := smtp.Connect(); err != nil {
		panic(err)
	}

	// send email
	if err := smtp.Send(from, to, title, body); err != nil {
		panic(err)
	}

	fmt.Printf("successful")
}
