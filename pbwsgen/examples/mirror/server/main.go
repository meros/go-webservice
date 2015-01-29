// Example server main using generated output from mirror.proto
//go:generate ../../../pbwsgen -req=MirrorReq -resp=MirrorResp -proto=../mirror.proto -out_server=lib

package main

import (
	"github.com/golang/protobuf/proto"
	slib "github.com/meros/go-webservice/pbwsgen/examples/mirror/server/lib"
	"net/http"
)

func reverseString(input string) string {
	// Get Unicode code points.
	n := 0
	runes := make([]rune, len(input))
	for _, r := range input {
		runes[n] = r
		n++
	}
	runes = runes[0:n]

	// Reverse
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	// Convert back to UTF-8.
	return string(runes)
}

func ws(req *slib.MirrorReq) *slib.MirrorResp {
	return &slib.MirrorResp{A: proto.String(reverseString(req.GetA()))}
}

func runServer() {
	handler := slib.CreateHttpHandler(ws)
	http.Handle("/mirror/ws", handler)
	http.ListenAndServe(":8080", nil)
}

func main() {
	runServer()
}
