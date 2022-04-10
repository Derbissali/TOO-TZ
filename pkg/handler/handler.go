package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"tidy/pkg/model"
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

}
func (h *Handlers) substrfind(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return

	}
	var M model.Substring
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
	M.Substring, err = h.services.SubstringService.MaxLength(&M.Substring)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	w.Write([]byte(M.Substring))
}

func (h *Handlers) home_page(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

}
func (h *Handlers) getuser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return

	}
	M := model.User{}
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
	err = h.services.UserService.Create(&M)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// temp.Execute(w, M)

}

func (h *Handlers) getOneUser(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/rest/user/") {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	var M model.User
	id := r.RequestURI[11:]
	switch r.Method {
	case "GET":
		M, err := h.services.UserService.ReadOne(id)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		w.Write([]byte("first_name: "))
		w.Write([]byte(M.Name))
		w.Write([]byte(", "))
		w.Write([]byte("last_name: "))
		w.Write([]byte(M.Surname))
		return
	case "PUT":
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = json.Unmarshal(b, &M)
		if err != nil {
			fmt.Println(err)
			fmt.Println("asd")
			return
		}
		err = h.services.UserService.Update(&M, id)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	case "DELETE":
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = json.Unmarshal(b, &M)
		if err != nil {
			fmt.Println(err)
			fmt.Println("asd")
			return
		}
		err = h.services.UserService.Delete(id)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}
}
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
