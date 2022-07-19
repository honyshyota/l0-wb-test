package nats

import (
	config "github.com/honyshyota/l0-wb-test/configs"
	"github.com/honyshyota/l0-wb-test/internal/service"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

func Connection(config *config.Config) (stan.Conn, error) {
	options := stan.NatsURL(config.NatsStreaming.Path + config.NatsStreaming.Host + config.NatsStreaming.Port)

	stanConn, err := stan.Connect(config.NatsStreaming.Cluster, config.NatsStreaming.ClientID, options)
	if err != nil {
		return nil, err
	}

	return stanConn, nil
}

func StanSub(ch chan struct{}, conn stan.Conn, service service.Service, config *config.Config) {
	stanOpt := stan.StartAtTimeDelta(config.NatsStreaming.Time)

	sub, err := conn.QueueSubscribe("wb", "wb", func(m *stan.Msg) {
		logrus.Println("get from nats-streaming: ", string(m.Data))
		service.Set(m.Data)
	}, stanOpt, stan.DurableName("wb"))
	if err != nil {
		logrus.Fatal(err.Error())
	}

	defer func() {
		sub.Unsubscribe()
		sub.Close()
	}()

	<-ch
}
