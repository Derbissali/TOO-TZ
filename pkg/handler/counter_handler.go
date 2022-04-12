package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handlers) AddCounter(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	idstr := r.RequestURI[18:]
	id, err := strconv.Atoi(idstr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	num := strconv.Itoa(id)
	err = h.services.CounterService.AddCounter(num)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%s", err)))
		return
	}
	w.WriteHeader(200)
}

func (h *Handlers) SubCounter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	idStr := r.RequestURI[18:]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	num := strconv.Itoa(id)
	err = h.services.CounterService.SubCounter(num)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%s", err)))
		return
	}
	w.WriteHeader(200)
}

func (h *Handlers) GetCounter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	res, err := h.services.CounterService.GetCounter()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%s", err)))
		return
	}
	w.Write([]byte(res))
}
