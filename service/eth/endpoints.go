package eth

import (
	"net/url"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/rpc"
	"gitlab.chainresearch.org/wallet/stone/common"
)

type endpoint struct {
	weight         int
	url            string
	isOk           bool
	interval       int
	intervalAmount int
}

func (e *endpoint) rpc(result interface{}, method string, args ...interface{}) error {
	client, err := rpc.Dial(e.url)
	if err != nil {
		common.Logger.Error("dial error in rpc: ", e.url)
		return err
	}
	err = client.Call(result, method, args...)
	if err != nil {
		common.Logger.Error(err)
		return err
	}
	return nil
}

func (e *endpoint) heartbeat() bool {
	var res string
	err := e.rpc(&res, "net_version")
	if err != nil {
		common.Logger.Info("heartbeat error: ", e.url)
		return false
	}
	return true
}

// EndpointsManager endpoints of ethereum
type EndpointsManager struct {
	endpoints       []*endpoint
	rAliveEndpoints []*endpoint
	rwMutex         sync.RWMutex
	exit            chan bool
	closed          chan bool
}

var endPoints *EndpointsManager

// NewEndPointsManager create a endPoint manager
func NewEndPointsManager() *EndpointsManager {
	endPoints = &EndpointsManager{
		endpoints:       []*endpoint{},
		rAliveEndpoints: []*endpoint{},
		exit:            make(chan bool),
		closed:          make(chan bool),
	}
	return endPoints
}

//GetEndPointsManager 节点管理
func GetEndPointsManager() *EndpointsManager {
	return endPoints
}

//AddEndPoint 增加监听节点
func (e *EndpointsManager) AddEndPoint(endpointURL string, weight int, interval int) {
	e.rwMutex.Lock()
	defer e.rwMutex.Unlock()
	endpoint := &endpoint{
		url:      endpointURL,
		weight:   weight,
		interval: interval,
	}
	e.endpoints = append(e.endpoints, endpoint)
	e.rAliveEndpoints = append(e.rAliveEndpoints, endpoint)
}

// Run endpoints run, monitor alive Endpoint
func (e *EndpointsManager) Run() {
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			e.watchAliveEndpoint()
		case <-e.exit:
			close(e.closed)
			common.Logger.Info("service done!!!")
			return
		}
	}
}

// Stop Stop
func (e *EndpointsManager) Stop() {
	close(e.exit)
	// wait for stop
	<-e.closed
}

func (e *EndpointsManager) watchAliveEndpoint() error {
	for _, item := range e.endpoints {
		common.Logger.Infof("endpoint %s, interval: %d, intervalAmount: %d", item.url, item.interval, item.intervalAmount)
		if item.intervalAmount == 0 {
			item.isOk = item.heartbeat()
			common.Logger.Infof("endpoint watch: %s, status: %t", item.url, item.isOk)
		}
		item.intervalAmount++
		if item.intervalAmount >= item.interval {
			item.intervalAmount = 0
		}
	}
	e.updateAliveEndpoint()

	common.Logger.Info("endpoint watch size: ", len(e.rAliveEndpoints))
	return nil
}

func (e *EndpointsManager) updateAliveEndpoint() {
	e.rwMutex.Lock()
	defer e.rwMutex.Unlock()
	res := []*endpoint{}
	for _, item := range e.endpoints {
		if item.isOk {
			res = append(res, item)
		}
	}
	e.rAliveEndpoints = res
}

// RPC rpc
func (e *EndpointsManager) RPC(result interface{}, method string, args ...interface{}) (err error) {
	e.rwMutex.RLock()
	defer e.rwMutex.RUnlock()
	for _, item := range e.rAliveEndpoints {
		err = item.rpc(result, method, args...)
		if _, ok := err.(*url.Error); ok {
			continue
		} else {
			break
		}
	}
	return
}
