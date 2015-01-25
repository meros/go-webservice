//go:generate go-bindata templates/...

package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

// Run protoc of file with go as output (requires go protobuf plugin in path)
func protoc(protoFile string, outDir string) {
	os.MkdirAll(outDir, 0777)

	cmd := exec.Command("protoc",
		strings.Join([]string{"--go_out=", outDir}, ""),
		strings.Join([]string{protoFile}, ""))

	err := cmd.Run()

	if err != nil {
		fmt.Println("Failed to run protoc:", err)
	}
}

type serverTemplateData struct {
	Name            string
	ProtocolPackage string
	Req             string
	Resp            string
}

func generateServer(protoFile string, reqName string, respName string, outDir string) error {
	protoc(strings.Join([]string{protoFile, ".proto"}, ""), outDir)

	file, err := os.Create(strings.Join([]string{outDir, "server.go"}, "/"))
	if err != nil {
		return err
	}

	defer file.Close()

	template := template.New("serverTemplate")

	templateData, err := Asset("templates/server/server.go.template")
	if err != nil {
		return err
	}

	template, err = template.Parse(string(templateData))
	if err != nil {
		return err
	}

	data := &serverTemplateData{
		Name: protoFile,
		Req:  reqName,
		Resp: respName}
	err = template.Execute(file, data)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	protoFilePtr := flag.String(
		"proto",
		"protocol",
		"proto file (wihtout .proto) containing req/resp definitions, this will also be the package name for the protocol")
	protoReqMessagePtr := flag.String(
		"req",
		"Req",
		"req message")
	protoRespMessagePtr := flag.String(
		"resp",
		"Resp",
		"resp message")
	outDirPtr := flag.String(
		"out",
		"out",
		"out dir")

	// TODO: flags for:
	// * generator type (go server, go client, js client)
	// * project name

	flag.Parse()

	fmt.Println("Using", *protoFilePtr,
		"with req message:", *protoReqMessagePtr,
		"and resp message:", *protoRespMessagePtr,
		"and out dir:", *outDirPtr)

	err := generateServer(*protoFilePtr, *protoReqMessagePtr, *protoRespMessagePtr, *outDirPtr)
	if err != nil {
		fmt.Println("Error while generating server:", err)
	}
}
