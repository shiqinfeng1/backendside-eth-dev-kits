package command

import (
	"github.com/labstack/gommon/log"

	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/db"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/eth"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/httpservice"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/nsqs"

	"github.com/shiqinfeng1/backendside-eth-dev-kits/api/v1"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/common"
	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "start a daemon backend service",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		common.Logger = log.New("daemon")
		db.InitMysql()
		if err := nsqs.InitConfig(); err != nil {
			return err
		}
		nsqs.Start()
		if err := eth.AttachEthNode(); err != nil {
			return err
		}
		if err := eth.CompileSolidity(); err != nil {
			return err
		}
		endpointsManager := eth.NewEndPointsManager()
		endpointsManager.AddEndPoint(common.Config().GetString("ethereum.endpoints"), 1, 0)
		go endpointsManager.Run()

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		httpServer := httpservice.InitHTTPService()
		v1.RegisterDevKitsAPI(httpServer)
		httpservice.RunHTTPService(httpServer)
		return
	},
}
