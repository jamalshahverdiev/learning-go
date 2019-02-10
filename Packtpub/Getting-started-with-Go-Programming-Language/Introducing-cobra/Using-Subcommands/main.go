package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// ./Using-Subcommands.exe get -u foo -p bar https://httpbin.org/basic-auth/foo/bar
// ./Using-Subcommands.exe post -c "Post content to check Code" https://httpbin.org/post
var cmd = &cobra.Command{
	Use:   "cobraintro",
	Short: "This tool gets a URL basic auth",
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatalln("must use a subcommand")
	},
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a URL",
	Run: func(cmd *cobra.Command, args []string) {
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

var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Post a URL",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalln("must set URL!")
		}
		client := http.Client{}
		var contentReader io.Reader
		if content != "" {
			contentReader = bytes.NewReader([]byte(content))
		}
		req, err := http.NewRequest("POST", args[0], contentReader)
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

var username, password, content string

func main() {
	cmd.PersistentFlags().StringVarP(&username, "username", "u", "", "Username for credentials")
	cmd.PersistentFlags().StringVarP(&password, "password", "p", "", "Password for credentials")
	postCmd.PersistentFlags().StringVarP(&content, "content", "c", "", "Content for POST")
	cmd.AddCommand(getCmd)
	cmd.AddCommand(postCmd)
	cmd.Execute()
}
