package jsonrpc2

import (
	"encoding/json"
	"errors"
	"net"
)

// Server is struct for the server
type Server struct{}

// Serve adheres to the server protocol
func (s Server) Serve(p string, hs map[string]IHandler) error {
	if hs == nil {
		return errors.New("nil handlers map")
	}
	l, err := net.Listen("tcp", p)
	if err != nil {
		return err
	}
	for {
		c, err := l.Accept()
		if err != nil {
			go s.HandleConnection(c, &hs)
		}
	}
}

// HandleConnection adheres to IConnectionHandler protocol
func (s Server) HandleConnection(c net.Conn, hs *map[string]IHandler) {
	jsonDecoder := json.NewDecoder(c)
	jsonEncoder := json.NewEncoder(c)

	req := Request{}
	err := jsonDecoder.Decode(&req)
	if err != nil {
		jsonEncoder.Encode(ErrorResponse(-1, 400, "", ""))
	} else {
		method := req.Method
		handler := (*hs)[method]
		if handler == nil {
			jsonEncoder.Encode(ErrorResponse(req.ID, 404, "Not Found", ""))
		}
		res := handler.Handle(req)
		jsonEncoder.Encode(res)
	}
}
