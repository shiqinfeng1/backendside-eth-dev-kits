package command

import (
	"errors"
	"fmt"

	"github.com/shiqinfeng1/chunyuyisheng/api/v1"
	"github.com/shiqinfeng1/chunyuyisheng/service/common"
	"github.com/shiqinfeng1/chunyuyisheng/service/db"
	"github.com/spf13/cobra"
)

var diagCmd = &cobra.Command{
	Short: "diagnosis",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}

		return fmt.Errorf("invalid color specified: %s", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		db.InitMysql()
		httpServer = common.InitHttpService()
		v1.RegisterDiagAPI(httpServer)
		return common.RunHttpService(httpServer)
	},
}
