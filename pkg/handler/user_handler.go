package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"tidy/pkg/model"
)

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
	M.ID, err = h.services.UserService.Create(&M)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("ID: "))
	w.Write([]byte(strconv.Itoa(M.ID)))

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
		var U model.UpdateU
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = json.Unmarshal(b, &U)
		if err != nil {
			fmt.Println(err)
			fmt.Println("asd")
			return
		}
		err = h.services.UserService.Update(&U, id)
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
