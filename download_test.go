package backkit_test

import (
	"encoding/json"
	"testing"

	"github.com/iIIusi0n/backkit"
)

func TestDownloadString(t *testing.T) {
	type testResponse struct {
		Url string `json:"url"`
	}
	testString, err := backkit.DownloadString("https://httpbin.org/get")
	if err != nil {
		t.Error(err.Error())
	}
	var httpbinResponse testResponse
	json.Unmarshal([]byte(testString), &httpbinResponse)
	if httpbinResponse.Url != "https://httpbin.org/get" {
		t.Error("wrong response from httpbin.org")
	}
}
