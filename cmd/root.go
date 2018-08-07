package main

import (
	"fmt"
	"os"

	"github.com/shiqinfeng1/backendside-eth-dev-kits/cmd/command"
	cmn "github.com/shiqinfeng1/backendside-eth-dev-kits/service/common"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/httpservice"
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "{ARCH}_ethsmart",
	Short: "backendside ethereum develop kits",
	Long:  `和以太坊交互的后端服务 `,
}

func init() {
	cobra.OnInitialize(initConfig)
	//Persistent类型的参数表示在子命令种也有效， 相当于全局参数。对应的，设置本地参数的方法是：Flags()
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is ./myConfig.yaml or $HOME/myConfig.yaml)")

}

func initConfig() {
	cmn.InitConfig(cfgFile)
}

func main() {

	command.AddCommands(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		if sterr, ok := err.(httpservice.StatusError); ok {
			if sterr.Status != "" {
				fmt.Println(sterr.Error())
			}
			// StatusError should only be used for errors, and all errors should
			// have a non-zero exit status, so never exit with 0
			if sterr.StatusCode == 0 {
				os.Exit(1)
			}
			os.Exit(sterr.StatusCode)
		}
		fmt.Println(err)
		os.Exit(1)
	}
}
