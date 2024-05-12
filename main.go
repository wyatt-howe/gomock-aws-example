package main

import (
	"log"
)

func main() {
	var SqsClient = createSQSClient()

	message, s, err := RetrieveSqsMessage(SqsClient, "somepath", false)
	if err != nil {
		log.Fatalf("unable to retrieve mesages: %v", err)
	}

	log.Println(message, s)
}
