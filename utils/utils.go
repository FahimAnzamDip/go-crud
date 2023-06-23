package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func CheckNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err != nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
