package common

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	//"github.com/go-playground/validator"
	"github.com/labstack/echo"
	emw "github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"
)

// EchoHTTPErrorHandler is a HTTP error handler. It sends a JSON response
// with status code.
func EchoHTTPErrorHandler(e *echo.Echo) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		var (
			code = http.StatusOK
			msg  interface{}
			rmsg string
		)

		errcode := ErrorCode9999
		if he, ok := err.(*echo.HTTPError); ok {
			msg = he.Message
		} else if be, ok := err.(*BizError); ok {
			errcode = be.Code
			msg = be.TMsg(c)
		} else if e.Debug {
			msg = err.Error()
		} else {
			msg = http.StatusText(http.StatusInternalServerError)
		}
		// 处理错误信息
		if v, ok := msg.(string); ok {
			rmsg = v
		} else {
			rmsg = fmt.Sprintf("%s", msg)
		}

		if !c.Response().Committed {
			if c.Request().Method == echo.HEAD { // Issue #608
				if err := c.NoContent(code); err != nil {
					goto ERROR
				}
			} else {
				// 统一封装返回值
				if err := c.JSON(code, ErrorReturns(c, errcode, rmsg)); err != nil {
					goto ERROR
				}
			}
		}
	ERROR:
		e.Logger.Error(err)
	}
}

// GetAcceptLanguage Get Accept-Language from request header
func GetAcceptLanguage(c echo.Context) string {
	return c.Request().Header.Get("Accept-Language")
}

func echoInit(e *echo.Echo) {
	if Config().GetString("debugLevel") == "disable" {
		e.Debug = false
	} else {
		e.Debug = true
	}
	LoggerInit(e, Config().GetString("debugLevel"))
	e.HTTPErrorHandler = EchoHTTPErrorHandler(e)
	e.Validator = &SimpleValidator{Validator: validator.New()}
}

//InitHTTPService 启动echo服务
func InitHTTPService() (e *echo.Echo) {
	e = echo.New()
	echoInit(e)
	e.Use(emw.CORSWithConfig(emw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	return e
}

//RunHTTPService 运行echo服务
func RunHTTPService(e *echo.Echo) {

	srvAddr := ":" + Config().GetString("httpSrvPort")

	Logger.Printf("Listening and serving HTTP on %s\n", srvAddr)
	// Start server
	go func() {
		if err := e.Start(srvAddr); err != nil {
			log.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGUSR1, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	Logger.Printf("Diagnosis Server exit\n")
}
