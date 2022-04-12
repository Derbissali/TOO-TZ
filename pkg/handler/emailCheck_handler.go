package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"tidy/pkg/model"
)

func (h *Handlers) iinCheck(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return

	}
	var M model.IinCheck
	var N model.IinCheckN

	var err error
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(b, &M)
	if err != nil {
		err = json.Unmarshal(b, &N)
		if err != nil {
			return
		}
		M.Iin = string(N.IinN)

	}

	M.Iin, err = h.services.EmailCheckService.IinCheck(M.Iin)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	w.Write([]byte(M.Iin))
}

func (h *Handlers) emailCheck(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return

	}
	var M model.EmailCheck
	var err error
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(b, &M)
	if err != nil {
		fmt.Println(err)
		return
	}
	M.Email, err = h.services.EmailCheckService.EmailCheck(M.Email) //h.emailService.EmailCheck(M.Email)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	w.Write([]byte(M.Email))
}
