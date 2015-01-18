package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/meros/go-webservice/pbws"
	"github.com/meros/go-webservice/mirror/protocol"
)

type MirrorClient struct {
}

func (self *MirrorClient) GetResults() (proto.Message, error) {
	return &protocol.MirrorResp{}, nil
}

func (self *MirrorClient) GetUrl() string {
	return "http://localhost:8080/mirror/ws"
}

func main() {
	arguments := &protocol.MirrorReq{A: proto.String("Testing123")}

	client := pbws.NewClient(&MirrorClient{})

	results, err := client.Call(arguments)

	if err != nil {
		fmt.Println("Failed to call web service, err = ", err)
		return
	}

	fmt.Println("All ok! Arguments:", arguments, "results:", results)
}
