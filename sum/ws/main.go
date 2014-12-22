package main

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/meros/go-webservice/pbws"
	"github.com/meros/go-webservice/sum/protocol"
	"net/http"
)

type SumService struct {
}

func (self *SumService) GetArguments() (proto.Message, error) {
	return &protocol.SumReq{}, nil
}

func (self *SumService) GetResults(arguments proto.Message) (proto.Message, error) {
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

	http.Handle("/ws/sum", pbws.New(&SumService{}))
	http.ListenAndServe(":8080", nil)
}
