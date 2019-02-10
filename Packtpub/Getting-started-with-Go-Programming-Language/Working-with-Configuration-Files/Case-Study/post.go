package wreck

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

// Post will perform a POST request
func Post(url string) (string, error) {
	client := http.Client{}
	var contentReader io.Reader
	content := UserConfig.GetString("content")
	if content != "" {
		contentReader = bytes.NewReader([]byte(content))
	}
	req, err := http.NewRequest("POST", url, contentReader)
	if err != nil {
		return "", err
	}
	username := UserConfig.GetString("username")
	password := UserConfig.GetString("password")
	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}
	commonHeaders := GlobalConfig.GetStringMapString("headers.common")
	for k, v := range commonHeaders {
		req.Header.Add(http.CanonicalHeaderKey(k), v)
	}
	postHeaders := GlobalConfig.GetStringMapString("headers.post")
	for k, v := range postHeaders {
		req.Header.Add(http.CanonicalHeaderKey(k), v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(respContent), nil
}
