package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func protoc(protoFile string, outDir string) {
	os.MkdirAll(outDir, 0777)

	cmd := exec.Command("protoc", strings.Join([]string{"--go_out=", outDir}, ""), protoFile)

	fmt.Println(cmd)

	err := cmd.Run()

	if err != nil {
		fmt.Println("Failed to run protoc:", err)
	}
}

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

	serviceNamePtr := flag.String(
		"name",
		"service",
		"service name")

	// TODO: flags for:
	// * generator type (go server, go client, js client)
	// * project name

	flag.Parse()

	fmt.Println("Using", *protoFilePtr,
		"with req message:", *protoReqMessagePtr,
		"and resp message:", *protoRespMessagePtr,
		"with service name:", *serviceNamePtr)

	outDir := strings.Join([]string{".", "out", *serviceNamePtr}, "/")

	protoc(*protoFilePtr, outDir)
}
