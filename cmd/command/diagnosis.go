package command

import (
	"github.com/shiqinfeng1/chunyuyisheng/api/v1"
	"github.com/shiqinfeng1/chunyuyisheng/service/common"
	"github.com/spf13/cobra"
)

var diagCmd = &cobra.Command{
	Use:   "consult",
	Short: "start a consult backend service",

	Run: func(cmd *cobra.Command, args []string) {
		//db.InitMysql()
		httpServer := common.InitHttpService()
		v1.RegisterDiagAPI(httpServer)
		common.RunHttpService(httpServer)
		return
	},
}
