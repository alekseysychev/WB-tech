package adapter

import (
	"dev_11/internal/dto"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func (a *adapter) setEventHandlers() {
	a.mux.Handle("/create_event", middleware(logger(a.createEvent())))
	a.mux.Handle("/update_event", middleware(logger(a.updateEvent())))
	a.mux.Handle("/delete_event", middleware(logger(a.deleteEvent())))
	a.mux.Handle("/events_for_day", middleware(logger(a.getEvents(time.Now().AddDate(0, 0, -1)))))
	a.mux.Handle("/events_for_week", middleware(logger(a.getEvents(time.Now().AddDate(0, 0, -7)))))
	a.mux.Handle("/events_for_month", middleware(logger(a.getEvents(time.Now().AddDate(0, -1, 0)))))
}

func (a *adapter) createEvent() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"неправильный тип запроса"}`))
			return
		}

		userId, ok := eventValidation(w, r, "user_id")
		if !ok {
			return
		}
		name, ok := eventValidation(w, r, "name")
		if !ok {
			return
		}
		text, ok := eventValidation(w, r, "text")
		if !ok {
			return
		}
		date, ok := eventValidation(w, r, "date")
		if !ok {
			return
		}

		err := a.service.Create(dto.Event{
			User: userId.(int),
			Name: name.(string),
			Text: text.(string),
			Date: date.(time.Time),
		})
		if err != nil {
		}
	})
}

func (a *adapter) updateEvent() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"неправильный тип запроса"}`))
			return
		}
		w.Write([]byte("Это моя домашняя страница"))
	})
}

func (a *adapter) deleteEvent() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"неправильный тип запроса"}`))
			return
		}
		w.Write([]byte("Это моя домашняя страница"))
	})
}

func (a *adapter) getEvents(start time.Time) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"неправильный тип запроса"}`))
			return
		}

		fmt.Println(r.Form)
		fmt.Println(r.ParseForm())
		fmt.Println(r.Form)

		w.Write([]byte("Это моя домашняя страница"))
	})
}

func eventValidation(w http.ResponseWriter, r *http.Request, fieldName string) (interface{}, bool) {
	fieldData, ok := r.Form[fieldName]
	if !ok {
		wrongRequest(w, "поле "+fieldName+" обязательно")
		return nil, false
	}

	if len(fieldData) != 1 {
		wrongRequest(w, "поле "+fieldName+" должно быть задано в единственном числе")
		return nil, false
	}

	if fieldData[0] == "" {
		wrongRequest(w, "поле "+fieldName+" не может быть пустым")
		return nil, false
	}

	var data interface{}
	var err error

	switch fieldName {
	case "id":
		fallthrough
	case "user_id":
		data, err = strconv.Atoi(fieldData[0])
		if err != nil {
			wrongRequest(w, "поле "+fieldName+" некорректно")
			return nil, false
		}
	case "name":
		fallthrough
	case "text":
		data = fieldData[0]
	case "date":
		data, err = time.Parse("2006-05-04", fieldData[0])
		if err != nil {
			wrongRequest(w, "поле "+fieldName+" должно быть задано в единственном числе")
			return nil, false
		}
	}
	return data, true
}
