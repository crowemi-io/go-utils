package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Client is the base AWS client containing configuration
type Client struct {
	Session *session.Session
}

func (client *Client) CreateSession(config *aws.Config) {
	var err error
	client.Session, err = session.NewSession(config)
	if err != nil {
		// handle session error, panic?
	}
}
