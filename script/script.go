package main

import (
	"io/ioutil"
	"strconv"

	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

func main() {
	options := stan.NatsURL("nats://localhost:4222")

	stanConn, err := stan.Connect("test-cluster", "pub", options)
	if err != nil {
		logrus.Println(err)
	}

	for i := 0; i < 4; i++ {
		s := strconv.Itoa(i)

		model, err := ioutil.ReadFile(s + ".json")
		if err != nil {
			break
		}

		err = stanConn.Publish("wb", model)
		if err != nil {
			logrus.Println(err)
		}

	}
}
