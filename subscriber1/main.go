package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ersa97/NATS/subscriber1/routes"

	"github.com/ersa97/nats-libs/models"

	"github.com/nats-io/nats.go"

	"github.com/ersa97/nats-libs/services"
)

const (
	subSubjectName = "TESTS.created"
	pubSubjectName = "TESTS.approved"
)

func main() {

	connection := models.NatsConnection{
		Ip:   "127.0.0.1",
		Port: "4222",
	}

	sub, ctx, err := services.SubscribeNats(connection, subSubjectName)
	if err != nil {
		log.Println(err)
	}

	log.Println("module is Ready")

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		msgs, _ := sub.Fetch(10, nats.Context(ctx))
		for _, msg := range msgs {
			msg.Ack()
			var request models.Request
			err := json.Unmarshal(msg.Data, &request)
			if err != nil {
				log.Println(err)
			}
			processData(request)
		}
	}

}

func processData(data models.Request) {
	defer func() { recover() }()

	fmt.Println("==========================")
	fmt.Println("Command : ", data.Command)
	fmt.Println("Param : ", data.Param)
	fmt.Println("Data : ", data.Data)
	fmt.Println("==========================")

	routes.HandleRequest(data.Command, data.Param, data.Data)

}
