// Package pbws provides client and server support for a HTTP/protobuf web service transport
package pbws

import (
	"bytes"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
)

type ClientDeps interface {
	// Return arguments message for pbws to fill in
	GetResults() (proto.Message, error)
	GetUrl() string
}

type Client struct {
	deps ClientDeps
}

func NewClient(wc ClientDeps) *Client {
	return &Client{wc}
}

// TODO: string url here? I don't like this
func (self *Client) Call(arguments proto.Message) (proto.Message, error) {
	argumentsData, err := proto.Marshal(arguments)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	request, err := http.NewRequest("POST", self.deps.GetUrl(), bytes.NewReader(argumentsData))
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	resultsData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	results, err := self.deps.GetResults()
	if err != nil {
		return nil, err
	}

	err = proto.Unmarshal(resultsData, results)
	if err != nil {
		return nil, err
	}

	return results, nil
}
