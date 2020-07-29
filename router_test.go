package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nszilard/prometheus-custom-metrics/config"
	"github.com/stretchr/testify/assert"
)

var ts *httptest.Server

func init() {
	var err error

	ts, err = getTestServer()
	if err != nil {
		panic(fmt.Sprintf("unable to start test server: %v", err))
	}
}

func TestMetrics(t *testing.T) {
	ts.Start()
	defer ts.Close()

	sendRequestToEndpoint(pathHome, 1)
	sendRequestToEndpoint(pathVersion, 2)
	sendRequestToEndpoint(pathHealthcheck, 3)
	sendRequestToEndpoint(pathError, 4)
	sendRequestToEndpoint(pathCreateConnection, 6)
	sendRequestToEndpoint(pathTerminateConnection, 5)
	sendRequestToEndpoint(pathRandom, 6)

	observedMetrics := getMetrics(t)

	expected := map[string][]string{
		"active_database_connections": {
			fmt.Sprintf(`active_database_connection{system="%s"} %v`, config.System, 1),
		},
		"application_error": {
			fmt.Sprintf(`application_error{code="%s",endpoint="%s",system="%s"} %v`, "404", pathError, config.System, 4),
		},
		"endpoint_accessed": {
			fmt.Sprintf(`endpoint_accessed{endpoint="%s",system="%s"} %v`, pathHome, config.System, 1),
			fmt.Sprintf(`endpoint_accessed{endpoint="%s",system="%s"} %v`, pathVersion, config.System, 2),
			fmt.Sprintf(`endpoint_accessed{endpoint="%s",system="%s"} %v`, pathHealthcheck, config.System, 3),
			fmt.Sprintf(`endpoint_accessed{endpoint="%s",system="%s"} %v`, pathError, config.System, 4),
			fmt.Sprintf(`endpoint_accessed{endpoint="%s",system="%s"} %v`, pathTerminateConnection, config.System, 5),
			fmt.Sprintf(`endpoint_accessed{endpoint="%s",system="%s"} %v`, pathCreateConnection, config.System, 6),
			fmt.Sprintf(`endpoint_accessed{endpoint="%s",system="%s"} %v`, pathRandom, config.System, 6),
		},
		"response_duration": {
			fmt.Sprintf(`response_duration_seconds_count{endpoint="%s",system="%s"} %v`, pathHome, config.System, 1),
			fmt.Sprintf(`response_duration_seconds_count{endpoint="%s",system="%s"} %v`, pathRandom, config.System, 6),
		},
	}

	for m, lines := range expected {
		for _, line := range lines {
			assert.Containsf(t, observedMetrics, line, "Metric: %s didn't return the expected value", m)
		}
	}
}

func sendRequestToEndpoint(endpoint string, repeat int) {
	for i := 0; i < repeat; i++ {
		http.Get(fmt.Sprintf("http://127.0.0.1:%v%v", config.TestPort, endpoint))
	}
}

func getMetrics(t *testing.T) string {
	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%v%v", config.TestPort, pathMetrics))
	if err != nil {
		t.Fatalf("get metrics: unexpected error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("get metrics: unexpected status code: %v", resp.StatusCode)
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		t.Fatalf("get metrics: unable to read response: %v", err)
	}

	return buf.String()
}
