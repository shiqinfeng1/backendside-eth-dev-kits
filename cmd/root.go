package main

import (
	"fmt"
	"os"

	"github.com/shiqinfeng1/chunyuyisheng/cmd/command"
	cmn "github.com/shiqinfeng1/chunyuyisheng/service/common"
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "{ARCH}_myconsult",
	Short: "the consultation platform",
	Long:  `'春雨医生' 和 '国医链' 的问诊对接平台 `,
}

func init() {
	cobra.OnInitialize(initConfig)
	//Persistent类型的参数表示在子命令种也有效， 相当于全局参数。对应的，设置本地参数的方法是：Flags()
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is ./myConsultConfig.yaml $HOME/myConsultConfig.yaml)")

	//rootCmd.MarkFlagRequired("config")

}

func initConfig() {
	cmn.InitConfig(cfgFile)
}

func main() {

	command.AddCommands(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		if sterr, ok := err.(cmn.StatusError); ok {
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
