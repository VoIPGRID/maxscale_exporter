package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/wearespindle/maxscale_exporter/maxscale"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const envPrefix = "MAXSCALE_EXPORTER"

// strflag is like flag.String, with value overridden by an environment
// variable (when present). e.g. with name token, the env var used as default
// is DAVE_TOKEN, if present in env.
func strflag(name string, value string, usage string) *string {
	if v, ok := os.LookupEnv(envPrefix + strings.ToUpper(name)); ok {
		return flag.String(name, v, usage)
	}
	return flag.String(name, value, usage)
}

var (
	address *string
	port    *string
)

func main() {
	log.SetFlags(0)

	address = strflag("address", "127.0.0.1:8003", "http json address to get maxscale statistics from")
	port = strflag("port", ":9195", "the port that the maxscale exporter listens on")
	flag.Parse()

	m, err := maxscale.New(*address)
	if err != nil {
		log.Fatalf("Failed to start maxscale exporter: %v\n", err)
	}

	http.HandleFunc("/metrics", metrics)
	err := http.ListenAndServe(listenAddress, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
