##Overview

This exports the following MaxScale metrics for Prometheus:

- Server connections
- Service session count
- Maxscale instance status and variables
- Event Times (Executed & Queued histograms)

### Requirements:

Your MaxScale instance needs to have the JSON HTTP listener enabled so this can pull the stats from your MaxScale server. You can read [here](https://mariadb.com/kb/en/mariadb-enterprise/mariadb-maxscale-14/maxinfo-plugin/#configuration) how to set this up.

## Installation

Quite simple: get [Go](https://golang.org/dl), set a `$GOPATH`, and run

    go get github.com/wearespindle/maxscale_exporter

## Use

Make sure `$GOPATH/bin` is in your `$PATH`.

    $ maxscale_exporter -h
    Usage of maxscale_exporter:
      -address string
            address to get maxscale statistics from (default "127.0.0.1:8003")
      -port string
            the port that the maxscale exporter listens on (default ":9195")

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
maxscale_events_executed_seconds_bucket{le="0.1"} 206
maxscale_events_executed_seconds_bucket{le="0.2"} 2
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
maxscale_events_executed_seconds_bucket{le="+Inf"} 208
maxscale_events_executed_seconds_sum 21
maxscale_events_executed_seconds_count 208
# HELP maxscale_events_queued_seconds Amount of events queued
# TYPE maxscale_events_queued_seconds histogram
maxscale_events_queued_seconds_bucket{le="0.1"} 209
maxscale_events_queued_seconds_bucket{le="0.2"} 0
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
maxscale_events_queued_seconds_bucket{le="+Inf"} 209
maxscale_events_queued_seconds_sum 20.900000000000002
maxscale_events_queued_seconds_count 209
# HELP maxscale_exporter_total_scrapes Current total MaxScale scrapes
# TYPE maxscale_exporter_total_scrapes counter
maxscale_exporter_total_scrapes 1
# HELP maxscale_server_connections Amount of connections to the server
# TYPE maxscale_server_connections gauge
maxscale_server_connections{address="1.3.3.7",server="db1"} 4
maxscale_server_connections{address="1.3.3.8",server="db2"} 2
# HELP maxscale_server_up Is the server up
# TYPE maxscale_server_up gauge
maxscale_server_up{address="1.3.3.7",server="db1"} 1
maxscale_server_up{address="1.3.3.8",server="db2"} 1
# HELP maxscale_service_current_sessions Amount of sessions currently active
# TYPE maxscale_service_current_sessions gauge
maxscale_service_current_sessions{name="CLI",router="cli"} 1
maxscale_service_current_sessions{name="MaxInfo",router="maxinfo"} 2
maxscale_service_current_sessions{name="db1",router="readconnroute"} 5
maxscale_service_current_sessions{name="db2",router="readconnroute"} 3
# HELP maxscale_service_total_sessions Total amount of sessions
# TYPE maxscale_service_total_sessions gauge
maxscale_service_total_sessions{name="CLI",router="cli"} 1
maxscale_service_total_sessions{name="MaxInfo",router="maxinfo"} 4
maxscale_service_total_sessions{name="db1",router="readconnroute"} 5
maxscale_service_total_sessions{name="db2",router="readconnroute"} 3
# HELP maxscale_status_accept_events How many accept events happened
# TYPE maxscale_status_accept_events gauge
maxscale_status_accept_events 10
# HELP maxscale_status_backend_connections How many backend connections there are
# TYPE maxscale_status_backend_connections gauge
maxscale_status_backend_connections 6
# HELP maxscale_status_client_connections How many client connections there are
# TYPE maxscale_status_client_connections gauge
maxscale_status_client_connections 7
# HELP maxscale_status_connections How many connections there are
# TYPE maxscale_status_connections gauge
maxscale_status_connections 17
# HELP maxscale_status_error_events How many error events happened
# TYPE maxscale_status_error_events gauge
maxscale_status_error_events 0
# HELP maxscale_status_event_queue_length How long the event queue is
# TYPE maxscale_status_event_queue_length gauge
maxscale_status_event_queue_length 1
# HELP maxscale_status_hangup_events How many hangup events happened
# TYPE maxscale_status_hangup_events gauge
maxscale_status_hangup_events 0
# HELP maxscale_status_internal_descriptors How many internal descriptors there are
# TYPE maxscale_status_internal_descriptors gauge
maxscale_status_internal_descriptors 13
# HELP maxscale_status_listeners How many listeners there are
# TYPE maxscale_status_listeners gauge
maxscale_status_listeners 4
# HELP maxscale_status_max_event_execution_time The max event execution time
# TYPE maxscale_status_max_event_execution_time gauge
maxscale_status_max_event_execution_time 1
# HELP maxscale_status_max_event_queue_length The max length of the event queue
# TYPE maxscale_status_max_event_queue_length gauge
maxscale_status_max_event_queue_length 2
# HELP maxscale_status_max_event_queue_time The max event queue time
# TYPE maxscale_status_max_event_queue_time gauge
maxscale_status_max_event_queue_time 0
# HELP maxscale_status_pending_events How many events are pending
# TYPE maxscale_status_pending_events gauge
maxscale_status_pending_events 0
# HELP maxscale_status_read_events How many read events happened
# TYPE maxscale_status_read_events gauge
maxscale_status_read_events 178
# HELP maxscale_status_threadpool_threads How many threadpool threads there are
# TYPE maxscale_status_threadpool_threads gauge
maxscale_status_threadpool_threads 1
# HELP maxscale_status_threads_connected How many threads are connected
# TYPE maxscale_status_threads_connected gauge
maxscale_status_threads_connected 11
# HELP maxscale_status_threads_created How many threads have been created
# TYPE maxscale_status_threads_created gauge
maxscale_status_threads_created 1
# HELP maxscale_status_threads_running How many threads are running
# TYPE maxscale_status_threads_running gauge
maxscale_status_threads_running 1
# HELP maxscale_status_uptime How long has the server been running
# TYPE maxscale_status_uptime gauge
maxscale_status_uptime 12
# HELP maxscale_status_uptime_since_flush_status How long the server has been up since flush status
# TYPE maxscale_status_uptime_since_flush_status gauge
maxscale_status_uptime_since_flush_status 12
# HELP maxscale_status_write_events How many write events happened
# TYPE maxscale_status_write_events gauge
maxscale_status_write_events 193
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
maxscale_variables_sessions 11
# HELP maxscale_variables_thread MAXSCALE_THREADS
# TYPE maxscale_variables_thread gauge
maxscale_variables_thread 1
```
