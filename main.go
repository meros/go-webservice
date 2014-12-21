package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/meros/go-webservice/protocol"
	"io"
	"io/ioutil"
	"net/http"
)

// TODO: move to another package
type ProtobufService interface {
	ProvideArgumentMessage() proto.Message
	Handle() proto.Message
}

// TODO: move to another package
type TestServer struct {
	arguments protocol.Arguments
}

func (self *TestServer) ProvideArgumentMessage() proto.Message {
	return &self.arguments
}

func (self *TestServer) Handle() proto.Message {

	fmt.Println("In testserver/handle!")

	return nil
}

func getTestServer() ProtobufService {
	return &TestServer{}
}

// TODO: change to Handler type
func handleProtobufServiceCall(w http.ResponseWriter, r *http.Request, getProtobufService func() ProtobufService) {

	service := getProtobufService()

	if r.Method != "POST" {
		io.WriteString(w, "Unsupported method!")
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		io.WriteString(w, "Failed to read body")
		return
	}

	arguments := service.ProvideArgumentMessage()
	err = proto.Unmarshal(body, arguments)

	if err != nil {
		io.WriteString(w, "Failed to unmarshal argument")
		return
	}

	results := service.Handle()
	io.WriteString(w, "TODO: Result here")
}

func main() {
	fmt.Println("Will start web server at 8080")

	http.HandleFunc("/ws/test", func(w http.ResponseWriter, r *http.Request) {
		handleProtobufServiceCall(w, r, getTestServer)
	})
	http.ListenAndServe(":8080", nil)

	fmt.Println("Exit?")
}
