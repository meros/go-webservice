// Example server main using generated output from sum.proto

package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/meros/go-webservice/pbwsgen/out/mirror"
	"net/http"
)

func ws(req *mirror.MirrorReq) *mirror.MirrorResp {

	result := []rune(*req.A)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	resp := &mirror.MirrorResp{A: proto.String(string(result))}
	return resp
}

func main() {
	handler := mirror.CreateHttpHandler(ws)
	http.Handle("/mirror/ws", handler)
	http.ListenAndServe(":8080", nil)
}
