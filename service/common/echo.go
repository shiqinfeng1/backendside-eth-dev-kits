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

type (
	// NoCacheConfig defines the config for nocache middleware.
	NoCacheConfig struct {
		// Skipper defines a function to skip middleware.
		Skipper emw.Skipper
	}
)

var (
	// Unix epoch time
	epoch = time.Unix(0, 0).Format(time.RFC1123)

	// Taken from https://github.com/mytrile/nocache
	noCacheHeaders = map[string]string{
		"Expires":         epoch,
		"Cache-Control":   "no-cache, private, max-age=0",
		"Pragma":          "no-cache",
		"X-Accel-Expires": "0",
	}
	etagHeaders = []string{
		"ETag",
		"If-Modified-Since",
		"If-Match",
		"If-None-Match",
		"If-Range",
		"If-Unmodified-Since",
	}
	// DefaultNoCacheConfig is the default nocache middleware config.
	DefaultNoCacheConfig = NoCacheConfig{
		Skipper: emw.DefaultSkipper,
	}
)

// NoCache is a simple piece of middleware that sets a number of HTTP headers to prevent
// a router (or subrouter) from being cached by an upstream proxy and/or client.
//
// As per http://wiki.nginx.org/HttpProxyModule - NoCache sets:
//      Expires: Thu, 01 Jan 1970 00:00:00 UTC
//      Cache-Control: no-cache, private, max-age=0
//      X-Accel-Expires: 0
//      Pragma: no-cache (for HTTP/1.0 proxies/clients)
func NoCache() echo.MiddlewareFunc {
	return NoCacheWithConfig(DefaultNoCacheConfig)
}

// NoCacheWithConfig returns a nocache middleware with config.
func NoCacheWithConfig(config NoCacheConfig) echo.MiddlewareFunc {
	// Defaults
	if config.Skipper == nil {
		config.Skipper = DefaultNoCacheConfig.Skipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if config.Skipper(c) {
				return next(c)
			}
			req := c.Request()
			// Delete any ETag headers that may have been set
			for _, v := range etagHeaders {
				if req.Header.Get(v) != "" {
					req.Header.Del(v)
				}
			}

			// Set our NoCache headers
			res := c.Response()
			for k, v := range noCacheHeaders {
				res.Header().Set(k, v)
			}

			return next(c)
		}
	}
}

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
	e.Pre(emw.RemoveTrailingSlash())
	e.Pre(NoCache())
	e.Use(emw.Recover())
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
