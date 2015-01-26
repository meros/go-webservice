// Example server main using generated output from sum.proto
//go:generate ../../pbwsgen -req=SumReq -resp=SumResp -proto=./sum.proto -out_server=server -out_client=client

package main

import (
	"github.com/golang/protobuf/proto"
	sum "github.com/meros/go-webservice/pbwsgen/examples/sum/server"
	"net/http"
)

func ws(req *sum.SumReq) *sum.SumResp {
	return &sum.SumResp{Sum: proto.Uint32(req.GetA() + req.GetB())}
}

func main() {
	handler := sum.CreateHttpHandler(ws)
	http.Handle("/sum/ws", handler)
	http.ListenAndServe(":8080", nil)
}
