package v2

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Endpoint string

func (endpoint Endpoint) Call(r Request) (*Response, error) {
	r.Version = "2.0"
	reqbuf, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	log.Print(string(reqbuf))
	res, err := http.Post(string(endpoint), "application/json", bytes.NewReader(reqbuf))
	if err != nil {
		return nil, err
	}
	resbuf, err := ioutil.ReadAll(res.Body)
	log.Print(string(resbuf))
	if err != nil {
		return nil, err
	}
	result := &Response{}
	if err := json.Unmarshal(resbuf, result); err != nil {
		return nil, err
	}
	return result, nil
}