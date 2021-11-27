package sqs

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type ConsumerTestSuite struct {
	suite.Suite
}

func TestConsumerTestSuite(t *testing.T) {
	suite.Run(t, new(ConsumerTestSuite))
}

func (suite *ConsumerTestSuite) TestReceiveMessage() {

	client := mockedClientForTest()

	messages, err := client.Receive("tzn-unittest")
	suite.Nil(err)
	suite.Len(messages, 1)
	suite.Nil(client.Ack("tzn-unittest", messages[0].ReceiptHandle))
}

func (suite *ConsumerTestSuite) TestSqsIntegration() {

	if runAtCI() {
		suite.T().Skip("Skip direct integration to SQS.")
	}

	consumer := NewConsumer(loadConfigForTest())
	suite.sendTestMessage(consumer.(*Client))
	time.Sleep(2)

	messages, err := consumer.Receive("tzn-unittest")
	suite.Nil(err)
	suite.True(len(messages) > 0)
	suite.Nil(consumer.Ack("tzn-unittest", messages[0].ReceiptHandle))
	suite.NotNil(consumer.Ack("xxx", messages[0].ReceiptHandle))

	messages2, err2 := consumer.Receive("xxx")
	suite.NotNil(err2)
	suite.Len(messages2, 0)

	consumer.(*Client).queueUrls["tzn-unittest"] = consumer.(*Client).queueUrls["tzn-unittest"] + "-xxx"
	messages3, err3 := consumer.Receive("tzn-unittest")
	suite.NotNil(err3)
	suite.Len(messages3, 0)
}

func (suite *ConsumerTestSuite) sendTestMessage(publisher Publisher) {
	_, err := publisher.Send(newTestMessage(), "tzn-unittest")
	suite.Nil(err)
}
