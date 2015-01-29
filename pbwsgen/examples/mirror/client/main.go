// Example client main using generated output from mirror.proto
//go:generate ../../../pbwsgen -req=MirrorReq -resp=MirrorResp -proto=../mirror.proto -out_client=lib
package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	clib "github.com/meros/go-webservice/pbwsgen/examples/mirror/client/lib"
)

func runClient() {
	client := clib.Create("http://localhost:8080/mirror/ws")
	_, err := client.Call(&clib.MirrorReq{A: proto.String("Hello world ☃")})
	if err != nil {
		fmt.Println("Error in web service call", err)
	}
}

func main() {
	runClient()
}
