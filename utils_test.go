package sqs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UtilsTestSuite struct {
	suite.Suite
}

func TestUtilsTestSuite(t *testing.T) {
	suite.Run(t, new(UtilsTestSuite))
}
func (suite *UtilsTestSuite) TestGetMessageUrl() {

	client := mockedClientForTest()

	expectedUrl := "https://sqs.eu-central-1.amazonaws.com/<AccountId>/test"
	url1, err1 := client.urlForMessageQueue("test")
	suite.Nil(err1)
	suite.Equal(expectedUrl, *url1)
	suite.Equal(1, client.sqsClient.(*sqsMock).callCount)

	url2, err2 := client.urlForMessageQueue("test")
	suite.Nil(err2)
	suite.Equal(expectedUrl, *url2)
	suite.Equal(1, client.sqsClient.(*sqsMock).callCount)

	url3, err3 := client.urlForMessageQueue("error")
	suite.NotNil(err3)
	suite.Nil(url3)
}

func (suite *UtilsTestSuite) TestGetAwsRegion() {

	suite.Equal("eu-central-1", getAwsRegion())

	awsRegion := "eu-west-5"
	os.Setenv("AWS_REGION", awsRegion)
	suite.Equal(awsRegion, getAwsRegion())
	os.Unsetenv("AWS_REGION")
}

func (suite *UtilsTestSuite) TestSqsIntegration() {

	if runAtCI() {
		suite.T().Skip("Skip direct integration to SQS.")
	}

	client := clientForTest()

	url1, err1 := client.urlForMessageQueue("tzn-unittest")
	suite.Nil(err1)
	suite.NotEqual("", *url1)

	url2, err2 := client.urlForMessageQueue("xxx")
	suite.NotNil(err2)
	suite.Nil(url2)
}
