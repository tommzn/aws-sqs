package sqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func (client *Client) urlForMessageQueue(queueName string) (*string, error) {

	if queueUrl, ok := client.queueUrls[queueName]; ok {
		return &queueUrl, nil
	}

	getQueueUrlInput := &sqs.GetQueueUrlInput{QueueName: aws.String(queueName)}
	getQueueUrlOutput, err := client.sqsClient.GetQueueUrl(getQueueUrlInput)
	if err == nil {
		client.queueUrls[queueName] = *getQueueUrlOutput.QueueUrl
		return getQueueUrlOutput.QueueUrl, nil
	}
	return nil, err
}
