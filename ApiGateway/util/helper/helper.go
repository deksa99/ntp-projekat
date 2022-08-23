package helper

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetJson(r *http.Response, target interface{}) error {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(r.Body)
	return json.NewDecoder(r.Body).Decode(target)
}
