// Example client main using generated output from sum.proto
//go:generate pbwsgen -req=SumReq -resp=SumResp -proto=../sum.proto -out_client=lib

package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	clib "github.com/meros/go-webservice/pbwsgen/examples/sum/client/lib"
)

func runClient() {
	client := clib.Create("http://localhost:8080/sum/ws")
	_, err := client.Call(&clib.SumReq{A: proto.Uint32(32), B: proto.Uint32(32)})
	if err != nil {
		fmt.Println("Error in web service call", err)
	}
}

func main() {
	runClient()
}
