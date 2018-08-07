package v1

import (
	"github.com/labstack/echo"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/common"
)

//Foo :
func Foo(c echo.Context) error {
	p := common.PadPayload{}
	if err := c.Bind(&p); err != nil {
		return err
	}
	if err := c.Echo().Validator.Validate(p); err != nil {
		return err
	}
	return common.JSONReturns(c, nil)
}
