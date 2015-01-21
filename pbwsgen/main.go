package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	protoFilePtr := flag.String(
		"proto",
		"protocol.proto",
		"proto file containing req/resp definitions")
	protoReqMessagePtr := flag.String(
		"req",
		"Req",
		"req message")
	protoRespMessagePtr := flag.String(
		"resp",
		"Resp",
		"resp message")

	// TODO: flags for:
	// * generator type (go server, go client, js client)
	// * project name

	flag.Parse()

	fmt.Println("Using", *protoFilePtr, "with req message:", *protoReqMessagePtr, "and resp message:", *protoRespMessagePtr)

	// Make sure we have the file specified
	_, err := os.Open(*protoFilePtr)
	if err != nil {
		fmt.Println(err)
		return
	}

	// TODO: based on generator chosen, take appropriate steps to generate code
}
