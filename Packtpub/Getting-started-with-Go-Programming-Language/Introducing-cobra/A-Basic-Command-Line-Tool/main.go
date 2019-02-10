package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// ./A-Basic-Command-Line-Tool.exe -u foo -p bar https://httpbin.org/basic-auth/foo/bar

var cmd = &cobra.Command{
	Use:   "cobraintro",
	Short: "This tool gets a URL basic auth",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("Hello!")
		if len(args) != 1 {
			log.Fatalln("must set URL!")
		}
		client := http.Client{}
		req, err := http.NewRequest("GET", args[0], nil)
		if err != nil {
			log.Fatalln("Unable to get request")
		}
		if username != "" && password != "" {
			req.SetBasicAuth(username, password)
		}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln("Unable to get response")
		}
		defer resp.Body.Close()
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln("unable to read body")
		}
		fmt.Println(string(content))

	},
}

var username, password string

func main() {
	cmd.PersistentFlags().StringVarP(&username, "username", "u", "", "Username for credentials")
	cmd.PersistentFlags().StringVarP(&password, "password", "p", "", "Password for credentials")
	cmd.Execute()
}
