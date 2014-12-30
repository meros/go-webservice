package pbws

import (
	"bytes"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
)

// TODO: simplify names
type ProtobufWebClient interface {
	// Return arguments message for pbws to fill in
	GetResults() (proto.Message, error)
}

type HttpClient struct {
	wc ProtobufWebClient
}

func NewClient(wc ProtobufWebClient) *HttpClient {
	return &HttpClient{wc}
}

// TODO: string url here? I don't like this
func (self *HttpClient) Call(url string, arguments proto.Message) (proto.Message, error) {
	argumentsData, err := proto.Marshal(arguments)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	request, err := http.NewRequest("POST", url, bytes.NewReader(argumentsData))
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

	results, err := self.wc.GetResults()
	if err != nil {
		return nil, err
	}

	err = proto.Unmarshal(resultsData, results)
	if err != nil {
		return nil, err
	}

	return results, nil
}
