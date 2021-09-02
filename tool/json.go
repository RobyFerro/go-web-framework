package tool

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

// DecodeJsonRequest in a selected struct
func DecodeJsonRequest(r *http.Request, interfaceRef interface{}) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, interfaceRef); err != nil {
		if err == io.EOF {
			return errors.New("missing request body")
		}

		return err
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(data))

	return nil
}
