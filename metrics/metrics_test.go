package metrics

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

	Register()
}

func TestCustomMetrics(t *testing.T) {
	ts.Start()
	defer ts.Close()

	cases := map[string]map[string]struct {
		endpoint string
		repeat   int
		add      int
		sub      int
		expected []string
	}{
		"Counter": {
			"endpoint_accessed": {
				endpoint: home,
				repeat:   18,
				expected: []string{fmt.Sprintf("endpoint_accessed{endpoint=\"%v\",system=\"%v\"} %v", home, config.System, 18)},
			},
			"application_error": {
				endpoint: errorNotFound,
				repeat:   15,
				expected: []string{fmt.Sprintf("application_error{code=\"%s\",endpoint=\"%v\",system=\"%v\"} %v", "404", errorNotFound, config.System, 15)},
			},
		},
		"Gauge": {
			"active_database_connection": {
				add:      8,
				sub:      6,
				expected: []string{fmt.Sprintf("active_database_connection{system=\"%v\"} %v", config.System, 2)},
			},
		},
		"Histogram": {
			"response_duration_seconds": {
				endpoint: responseDuration,
				repeat:   8,
				expected: []string{
					fmt.Sprintf(`response_duration_seconds_bucket{endpoint="%s",system="%s",le="0.01"}`, responseDuration, config.System),
					fmt.Sprintf(`response_duration_seconds_bucket{endpoint="%s",system="%s",le="0.04"}`, responseDuration, config.System),
					fmt.Sprintf(`response_duration_seconds_bucket{endpoint="%s",system="%s",le="0.16"}`, responseDuration, config.System),
					fmt.Sprintf(`response_duration_seconds_bucket{endpoint="%s",system="%s",le="0.64"}`, responseDuration, config.System),
					fmt.Sprintf(`response_duration_seconds_bucket{endpoint="%s",system="%s",le="2.56"}`, responseDuration, config.System),
					fmt.Sprintf(`response_duration_seconds_bucket{endpoint="%s",system="%s",le="10.24"}`, responseDuration, config.System),
					fmt.Sprintf(`response_duration_seconds_bucket{endpoint="%s",system="%s",le="40.96"}`, responseDuration, config.System),
					fmt.Sprintf(`response_duration_seconds_bucket{endpoint="%s",system="%s",le="+Inf"}`, responseDuration, config.System),
					fmt.Sprintf(`response_duration_seconds_sum{endpoint="%s",system="%s"}`, responseDuration, config.System),
					fmt.Sprintf(`response_duration_seconds_count{endpoint="%s",system="%s"}`, responseDuration, config.System),
				},
			},
		},
	}

	for metricType, config := range cases {
		switch metricType {
		case "Counter":
			for testCase, c := range config {
				sendRequestToEndpoint(c.endpoint, c.repeat)

				actual := getMetrics(t)
				for _, e := range c.expected {
					assert.Containsf(t, actual, e, "Case: %s didn't contain the expected value.", testCase)
				}
			}
		case "Gauge":
			for testCase, c := range config {
				sendRequestToEndpoint(add, c.add)
				sendRequestToEndpoint(sub, c.sub)

				actual := getMetrics(t)
				for _, e := range c.expected {
					assert.Containsf(t, actual, e, "Case: %s didn't contain the expected value.", testCase)
				}
			}
		case "Histogram":
			for testCase, c := range config {
				sendRequestToEndpoint(responseDuration, c.repeat)

				actual := getMetrics(t)
				for _, e := range c.expected {
					assert.Containsf(t, actual, e, "Case: %s didn't contain the expected value.", testCase)
				}
			}
		}
	}
}

func sendRequestToEndpoint(endpoint string, repeat int) {
	for i := 0; i < repeat; i++ {
		http.Get(fmt.Sprintf("http://%v:%v%v", host, port, endpoint))
	}
}

func getMetrics(t *testing.T) string {
	resp, err := http.Get(fmt.Sprintf("http://%v:%v%v", host, port, metrics))
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
