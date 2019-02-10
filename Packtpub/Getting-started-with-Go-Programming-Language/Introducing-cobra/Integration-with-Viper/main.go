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

// ./Integration-with-Viper.exe get -u foo -p bar https://httpbin.org/basic-auth/foo/bar
// ./Integration-with-Viper.exe get https://httpbin.org/basic-auth/Jamal/gizli
// ./Integration-with-Viper.exe post -c "Post content to check Code" https://httpbin.org/post
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
		username := viper.GetString("username")
		password := viper.GetString("password")
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
		content := viper.GetString("content")
		if content != "" {
			contentReader = bytes.NewReader([]byte(content))
		}
		req, err := http.NewRequest("POST", args[0], contentReader)
		if err != nil {
			log.Fatalln("Unable to get request")
		}
		username := viper.GetString("username")
		password := viper.GetString("password")
		if username != "" && password != "" {
			req.SetBasicAuth(username, password)
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

func main() {
	cmd.PersistentFlags().StringP("username", "u",
		viper.GetString("credentials.username"), "Username for credentials")
	cmd.PersistentFlags().StringP("password", "p",
		viper.GetString("credentials.password"), "Password for credentials")
	viper.BindPFlag("username", cmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("password", cmd.PersistentFlags().Lookup("password"))
	postCmd.PersistentFlags().StringP("content", "c", "", "Content for POST")
	viper.BindPFlag("content", postCmd.PersistentFlags().Lookup("content"))
	cmd.AddCommand(getCmd)
	cmd.AddCommand(postCmd)
	cmd.Execute()
}

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("cobra")
	viper.ReadInConfig()
}
