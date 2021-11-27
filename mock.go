package sqs

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// sqsMock mocks access to AWS SQS for testing.
type sqsMock struct {
	callCount int
}

// newSqsMock creates a new mock for AWS SQS.
func newSqsMock() *sqsMock {
	return &sqsMock{callCount: 0}
}

func (mock *sqsMock) SendMessage(input *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	mock.callCount++
	return &sqs.SendMessageOutput{
		MessageId: aws.String("MessageId"),
	}, nil
}

func (mock *sqsMock) ReceiveMessage(input *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
	mock.callCount++
	messageBody, _ := json.Marshal(newTestMessage())
	return &sqs.ReceiveMessageOutput{
		Messages: []*sqs.Message{
			&sqs.Message{
				MessageId:     aws.String("1"),
				ReceiptHandle: aws.String("xxx"),
				Body:          aws.String(string(messageBody)),
			},
		},
	}, nil
}

func (mock *sqsMock) DeleteMessage(input *sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error) {
	return &sqs.DeleteMessageOutput{}, nil
}

// GetQueueUrl will return a local generated sqs queue url.
// If you send "error" as queue name it will return with an error message.
func (mock *sqsMock) GetQueueUrl(input *sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error) {

	mock.callCount++
	if *input.QueueName == "error" {
		return nil, errors.New("Not Found!")
	}
	return &sqs.GetQueueUrlOutput{
		QueueUrl: aws.String("https://sqs.eu-central-1.amazonaws.com/<AccountId>/" + *input.QueueName),
	}, nil

}
