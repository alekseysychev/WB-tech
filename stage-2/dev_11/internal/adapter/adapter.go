package adapter

import (
	"dev_11/internal/service"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Adapter interface {
	Start(chan error) // запуск сервера
}

type adapter struct {
	addr    string
	mux     *http.ServeMux
	service service.Service
}

func New(addr string, srv service.Service) (Adapter, error) {
	return &adapter{
		addr:    addr,
		mux:     http.NewServeMux(),
		service: srv,
	}, nil
}

func (a *adapter) Start(chanError chan error) {
	a.setEventHandlers()

	a.mux.Handle("/", middleware(logger(a.errorPage())))

	server := http.Server{
		Addr:    a.addr,
		Handler: a.mux,
	}

	go func() {
		logrus.Infof("http сервер запущен по адресу: %s", a.addr)
		chanError <- server.ListenAndServe()
	}()
}

func (a *adapter) errorPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error":"неправильный адрес запроса"}`))
	})
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte(`{"error":"неправильный Content-Type запроса"}`))
			return
		}
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"ошибка разбора запроса"}`))
		}
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func wrongRequest(w http.ResponseWriter, s string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(`{"error":"` + s + `"}`))
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
		logrus.Infof("%s %s", r.Method, r.RequestURI)
	})
}
