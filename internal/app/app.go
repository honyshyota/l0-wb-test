package app

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	config "github.com/honyshyota/l0-wb-test/configs"
	"github.com/honyshyota/l0-wb-test/internal/nats"
	"github.com/honyshyota/l0-wb-test/internal/service"
	"github.com/sirupsen/logrus"
)

func Run(config *config.Config) error {
	NatsConn, err := nats.Connection(config)
	if err != nil {
		return err
	}
	logrus.Println("nats-streaming connected")

	service := service.NewService(config)
	logrus.Println("service running")

	err = service.GetAll()
	if err != nil {
		return err
	}

	server := newServer(config, service)

	ch := make(chan struct{})
	go nats.StanSub(ch, NatsConn, service, config)

	go gracefulShutdown(ch,
		[]os.Signal{
			syscall.SIGABRT,
			syscall.SIGQUIT,
			syscall.SIGHUP,
			os.Interrupt,
			syscall.SIGTERM,
		},
		server,
		NatsConn,
		config.DB.DB,
	)

	return server.ListenAndServe()
}

func gracefulShutdown(ch chan struct{}, signals []os.Signal, closeItems ...io.Closer) {
	sign := make(chan os.Signal, 1)
	signal.Notify(sign, signals...)
	sig := <-sign
	log.Printf("Caught signal %s. Shutting down...", sig)
	ch <- struct{}{}
	for _, closer := range closeItems {
		err := closer.Close()
		if err != nil {
			fmt.Printf("failed to close %v: %v", closer, err)
		}
	}
}
