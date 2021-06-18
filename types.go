package sqs

// RawMessage is used to return messages read from a queue in AWS SQS.
type RawMessage struct {

	// MessageId is the id of a message in AWS SQS.
	MessageId *string

	// ReceiptHandle is the receipt handle in AWS SQS.
	ReceiptHandle *string

	// Body contains the message body read from AWS SQS.
	Body *string
}

// SqsClient provides access to AWS SQS to send and receive messages.
type Client struct {

	// sqsClient a client to connect to AWS SQS.
	sqsClient sqsClient

	// queueUrls is an internal cache for queue urls to avoid
	// requesting AWS each time.
	queueUrls map[string]string
}
