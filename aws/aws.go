package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type AwsClient struct {
	Session *session.Session
}

func (client *AwsClient) CreateSession(config *aws.Config) {
	client.Session, _ = session.NewSession(config)
}

func (client *AwsClient) GetSES(session session.Session) *ses.SES {
	// if session passed in, use session else use AwsClient.Session
	return ses.New(client.Session)
}
