package handler

import (
	"net/http"
	"tidy/pkg/service"
)

type Handlers struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handlers {
	return &Handlers{services: services}
}

func (h *Handlers) Register(router *http.ServeMux) {
	router.HandleFunc("/", h.home_page)
	router.HandleFunc("/rest/substr/find", h.substrfind)
	router.HandleFunc("/rest/email/check", h.emailCheck)
	router.HandleFunc("/rest/iin/check", h.iinCheck)
	router.HandleFunc("/rest/user", h.getuser)
	router.HandleFunc("/rest/user/", h.getOneUser)
	router.HandleFunc("/rest/counter/add/", h.AddCounter)
	router.HandleFunc("/rest/counter/sub/", h.SubCounter)
	router.HandleFunc("/rest/counter/val", h.GetCounter)
	router.HandleFunc("/rest/hash/calc", h.HashCalc)
	router.HandleFunc("/rest/hash/result/", h.GetHash)

}

func (h *Handlers) home_page(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

}
