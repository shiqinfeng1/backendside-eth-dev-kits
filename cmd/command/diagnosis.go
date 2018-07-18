package command

import (
	"github.com/labstack/gommon/log"

	"github.com/shiqinfeng1/chunyuyisheng/service/nsqs"

	"github.com/shiqinfeng1/chunyuyisheng/api/v1"
	"github.com/shiqinfeng1/chunyuyisheng/service/common"
	"github.com/spf13/cobra"
)

var diagCmd = &cobra.Command{
	Use:   "consult",
	Short: "start a consult backend service",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		common.Logger = log.New("consult")
		//db.InitMysql()
		if err := nsqs.InitConfig(); err != nil {
			return err
		}
		nsqs.Start()
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		httpServer := common.InitHTTPService()
		v1.RegisterDiagAPI(httpServer)
		common.RunHTTPService(httpServer)
		return
	},
}
