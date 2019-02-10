package wreck

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Get will perform a GET request
func Get(url string) (string, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln("Unable to get request")
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
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
