package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CheckNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func ParseBody(w http.ResponseWriter, r *http.Request, x interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &x); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
