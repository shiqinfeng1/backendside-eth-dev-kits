package command

import (
	"math/big"

	"github.com/shiqinfeng1/backendside-eth-dev-kits/api/v1"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/accounts"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/common"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/contracts"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/db"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/endpoints"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/eth"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/httpservice"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/nsqs"
	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "start a daemon backend service",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		db.InitMysql()
		if err := nsqs.InitConfig(); err != nil {
			return err
		}
		common.LoggerInit(common.Config().GetString("debugLevel"))
		nsqs.Start()
		if err := contracts.CompileContracts(); err != nil {
			return err
		}
		go endpoints.NewEndPointsManager().Run()
		go eth.NewPendingTransactionManager().Run()
		accounts.NewRootHDWallet()

		/***for test*****************/
		accounts.NewAccount("15422339579")

		auth, err := contracts.GetUserAuth("15422339579")
		if err != nil {
			common.Logger.Error("GetUserAuth: 15422339579 fail.")
			return nil
		}
		userAddress, _ := accounts.GetUserAddress("15422339579")
		nonce, _ := eth.ConnectEthNodeForWeb3("ethereum").EthGetNonce(userAddress)
		auth.Nonce = nonce.ToInt()
		common.Logger.Errorf("1.nonce=%v", auth.Nonce)

		if common.Config().GetString("ethereum.omcaddress") == "" {
			contracts.DeployOMCToken("ethereum", "15422339579", auth)
			auth.Nonce.Add(auth.Nonce, big.NewInt(1))
			common.Logger.Errorf("2.nonce=%v", auth.Nonce)
		}

		contracts.OMCTokenTransfer("ethereum", "15422339579", auth, "0x1dcef12e93b0abf2d36f723e8b59cc762775d513", 100000)
		auth.Nonce.Add(auth.Nonce, big.NewInt(1))
		common.Logger.Errorf("3.nonce=%v", auth.Nonce)

		if common.Config().GetString("poa.pointsaddress") == "" {
			contracts.DeployPointCoin("poa", "15422339579", auth)
			auth.Nonce.Add(auth.Nonce, big.NewInt(1))
			common.Logger.Errorf("4.nonce=%v", auth.Nonce)
		}

		contracts.PointsBuy("poa", "15422339579", auth, "0x1dcef12e93b0abf2d36f723e8b59cc762775d513", 100000)
		/****************************/

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		httpServer := httpservice.InitHTTPService()
		v1.RegisterDevKitsAPI(httpServer)
		httpservice.RunHTTPService(httpServer)
		return
	},
}
