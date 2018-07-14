package command

import (
	"github.com/shiqinfeng1/chunyuyisheng/api/v1"
	"github.com/shiqinfeng1/chunyuyisheng/service/common"
	"github.com/shiqinfeng1/chunyuyisheng/service/nsqs"
	"github.com/spf13/cobra"
)

var umengPusherCmd = &cobra.Command{
	Use:   "umengpusher",
	Short: "start a msg pusher service",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := nsqs.InitConfig(); err != nil {
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		httpServer := common.InitHttpService()
		v1.RegisterDiagAPI(httpServer)
		common.RunHttpService(httpServer)
		return
	},
}
