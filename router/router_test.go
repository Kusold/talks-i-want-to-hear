package router

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestHomeRoute(t *testing.T) {
	Router()
	testServer := httptest.NewServer(nil)
	defer testServer.Close()

	client := &http.Client{}

	req, err := http.NewRequest("GET", testServer.URL, nil)
	if err != nil {
		t.Fatal("Error creating request")
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("Error sending the request")
	}
	if resp.StatusCode != 200 {
		t.Error("Unable to find route. Expected 200, Received", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Unable to read body")
	}
	defer resp.Body.Close()

	if !strings.Contains(string(body), "Hello World") {
		t.Error("Unexpected body")
	}
}

func TestHomeRouteWithGetQueryParameters(t *testing.T) {
	testServer := httptest.NewServer(nil)
	defer testServer.Close()

	client := &http.Client{}

	parameters := url.Values{}
	parameters.Set("query", "bouldergolang")
	req, err := http.NewRequest("GET", testServer.URL+"/?"+parameters.Encode(), nil)
	if err != nil {
		t.Fatal("Error creating request")
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("Error sending the request")
	}
	if resp.StatusCode != 200 {
		t.Error("Unable to find route. Expected 200, Received", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Unable to read body")
	}
	defer resp.Body.Close()

	if !strings.Contains(string(body), "bouldergolang") {
		t.Error("Unexpected body: " + string(body))
	}
}
