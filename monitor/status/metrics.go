package status

import (
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/prometheus"
	"github.com/irisnet/irishub-sync/logger"
	"github.com/irisnet/irishub-sync/store/document"
	"github.com/irisnet/irishub-sync/util/helper"
	prom "github.com/prometheus/client_golang/prometheus"
	"time"
)

const (
	NodeStatusNotReachable = 1
	NodeStatusCatchingUp   = 2
	NodeStatusSyncing      = 3
)

type Metrics struct {
	NodeHeight metrics.Gauge
	DbHeight   metrics.Gauge
	NodeStatus metrics.Gauge
}

func PrometheusMetrics() *Metrics {
	return &Metrics{
		NodeHeight: prometheus.NewGaugeFrom(prom.GaugeOpts{
			Namespace: "sync",
			Subsystem: "status",
			Name:      "node_height",
			Help:      "full node latest block height",
		}, []string{}),
		DbHeight: prometheus.NewGaugeFrom(prom.GaugeOpts{
			Namespace: "sync",
			Subsystem: "status",
			Name:      "db_height",
			Help:      "sync system database max block height",
		}, []string{}),
		NodeStatus: prometheus.NewGaugeFrom(prom.GaugeOpts{
			Namespace: "sync",
			Subsystem: "status",
			Name:      "node_status",
			Help:      "full node status(1:NotReachable,2:CatchingUp,3:Syncing)",
		}, []string{}),
	}
}

func (cs *Metrics) Report() {
	client, err := helper.GetClientWithTimeout(10 * time.Second)
	if err != nil {
		logger.Error("rpc node connection exception", logger.String("error", err.Error()))
		cs.NodeStatus.Set(float64(NodeStatusNotReachable))
		return
	}
	defer func() {
		client.Release()
	}()
	status, err := client.Status()
	if err != nil {
		logger.Error("rpc node connection exception", logger.String("error", err.Error()))
		cs.NodeStatus.Set(float64(NodeStatusNotReachable))
		return
	}
	// node height
	cs.NodeHeight.Set(float64(status.SyncInfo.LatestBlockHeight))
	// db height
	height, err := document.Block{}.GetMaxBlockHeight()
	if err != nil {
		logger.Error("query block exception", logger.String("error", err.Error()))
	}
	cs.DbHeight.Set(float64(height))
	if status.SyncInfo.CatchingUp {
		cs.NodeStatus.Set(float64(NodeStatusCatchingUp))
	} else {
		cs.NodeStatus.Set(float64(NodeStatusSyncing))
	}
}
