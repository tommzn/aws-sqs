// Package sqs provides publisher and receiver to interact with AWS SQS.
package sqs

import (
	config "github.com/tommzn/go-config"
)

// newClient returns a new client with passed sqs client.
func newClient(sqsClient sqsClient) *Client {
	return &Client{sqsClient: sqsClient, queueUrls: make(map[string]string)}
}

// newClientFromConfig creates a new client with settings obtained
// from passed config.
func newClientFromConfig(conf config.Config) *Client {
	return newClient(newSQSClient(conf))
}
