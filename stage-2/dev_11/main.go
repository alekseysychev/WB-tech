package main

import (
	"dev_11/internal/adapter"
	"dev_11/internal/repository"
	"dev_11/internal/service"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func main() {
	repo, err := repository.New()
	if err != nil {
		logrus.Fatalln(err)
	}

	srv, err := service.New(repo)
	if err != nil {
		logrus.Fatalln(err)
	}

	api, err := adapter.New(":9090", srv)
	if err != nil {
		logrus.Fatalln(err)
	}

	// сигналы на прерывание приложения
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)

	// канал для ошибок
	errChannel := make(chan error)

	api.Start(errChannel)

	// ожидание ощибки или выхода
	select {
	case err = <-errChannel:
		if err != nil {
			logrus.Error(err)
		}
	case <-osSignals:
		logrus.Info("osSignals exit")
	}
}
