package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/vault/api"
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

type authConfig struct {
	Token     string `yaml:"token"`
	VaultAddr string `yaml:"vault_addr"`
}

func main() {
	key := flag.String("k", "", "key of secret to be retrieved")
	prop := flag.String("p", "", "property to be retrieved from secret value")
	flag.Parse()

	if *key == "" {
		fmt.Println("Must provide key")
		return
	}

	cfg, err := readConfig()
	if err != nil {
		fmt.Printf("Failed to load configuration: %v", err)
		return
	}

	c, err := api.NewClient(&api.Config{
		Address: cfg.VaultAddr,
	})
	if err != nil {
		fmt.Printf("Failed to create Vault client: %v", err)
		return
	}

	c.SetToken(cfg.Token)

	sec, err := c.Logical().Read(*key)
	if err != nil {
		fmt.Printf("Failed to get secret: %v", err)
		return
	}

	if sec == nil || sec.Data == nil {
		fmt.Printf("No data for key %s\n", *key)
		return
	}

	if *prop == "" {
		for _, v := range sec.Data {

			f, err := os.Create("./vault_pass.txt")
			if err != nil {
				fmt.Println(err)
				return
			}
			l, err := f.WriteString(fmt.Sprintf("%v", v))
			if err != nil {
				fmt.Println(err)
				f.Close()
				return
			}
			fmt.Println(l, "bytes for the password written successfully")
			err = f.Close()
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	} else {
		fmt.Printf("%s:%s -> %v\n", *key, *prop, sec.Data[*prop])
	}
}

func readConfig() (authConfig, error) {
	configPath := os.Getenv("CONFIG_FILE")
	if configPath == "" {
		configPath = ".auth.yaml"
	}

	bs, err := ioutil.ReadFile(configPath)
	if err != nil {
		return authConfig{}, errors.Wrap(err, "failed to read configuration file")
	}

	var cfg authConfig
	if err := yaml.Unmarshal(bs, &cfg); err != nil {
		return authConfig{}, errors.Wrap(err, "failed to parse configuration file")
	}

	return cfg, nil
}
