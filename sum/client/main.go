package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/meros/go-webservice/pbws"
	"github.com/meros/go-webservice/sum/protocol"
)

type SumClient struct {
}

func (self *SumClient) GetResults() (proto.Message, error) {
	return &protocol.SumResp{}, nil
}

func (self *SumClient) GetUrl() string {
	return "http://localhost:8080/sum/ws"
}

func main() {
	arguments := &protocol.SumReq{A: proto.Uint32(42), B: proto.Uint32(42)}

	client := pbws.NewClient(&SumClient{})

	results, err := client.Call(arguments)

	if err != nil {
		fmt.Println("Failed to call web service, err = ", err)
		return
	}

	fmt.Println("All ok! Arguments:", arguments, "results:", results)
}
