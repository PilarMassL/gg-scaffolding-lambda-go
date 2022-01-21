package adapter

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type MyIpClientUsingAws struct{}

var (
	// DefaultHTTPGetAddress Default Address
	defaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	errNoIP = errors.New("no IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	errNon200Response = errors.New("non 200 Response found")
)

func NewMyIpClientUsingAws() *MyIpClientUsingAws {
	return &MyIpClientUsingAws{}
}

func (s *MyIpClientUsingAws) GetIp() (string, error) {
	resp, err := http.Get(defaultHTTPGetAddress)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errNon200Response
	}

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if len(ip) == 0 {
		return "", errNoIP
	}

	return fmt.Sprintf("Hello, %v", string(ip)), nil
}
