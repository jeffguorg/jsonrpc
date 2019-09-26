package jsonrpc_test

import (
	"github.com/jeffguorg/jsonrpc"
	"testing"
)

var (
	endpoint = jsonrpc.Endpoint("http://127.0.0.1:6800/jsonrpc")
)

func TestEndpoint_Call(t *testing.T) {
	if res, err := endpoint.Call(jsonrpc.Request{
		Method: "aria2.addUri",
		ID:     1,
		Parameters: []interface{}{
			[]string{"https://mirrors.tuna.tsinghua.edu.cn/archlinux/iso/latest/archlinux-2019.09.01-x86_64.iso"},
		},
	}); err != nil {
		panic(err)
	} else {
		t.Log(res)
	}
}
