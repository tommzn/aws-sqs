package sqs

import (
	"github.com/aws/aws-sdk-go/service/sqs"
)

// Publisher provides message sending functionality.
type Publisher interface {

	// Send will marshal given message to JSON if it's not already a byte array,
	// try to send it to AWS SQS and will return the message id in case of successful delivering.
	Send(interface{}, string) (*string, error)

	// SendAttributedMessage will marshal given message to JSON if it's not already a byte array,
	// try to send it to AWS SQS using passed message attributes and will return the message id
	// in case of successful delivering.
	SendAttributedMessage(interface{}, string, map[string]string) (*string, error)
}

// Consumer provides read access to message queues in AWS SQS.
type Consumer interface {

	// Receive will read messages from given queue in AWS SQS.
	Receive(string) ([]RawMessage, error)
}

// sqsClient is an interface for a client interacting with AWS SQS.
type sqsClient interface {

	// SendMessage will deliver a message to AWS SQS.
	SendMessage(input *sqs.SendMessageInput) (*sqs.SendMessageOutput, error)

	// ReceiveMessage will try to read messages from AWS SQS.
	ReceiveMessage(input *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error)

	// GetQueueUrl will try to retrieve the url fr given queue from AWS SQS.
	GetQueueUrl(input *sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error)
}
