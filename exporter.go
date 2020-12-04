package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "maxscale"
)

var (
	descServerUp                = newDesc("server", "up", "Is the server up", serverLabelNames, prometheus.GaugeValue)
	descServerMaster            = newDesc("server", "master", "Is the server master", serverLabelNames, prometheus.GaugeValue)
	descServerConnections       = newDesc("server", "connections", "Current number of connections to the server", serverLabelNames, prometheus.GaugeValue)
	descServerTotalConnections  = newDesc("server", "total_connections", "Total connections", serverLabelNames, prometheus.CounterValue)
	descServerReusedConnections = newDesc("server", "reused_connections", "Reused connections", serverLabelNames, prometheus.CounterValue)
	descServerActiveOperations  = newDesc("server", "active_operations", "Curren number of active operations", serverLabelNames, prometheus.GaugeValue)
	descServiceCurrentSessions  = newDesc("service", "current_sessions", "Amount of sessions currently active", serviceLabelNames, prometheus.GaugeValue)
	descServiceSessionsTotal    = newDesc("service", "total_sessions", "Total amount of sessions", serviceLabelNames, prometheus.CounterValue)
	descQueryStatisticsRead     = newDesc("query_statistics", "read", "Total reads", queryStatisticsLabelNames, prometheus.CounterValue)
	descQueryStatisticsWrite    = newDesc("query_statistics", "write", "Total writes", queryStatisticsLabelNames, prometheus.CounterValue)
)

type Exporter struct {
	Address string // address of the maxscale instance
	ctx     context.Context
	up      prometheus.Gauge
}

var (
	serverLabelNames          = []string{"server"}
	serviceLabelNames         = []string{"service"}
	queryStatisticsLabelNames = []string{"service", "server"}
)

func NewExporter(ctx context.Context, address string) *Exporter {
	return &Exporter{
		Address: address,
		ctx:     ctx,
		up: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "up",
			Help:      "Was the last scrape of MaxScale successful?",
		}),
	}
}

// Describe describes all the metrics ever exported by the MaxScale exporter. It
// implements prometheus.Collector.
func (m *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- descServerUp.Desc
	ch <- descServerMaster.Desc
	ch <- descServerConnections.Desc
	ch <- descServerTotalConnections.Desc
	ch <- descServerReusedConnections.Desc
	ch <- descServerActiveOperations.Desc
	ch <- descServiceCurrentSessions.Desc
	ch <- descServiceSessionsTotal.Desc
	ch <- descQueryStatisticsRead.Desc
	ch <- descQueryStatisticsWrite.Desc

	ch <- m.up.Desc()
}

// Collect fetches the stats from configured MaxScale location and delivers them
// as Prometheus metrics. It implements prometheus.Collector.
func (m *Exporter) Collect(ch chan<- prometheus.Metric) {
	parseErrors := false

	if err := m.parseServers(ch); err != nil {
		parseErrors = true
		log.Print(err)
	}

	if err := m.parseServices(ch); err != nil {
		parseErrors = true
		log.Print(err)
	}

	m.up.Set(boolToFloat(!parseErrors))
	ch <- m.up
}

const contentTypeJSON = "application/json"

func (m *Exporter) fetchJSON(path string, v interface{}) error {
	url := "http://" + m.Address + "/v1" + path
	// build the request
	req, err := http.NewRequestWithContext(m.ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", contentTypeJSON)

	// execute the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error while getting %v: %w", url, err)
	}
	defer resp.Body.Close()

	// check status code
	if status := resp.StatusCode; status != http.StatusOK {
		return fmt.Errorf("unexpected status code %v for url %v", status, url)
	}

	// check content type
	if contentType := resp.Header.Get("Content-Type"); contentType != contentTypeJSON {
		return fmt.Errorf("unexpected content type %v", contentType)
	}

	// decode JSON body
	respObj := Response{}
	err = json.NewDecoder(resp.Body).Decode(&respObj)
	if err != nil {
		return err
	}

	return json.Unmarshal(respObj.Data, v)
}

func boolToFloat(value bool) float64 {
	if value {
		return 1
	}

	return 0
}

func (m *Exporter) parseServers(ch chan<- prometheus.Metric) error {
	var servers []ServerData
	err := m.fetchJSON("/servers", &servers)
	if err != nil {
		return err
	}

	for _, server := range servers {
		ch <- descServerConnections.new(
			float64(server.Attributes.Statistics.Connections),
			server.Attributes.Name,
		)

		ch <- descServerReusedConnections.new(
			float64(server.Attributes.Statistics.ReusedConnections),
			server.Attributes.Name,
		)
		ch <- descServerTotalConnections.new(
			float64(server.Attributes.Statistics.TotalConnections),
			server.Attributes.Name,
		)
		ch <- descServerActiveOperations.new(
			float64(server.Attributes.Statistics.ActiveOperations),
			server.Attributes.Name,
		)

		ch <- descServerMaster.new(
			boolToFloat(strings.HasPrefix(server.Attributes.State, "Master,")),
			server.Attributes.Name,
		)
		ch <- descServerUp.new(
			boolToFloat(strings.HasSuffix(server.Attributes.State, ", Running")),
			server.Attributes.Name,
		)
	}

	return nil
}

func (m *Exporter) parseServices(ch chan<- prometheus.Metric) error {
	var services []ServiceData
	err := m.fetchJSON("/services", &services)
	if err != nil {
		return err
	}

	for _, service := range services {
		ch <- descServiceCurrentSessions.new(
			float64(service.Attributes.Statistics.Connections),
			service.ID,
		)

		for _, statistics := range service.Attributes.RouterDiagnostics.ServerQueryStatistics {
			labelValues := []string{service.ID, statistics.ID}

			ch <- descQueryStatisticsRead.new(
				float64(statistics.Read),
				labelValues...,
			)
			ch <- descQueryStatisticsWrite.new(
				float64(statistics.Write),
				labelValues...,
			)
		}
	}

	return nil
}
