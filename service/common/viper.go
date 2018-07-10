package common

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var (
	cfgData = viper.New()
)

func InitConfig(cfgFile string) {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		cfgData.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra" (without extension).
		cfgData.AddConfigPath(home)
		cfgData.SetConfigName("myConsultConfig.yaml")
	}
	cfgData.AutomaticEnv()
	if err := cfgData.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
	/*
		get method:
		Get(key string) : interface{}
		GetBool(key string) : bool
		GetFloat64(key string) : float64
		GetInt(key string) : int
		GetString(key string) : string
		GetStringMap(key string) : map[string]interface{}
		GetStringMapString(key string) : map[string]string
		GetStringSlice(key string) : []string
		GetTime(key string) : time.Time
		GetDuration(key string) : time.Duration
		IsSet(key string) : bool
	*/

}

func Config() {
	return cfgData
}
