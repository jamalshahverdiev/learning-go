package wreck

import "testing"

func TestPostCommand(t *testing.T) {
	content, err := Post("https://httpbin.org/post")
	if err != nil {
		t.Fail()
	}
	if content == "" {
		t.Fail()
	}
}
