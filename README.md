[![Docker Pulls](https://img.shields.io/docker/pulls/rubenhoms/maxscale_exporter.svg)](https://hub.docker.com/r/rubenhoms/maxscale_exporter/)
[![Docker build](https://img.shields.io/docker/cloud/build/rubenhoms/maxscale_exporter)](https://hub.docker.com/r/rubenhoms/maxscale_exporter/builds)

## Overview
This exports the following MaxScale metrics for Prometheus:

- Server connections
- Service session count
- Maxscale instance status and variables
- Event Times (Executed & Queued histograms)

### Requirements:
Your MaxScale instance needs to have the JSON HTTP listener enabled so this can pull the stats from your MaxScale server. You can read [here](https://mariadb.com/kb/en/mariadb-enterprise/mariadb-maxscale-14/maxinfo-plugin/#configuration) how to set this up.

**NOTE:** This will work with all MaxScale versions < 2.4. For MaxScale versions > 2.4, you should use the [maxctrl_exporter](https://github.com/Vetal1977/maxctrl_exporter). This is due to support being dropped for the MaxAdmin & MaxInfo in the newer versions.

## Installation
Quite simple: get [Go](https://golang.org/dl), set a `$GOPATH`, and run

    go get github.com/RubenHoms/maxscale_exporter

## Use
Make sure `$GOPATH/bin` is in your `$PATH`.

    $ maxscale_exporter -h
    Usage of maxscale_exporter:
      -address string
        	address to get maxscale statistics from (default "127.0.0.1:8003")
      -pidfile string
        	the pid file for maxscale to monitor process statistics
      -port string
        	the port that the maxscale exporter listens on (default "9195")

## Process metrics
This exporter exposes two different sets of process metrics but only one of them is enabled by default.
The metrics of the exporter process itself are exposed in the `process` namepsace (e.g. `process_start_time_seconds`), this happens automatically and no further setup is needed.

However if you want to expose metrics about the MaxScale process itself, you can do that by specifying the `-pidfile` flag.
However this only works in Linux systems where `/proc` is available. For example you could set this up by specifying `-pidfile /run/maxscale/maxscale.pid` as a flag.
Note that the user that runs the exporter process needs to have read access to the pidfile in order for this to work.

## Testing locally
If you want to try out the exporter locally, a Dockerfile is provided to run a MaxScale instance with the HTTP JSON listener enabled. To test it out locally you run:

```
$ cd maxscale_docker
$ docker build . -t maxscale_maxinfo:latest
$ docker run -d -p 8003:8003 maxscale_maxinfo:latest
```
If you then run the maxscale_exporter it should use the default settings and pull statistics from the Docker container.

### Example output
```
# HELP maxscale_events_executed_seconds Amount of events executed
# TYPE maxscale_events_executed_seconds histogram
maxscale_events_executed_seconds_bucket{le="0.1"} 999041
maxscale_events_executed_seconds_bucket{le="0.2"} 337
maxscale_events_executed_seconds_bucket{le="0.3"} 0
maxscale_events_executed_seconds_bucket{le="0.4"} 0
maxscale_events_executed_seconds_bucket{le="0.5"} 0
maxscale_events_executed_seconds_bucket{le="0.6"} 0
maxscale_events_executed_seconds_bucket{le="0.7"} 0
maxscale_events_executed_seconds_bucket{le="0.8"} 0
maxscale_events_executed_seconds_bucket{le="0.9"} 0
maxscale_events_executed_seconds_bucket{le="1"} 0
maxscale_events_executed_seconds_bucket{le="1.1"} 0
maxscale_events_executed_seconds_bucket{le="1.2"} 0
maxscale_events_executed_seconds_bucket{le="1.3"} 0
maxscale_events_executed_seconds_bucket{le="1.4"} 0
maxscale_events_executed_seconds_bucket{le="1.5"} 0
maxscale_events_executed_seconds_bucket{le="1.6"} 0
maxscale_events_executed_seconds_bucket{le="1.7"} 0
maxscale_events_executed_seconds_bucket{le="1.8"} 0
maxscale_events_executed_seconds_bucket{le="1.9"} 0
maxscale_events_executed_seconds_bucket{le="2"} 0
maxscale_events_executed_seconds_bucket{le="2.1"} 0
maxscale_events_executed_seconds_bucket{le="2.2"} 0
maxscale_events_executed_seconds_bucket{le="2.3"} 0
maxscale_events_executed_seconds_bucket{le="2.4"} 0
maxscale_events_executed_seconds_bucket{le="2.5"} 0
maxscale_events_executed_seconds_bucket{le="2.6"} 0
maxscale_events_executed_seconds_bucket{le="2.7"} 0
maxscale_events_executed_seconds_bucket{le="2.8"} 0
maxscale_events_executed_seconds_bucket{le="2.9"} 0
maxscale_events_executed_seconds_bucket{le="+Inf"} 999378
maxscale_events_executed_seconds_sum 99971.5
maxscale_events_executed_seconds_count 999378
# HELP maxscale_events_queued_seconds Amount of events queued
# TYPE maxscale_events_queued_seconds histogram
maxscale_events_queued_seconds_bucket{le="0.1"} 999306
maxscale_events_queued_seconds_bucket{le="0.2"} 73
maxscale_events_queued_seconds_bucket{le="0.3"} 0
maxscale_events_queued_seconds_bucket{le="0.4"} 0
maxscale_events_queued_seconds_bucket{le="0.5"} 0
maxscale_events_queued_seconds_bucket{le="0.6"} 0
maxscale_events_queued_seconds_bucket{le="0.7"} 0
maxscale_events_queued_seconds_bucket{le="0.8"} 0
maxscale_events_queued_seconds_bucket{le="0.9"} 0
maxscale_events_queued_seconds_bucket{le="1"} 0
maxscale_events_queued_seconds_bucket{le="1.1"} 0
maxscale_events_queued_seconds_bucket{le="1.2"} 0
maxscale_events_queued_seconds_bucket{le="1.3"} 0
maxscale_events_queued_seconds_bucket{le="1.4"} 0
maxscale_events_queued_seconds_bucket{le="1.5"} 0
maxscale_events_queued_seconds_bucket{le="1.6"} 0
maxscale_events_queued_seconds_bucket{le="1.7"} 0
maxscale_events_queued_seconds_bucket{le="1.8"} 0
maxscale_events_queued_seconds_bucket{le="1.9"} 0
maxscale_events_queued_seconds_bucket{le="2"} 0
maxscale_events_queued_seconds_bucket{le="2.1"} 0
maxscale_events_queued_seconds_bucket{le="2.2"} 0
maxscale_events_queued_seconds_bucket{le="2.3"} 0
maxscale_events_queued_seconds_bucket{le="2.4"} 0
maxscale_events_queued_seconds_bucket{le="2.5"} 0
maxscale_events_queued_seconds_bucket{le="2.6"} 0
maxscale_events_queued_seconds_bucket{le="2.7"} 0
maxscale_events_queued_seconds_bucket{le="2.8"} 0
maxscale_events_queued_seconds_bucket{le="2.9"} 0
maxscale_events_queued_seconds_bucket{le="+Inf"} 999379
maxscale_events_queued_seconds_sum 99945.20000000001
maxscale_events_queued_seconds_count 999379
# HELP maxscale_exporter_total_scrapes Current total MaxScale scrapes
# TYPE maxscale_exporter_total_scrapes counter
maxscale_exporter_total_scrapes 1
# HELP maxscale_process_cpu_seconds_total Total user and system CPU time spent in seconds.
# TYPE maxscale_process_cpu_seconds_total counter
maxscale_process_cpu_seconds_total 70.98
# HELP maxscale_process_max_fds Maximum number of open file descriptors.
# TYPE maxscale_process_max_fds gauge
maxscale_process_max_fds 65535
# HELP maxscale_process_open_fds Number of open file descriptors.
# TYPE maxscale_process_open_fds gauge
maxscale_process_open_fds 49
# HELP maxscale_process_resident_memory_bytes Resident memory size in bytes.
# TYPE maxscale_process_resident_memory_bytes gauge
maxscale_process_resident_memory_bytes 5.6098816e+07
# HELP maxscale_process_start_time_seconds Start time of the process since unix epoch in seconds.
# TYPE maxscale_process_start_time_seconds gauge
maxscale_process_start_time_seconds 1.52275537227e+09
# HELP maxscale_process_virtual_memory_bytes Virtual memory size in bytes.
# TYPE maxscale_process_virtual_memory_bytes gauge
maxscale_process_virtual_memory_bytes 5.81009408e+08
# HELP maxscale_server_connections Amount of connections to the server
# TYPE maxscale_server_connections gauge
maxscale_server_connections{address="1.3.3.7",server="db0"} 11
maxscale_server_connections{address="1.3.3.7",server="db1"} 6
# HELP maxscale_server_up Is the server up
# TYPE maxscale_server_up gauge
maxscale_server_up{address="1.3.3.7",server="db0"} 1
maxscale_server_up{address="1.3.3.8",server="db1"} 1
# HELP maxscale_service_current_sessions Amount of sessions currently active
# TYPE maxscale_service_current_sessions gauge
maxscale_service_current_sessions{name="CLI",router="cli"} 1
maxscale_service_current_sessions{name="MaxInfo",router="maxinfo"} 2
maxscale_service_current_sessions{name="db0",router="readconnroute"} 12
maxscale_service_current_sessions{name="db1",router="readconnroute"} 7
# HELP maxscale_service_total_sessions Total amount of sessions
# TYPE maxscale_service_total_sessions counter
maxscale_service_total_sessions{name="CLI",router="cli"} 5558
maxscale_service_total_sessions{name="MaxInfo",router="maxinfo"} 495
maxscale_service_total_sessions{name="db0",router="readconnroute"} 771
maxscale_service_total_sessions{name="db1",router="readconnroute"} 9
# HELP maxscale_status_accept_events How many accept events happened
# TYPE maxscale_status_accept_events counter
maxscale_status_accept_events 6830
# HELP maxscale_status_backend_connections How many backend connections there are
# TYPE maxscale_status_backend_connections gauge
maxscale_status_backend_connections 17
# HELP maxscale_status_client_connections How many client connections there are
# TYPE maxscale_status_client_connections gauge
maxscale_status_client_connections 18
# HELP maxscale_status_connections How many connections there are
# TYPE maxscale_status_connections gauge
maxscale_status_connections 39
# HELP maxscale_status_error_events How many error events happened
# TYPE maxscale_status_error_events counter
maxscale_status_error_events 0
# HELP maxscale_status_event_queue_length How long the event queue is
# TYPE maxscale_status_event_queue_length gauge
maxscale_status_event_queue_length 1
# HELP maxscale_status_hangup_events How many hangup events happened
# TYPE maxscale_status_hangup_events counter
maxscale_status_hangup_events 5817
# HELP maxscale_status_internal_descriptors How many internal descriptors there are
# TYPE maxscale_status_internal_descriptors gauge
maxscale_status_internal_descriptors 35
# HELP maxscale_status_listeners How many listeners there are
# TYPE maxscale_status_listeners gauge
maxscale_status_listeners 4
# HELP maxscale_status_max_event_execution_time The max event execution time
# TYPE maxscale_status_max_event_execution_time gauge
maxscale_status_max_event_execution_time 1
# HELP maxscale_status_max_event_queue_length The max length of the event queue
# TYPE maxscale_status_max_event_queue_length gauge
maxscale_status_max_event_queue_length 4
# HELP maxscale_status_max_event_queue_time The max event queue time
# TYPE maxscale_status_max_event_queue_time gauge
maxscale_status_max_event_queue_time 1
# HELP maxscale_status_pending_events How many events are pending
# TYPE maxscale_status_pending_events gauge
maxscale_status_pending_events 0
# HELP maxscale_status_read_events How many read events happened
# TYPE maxscale_status_read_events counter
maxscale_status_read_events 985046
# HELP maxscale_status_threadpool_threads How many threadpool threads there are
# TYPE maxscale_status_threadpool_threads gauge
maxscale_status_threadpool_threads 1
# HELP maxscale_status_threads_connected How many threads are connected
# TYPE maxscale_status_threads_connected gauge
maxscale_status_threads_connected 22
# HELP maxscale_status_threads_created How many threads have been created
# TYPE maxscale_status_threads_created counter
maxscale_status_threads_created 1
# HELP maxscale_status_threads_running How many threads are running
# TYPE maxscale_status_threads_running gauge
maxscale_status_threads_running 1
# HELP maxscale_status_uptime How long has the server been running
# TYPE maxscale_status_uptime counter
maxscale_status_uptime 79379
# HELP maxscale_status_uptime_since_flush_status How long the server has been up since flush status
# TYPE maxscale_status_uptime_since_flush_status counter
maxscale_status_uptime_since_flush_status 79379
# HELP maxscale_status_write_events How many write events happened
# TYPE maxscale_status_write_events counter
maxscale_status_write_events 992543
# HELP maxscale_status_zombie_connections How many zombie connetions there are
# TYPE maxscale_status_zombie_connections gauge
maxscale_status_zombie_connections 0
# HELP maxscale_up Was the last scrape of MaxScale successful?
# TYPE maxscale_up gauge
maxscale_up 1
# HELP maxscale_variables_nbpolls MAXSCALE_NBPOLLS
# TYPE maxscale_variables_nbpolls gauge
maxscale_variables_nbpolls 3
# HELP maxscale_variables_pollsleep MAXSCALE_POLLSLEEP
# TYPE maxscale_variables_pollsleep gauge
maxscale_variables_pollsleep 1000
# HELP maxscale_variables_sessions MAXSCALE_SESSIONS
# TYPE maxscale_variables_sessions gauge
maxscale_variables_sessions 22
# HELP maxscale_variables_thread MAXSCALE_THREADS
# TYPE maxscale_variables_thread gauge
maxscale_variables_thread 1

```
