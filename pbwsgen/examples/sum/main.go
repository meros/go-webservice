// Example server main using generated output from sum.proto
//go:generate ../../pbwsgen -req=SumReq -resp=SumResp -proto=./sum.proto -out_server=server -out_client=client

package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	clib "github.com/meros/go-webservice/pbwsgen/examples/sum/client"
	slib "github.com/meros/go-webservice/pbwsgen/examples/sum/server"
	"net/http"
)

func ws(req *slib.SumReq) *slib.SumResp {
	return &slib.SumResp{Sum: proto.Uint32(req.GetA() + req.GetB())}
}

func runServer() {
	handler := slib.CreateHttpHandler(ws)
	http.Handle("/sum/ws", handler)
	http.ListenAndServe(":8080", nil)
}

func runClient() {
	client := clib.Create("http://localhost:8080/sum/ws")
	_, err := client.Call(&clib.SumReq{A: proto.Uint32(32), B: proto.Uint32(32)})
	if err != nil {
		fmt.Println("Error in web service call", err)
	}
}

func main() {
	go runServer()
	runClient()
}
