package jsonrpc2

import (
	"encoding/json"
	"net"
	"testing"
	"time"
)

func TestServe_Should_Return_Error_With_Nil_Handlers_Map(t *testing.T) {
	s := Server{}
	err := s.Serve("", nil)
	if err == nil {
		t.Fail()
	}
}

type mockConnection struct {
	Response *Response
}

func (m mockConnection) Read(b []byte) (n int, err error) {
	bb, err := json.Marshal(Request{
		ID:      10,
		Params:  "",
		Method:  "a",
		JSONRpc: JSONRPCVersion,
	})
	copy(b, bb)
	return len(bb), err
}
func (m mockConnection) Write(b []byte) (n int, err error) {
	err = json.Unmarshal(b, m.Response)
	return len(b), err
}
func (m mockConnection) Close() error {
	return nil
}
func (m mockConnection) LocalAddr() net.Addr {
	return nil
}
func (m mockConnection) RemoteAddr() net.Addr {
	return nil
}
func (m mockConnection) SetDeadline(t time.Time) error {
	return nil
}
func (m mockConnection) SetReadDeadline(t time.Time) error {
	return nil
}
func (m mockConnection) SetWriteDeadline(t time.Time) error {
	return nil
}
func (m mockConnection) Handle(r Request) Response {
	return NewResponse(r.ID, "", Error{})
}

func newMockConnection() mockConnection {
	c := mockConnection{}
	c.Response = &Response{}
	return c
}

func TestHandleConnection_Should_Return_Error_Response(t *testing.T) {
	s := Server{}
	c := newMockConnection()
	hs := make(map[string]IHandler)

	hs["a"] = c
	s.HandleConnection(c, &hs)
	if c.Response.ID != 10 {
		t.Fail()
	}
}
