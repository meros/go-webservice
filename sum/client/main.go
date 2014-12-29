package main

import (
	"bytes"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/meros/go-webservice/sum/protocol"
	"io/ioutil"
	"net/http"
)

func main() {
	// TODO: move most of the code here to pbws/client.go
	arguments := &protocol.SumReq{A: proto.Uint32(42), B: proto.Uint32(42)}
	argumentsData, err := proto.Marshal(arguments)
	if err != nil {
		fmt.Println("Failed to marshal arguments")
		return
	}

	client := &http.Client{}

	request, err := http.NewRequest("POST", "http://localhost:8080/sum/ws", bytes.NewReader(argumentsData))
	if err != nil {
		fmt.Println("Failed to create http request")
		return
	}

	response, err := client.Do(request)

	if err != nil {
		fmt.Println("Failed to do request")
		return
	}

	resultsData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Failed to read response")
		return
	}

	results := &protocol.SumResp{}
	err = proto.Unmarshal(resultsData, results)
	if err != nil {
		fmt.Println("Failed to unmarshal results")
		return
	}

	fmt.Println("All ok! Arguments:", arguments, "results:", results)
}
