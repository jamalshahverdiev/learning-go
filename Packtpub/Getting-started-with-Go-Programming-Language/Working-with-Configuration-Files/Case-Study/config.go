package wreck

import (
	"os"

	"github.com/spf13/viper"
)

// GlobalConfig stores global configurations
var GlobalConfig = viper.New()

// UserConfig stores user specific configurations
var UserConfig = viper.New()

func init() {
	GlobalConfig.AddConfigPath(".")
	GlobalConfig.AddConfigPath("etc/")
	GlobalConfig.SetConfigName("wreck")
	GlobalConfig.SetDefault("headers.common.content-type", "text/plain")
	GlobalConfig.SetDefault("headers.common.user-agent", "wreck/1.0")
	GlobalConfig.SetDefault("headers.post.content-type", "application/json")
	GlobalConfig.ReadInConfig()
	UserConfig.AddConfigPath(".")
	UserConfig.AddConfigPath(os.Getenv("HOME"))
	UserConfig.SetConfigName(".wreck")
	UserConfig.ReadInConfig()
}
