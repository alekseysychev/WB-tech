package adapter

import (
	"context"
	"dev11/config"
	"dev11/internal/service"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Adapter interface {
	Start() chan error // запуск адаптера
	Stop() error       // остановка адаптера
}

type adapter struct {
	service service.Service
	server  http.Server
}

func New(cfg config.Config, srv service.Service) Adapter {
	mux := http.NewServeMux()

	handlers := handler{
		service: srv,
	}
	mux.Handle("/events", log(get(handlers.events())))
	mux.Handle("/events_for_day", log(get(handlers.eventsForDay())))
	mux.Handle("/events_for_week", log(get(handlers.eventsForWeek())))
	mux.Handle("/events_for_month", log(get(handlers.eventsForMonth())))
	mux.Handle("/create_event", log(post(handlers.eventCreate())))
	mux.Handle("/update_event", log(post(handlers.eventUpdate())))
	mux.Handle("/delete_event", log(post(handlers.eventDelete())))

	return &adapter{
		service: srv,
		server: http.Server{
			Addr:    cfg.Address,
			Handler: mux,
		},
	}
}

func (a *adapter) Start() chan error {
	// канал для получения ошибки в процессе запуска/работы адаптера
	errChannel := make(chan error)
	// старт сервера
	go func() {
		errChannel <- a.server.ListenAndServe()
	}()
	return errChannel
}

func (a *adapter) Stop() error {
	return a.server.Shutdown(context.Background())
}

func log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Infof("%s %s", r.Method, r.URL)
		if r.Method == http.MethodPost {
			r.ParseForm()
			logrus.Infof("%v", r.Form)
		}
		next.ServeHTTP(w, r)
	})
}

func get(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func post(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	})
}
