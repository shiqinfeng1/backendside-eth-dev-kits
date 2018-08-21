package nsqs

import (
	"fmt"
	"time"

	cmn "github.com/shiqinfeng1/backendside-eth-dev-kits/service/common"
	"github.com/shiqinfeng1/gorequest"
)

// ShootMessage Shoot message
func ShootMessage(address, topic string, payload interface{}) error {
	request := gorequest.New().Timeout(10 * time.Second)
	request.SetDebug(cmn.Config().GetBool("nsq.debug"))
	//request.SetLogger(cmn.Logger)
	_, _, errs := request.Post(fmt.Sprintf("http://%s/pub?topic=%s", address, topic)).
		Send(payload).
		End()
	if errs != nil {
		err := fmt.Errorf("nsqs.ShootMessage error: %q", errs)
		return err
	}
	return nil
}
