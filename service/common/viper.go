package common

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgData = viper.New()
)

//InitConfig 读入配置
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
		cfgData.AddConfigPath(".")
		cfgData.AddConfigPath(home)
		cfgData.SetConfigName("myConsultConfig")
	}
	cfgData.AutomaticEnv()
	if err := cfgData.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
	fmt.Printf("\n********config data********\n")
	for k, v := range cfgData.AllSettings() {
		if he, ok := v.(map[string]interface{}); ok {
			fmt.Printf("%s:\n", k)
			for k2, v2 := range he {
				fmt.Printf("  %-18s: %v\n", k2, v2)
			}
		} else {
			fmt.Printf("%-18s: %v\n", k, v)
		}
	}
	fmt.Printf("********config data********\n\n")
	cfgData.WatchConfig()
	cfgData.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err := cfgData.ReadInConfig(); err != nil {
			fmt.Println("Can't read config:", err)
		}
	})
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

//Config 获得配置
func Config() *viper.Viper {
	return cfgData
}
