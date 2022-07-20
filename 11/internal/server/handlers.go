package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	log "11/internal/logging"
	event "11/internal/model"
	"11/internal/usecase"
	"11/internal/validation"
)

type Handler struct {
	repo   usecase.Repository
	logger log.LoggerEx
}

func New(repo usecase.Repository, logger log.LoggerEx) Handler {
	return Handler{repo: repo, logger: logger}
}

func Register(repo usecase.Repository, logger log.LoggerEx) *http.ServeMux {
	h := New(repo, logger)

	mux := http.NewServeMux()
	mux.HandleFunc("/create", h.Middleware(h.Create, logger))
	mux.HandleFunc("/delete", h.Middleware(h.Delete, logger))
	mux.HandleFunc("/update", h.Middleware(h.Update, logger))
	mux.HandleFunc("/today", h.Middleware(h.Today, logger))
	mux.HandleFunc("/week", h.Middleware(h.Week, logger))
	mux.HandleFunc("/month", h.Middleware(h.Month, logger))

	return mux
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

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

	model, err := h.repo.Create(m.ID, m.Date)
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

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
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

	err = h.repo.Update(int64(id), t, newTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(200)
	_, err = w.Write([]byte("All events with updated"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, time, err := validation.ParseParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.repo.Delete(int64(id), time)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = w.Write([]byte("Event deleted"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) Today(w http.ResponseWriter, r *http.Request) {

	id, _, err := validation.ParseParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.repo.Today(int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = w.Write(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) Week(w http.ResponseWriter, r *http.Request) {

	id, _, err := validation.ParseParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.repo.Week(int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = w.Write(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) Month(w http.ResponseWriter, r *http.Request) {

	id, _, err := validation.ParseParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.repo.Month(int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = w.Write(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) Middleware(next http.HandlerFunc, l log.LoggerEx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(r.URL.Query()["id"][0])
		if err != nil {
			h.logger.WriteInfo(err.Error())
			http.Error(w, "No id in request", http.StatusBadRequest)
		}

		if err = validation.ValidateID(id); err != nil {
			http.Error(w, "id must be 0 < 1 < 2^63", http.StatusBadRequest)
		}

		reqInfo := fmt.Sprintf("Method: %s, id: %s", r.Method, r.URL.Query()["id"][0])
		err = l.WriteInfo(reqInfo)
		if err != nil {
			h.logger.WriteErr(err)
		}

		h.repo.Check(int64(id))

		next(w, r)
	}
}
