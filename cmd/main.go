package main

import (
	"fmt"
	"github.com/crowemi-io/go-utils/pkg/aws"
)

func main() {

	var toAddresses []string
	toAddresses = append(toAddresses, "crowemi@hotmail.com")

	var s = aws.SesClient{
		Client: aws.Client{
			Session: aws.CreateSession("us-west-2"),
		},
		FromAddress:    "hello@crowemi.com",
		ReplyToAddress: "no-reply@crowemi.com",
		Subject:        "Hello World!",
		HtmlBody:       "<h1>Hello World! From GoLang!</h1>",
		TextBody:       "Hello World! From GoLang!",
	}
	s.AddDestinationTo(toAddresses)

	o, e := s.SendMail()
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(o)

}
