package sqs

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type PublisherTestSuite struct {
	suite.Suite
}

func TestPublisherTestSuite(t *testing.T) {
	suite.Run(t, new(PublisherTestSuite))
}

func (suite *PublisherTestSuite) TestSendMessage() {

	client := mockedClientForTest()
	message := newTestMessage()

	messageId1, err1 := client.Send(message, "tzn-unittests")
	suite.Nil(err1)
	suite.NotNil(messageId1)
	suite.Equal(2, client.sqsClient.(*sqsMock).callCount)

	messageId2, err2 := client.Send([]byte("TestMessage"), "tzn-unittests")
	suite.Nil(err2)
	suite.NotNil(messageId2)
	suite.Equal(3, client.sqsClient.(*sqsMock).callCount)

	messageId3, err3 := client.SendAttributedMessage(message, "tzn-unittests", map[string]string{"Key1": "Val1"})
	suite.Nil(err3)
	suite.NotNil(messageId3)
	suite.Equal(4, client.sqsClient.(*sqsMock).callCount)
}

func (suite *PublisherTestSuite) TestSqsIntegration() {

	if runAtCI() {
		suite.T().Skip("Skip direct integration to SQS.")
	}

	publisher := NewPublisher(loadConfigForTest())
	message := newTestMessage()

	messageId, err := publisher.Send(message, "tzn-unittest")
	suite.Nil(err)
	suite.NotNil(messageId)

	messageId1, err1 := publisher.SendAttributedMessage(message, "tzn-unittest", map[string]string{"Key1": "Val1"})
	suite.Nil(err1)
	suite.NotNil(messageId1)

	messageId2, err2 := publisher.Send(message, "xxx")
	suite.NotNil(err2)
	suite.Nil(messageId2)

	publisher.(*Client).queueUrls["tzn-unittest"] = publisher.(*Client).queueUrls["tzn-unittest"] + "-xxx"
	messageId3, err3 := publisher.Send(message, "tzn-unittest")
	suite.NotNil(err3)
	suite.Nil(messageId3)
}
