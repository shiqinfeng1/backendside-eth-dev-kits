package common

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

var (
	// Logger debug logger
	Logger echo.Logger
)

// LoggerInit init globel logger
func LoggerInit(e *echo.Echo, debuglevel string) {
	Logger = e.Logger
	if debuglevel == "disable" {
		Logger.SetLevel(log.OFF)
	} else if debuglevel == "debug" {
		Logger.SetLevel(log.DEBUG)
	} else if debuglevel == "info" {
		Logger.SetLevel(log.INFO)
	} else if debuglevel == "warn" {
		Logger.SetLevel(log.WARN)
	} else if debuglevel == "error" {
		Logger.SetLevel(log.ERROR)
	}
}
