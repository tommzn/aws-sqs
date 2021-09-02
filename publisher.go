package sqs

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	config "github.com/tommzn/go-config"
)

// NewPublisher returns a new publisher used to send messages to AWS SQS.
func NewPublisher(conf config.Config) Publisher {
	return newClientFromConfig(conf)
}

// Send will marshal passed message and try to deliver it to AWS SQS.
// If successful it returns the message id.
func (client *Client) Send(message interface{}, queueName string) (*string, error) {

	var messageBody []byte
	if msg, ok := message.([]byte); ok {
		messageBody = msg
	} else {
		messageBody, _ = json.Marshal(message)
	}

	qURL, err := client.urlForMessageQueue(queueName)
	if err != nil {
		return nil, err
	}

	result, err := client.sqsClient.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(messageBody)),
		QueueUrl:    qURL,
	})
	if err != nil {
		return nil, err
	}
	return result.MessageId, nil
}
