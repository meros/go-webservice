// Example server main using generated output from sum.proto
//go:generate ../../../pbwsgen -req=SumReq -resp=SumResp -proto=../sum.proto -out_server=lib

package main

import (
	"github.com/golang/protobuf/proto"
	slib "github.com/meros/go-webservice/pbwsgen/examples/sum/server/lib"
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

func main() {
	runServer()
}
