package sqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	config "github.com/tommzn/go-config"
)

// NewConsumer returns a client to read messages from an AWS SQS queue.
func NewConsumer(conf config.Config) Consumer {
	return newClientFromConfig(conf)
}

// Receive will try to read message from given queue.
func (client *Client) Receive(queueName string) ([]RawMessage, error) {

	messsages := []RawMessage{}
	qURL, err := client.urlForMessageQueue(queueName)
	if err != nil {
		return messsages, err
	}

	receiveMessageOutput, err := client.sqsClient.ReceiveMessage(&sqs.ReceiveMessageInput{
		MaxNumberOfMessages: aws.Int64(10),
		QueueUrl:            qURL,
	})
	if err != nil {
		return messsages, err
	}

	for _, message := range receiveMessageOutput.Messages {
		messsages = append(messsages, RawMessage{
			MessageId:     message.MessageId,
			ReceiptHandle: message.ReceiptHandle,
			Body:          message.Body,
		})
	}
	return messsages, err
}
