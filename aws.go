package sqs

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	config "github.com/tommzn/go-config"
)

// newSQSClient creates a new AWS SQS client with given config.
// Config can contain the AWS region this client should talk to.
// If no region is specified it will lookup for AWS_REGION env var
// or use default region eu-central if it's not set.
func newSQSClient(conf config.Config) *sqs.SQS {

	awsRegion := conf.Get("aws.sqs.region", config.AsStringPtr(getAwsRegion()))
	return sqs.New(session.Must(session.NewSession(&aws.Config{Region: awsRegion})))
}

// getAwsRegion will try to retieve AWS region from envrionment and
// will return default region eu-central-1 if it's not set.
func getAwsRegion() string {
	if region, ok := os.LookupEnv("AWS_REGION"); ok {
		return region
	}
	return "eu-central-1"
}
