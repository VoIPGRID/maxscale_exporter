package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	envPrefix   = "MAXSCALE_EXPORTER_"
	metricsPath = "/metrics"
)

// Flags for CLI invocation.
var (
	address = strflag("address", "127.0.0.1:8989", "address to get maxscale statistics from")
	port    = strflag("port", "9195", "the port that the maxscale exporter listens on")
)

func main() {
	log.SetFlags(log.Lshortfile)
	flag.Parse()

	log.Print("Starting MaxScale exporter")
	log.Printf("Scraping MaxScale JSON API at: %v", *address)

	http.HandleFunc(metricsPath, func(w http.ResponseWriter, r *http.Request) {
		reg := prometheus.NewRegistry()

		if r.FormValue("runtime") != "false" {
			reg.MustRegister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))
			reg.MustRegister(prometheus.NewGoCollector())
		}

		reg.MustRegister(NewExporter(r.Context(), *address))

		h := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
		h.ServeHTTP(w, r)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>MaxScale Exporter</title></head>
			<body>
			<h1>MaxScale Exporter</h1>
			<p><a href="` + metricsPath + `">Metrics</a></p>
			</body>
			</html>`))
	})
	log.Printf("Started MaxScale exporter, listening on port: %v", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

// strflag is like flag.String, with value overridden by an environment
// variable (when present). e.g. with address, the env var used as default
// is MAXSCALE_EXPORTER_ADDRESS, if present in env.
func strflag(name, value, usage string) *string {
	if v, ok := os.LookupEnv(envPrefix + strings.ToUpper(name)); ok {
		return flag.String(name, v, usage)
	}
	return flag.String(name, value, usage)
}
