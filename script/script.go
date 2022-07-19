package main

import (
	"io/ioutil"

	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

func main() {
	model, err := ioutil.ReadFile("model.json")
	if err != nil {
		logrus.Fatalln(err)
	}

	model1, err := ioutil.ReadFile("1.json")
	if err != nil {
		logrus.Fatalln(err)
	}

	options := stan.NatsURL("nats://localhost:4222")

	stanConn, err := stan.Connect("test-cluster", "pub", options)
	if err != nil {
		logrus.Println(err)
	}

	err = stanConn.Publish("wb", model)
	if err != nil {
		logrus.Println(err)
	}

	err = stanConn.Publish("wb", model1)
	if err != nil {
		logrus.Println(err)
	}
}
