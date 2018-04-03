package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

const (
	envPrefix   = "MAXSCALE_EXPORTER"
	metricsPath = "/metrics"
)

// strflag is like flag.String, with value overridden by an environment
// variable (when present). e.g. with name token, the env var used as default
// is DAVE_TOKEN, if present in env.
func strflag(name string, value string, usage string) *string {
	if v, ok := os.LookupEnv(envPrefix + strings.ToUpper(name)); ok {
		return flag.String(name, v, usage)
	}
	return flag.String(name, value, usage)
}

// Flags for CLI invocation
var (
	address *string
	port    *string
)

const namespace = "maxscale"

type MaxScale struct {
	Address         string
	up              prometheus.Gauge
	totalScrapes    prometheus.Counter
	serverMetrics   map[string]*prometheus.GaugeVec
	serviceMetrics  map[string]*prometheus.GaugeVec
	statusMetrics   map[string]*prometheus.GaugeVec
	variableMetrics map[string]*prometheus.GaugeVec
	eventsExecuted  prometheus.Metric
	eventsQueued    prometheus.Metric
	mutex           sync.RWMutex
}

type Server struct {
	Server      string
	Address     string
	Port        int
	Connections float64
	Status      string
}

type Service struct {
	Name          string  `json:"Service Name"`
	Router        string  `json:"Router Module"`
	Sessions      float64 `json:"No. Sessions"`
	TotalSessions float64 `json:"Total Sessions"`
}

type Status struct {
	Name  string  `json:"Variable_name"`
	Value float64 `json:"Value"`
}

type Variable struct {
	Name  string      `json:"Variable_name"`
	Value json.Number `json:"Value,num_integer"`
}

type Event struct {
	Duration string `json:"Duration"`
	Queued   uint64 `json:"No. Events Queued"`
	Executed uint64 `json:"No. Events Executed"`
}

var (
	serverLabelNames    = []string{"server", "address"}
	serviceLabelNames   = []string{"name", "router"}
	statusLabelNames    = []string{}
	variablesLabelNames = []string{}
)

func newMetric(metricName string, docString string, constLabels prometheus.Labels, labelNames []string) *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace:   namespace,
			Name:        metricName,
			Help:        docString,
			ConstLabels: constLabels,
		},
		labelNames,
	)
}

type metrics map[string]*prometheus.GaugeVec

var (
	serverMetrics = metrics{
		"server_connections": newMetric("server_connections", "Amount of connections to the server", nil, serverLabelNames),
		"server_up":          newMetric("server_up", "Is the server up", nil, serverLabelNames),
	}
	serviceMetrics = metrics{
		"service_current_sessions": newMetric("service_current_sessions", "Amount of sessions currently active", nil, serviceLabelNames),
		"service_total_sessions":   newMetric("service_total_sessions", "Total amount of sessions", nil, serviceLabelNames),
	}

	statusMetrics = metrics{
		"status_uptime":                    newMetric("status_uptime", "How long has the server been running", nil, statusLabelNames),
		"status_uptime_since_flush_status": newMetric("status_uptime_since_flush_status", "How long the server has been up since flush status", nil, statusLabelNames),
		"status_threads_created":           newMetric("status_threads_created", "How many threads have been created", nil, statusLabelNames),
		"status_threads_running":           newMetric("status_threads_running", "How many threads are running", nil, statusLabelNames),
		"status_threadpool_threads":        newMetric("status_threadpool_threads", "How many threadpool threads there are", nil, statusLabelNames),
		"status_threads_connected":         newMetric("status_threads_connected", "How many threads are connected", nil, statusLabelNames),
		"status_connections":               newMetric("status_connections", "How many connections there are", nil, statusLabelNames),
		"status_client_connections":        newMetric("status_client_connections", "How many client connections there are", nil, statusLabelNames),
		"status_backend_connections":       newMetric("status_backend_connections", "How many backend connections there are", nil, statusLabelNames),
		"status_listeners":                 newMetric("status_listeners", "How many listeners there are", nil, statusLabelNames),
		"status_zombie_connections":        newMetric("status_zombie_connections", "How many zombie connetions there are", nil, statusLabelNames),
		"status_internal_descriptors":      newMetric("status_internal_descriptors", "How many internal descriptors there are", nil, statusLabelNames),
		"status_read_events":               newMetric("status_read_events", "How many read events happened", nil, statusLabelNames),
		"status_write_events":              newMetric("status_write_events", "How many write events happened", nil, statusLabelNames),
		"status_hangup_events":             newMetric("status_hangup_events", "How many hangup events happened", nil, statusLabelNames),
		"status_error_events":              newMetric("status_error_events", "How many error events happened", nil, statusLabelNames),
		"status_accept_events":             newMetric("status_accept_events", "How many accept events happened", nil, statusLabelNames),
		"status_event_queue_length":        newMetric("status_event_queue_length", "How long the event queue is", nil, statusLabelNames),
		"status_max_event_queue_length":    newMetric("status_max_event_queue_length", "The max length of the event queue", nil, statusLabelNames),
		"status_max_event_queue_time":      newMetric("status_max_event_queue_time", "The max event queue time", nil, statusLabelNames),
		"status_max_event_execution_time":  newMetric("status_max_event_execution_time", "The max event execution time", nil, statusLabelNames),
	}

	variableMetrics = metrics{
		"variables_maxscale_threads":   newMetric("variables_thread", "MAXSCALE_THREADS", nil, variablesLabelNames),
		"variables_maxscale_nbpolls":   newMetric("variables_nbpolls", "MAXSCALE_NBPOLLS", nil, variablesLabelNames),
		"variables_maxscale_pollsleep": newMetric("variables_pollsleep", "MAXSCALE_POLLSLEEP", nil, variablesLabelNames),
		"variables_maxscale_sessions":  newMetric("variables_sessions", "MAXSCALE_SESSIONS", nil, variablesLabelNames),
	}
)

func NewExporter(address string) (*MaxScale, error) {
	return &MaxScale{
		Address: address,
		up: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "up",
			Help:      "Was the last scrape of MaxScale successful?",
		}),
		totalScrapes: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "exporter_total_scrapes",
			Help:      "Current total MaxScale scrapes",
		}),
		serverMetrics:   serverMetrics,
		serviceMetrics:  serviceMetrics,
		statusMetrics:   statusMetrics,
		variableMetrics: variableMetrics,
	}, nil
}

// Describe describes all the metrics ever exported by the MaxScale exporter. It
// implements prometheus.Collector.
func (m *MaxScale) Describe(ch chan<- *prometheus.Desc) {
	for _, metric := range m.serverMetrics {
		metric.Describe(ch)
	}

	for _, metric := range m.serviceMetrics {
		metric.Describe(ch)
	}

	for _, metric := range m.statusMetrics {
		metric.Describe(ch)
	}

	for _, metric := range m.variableMetrics {
		metric.Describe(ch)
	}

	ch <- m.up.Desc()
	ch <- m.totalScrapes.Desc()
}

// Collect fetches the stats from configured MaxScale location and delivers them
// as Prometheus metrics. It implements prometheus.Collector.
func (m *MaxScale) Collect(ch chan<- prometheus.Metric) {
	m.mutex.Lock() // To protect metrics from concurrent collects.
	defer m.mutex.Unlock()

	m.resetMetrics()
	m.scrape()

	ch <- m.up
	ch <- m.totalScrapes
	ch <- m.eventsExecuted
	ch <- m.eventsQueued

	m.collectMetrics(ch)
}

func (m *MaxScale) resetMetrics() {
	for _, metric := range m.serverMetrics {
		metric.Reset()
	}

	for _, metric := range m.serviceMetrics {
		metric.Reset()
	}

	for _, metric := range m.statusMetrics {
		metric.Reset()
	}

	for _, metric := range m.variableMetrics {
		metric.Reset()
	}
}

func (m *MaxScale) collectMetrics(metrics chan<- prometheus.Metric) {
	for _, metric := range m.serverMetrics {
		metric.Collect(metrics)
	}

	for _, metric := range m.serviceMetrics {
		metric.Collect(metrics)
	}

	for _, metric := range m.statusMetrics {
		metric.Collect(metrics)
	}

	for _, metric := range m.variableMetrics {
		metric.Collect(metrics)
	}
}

func (m *MaxScale) scrape() {
	m.totalScrapes.Inc()

	if err := m.fetch(); err != nil {
		log.Fatal(err)
		m.up.Set(0)
		return
	}

	m.up.Set(1)
}

func (m *MaxScale) url(path string) string {
	return "http://" + m.Address + path
}

func (m *MaxScale) getStatistics(path string) ([]byte, error) {
	url := m.url(path)
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("Error while getting %v: %v\n", path, err)
	}

	jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("Error while reading response from %v: %v\n", path, err)
	}

	return jsonDataFromHttp, nil
}

func (m *MaxScale) fetch() error {
	if err := m.parseServers(); err != nil {
		return err
	}

	if err := m.parseServices(); err != nil {
		return err
	}

	if err := m.parseStatus(); err != nil {
		return err
	}

	if err := m.parseVariables(); err != nil {
		return err
	}

	if err := m.parseEvents(); err != nil {
		return err
	}

	return nil
}

func (m *MaxScale) parseServers() error {
	response, err := m.getStatistics("/servers")

	if err != nil {
		return err
	}

	var servers []Server

	err = json.Unmarshal([]byte(response), &servers)

	if err != nil {
		return fmt.Errorf("Error while unmarshaling json: %v\n", err)
	}

	for _, server := range servers {
		m.serverMetrics["server_connections"].WithLabelValues(server.Server, server.Address).Set(server.Connections)
		m.serverMetrics["server_up"].WithLabelValues(server.Server, server.Address).Set(serverUp(server.Status))
	}

	return nil
}

func serverUp(status string) float64 {
	switch status {
	case "Down":
		return 0
	case "Running":
		return 1
	default:
		// Unsure about other status messages, return false just in case
		return 0
	}
}

func (m *MaxScale) parseServices() error {
	response, err := m.getStatistics("/services")

	if err != nil {
		return err
	}

	var services []Service

	err = json.Unmarshal([]byte(response), &services)
	if err != nil {
		return fmt.Errorf("Error while unmarshaling json: %v\n", err)
	}

	for _, service := range services {
		m.serviceMetrics["service_current_sessions"].WithLabelValues(service.Name, service.Router).Set(service.Sessions)
		m.serviceMetrics["service_total_sessions"].WithLabelValues(service.Name, service.Router).Set(service.TotalSessions)
	}

	return nil
}

func (m *MaxScale) parseStatus() error {
	response, err := m.getStatistics("/status")

	if err != nil {
		return err
	}

	var status []Status

	err = json.Unmarshal([]byte(response), &status)
	if err != nil {
		return fmt.Errorf("Error while unmarshaling json: %v\n", err)
	}

	for _, element := range status {
		var metric = "status_" + strings.ToLower(element.Name)
		m.statusMetrics[metric].WithLabelValues().Set(element.Value)
	}

	return nil
}
func (m *MaxScale) parseVariables() error {
	response, err := m.getStatistics("/variables")

	if err != nil {
		return err
	}

	var variables []Variable

	err = json.Unmarshal([]byte(response), &variables)
	if err != nil {
		log.Print(err)
		return err
	}

	for _, element := range variables {
		metric := "variables_" + strings.ToLower(element.Name)
		if _, ok := m.variableMetrics[metric]; ok {
			value, err := element.Value.Float64()
			if err != nil {
				return err
			}
			m.variableMetrics[metric].WithLabelValues().Set(value)
		}
	}

	return nil
}
func (m *MaxScale) parseEvents() error {
	response, err := m.getStatistics("/event/times")

	if err != nil {
		return err
	}

	var events []Event

	err = json.Unmarshal([]byte(response), &events)
	if err != nil {
		return err
	}

	eventExecutedBuckets := map[float64]uint64{
		0.1: 0,
		0.2: 0,
		0.3: 0,
		0.4: 0,
		0.5: 0,
		0.6: 0,
		0.7: 0,
		0.8: 0,
		0.9: 0,
		1.0: 0,
		1.1: 0,
		1.2: 0,
		1.3: 0,
		1.4: 0,
		1.5: 0,
		1.6: 0,
		1.7: 0,
		1.8: 0,
		1.9: 0,
		2.0: 0,
		2.1: 0,
		2.2: 0,
		2.3: 0,
		2.4: 0,
		2.5: 0,
		2.6: 0,
		2.7: 0,
		2.8: 0,
		2.9: 0,
	}
	executedSum := float64(0)
	executedCount := uint64(0)
	executedTime := 0.1
	for _, element := range events {
		executedCount += element.Executed
		executedSum = executedSum + (float64(element.Executed) * executedTime)
		executedTime += 0.1
		switch element.Duration {
		case "< 100ms":
			eventExecutedBuckets[0.1] = element.Executed
		case "> 3000ms":
			break // Do nothing as these will get accumulated in the +Inf bucket
		default:
			durationf := strings.Split(element.Duration, " ")
			ad := strings.Trim(durationf[len(durationf)-1], "ms")
			milliseconds, _ := strconv.ParseFloat(ad, 64)
			seconds := milliseconds / 1000
			eventExecutedBuckets[seconds] = element.Executed
		}
	}

	desc := prometheus.NewDesc(
		"maxscale_events_executed_seconds",
		"Amount of events executed",
		[]string{},
		prometheus.Labels{},
	)

	// Create a constant histogram from values we got from a 3rd party telemetry system.
	eventsExecuted := prometheus.MustNewConstHistogram(
		desc,
		executedCount, executedSum,
		eventExecutedBuckets,
	)

	m.eventsExecuted = eventsExecuted

	eventQueuedBuckets := map[float64]uint64{
		0.1: 0,
		0.2: 0,
		0.3: 0,
		0.4: 0,
		0.5: 0,
		0.6: 0,
		0.7: 0,
		0.8: 0,
		0.9: 0,
		1.0: 0,
		1.1: 0,
		1.2: 0,
		1.3: 0,
		1.4: 0,
		1.5: 0,
		1.6: 0,
		1.7: 0,
		1.8: 0,
		1.9: 0,
		2.0: 0,
		2.1: 0,
		2.2: 0,
		2.3: 0,
		2.4: 0,
		2.5: 0,
		2.6: 0,
		2.7: 0,
		2.8: 0,
		2.9: 0,
	}

	queuedSum := float64(0)
	queuedCount := uint64(0)
	queuedTime := 0.1
	for _, element := range events {
		queuedCount += element.Queued
		queuedSum = queuedSum + (float64(element.Queued) * queuedTime)
		queuedTime += 0.1
		switch element.Duration {
		case "< 100ms":
			eventQueuedBuckets[0.1] = element.Queued
		case "> 3000ms":
			break // Do nothing as this gets accumulated in the +Inf bucket
		default:
			durationf := strings.Split(element.Duration, " ")
			ad := strings.Trim(durationf[len(durationf)-1], "ms")
			milliseconds, _ := strconv.ParseFloat(ad, 64)
			seconds := milliseconds / 1000
			eventQueuedBuckets[seconds] = element.Queued
		}
	}

	queuedDesc := prometheus.NewDesc(
		"maxscale_events_queued_seconds",
		"Amount of events queued",
		[]string{},
		prometheus.Labels{},
	)

	// Create a constant histogram from values we got from a 3rd party telemetry system.
	eventsQueued := prometheus.MustNewConstHistogram(
		queuedDesc,
		queuedCount, queuedSum,
		eventQueuedBuckets,
	)

	m.eventsQueued = eventsQueued

	return nil
}

func main() {
	log.SetFlags(0)

	address = strflag("address", "127.0.0.1:8003", "address to get maxscale statistics from")
	port = strflag("port", ":9195", "the port that the maxscale exporter listens on")
	flag.Parse()

	log.Print("Starting MaxScale exporter")
	exporter, err := NewExporter(*address)
	if err != nil {
		log.Fatalf("Failed to start maxscale exporter: %v\n", err)
	}

	prometheus.MustRegister(exporter)
	http.Handle(metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>MaxScale Exporter</title></head>
			<body>
			<h1>MaxScale Exporter</h1>
			<p><a href="` + metricsPath + `">Metrics</a></p>
			</body>
			</html>`))
	})
	log.Fatal(http.ListenAndServe(*port, nil))
}
