package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// ./Configuration-File-Structure.exe get https://httpbin.org/get
// ./Configuration-File-Structure.exe get -u foo -p bar https://httpbin.org/basic-auth/foo/bar
// ./Configuration-File-Structure.exe get https://httpbin.org/basic-auth/Jamal/gizli
// ./Configuration-File-Structure.exe post -c "Post content to check Code" https://httpbin.org/post
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
		username := userConfig.GetString("username")
		password := userConfig.GetString("password")
		if username != "" && password != "" {
			req.SetBasicAuth(username, password)
		}
		commonHeaders := globalConfig.GetStringMapString("headers.common")
		for k, v := range commonHeaders {
			req.Header.Add(http.CanonicalHeaderKey(k), v)
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
		content := userConfig.GetString("content")
		if content != "" {
			contentReader = bytes.NewReader([]byte(content))
		}
		req, err := http.NewRequest("POST", args[0], contentReader)
		if err != nil {
			log.Fatalln("Unable to get request")
		}
		username := userConfig.GetString("username")
		password := userConfig.GetString("password")
		if username != "" && password != "" {
			req.SetBasicAuth(username, password)
		}
		commonHeaders := globalConfig.GetStringMapString("headers.common")
		for k, v := range commonHeaders {
			req.Header.Add(http.CanonicalHeaderKey(k), v)
		}
		postHeaders := globalConfig.GetStringMapString("headers.post")
		for k, v := range postHeaders {
			req.Header.Add(http.CanonicalHeaderKey(k), v)
		}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln("Unable to get response")
		}
		defer resp.Body.Close()
		respContent, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln("unable to read body")
		}
		fmt.Println(string(respContent))
	},
}

var globalConfig = viper.New()
var userConfig = viper.New()

func main() {
	cmd.PersistentFlags().StringP("username", "u",
		userConfig.GetString("credentials.username"), "Username for credentials")
	cmd.PersistentFlags().StringP("password", "p",
		userConfig.GetString("credentials.password"), "Password for credentials")
	userConfig.BindPFlag("username", cmd.PersistentFlags().Lookup("username"))
	userConfig.BindPFlag("password", cmd.PersistentFlags().Lookup("password"))
	postCmd.PersistentFlags().StringP("content", "c", "", "Content for POST")
	userConfig.BindPFlag("content", postCmd.PersistentFlags().Lookup("content"))
	cmd.AddCommand(getCmd)
	cmd.AddCommand(postCmd)
	cmd.Execute()
}

func init() {
	globalConfig.AddConfigPath(".")
	globalConfig.SetConfigName("global")
	globalConfig.ReadInConfig()
	userConfig.AddConfigPath(".")
	userConfig.SetConfigName("user")
	userConfig.ReadInConfig()
}
