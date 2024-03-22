package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
)

const CharSet = "UTF-8"

// SesClient is a wrapper for aws.SES with methods for sending an email
type SesClient struct {
	Client      Client
	SesClient   *ses.SES
	destination *ses.Destination
	message     *ses.Message
	body        *ses.Body
	subject     *ses.Content
	Email       EmailContent
}

// EmailContent is bound to ts-utils EmailContent (url: )
type EmailContent struct {
	FromAddress    string
	ReplyToAddress string
	ToAddresses    []*string
	CcAddresses    []*string
	BccAddresses   []*string
	Subject        string
	HtmlBody       string
	TextBody       string
}

func (s *SesClient) CreateSesClient() *ses.SES {
	// if session passed in, use session else use AwsClient.Session
	s.SesClient = ses.New(s.Client.Session)
	return s.SesClient
}

// AddDestination adds a single recipient to the recipients list.
func (s *SesClient) AddDestination(DestinationTo []string, DestinationBcc []string, DestinationCc []string) {
	s.AddDestinationTo(DestinationTo)
	s.AddDestinationCc(DestinationCc)
	s.AddDestinationBcc(DestinationBcc)
}

// SetFromAddress adds a sender to the email.
func (s *SesClient) SetFromAddress(sender string) {
	s.Email.FromAddress = sender
}

// SetSubject adds a subject to the email.
func (s *SesClient) SetSubject(subject string) {
	s.Email.Subject = subject
}

func (s *SesClient) AddDestinationBcc(addresses []string) {
	for _, v := range addresses {
		s.Email.BccAddresses = append(s.Email.BccAddresses, &v)
	}
}
func (s *SesClient) AddDestinationCc(addresses []string) {
	for _, v := range addresses {
		s.Email.CcAddresses = append(s.Email.CcAddresses, &v)
	}
}
func (s *SesClient) AddDestinationTo(addresses []string) {
	for _, v := range addresses {
		s.Email.ToAddresses = append(s.Email.ToAddresses, &v)
	}
}

// CreateDestination creates and returns an AWS destination object while assigning it to the SesClient.
func (s *SesClient) CreateDestination() ses.Destination {
	destination := ses.Destination{
		ToAddresses:  s.Email.ToAddresses,
		CcAddresses:  s.Email.CcAddresses,
		BccAddresses: s.Email.BccAddresses,
	}
	s.destination = &destination
	return destination
}

// CreateMessage creates and returns an AWS message object while assigning it to the SesClient.
func (s *SesClient) CreateMessage() ses.Message {
	body := ses.Body{
		Html: &ses.Content{
			Charset: aws.String(CharSet),
			Data:    aws.String(s.Email.HtmlBody),
		},
		Text: &ses.Content{
			Charset: aws.String(CharSet),
			Data:    aws.String(s.Email.TextBody),
		},
	}
	subject := ses.Content{
		Charset: aws.String(CharSet),
		Data:    aws.String(s.Email.Subject),
	}
	message := ses.Message{
		Body:    &body,
		Subject: &subject,
	}

	s.body = &body
	s.subject = &subject
	s.message = &message

	return message
}

func (s *SesClient) CreateSendEmailInput() ses.SendEmailInput {
	var sendEmailInput ses.SendEmailInput

	destination := s.CreateDestination()
	message := s.CreateMessage()

	sendEmailInput.Destination = &destination
	sendEmailInput.Message = &message
	sendEmailInput.Source = &s.Email.FromAddress

	return sendEmailInput

}
func (s *SesClient) SendMail() (*ses.SendEmailOutput, error) {
	// create a session
	client := s.CreateSesClient()
	// assemble the email
	input := s.CreateSendEmailInput()
	// send the email
	var o *ses.SendEmailOutput
	var err error
	if o, err = client.SendEmail(&input); err != nil {
		return nil, fmt.Errorf("failed to send email: %w", err)
	}
	return o, err
}
