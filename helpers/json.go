package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

// DecodeJSONRequest in a selected struct
func DecodeJSONRequest(r *http.Request, interfaceRef interface{}) error {
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
