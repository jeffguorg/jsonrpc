package v2

type Error struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Request struct {
	Version    string `json:"jsonrpc"`
	Method     string `json:"method"`
	ID         interface{} `json:"id"`
	Parameters interface{} `json:"params,omitempty"`
}

type Response struct {
	Version string
	Error   *Error
	ID      interface{}
	Result  interface{}
}
