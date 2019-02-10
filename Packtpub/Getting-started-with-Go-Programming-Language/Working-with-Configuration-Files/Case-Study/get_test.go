package wreck

import "testing"

func TestGetCommand(t *testing.T) {
	content, err := Get("https://httpbin.org/get")
	if err != nil {
		t.Fail()
	}
	if content == "" {
		t.Fail()
	}
}
