package main

import (
	"fmt"
	"net"
	"net/http/httptest"

	"github.com/nszilard/prometheus-custom-metrics/config"
)

func getTestServer() (*httptest.Server, error) {
	l, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%v", config.TestPort))
	if err != nil {
		return nil, err
	}

	ts := httptest.NewUnstartedServer(&requestHandler{})
	ts.Listener = l
	return ts, nil
}
