package pbws

import (
	"bytes"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
)

type ClientDeps interface {
	// Return a message instance of correct type wrapped in proto.Message interface for pbws to decode incoming result into
	GetResults() (proto.Message, error)
	// Return an url to connect to
	GetUrl() string
}

type Client struct {
	deps ClientDeps
}

// Create a new client
func NewClient(wc ClientDeps) *Client {
	return &Client{wc}
}

// Call web service
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
