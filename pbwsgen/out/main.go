// Example server main using generated output from sum.proto

package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/meros/go-webservice/pbwsgen/out/sum"
)

func ws(req *sum.SumReq) *sum.SumResp {
	return &sum.SumResp{Sum: proto.Uint32(req.GetA() + req.GetB())}
}

func main() {
	sum.Run(ws)
}
