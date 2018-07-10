package main

import (
	"fmt"
	"os"

	"github.com/shiqinfeng1/chunyuyisheng/cmd/command"
	"github.com/spf13/cobra"
)

func newRootCommand() *cobra.Command {
	var cfgFile string
	rootCmd := &cobra.Command{
		Use:   "myconsult",
		Short: "myconsult is the consultation platform connected to ‘春雨医生’ and ‘国医链’",
		Long:  `myconsult is the consultation platform connected to ‘春雨医生’ and ‘国医链’`,
	}

	//Persistent类型的参数表示在子命令种也有效， 相当于全局参数。对应的，设置本地参数的方法是：Flags()
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/myConsultConfig.yaml)")
	//rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")

	//将输入参数和配置文件中的字段绑定，如果命令参数携带 --viper，则从配置文件中读入该对应的字段
	//cfgData.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	//fmt.Fprintln("useViper from configFile =", cfgData.GetBool("useViper"))

	common.InitConfig(cfgFile)

	return rootCmd
}

func main() {

	rootcmd := newRootCommand()
	command.AddCommands(rootcmd)

	if err := rootcmd.Execute(); err != nil {
		if sterr, ok := err.(cli.StatusError); ok {
			if sterr.Status != "" {
				fmt.Fprintln(stderr, sterr.Status)
			}
			// StatusError should only be used for errors, and all errors should
			// have a non-zero exit status, so never exit with 0
			if sterr.StatusCode == 0 {
				os.Exit(1)
			}
			os.Exit(sterr.StatusCode)
		}
		fmt.Fprintln(stderr, err)
		os.Exit(1)
	}
}
