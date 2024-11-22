package main

import (
	"context"
	"log"
	"log-service/data"
	"time"
)

// declare tpe rpc server
type RPCServer struct {
}

// declare data that we received for any method that ties
type RPCPayload struct {
	Name string
	Data string
}

func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Println("error writing to Mongo", err)
		return err
	}

	*resp = "Process payload via RPC:" + payload.Name
	return nil
}
