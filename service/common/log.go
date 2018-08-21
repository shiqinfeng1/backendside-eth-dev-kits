package common

import (
	"encoding/json"
	"os"
	"runtime"

	"github.com/ethereum/go-ethereum/common"
	"github.com/natefinch/lumberjack"
	logging "github.com/shiqinfeng1/go-logging"
)

var format = logging.MustStringFormatter(
	`%{color}%{time} %{callpath} > %{level:.4s} %{id:03x} %{message}`, //%{color:reset}
)

//Logger 日志记录器
var Logger = logging.MustGetLogger("ethsmart")

//LogOutpot log日志记录
var LogOutpot = &lumberjack.Logger{
	Filename:   "./ethsmart.log",
	MaxSize:    10, // megabytes
	MaxBackups: 3,
	MaxAge:     30,   //days
	Compress:   true, // disabled by default
	LocalTime:  true,
}

// LoggerInit init globel logger
func LoggerInit(debuglevel string) {
	//定义log后端
	backend1 := logging.NewLogBackend(LogOutpot, "[poa后端服务]", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	//设置log后端的打印级别
	backend1Level := logging.AddModuleLevel(backend1)
	backend1Level.SetLevel(logging.DEBUG, "")
	backend1Formatter := logging.NewBackendFormatter(backend1, format)

	backend2Level := logging.AddModuleLevel(backend2)
	if debuglevel == "debug" {
		backend2Level.SetLevel(logging.DEBUG, "")
	} else if debuglevel == "info" {
		backend2Level.SetLevel(logging.INFO, "")
	} else if debuglevel == "warn" {
		backend2Level.SetLevel(logging.WARNING, "")
	} else if debuglevel == "error" {
		backend2Level.SetLevel(logging.ERROR, "")
	}
	backend2Formatter := logging.NewBackendFormatter(backend2Level, format)

	//设置log后端
	logging.SetBackend(backend1Formatter, backend2Formatter)

	Logger.Debug("this is debug msg demo")
	Logger.Info("this is info msg demo")
	Logger.Warning("this is warn msg demo")
	Logger.Error("this is error msg demo")
}

//PrintDeployContactInfo 打印部署合约的信息
func PrintDeployContactInfo(addr common.Address, txn interface{}, err error) {
	var errstr = "-"
	if err != nil {
		errstr = err.Error()
	}
	if Config().GetBool("common.debug") {
		pc, _, _, _ := runtime.Caller(2)
		p2, _ := json.MarshalIndent(&txn, "", "  ")
		output :=
			"\n" +
				runtime.FuncForPC(pc).Name() +
				"\nContract Address:" + "\n" +
				addr.Hex() + "\n" +
				"TransactionInfo:" + "\n" +
				string(p2) + "\n" +
				"Deploy error:" + "\n" +
				errstr + "\n" +
				"--------\n"

		Logger.Info(output)
	}
}
