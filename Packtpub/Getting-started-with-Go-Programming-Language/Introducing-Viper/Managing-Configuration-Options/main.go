package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var username, password string

// func useCredentials(username, password string) {
func useCredentials() {
	fmt.Printf("Your username is %s\n", viper.GetString("credentials.username"))
	fmt.Printf("Your password is %s\n", viper.GetString("credentials.password"))
}

func main() {
	flag.StringVar(&username, "user", "", "username to use")
	flag.StringVar(&password, "password", "", "password to use")
	flag.Parse()
	viper.AddConfigPath(".")
	viper.SetConfigName("creds")
	// if err := viper.ReadInConfig(); err != nil {
	// 	log.Fatalln("unable to read config")
	// }
	viper.ReadInConfig()
	// if username == "" || password == "" {
	// 	log.Fatalln("must pass credentials!")
	// }
	if username != "" {
		viper.Set("credentials.username", username)
	}
	if password != "" {
		viper.Set("credentials.password", password)
	}
	if !viper.IsSet("credentials.username") || !viper.IsSet("credentials.password") {
		log.Fatalln("no credentials supllied")
	}
	// viper.Set("credentials.username", username)
	// viper.Set("credentials.password", password)
	//useCredentials(username, password)
	useCredentials()
}
