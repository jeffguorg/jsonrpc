package v2_test

import (
	jsonrpc "github.com/jeffguorg/jsonrpc/v2"
	"testing"
)

var (
	endpoint = jsonrpc.Endpoint("http://127.0.0.1:6800/jsonrpc")
)

func TestEndpoint_Call(t *testing.T) {
	if res, err := endpoint.Call(jsonrpc.Request{
		Method:     "aria2.addUri",
		ID:         1,
		Parameters: [][]string{
			[]string{"https://mirrors.tuna.tsinghua.edu.cn/archlinux/iso/latest/archlinux-2019.09.01-x86_64.iso"},
		},
	}); err != nil {
		panic(err)
	} else {
		t.Log(res)
	}
}