package adapter

import (
	"dev11/internal/dto"
	"dev11/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

type handler struct {
	service service.Service
}

func (h handler) events() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data []byte
		var err error
		events := h.service.EventsList()
		if data, err = json.Marshal(events); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logrus.Error(err)
			return
		}
		if _, err = w.Write(data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logrus.Error(err)
			return
		}
	})
}

func (h handler) eventsForDay() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data []byte
		var err error
		var userId int
		var date time.Time

		if userId, err = strconv.Atoi(r.URL.Query().Get("user_id")); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logrus.Error(err)
			return
		}

		if date, err = time.Parse("2006-01-02", r.URL.Query().Get("date")); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logrus.Error(err)
			return
		}

		events := h.service.EventsGetForDay(userId, date)

		if data, err = json.Marshal(events); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logrus.Error(err)
			return
		}
		if _, err = w.Write(data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logrus.Error(err)
			return
		}
	})
}

func (h handler) eventsForWeek() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data []byte
		var err error
		var userId int
		var date time.Time

		if userId, err = strconv.Atoi(r.URL.Query().Get("user_id")); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logrus.Error(err)
			return
		}

		if date, err = time.Parse("2006-01-02", r.URL.Query().Get("date")); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logrus.Error(err)
			return
		}

		events := h.service.EventsGetForWeek(userId, date)

		if data, err = json.Marshal(events); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logrus.Error(err)
			return
		}
		if _, err = w.Write(data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logrus.Error(err)
			return
		}
	})
}

func (h handler) eventsForMonth() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data []byte
		var err error
		var userId int
		var date time.Time

		if userId, err = strconv.Atoi(r.URL.Query().Get("user_id")); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logrus.Error(err)
			return
		}

		if date, err = time.Parse("2006-01-02", r.URL.Query().Get("date")); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logrus.Error(err)
			return
		}

		events := h.service.EventsGetForMonth(userId, date)

		if data, err = json.Marshal(events); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logrus.Error(err)
			return
		}
		if _, err = w.Write(data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logrus.Error(err)
			return
		}
	})
}

func (h handler) eventCreate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		var err error
		event := dto.Event{}
		if event.Date, err = time.Parse("2006-01-02", r.FormValue("date")); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logrus.Error(err)
			return
		}
		if event.UserId, err = strconv.Atoi(r.FormValue("user_id")); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logrus.Error(err)
			return
		}
		event.Text = r.FormValue("text")

		h.service.EventCreate(event)
	})
}

func (h handler) eventUpdate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error

		if err = r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logrus.Error(err)
			return
		}

		event := dto.Event{}
		if event.Id, err = strconv.Atoi(r.FormValue("id")); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logrus.Error(err)
			return
		}
		if event.Date, err = time.Parse("2006-01-02", r.FormValue("date")); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logrus.Error(err)
			return
		}
		if event.UserId, err = strconv.Atoi(r.FormValue("user_id")); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logrus.Error(err)
			return
		}
		event.Text = r.FormValue("text")

		h.service.EventUpdate(event)
	})
}

func (h handler) eventDelete() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		var err error
		event := dto.Event{}
		if event.Id, err = strconv.Atoi(r.FormValue("id")); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logrus.Error(err)
			return
		}
		h.service.EventDelete(event)
	})
}

func (h handler) mock() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`test`))
	})
}
