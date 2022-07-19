package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	log "11/internal/logging"
	event "11/internal/model"
	"11/internal/usercases"
	"11/internal/validation"
)

type Implementation struct {
	repo   usercases.Repository
	logger log.LoggerEx
}

func New(repo usercases.Repository, logger log.LoggerEx) Implementation {
	return Implementation{repo: repo, logger: logger}
}

func NewServ(repo usercases.Repository, logger log.LoggerEx) *http.ServeMux {
	mux := http.NewServeMux()

	impl := New(repo, logger)

	mux.HandleFunc("/create", impl.Middleware(impl.Create, logger))

	mux.HandleFunc("/delete", impl.Middleware(impl.Delete, logger))

	mux.HandleFunc("/update", impl.Middleware(impl.Update, logger))

	mux.HandleFunc("/today", impl.Middleware(impl.Today, logger))

	mux.HandleFunc("/week", impl.Middleware(impl.Week, logger))

	mux.HandleFunc("/month", impl.Middleware(impl.Month, logger))

	return mux
}

func (i *Implementation) Create(w http.ResponseWriter, r *http.Request) {

	id, date, err := validation.ParseParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	m, err := event.NewEvent(int64(id), date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	model, err := i.repo.Create(m.ID, m.Date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	status, err := w.Write(res)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

}

func (i *Implementation) Update(w http.ResponseWriter, r *http.Request) {
	id, t, err := validation.ParseParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if r.URL.Query()["newTime"][0] == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	newTime, err := validation.ValidateTime(r.URL.Query()["newTime"][0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = i.repo.Update(int64(id), t, newTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(200)
	_, err = w.Write([]byte("All events with updated"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (i *Implementation) Delete(w http.ResponseWriter, r *http.Request) {
	id, time, err := validation.ParseParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = i.repo.Delete(int64(id), time)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = w.Write([]byte("Event deleted"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (i *Implementation) Today(w http.ResponseWriter, r *http.Request) {

	id, _, err := validation.ParseParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := i.repo.Today(int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = w.Write(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (i *Implementation) Week(w http.ResponseWriter, r *http.Request) {

	id, _, err := validation.ParseParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := i.repo.Week(int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = w.Write(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (i *Implementation) Month(w http.ResponseWriter, r *http.Request) {

	id, _, err := validation.ParseParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := i.repo.Month(int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = w.Write(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (i *Implementation) Middleware(next http.HandlerFunc, l log.LoggerEx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(r.URL.Query()["id"][0])
		if err != nil {
			i.logger.WriteInfo(err.Error())
			http.Error(w, "No id in request", http.StatusBadRequest)
		}

		if err = validation.ValidateID(id); err != nil {
			http.Error(w, "id must be 0 < 1 < 2^63", http.StatusBadRequest)
		}

		reqInfo := fmt.Sprintf("Method: %s, id: %s", r.Method, r.URL.Query()["id"][0])
		err = l.WriteInfo(reqInfo)
		if err != nil {
			i.logger.WriteErr(err)
		}

		i.repo.Check(int64(id))

		next(w, r)
	}
}
