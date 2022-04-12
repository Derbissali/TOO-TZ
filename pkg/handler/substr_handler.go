package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"tidy/pkg/model"
)

func (h *Handlers) substrfind(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return

	}
	var M model.Substring
	var N model.SubstringN

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
		M.Substring = string(N.SubstringN)

	}
	M.Substring, err = h.services.SubstringService.MaxLength(&M.Substring)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	w.Write([]byte(M.Substring))
}
