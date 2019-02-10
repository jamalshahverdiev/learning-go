package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

// type Response struct {
// 	URL     string            `json:"url"`
// 	Headers map[string]string `json:"headers"`
// }

func main() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalln("unable to get response")
	}
	defer resp.Body.Close()
	viper.SetConfigType("json")
	if err := viper.ReadConfig(resp.Body); err != nil {
		log.Fatalln("unable to read body")
	}
	// fmt.Println(viper.AllSettings())
	m := viper.GetStringMapString("headers")
	for k, v := range m {
		fmt.Printf("%s: %s\n", k, v)
	}

	// content, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln("unable to read content")
	// }
	// respContent := Response{}
	// json.Unmarshal(content, &respContent)
	// fmt.Printf("From: %s\n", respContent.URL)
	// for k, v := range respContent.Headers {
	// 	fmt.Printf("%s: %s\n", k, v)
	// }
}
