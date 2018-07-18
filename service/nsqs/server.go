package nsqs

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"

	"github.com/labstack/gommon/color"
	nsq "github.com/nsqio/go-nsq"
	cmn "github.com/shiqinfeng1/chunyuyisheng/service/common"
)

// NsqStopable which has Stop() method for grancful stop, eg: nsq consumer/producer
type NsqStopable interface {
	Stop()
}

var stopables []NsqStopable
var access sync.Mutex
var started bool
var wg sync.WaitGroup

// GlobalConfig global config
var globalConfig = &SimpleConfig{}

// InitConfig initialize global emmiter
func InitConfig() error {

	globalConfig.NsqAddress = cmn.Config().GetString("nsq.nodeAddress")
	globalConfig.MaxInFlight = cmn.Config().GetInt("nsq.maxInFlight")
	globalConfig.Lookups = []string{cmn.Config().GetString("nsq.LookupsAddress")}

	return nil
}

// addNsqStopable Add a nsq consumer/producer
func addNsqStopable(ns NsqStopable) {
	access.Lock()
	defer access.Unlock()

	wg.Add(1)
	stopables = append(stopables, ns)
}

// Quit quit server
func quit() {
	access.Lock()
	defer access.Unlock()

	if !started {
		return
	}
	started = false
	cmn.Logger.Info("Stop nsqs ...")

	for _, ns := range stopables {
		go func(ns NsqStopable) {
			defer wg.Done()
			ns.Stop()
		}(ns)
	}
}

func handleSignals() {
	quitSignal := make(chan os.Signal)
	signal.Notify(quitSignal, syscall.SIGUSR1, syscall.SIGINT, syscall.SIGTERM)
	<-quitSignal

	quit()
}

func waitForExit() {
	wg.Wait()
}

// Start start nsqs
func Start() {
	cmn.Logger.Info("Start nsqs ...")
	access.Lock()
	defer access.Unlock()

	if started {
		return
	}
	started = true
}

// Run run server
func Run() {
	Start()
	go handleSignals()
	waitForExit()
}

// Stop graceful stop server
func Stop() {
	quit()
	waitForExit()
}

// HandlerFunc handler function
type HandlerFunc func(m *nsq.Message) error

func recoverableExec(handler HandlerFunc, message *nsq.Message) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch r := r.(type) {
			case error:
				err = r
			default:
				err = fmt.Errorf("%v", r)
			}
			stack := make([]byte, 4<<10)
			length := runtime.Stack(stack, false)
			cmn.Logger.Debugf("[%s] %s %s\n", color.Red("PANIC RECOVER"), err, stack[:length])

		}
	}()
	err = handler(message)
	return
}

//对handler的回调函数，封装了一层异常恢复流程，是一个闭包函数。
func handleMessage(handler HandlerFunc) nsq.HandlerFunc {

	return nsq.HandlerFunc(
		func(message *nsq.Message) (err error) {
			return recoverableExec(handler, message)
		})
}

// AddTopicLisenter is register a topic listener
func AddTopicLisenter(topic, channel string, handler HandlerFunc, concurrency int) (err error) {
	err = On(ListenerConfig{
		Lookup:             globalConfig.Lookups,
		Topic:              topic,
		Channel:            channel,
		HandlerConcurrency: concurrency,
	}, handleMessage(handler))
	return err
}

// AddTopicLisenterDefault is register a topic listener with concurrency process and default configuration.
func AddTopicLisenterDefault(topic string, handler HandlerFunc) (err error) {
	return AddTopicLisenter(topic, "default", handler, 10)
}

// PostTopic post topic
func PostTopic(topic string, payload interface{}) (err error) {
	err = ShootMessage(globalConfig.NsqAddress, topic, payload)
	cmn.Logger.Debug("nsqs post topic [", topic, "] playload: [", payload, "] error:", err)
	return
}
