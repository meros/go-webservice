// Package pbws provides client and server support for a HTTP/protobuf web service transport
package pbws

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io"
	"io/ioutil"
	"net/http"
)

type ServerDeps interface {
	// Return a message instance of correct type wrapped in proto.Message interface for pbws to decode incoming request into
	GetArguments() (proto.Message, error)

	// Return a message instance of correct type and values wrapped in proto.Message interface for pbws to encode and send back to client
	GetResults(arguments proto.Message) (proto.Message, error)
}

type Server struct {
	deps ServerDeps
}

// Create a new server, the returned instance complies with interface http.Handler
func NewServer(deps ServerDeps) *Server {
	return &Server{deps}
}

func (self *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		io.WriteString(w, fmt.Sprintln("Wrong method, only POST accepted, got: ", r.Method))
		return
	}

	argumentsData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		io.WriteString(w, fmt.Sprintln("Failed to read body, err: ", err))
		return
	}

	resultsData, err := self.handle(argumentsData)
	if err != nil {
		io.WriteString(w, fmt.Sprintln("Failed to get response from ws, err:", err))
		return
	}

	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}

	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	_, err = w.Write(resultsData)
	if err != nil {
		io.WriteString(w, fmt.Sprintln("Failed to write response, err:", err))
		return
	}
}

func (self *Server) handle(argumentsData []byte) ([]byte, error) {
	arguments, err := self.deps.GetArguments()
	if err != nil {
		return nil, err
	}

	err = proto.Unmarshal(argumentsData, arguments)
	if err != nil {
		return nil, err
	}

	results, err := self.deps.GetResults(arguments)
	if err != nil {
		return nil, err
	}

	resultsData, err := proto.Marshal(results)
	if err != nil {
		return nil, err
	}

	return resultsData, nil
}
