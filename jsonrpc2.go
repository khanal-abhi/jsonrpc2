package jsonrpc2

import (
	"net"
)

// RPCServerRelated constants
const (
	JSONRPCVersion = "2.0"
)

// ErrorResponse creates a error response for the request
func ErrorResponse(i int64, c int64, m string, d string) Response {
	return Response{
		ID:      i,
		JSONRpc: JSONRPCVersion,
		Result:  "",
		Error: Error{
			Code:    c,
			Message: m,
			Data:    d,
		},
	}
}

// NewResponse creates a new valid response
func NewResponse(i int64, r string, e Error) Response {
	return Response{
		ID:      i,
		JSONRpc: JSONRPCVersion,
		Result:  r,
		Error:   e,
	}
}

// IConnectionHandler handles socket connections
type IConnectionHandler interface {
	HandleConnection(net.Conn, map[string]IHandler)
}

// IServer to serve JSONRPC2 communication
type IServer interface {
	Serve(string, map[string]IHandler) error
}

// IHandler defines an interface for handling JSONRPC2 requests
type IHandler interface {
	Handle(Request) Response
}

// Request is the incoming JSONRPC2 request
type Request struct {
	ID      int64  `json:"id,omitempty"`
	JSONRpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  string `json:"params,omitempty"`
}

// Error is the outgoing error
type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data,omitempty"`
}

// Response is the outgoing JSONRPC2 response
type Response struct {
	ID      int64  `json:"id"`
	JSONRpc string `json:"jsonrpc"`
	Result  string `json:"method,omitempty"`
	Error   Error  `json:"error,omitempty"`
}
