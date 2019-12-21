package jsonrpc2

import "testing"

func TestErrorResponse_Should_Return_Valid_Response(t *testing.T) {
	err := ErrorResponse(-1, 400, "Bad Request", "")
	if err.ID != -1 || err.JSONRpc != "2.0" ||
		err.Error.Code != 400 ||
		err.Error.Message != "Bad Request" {
		t.Fail()
	}
}

func TestNewResponse_Should_Return_Valid_Response(t *testing.T) {
	res := NewResponse(-1, "", Error{})
	if res.ID != -1 || res.JSONRpc != "2.0" ||
		res.Result != "" || res.Error.Code != 0 {
		t.Fail()
	}
}
