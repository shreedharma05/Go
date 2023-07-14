package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func parseBody(r *http.Request, x interface{}) {
	if data, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(data), x); err != nil {
			return
		}
	}
}
