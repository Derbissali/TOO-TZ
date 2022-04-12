package handler

import (
	"fmt"
	"net/http"
	"tidy/pkg/model"
)

func (h *Handlers) HashCalc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return

	}
	var M model.Crc64
	var err error

	M.ID, err = h.services.HashCalcService.GetID()
	if err != nil {
		return
	}
	w.Write([]byte("ID: "))
	w.Write([]byte(M.ID))
}
func (h *Handlers) GetHash(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	var M model.Crc64
	var err error
	M.ID = r.RequestURI[18:]
	M.Hash, err = h.services.HashCalcService.GetHash(M.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write([]byte("Hash: "))
	w.Write([]byte(M.Hash))
}
