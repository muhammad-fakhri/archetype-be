package httputil

import (
	"io/ioutil"
	"net/http"

	json "github.com/json-iterator/go"
)

func ReadRequestBody(r *http.Request, model interface{}) (err error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	defer r.Body.Close()

	err = json.Unmarshal(body, model)
	if err != nil {
		return err
	}

	return nil
}
