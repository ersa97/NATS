package controllers

import (
	"log"
	"net/http"

	"github.com/ersa97/nats-libs/models"

	"github.com/ersa97/nats-libs/services"
)

const (
	streamName     = "TESTS"
	streamSubjects = "TESTS.*"
	subjectName    = "TESTS.created"
)

func TestApi(response http.ResponseWriter, request *http.Request) {

	//ADD LOGIC HERE

	connection := models.NatsConnection{
		Ip:   "127.0.0.1",
		Port: "4222",
	}

	err := services.PublishMessage(models.NatsConnection(connection), models.StreamSubject{
		StreamName:     streamName,
		StreamSubjects: streamSubjects,
		SubjectName:    subjectName,
	}, models.Request{
		Command: "testing",
		Param:   nil,
		Data:    "i'm here",
	})

	log.Println(err)
}
