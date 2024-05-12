package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"log"
)

func createSQSClient() *sqs.Client {
	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load aws config: %v", err)
	}
	return sqs.NewFromConfig(config)
}

func RetrieveSqsMessage(client *sqs.Client, inputQueueUrl string, deleteAfterReceive bool) (string, string, error) {

	var err error
	err = nil
	receiveMessageInput := sqs.ReceiveMessageInput{
		QueueUrl:            &inputQueueUrl,
		WaitTimeSeconds:     20,
		MaxNumberOfMessages: 1,
	}
	message, err := client.ReceiveMessage(context.TODO(), &receiveMessageInput)
	if err != nil {
		return "", "", err
	}

	// no messages in queue
	if len(message.Messages) == 0 {
		return "", "", nil
	}

	receiptId := *message.Messages[0].ReceiptHandle

	return *message.Messages[0].Body, receiptId, err
}
