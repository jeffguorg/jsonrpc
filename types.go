package jsonrpc

import "fmt"

type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

var _ error = Error{}

func (err Error) Error() string {
	return fmt.Sprintf("JsonRPC Error(%d): %s(%+v)", err.Code, err.Message, err.Data)
}

type Request struct {
	Version    string        `json:"jsonrpc"`
	Method     string        `json:"method"`
	ID         interface{}   `json:"id"`
	Parameters []interface{} `json:"params,omitempty"`
}

type Response struct {
	Version string
	Error   *Error
	ID      interface{}
	Result  interface{}
}
