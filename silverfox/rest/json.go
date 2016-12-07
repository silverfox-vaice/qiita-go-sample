package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type JsonParam struct {
	Body map[string]interface{}
}

func ParseHTTPBody(r *http.Request) (JsonParam, error) {
	jParams := JsonParam{}

	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		return jParams, errors.New(fmt.Sprintf("invalid 'Content-Length' header. err = %#v", err))
	}

	if length == 0 {
		return jParams, nil
	}

	if r.Header.Get("Content-Type") != "application/json" {
		return jParams, errors.New("invalid 'Content-Type' header. ｊson is unsupported.")
	}

	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		return jParams, err
	}

	var jsonBody map[string]interface{}
	err = json.Unmarshal(body[:length], &jsonBody)
	if err != nil {
		return jParams, err
	}

	jParams.Body = jsonBody
	return jParams, nil
}
