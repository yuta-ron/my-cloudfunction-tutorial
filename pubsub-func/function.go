// Package p contains a Pub/Sub Cloud Function.
// Package p contains a Pub/Sub Cloud Function.
package p

import (
	"context"
	"log"
)

// MyScheduledFunction is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// HelloPubSub consumes a Pub/Sub message.
func MyScheduledFunction(ctx context.Context, m PubSubMessage) error {
	log.Println(string("scheduled function"))
	return nil
}
