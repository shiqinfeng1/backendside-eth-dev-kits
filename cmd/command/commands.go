package command

import (
	"github.com/spf13/cobra"
)

//AddCommands :注册命令行命令
func AddCommands(root *cobra.Command) {
	root.AddCommand(daemonCmd)
}
