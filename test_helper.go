package sqs

import (
	"os"
	"time"

	config "github.com/tommzn/go-config"
)

// runAtCI will have a look if env var CI is set.
func runAtCI() bool {
	_, ok := os.LookupEnv("CI")
	return ok
}

// mockedClientForTest returns a new client with mocked sqs client.
func mockedClientForTest() *Client {
	return newClient(newSqsMock())
}

// clientForTest read the test config and returns a new client.
func clientForTest() *Client {
	return newClientFromConfig(loadConfigForTest())
}

// loadConfigForTest returns test config from file testconfig.yml.
func loadConfigForTest() config.Config {

	configFile := "testconfig.yml"
	source := config.NewFileConfigSource(&configFile)
	conf, _ := source.Load()
	return conf
}

// testMessage is a struct used for publisher and consumer tests.
type testMessage struct {
	Val1 string
	Val2 int
	Val3 time.Time
}

// newTestMessage creates a new message with dummy values.
func newTestMessage() testMessage {
	return testMessage{
		Val1: "Val1",
		Val2: 2536456,
		Val3: time.Now(),
	}
}
