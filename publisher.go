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

	return client.SendAttributedMessage(message, queueName, make(map[string]string))
}

// SendAttributedMessage will marshal given message to JSON if it's not already a byte array,
// try to send it to AWS SQS using passed message attributes and will return the message id
// in case of successful delivering.
func (client *Client) SendAttributedMessage(message interface{}, queueName string, attributes map[string]string) (*string, error) {

	messageBody := toMessageBody(message)
	qURL, err := client.urlForMessageQueue(queueName)
	if err != nil {
		return nil, err
	}

	sendMessageInput := &sqs.SendMessageInput{
		MessageBody: aws.String(messageBody),
		QueueUrl:    qURL,
	}
	if len(attributes) > 0 {
		sendMessageInput.MessageAttributes = toSqsMessageAttributes(attributes)
	}

	result, err := client.sqsClient.SendMessage(sendMessageInput)
	if err != nil {
		return nil, err
	}
	return result.MessageId, nil
}

// toMessageBody converts given value to a byte array using JSON marshal.
// If passed value is already a byte array it will be returns as it is.
func toMessageBody(message interface{}) string {
	switch v := message.(type) {
	case []byte:
		return string(v)
	case string:
		return v
	default:
		messageBody, _ := json.Marshal(message)
		return string(messageBody)
	}
}

// toSqsMessageAttributes converts given map of message atributes to a suitable map for SQS.
func toSqsMessageAttributes(attributes map[string]string) map[string]*sqs.MessageAttributeValue {

	messageAttributes := make(map[string]*sqs.MessageAttributeValue)
	for attributeKey, attributeValue := range attributes {
		messageAttributes[attributeKey] = &sqs.MessageAttributeValue{
			DataType:    aws.String("String"),
			StringValue: aws.String(attributeValue),
		}
	}
	return messageAttributes
}
