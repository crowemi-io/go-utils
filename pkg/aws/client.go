package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Client is the base AWS client containing configuration
type Client struct {
	Session *session.Session
	Config  *aws.Config
}

// CreateSession creates a new AWS session
func CreateSession(region string) *session.Session {
	var err error
	var s *session.Session
	// defaults region to us-west-2
	if region == "" {
		region = "us-west-2"
	}
	// use default credential chain
	if s, err = session.NewSession(&aws.Config{Region: aws.String(region)}); err != nil {
		return nil
	}
	return s
}

func (client *Client) CreateConfig() {}
