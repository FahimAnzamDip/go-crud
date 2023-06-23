package utils

import (
	"encoding/json"
	"net/http"
)

func CheckNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func ParseBody(r *http.Request, x interface{}) {
	if err := json.NewDecoder(r.Body).Decode(&x); err != nil {
		return
	}
}
