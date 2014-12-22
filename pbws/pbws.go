package pbws

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io"
	"io/ioutil"
	"net/http"
)

type HttpHandler struct {
	ws ProtobufWebService
}

func New(ws ProtobufWebService) *HttpHandler {
	return &HttpHandler{ws}
}

func (self *HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	_, err = w.Write(resultsData)
	if err != nil {
		io.WriteString(w, fmt.Sprintln("Failed to write response, err:", err))
		return
	}
}

func (self *HttpHandler) handle(argumentsData []byte) ([]byte, error) {
	arguments, err := self.ws.GetArguments()
	if err != nil {
		return nil, err
	}

	err = proto.Unmarshal(argumentsData, arguments)
	if err != nil {
		return nil, err
	}

	results, err := self.ws.GetResults(arguments)
	if err != nil {
		return nil, err
	}

	resultsData, err := proto.Marshal(results)
	if err != nil {
		return nil, err
	}

	return resultsData, nil
}

type ProtobufWebService interface {
	// Return arguments message for pbws to fill in
	GetArguments() (proto.Message, error)

	// Return results to send as a response to provided arguments
	GetResults(arguments proto.Message) (proto.Message, error)
}
