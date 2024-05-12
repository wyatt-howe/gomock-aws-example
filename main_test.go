package main_test

import (
	"errors"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gomock-aws-example"
	mock_main "gomock-aws-example/mocks"
	"log"
	"testing"
)

type Client_Interface interface {
	RetrieveSqsMessage(*sqs.Client, string, bool) (string, string, error)
}

func TestUseClient(t *testing.T) {

	// Create a new instance of the GoMock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock object using the mock interface
	mockClient := mock_main.NewMockClient_Interface(ctrl)

	// Set up expectations for the useClient method
	//var expectedMessage = ("{'finaloutput':{'externaldata':'1a'}}", "secondstring", errors.New("newerrorexpectedout"))
	// Expected result after formatting
	mockClient.EXPECT().RetrieveSqsMessage(gomock.Any(), gomock.Any(), gomock.Any()).Return("{'finaloutput':{'externaldata':'1a'}}", "secondstring", errors.New("newerrorexpectedout")) //expectedMessage)

	var observedMessage, _, err = main.RetrieveSqsMessage(mockClient, "somepath", false)
	if err != nil {
		log.Fatalf("unable to retrieve mesages: %v", err)
	}

	// This compares the result of the mock with the expectedMessage.
	assert.Equal(t, observedMessage, "{'finaloutput':{'externaldata':'1a'}}", "secondstring", errors.New("newerrorexpectedout")) //expectedMessage)

	print("You are good! Your code works as expected")
}
