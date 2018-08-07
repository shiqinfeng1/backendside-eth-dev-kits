package eth

import (
	"net/url"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/shiqinfeng1/backendside-eth-dev-kits/service/common"
)

type endpoint struct {
	nodeType       string
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
		// fmt.Println("---------------------------------")
		// fmt.Printf("RPC to NODE:%s fail. method = %s\n\n", e.url, method)
		// fmt.Print(err)
		// fmt.Println("---------------------------------")
		return err
	}
	return nil
}

func (e *endpoint) heartbeat() bool {
	var res string
	err := e.rpc(&res, "net_version")
	if err != nil {
		//common.Logger.Error(e.url, "  heartbeat error: connect fail.")
		return false
	}
	return true
}

// EndpointsManager endpoints of ethereum
type EndpointsManager struct {
	endpoints      []*endpoint
	AliveEndpoints []*endpoint
	rwMutex        sync.RWMutex
	exit           chan bool
	closed         chan bool
}

var endPoints *EndpointsManager

// NewEndPointsManager create a endPoint manager
func NewEndPointsManager() *EndpointsManager {
	endPoints = &EndpointsManager{
		endpoints:      []*endpoint{},
		AliveEndpoints: []*endpoint{},
		exit:           make(chan bool),
		closed:         make(chan bool),
	}
	return endPoints
}

//GetEndPointsManager 节点管理
func GetEndPointsManager() *EndpointsManager {
	return endPoints
}

func (e *EndpointsManager) updateEndPoint() {
	e.rwMutex.Lock()
	defer e.rwMutex.Unlock()
	var endpoints []*endpoint
	endpointURLs := common.Config().GetStringSlice("ethereum.endpoints")

	for _, endpointURL := range endpointURLs {
		endpoint := &endpoint{
			nodeType: "ethereum",
			url:      endpointURL,
			weight:   1,
			interval: 0,
		}
		endpoints = append(endpoints, endpoint)
	}
	endpointURLs = common.Config().GetStringSlice("poa.endpoints")
	for _, endpointURL := range endpointURLs {
		endpoint := &endpoint{
			nodeType: "poa",
			url:      endpointURL,
			weight:   1,
			interval: 0,
		}
		endpoints = append(endpoints, endpoint)
	}

	e.endpoints = endpoints
}

// Run endpoints run, monitor alive Endpoint
func (e *EndpointsManager) Run() {
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	e.updateEndPoint()
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
		if item.intervalAmount == 0 {
			item.isOk = item.heartbeat()
		}
		item.intervalAmount++
		if item.intervalAmount >= item.interval {
			item.intervalAmount = 0
		}
	}
	e.updateAliveEndpoint()
	e.updateEndPoint()
	return nil
}

func (e *EndpointsManager) updateAliveEndpoint() {
	e.rwMutex.Lock()
	defer e.rwMutex.Unlock()
	res := []*endpoint{}
	for _, item := range e.endpoints {
		if item.isOk {
			res = append(res, item)
		} else {
			common.Logger.Error(item.url, ": Not In Service!!!")
		}
	}
	e.AliveEndpoints = res
	if len(e.AliveEndpoints) == 0 {
		common.Logger.Error("No Node In Service!!!")
	}
}

// RPC rpc
func (e *EndpointsManager) RPC(result interface{}, method string, args ...interface{}) (err error) {
	e.rwMutex.RLock()
	defer e.rwMutex.RUnlock()
	for _, item := range e.AliveEndpoints {
		err = item.rpc(result, method, args...)
		if _, ok := err.(*url.Error); ok {
			continue
		} else {
			break
		}
	}
	return
}
