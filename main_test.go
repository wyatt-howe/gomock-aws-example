package main_test

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	mock_main "gomock-aws-example/mocks"
	"testing"
)

type Client_Interface interface {
	RetrieveSqsMessage(*sqs.Client, string, bool) (string, string, error)
	ReceiveMessage(ctx context.Context, params *sqs.ReceiveMessageInput, optFns ...func(*sqs.Options)) (*sqs.ReceiveMessageOutput, error)
}

// This test validates the behavior of Retrieve Sqs message Test. Uses a mock-client's Receive message.
func TestRetrieveSqsMessageTest(t *testing.T) {

	// Create a new instance of the GoMock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_main.NewMockClient_Interface(ctrl)

	mockClient.EXPECT().ReceiveMessage(gomock.Any(), gomock.Any()).Return(nil, errors.New("newerrorexpectedout")) //expectedMessage)

	var err error
	err = nil
	var inputQueueUrl = "some"
	receiveMessageInput := sqs.ReceiveMessageInput{
		QueueUrl:            &inputQueueUrl,
		WaitTimeSeconds:     20,
		MaxNumberOfMessages: 1,
	}
	// Using the mocked sqs client.
	message, err := mockClient.ReceiveMessage(context.TODO(), &receiveMessageInput)
	_ = err
	assert.Equal(t, nil, errors.New("newerrorexpectedout"), message)

}

//
//func TestUseClient(t *testing.T) {
//
//	// Create a new instance of the GoMock controller
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	// Create a mock object using the mock interface
//	mockClient := mock_main.NewMockClient_Interface(ctrl)
//
//	// Set up expectations for the useClient method
//	//var expectedMessage = ("{'finaloutput':{'externaldata':'1a'}}", "secondstring", errors.New("newerrorexpectedout"))
//	// Expected result after formatting
//	mockClient.EXPECT().RetrieveSqsMessage(gomock.Any(), gomock.Any(), gomock.Any()).Return("{'finaloutput':{'externaldata':'1a'}}", "secondstring", errors.New("newerrorexpectedout")) //expectedMessage)
//
//	var observedMessage, _, err = main.RetrieveSqsMessage(mockClient, "somepath", false)
//	if err != nil {
//		log.Fatalf("unable to retrieve mesages: %v", err)
//	}
//
//	// This compares the result of the mock with the expectedMessage.
//	assert.Equal(t, observedMessage, "{'finaloutput':{'externaldata':'1a'}}", "secondstring", errors.New("newerrorexpectedout")) //expectedMessage)
//
//	print("You are good! Your code works as expected")
//}
