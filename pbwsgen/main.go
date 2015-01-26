//go:generate go-bindata templates/...

package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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

type commonTemplateData struct {
	Name            string
	ProtocolPackage string
	Req             string
	Resp            string
}

func generateServer(packageName string, protoFile string, reqName string, respName string, outDir string) error {
	protoc(protoFile, outDir)

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

	data := &commonTemplateData{
		Name: packageName,
		Req:  reqName,
		Resp: respName}
	err = template.Execute(file, data)
	if err != nil {
		return err
	}

	return nil
}

func generateClient(packageName string, protoFile string, reqName string, respName string, outDir string) error {
	protoc(protoFile, outDir)

	file, err := os.Create(strings.Join([]string{outDir, "client.go"}, "/"))
	if err != nil {
		return err
	}

	defer file.Close()

	template := template.New("clientTemplate")

	templateData, err := Asset("templates/client/client.go.template")
	if err != nil {
		return err
	}

	template, err = template.Parse(string(templateData))
	if err != nil {
		return err
	}

	data := &commonTemplateData{
		Name: packageName,
		Req:  reqName,
		Resp: respName}
	err = template.Execute(file, data)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var protoFile string
	var protoReqMessage string
	var protoRespMessage string
	var outDirServer string
	var outDirClient string

	flag.StringVar(
		&protoFile,
		"proto",
		"protocol file",
		"proto file (wihtout .proto) containing req/resp definitions, this stripped name of this file will also be the package name for the protocol")
	flag.StringVar(
		&protoReqMessage,
		"req",
		"Req",
		"req message")
	flag.StringVar(
		&protoRespMessage,
		"resp",
		"Resp",
		"resp message")
	flag.StringVar(
		&outDirServer,
		"out_server",
		"",
		"out dir for server lib (include to enable generation of server)")
	flag.StringVar(
		&outDirClient,
		"out_client",
		"",
		"out dir for client lib (include to enable generation of client)")

	flag.Parse()

	if outDirClient == "" && outDirServer == "" {
		flag.PrintDefaults()
		return
	}

	packageName := filepath.Base(protoFile)
	packageName = packageName[0 : len(packageName)-len(filepath.Ext(packageName))]

	fmt.Println("Using", protoFile,
		"with req message:", protoReqMessage,
		"and resp message:", protoRespMessage,
		"and package name:", packageName)

	if outDirServer != "" {
		err := generateServer(packageName, protoFile, protoReqMessage, protoRespMessage, outDirServer)
		if err != nil {
			fmt.Println("Error while generating server:", err)
		}
	}

	if outDirClient != "" {
		err := generateClient(packageName, protoFile, protoReqMessage, protoRespMessage, outDirClient)
		if err != nil {
			fmt.Println("Error while generating client:", err)
		}
	}
}
