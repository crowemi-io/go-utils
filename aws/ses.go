package aws

import "github.com/aws/aws-sdk-go/service/ses"

type Ses struct {
	client *Client

	FromAddress  string
	ToAddresses  []string
	Subject      string
	CcAddresses  []string
	BccAddresses []string
}

func (client *Client) GetSES() *ses.SES {
	// if session passed in, use session else use AwsClient.Session
	return ses.New(client.Session)
}

// SetDestination adds a single recipient to the recipients list.
func (s Ses) SetDestination(recipient string) {

}

// SetFromAddress adds a sender to the email.
func (s Ses) SetFromAddress(sender string) {
	s.FromAddress = sender
}

// SetSubject adds a subject to the email.
func (s Ses) SetSubject(subject string) {
	s.Subject = subject
}
func (s Ses) AddBcc() {}
func (s Ses) AddCc()  {}

func (s Ses) SendMail() *ses.SendEmailOutput {
	// create a session
	client := s.client.GetSES()
	// assemble the email
	input := &ses.SendEmailInput{}
	// send the email
	output, err := client.SendEmail(input)
	if err != nil {
		// handle errors
	}
	return output
}
