// Example webservice, using pbws

package main

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/meros/go-webservice/mirror/protocol"
	"github.com/meros/go-webservice/pbws"
	"net/http"
)

type MirrorServer struct {
}

func (self *MirrorServer) GetArguments() (proto.Message, error) {
	return &protocol.MirrorReq{}, nil
}

func (self *MirrorServer) GetResults(arguments proto.Message) (proto.Message, error) {
	req, ok := arguments.(*protocol.MirrorReq)
	if !ok {
		return nil, errors.New("Wrong type of message sent to GetResults!")
	}

	result := []rune(*req.A)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	results := &protocol.MirrorResp{A: proto.String(string(result))}
	fmt.Println("Results: ", results)
	return results, nil
}

func main() {
	fmt.Println("Will start web server at 8080")

	// Server the web service
	http.Handle("/mirror/ws", pbws.NewServer(&MirrorServer{}))

	// Serve the js client
	webclientPrefix := "/mirror/client/"
	fs := http.FileServer(http.Dir("static"))
	http.Handle(webclientPrefix, http.StripPrefix(webclientPrefix, fs))

	http.ListenAndServe(":8080", nil)
}
