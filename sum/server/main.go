// Example webservice, using pbws

package main

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/meros/go-webservice/pbws"
	"github.com/meros/go-webservice/sum/protocol"
	"net/http"
)

type SumServer struct {
}

func (self *SumServer) GetArguments() (proto.Message, error) {
	return &protocol.SumReq{}, nil
}

func (self *SumServer) GetResults(arguments proto.Message) (proto.Message, error) {
	sumReq, ok := arguments.(*protocol.SumReq)
	if !ok {
		return nil, errors.New("Wrong type of message sent to GetResults!")
	}

	results := &protocol.SumResp{Sum: proto.Uint32(sumReq.GetA() + sumReq.GetB())}
	fmt.Println("Results: ", results)
	return results, nil
}

func main() {
	fmt.Println("Will start web server at 8080")

	// Server the web service
	http.Handle("/sum/ws", pbws.NewServer(&SumServer{}))

	// Serve the js client
	webclientPrefix := "/sum/client/"
	fs := http.FileServer(http.Dir("static"))
	http.Handle(webclientPrefix, http.StripPrefix(webclientPrefix, fs))

	http.ListenAndServe(":8080", nil)
}
