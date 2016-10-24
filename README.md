This exports the following MaxScale metrics for Prometheus:

- Server connections (gauges)
- Service session count (gauges)
- Maxscale instance status and variables (gauge & counter mix)
- Event Times (Executed & Queued histograms)

Environment:

Your MaxScale instance needs to have the JSON HTTP listener enabled so this can pull the stats from your MaxScale server, and you need to specify the listen address for the exporter:

````
MAXSCALE_MAXINFO_JSON_LISTENER_TCP_ADDR=maxscale:8003
MAXSCALE_EXPORTER_LISTEN_ADDR=:9195
```

Example output:

```
#TYPE maxscale_servers_connections gauge
#HELP maxscale_servers_connections Server connections
maxscale_servers_connections{server="mariadb1", address="mariadb1", port="3306"} 0
maxscale_servers_connections{server="mariadb2", address="mariadb2", port="3306"} 0
maxscale_servers_connections{server="mariadb3", address="mariadb3", port="3306"} 0
maxscale_servers_connections{server="mariadb4", address="mariadb4", port="3306"} 0
maxscale_servers_connections{server="mariadb5", address="mariadb5", port="3306"} 0

#TYPE maxscale_services_sessions gauge
#HELP maxscale_services_sessions Service Sessions
maxscale_services_sessions{name="Splitter Service", router="readwritesplit"} 1
maxscale_services_sessions{name="Read-Only Service", router="readconnroute"} 1
maxscale_services_sessions{name="MaxAdmin Service", router="cli"} 1
maxscale_services_sessions{name="MaxInfo", router="maxinfo"} 3

#TYPE maxscale_status_uptime counter
#HELP maxscale_status_uptime Uptime
maxscale_status_uptime 812

#TYPE maxscale_status_uptime_since_flush_status counter
#HELP maxscale_status_uptime_since_flush_status Uptime_since_flush_status
maxscale_status_uptime_since_flush_status 812

#TYPE maxscale_status_threads_created counter
#HELP maxscale_status_threads_created Threads_created
maxscale_status_threads_created 1

#TYPE maxscale_status_threads_running gauge
#HELP maxscale_status_threads_running Threads_running
maxscale_status_threads_running 1

#TYPE maxscale_status_threadpool_threads counter
#HELP maxscale_status_threadpool_threads Threadpool_threads
maxscale_status_threadpool_threads 1

#TYPE maxscale_status_threads_connected gauge
#HELP maxscale_status_threads_connected Threads_connected
maxscale_status_threads_connected 6

#TYPE maxscale_status_connections gauge
#HELP maxscale_status_connections Connections
maxscale_status_connections 6

#TYPE maxscale_status_client_connections gauge
#HELP maxscale_status_client_connections Client_connections
maxscale_status_client_connections 1

#TYPE maxscale_status_backend_connections gauge
#HELP maxscale_status_backend_connections Backend_connections
maxscale_status_backend_connections 0

#TYPE maxscale_status_listeners gauge
#HELP maxscale_status_listeners Listeners
maxscale_status_listeners 5

#TYPE maxscale_status_zombie_connections gauge
#HELP maxscale_status_zombie_connections Zombie_connections
maxscale_status_zombie_connections 0

#TYPE maxscale_status_internal_descriptors gauge
#HELP maxscale_status_internal_descriptors Internal_descriptors
maxscale_status_internal_descriptors 1

#TYPE maxscale_status_read_events counter
#HELP maxscale_status_read_events Read_events
maxscale_status_read_events 718

#TYPE maxscale_status_write_events counter
#HELP maxscale_status_write_events Write_events
maxscale_status_write_events 1396

#TYPE maxscale_status_hangup_events counter
#HELP maxscale_status_hangup_events Hangup_events
maxscale_status_hangup_events 0

#TYPE maxscale_status_error_events counter
#HELP maxscale_status_error_events Error_events
maxscale_status_error_events 0

#TYPE maxscale_status_accept_events counter
#HELP maxscale_status_accept_events Accept_events
maxscale_status_accept_events 718

#TYPE maxscale_status_event_queue_length gauge
#HELP maxscale_status_event_queue_length Event_queue_length
maxscale_status_event_queue_length 1

#TYPE maxscale_status_pending_events gauge
#HELP maxscale_status_pending_events Pending_events
maxscale_status_pending_events 0

#TYPE maxscale_status_max_event_queue_length gauge
#HELP maxscale_status_max_event_queue_length Max_event_queue_length
maxscale_status_max_event_queue_length 1

#TYPE maxscale_status_max_event_queue_time gauge
#HELP maxscale_status_max_event_queue_time Max_event_queue_time
maxscale_status_max_event_queue_time 0

#TYPE maxscale_status_max_event_execution_time gauge
#HELP maxscale_status_max_event_execution_time Max_event_execution_time
maxscale_status_max_event_execution_time 1

#TYPE maxscale_variables_maxscale_threads gauge
#HELP maxscale_variables_maxscale_threads MAXSCALE_THREADS
maxscale_variables_maxscale_threads 1

#TYPE maxscale_variables_maxscale_nbpolls gauge
#HELP maxscale_variables_maxscale_nbpolls MAXSCALE_NBPOLLS
maxscale_variables_maxscale_nbpolls 3

#TYPE maxscale_variables_maxscale_pollsleep gauge
#HELP maxscale_variables_maxscale_pollsleep MAXSCALE_POLLSLEEP
maxscale_variables_maxscale_pollsleep 1000

#TYPE maxscale_variables_maxscale_uptime counter
#HELP maxscale_variables_maxscale_uptime MAXSCALE_UPTIME
maxscale_variables_maxscale_uptime 812

#TYPE maxscale_variables_maxscale_sessions gauge
#HELP maxscale_variables_maxscale_sessions MAXSCALE_SESSIONS
maxscale_variables_maxscale_sessions 6

#TYPE maxscale_events_executed_seconds histogram
#HELP maxscale_events_executed_seconds Events Executed
maxscale_events_executed_seconds_bucket{le="0.100000"} 2103
maxscale_events_executed_seconds_bucket{le="0.200000"} 2119
maxscale_events_executed_seconds_bucket{le="0.300000"} 2119
maxscale_events_executed_seconds_bucket{le="0.400000"} 2119
maxscale_events_executed_seconds_bucket{le="0.500000"} 2119
maxscale_events_executed_seconds_bucket{le="0.600000"} 2119
maxscale_events_executed_seconds_bucket{le="0.700000"} 2119
maxscale_events_executed_seconds_bucket{le="0.800000"} 2119
maxscale_events_executed_seconds_bucket{le="0.900000"} 2119
maxscale_events_executed_seconds_bucket{le="1.000000"} 2119
maxscale_events_executed_seconds_bucket{le="1.100000"} 2119
maxscale_events_executed_seconds_bucket{le="1.200000"} 2119
maxscale_events_executed_seconds_bucket{le="1.300000"} 2119
maxscale_events_executed_seconds_bucket{le="1.400000"} 2119
maxscale_events_executed_seconds_bucket{le="1.500000"} 2119
maxscale_events_executed_seconds_bucket{le="1.600000"} 2119
maxscale_events_executed_seconds_bucket{le="1.700000"} 2119
maxscale_events_executed_seconds_bucket{le="1.800000"} 2119
maxscale_events_executed_seconds_bucket{le="1.900000"} 2119
maxscale_events_executed_seconds_bucket{le="2.000000"} 2119
maxscale_events_executed_seconds_bucket{le="2.100000"} 2119
maxscale_events_executed_seconds_bucket{le="2.200000"} 2119
maxscale_events_executed_seconds_bucket{le="2.300000"} 2119
maxscale_events_executed_seconds_bucket{le="2.400000"} 2119
maxscale_events_executed_seconds_bucket{le="2.500000"} 2119
maxscale_events_executed_seconds_bucket{le="2.600000"} 2119
maxscale_events_executed_seconds_bucket{le="2.700000"} 2119
maxscale_events_executed_seconds_bucket{le="2.800000"} 2119
maxscale_events_executed_seconds_bucket{le="2.900000"} 2119
maxscale_events_executed_seconds_bucket{le="+Inf"} 2119
maxscale_events_executed_seconds_sum 213
maxscale_events_executed_seconds_count 2119

#TYPE maxscale_events_queued_seconds histogram
#HELP maxscale_events_queued_seconds Events Queued
maxscale_events_queued_seconds_bucket{le="0.100000"} 2120
maxscale_events_queued_seconds_bucket{le="0.200000"} 2120
maxscale_events_queued_seconds_bucket{le="0.300000"} 2120
maxscale_events_queued_seconds_bucket{le="0.400000"} 2120
maxscale_events_queued_seconds_bucket{le="0.500000"} 2120
maxscale_events_queued_seconds_bucket{le="0.600000"} 2120
maxscale_events_queued_seconds_bucket{le="0.700000"} 2120
maxscale_events_queued_seconds_bucket{le="0.800000"} 2120
maxscale_events_queued_seconds_bucket{le="0.900000"} 2120
maxscale_events_queued_seconds_bucket{le="1.000000"} 2120
maxscale_events_queued_seconds_bucket{le="1.100000"} 2120
maxscale_events_queued_seconds_bucket{le="1.200000"} 2120
maxscale_events_queued_seconds_bucket{le="1.300000"} 2120
maxscale_events_queued_seconds_bucket{le="1.400000"} 2120
maxscale_events_queued_seconds_bucket{le="1.500000"} 2120
maxscale_events_queued_seconds_bucket{le="1.600000"} 2120
maxscale_events_queued_seconds_bucket{le="1.700000"} 2120
maxscale_events_queued_seconds_bucket{le="1.800000"} 2120
maxscale_events_queued_seconds_bucket{le="1.900000"} 2120
maxscale_events_queued_seconds_bucket{le="2.000000"} 2120
maxscale_events_queued_seconds_bucket{le="2.100000"} 2120
maxscale_events_queued_seconds_bucket{le="2.200000"} 2120
maxscale_events_queued_seconds_bucket{le="2.300000"} 2120
maxscale_events_queued_seconds_bucket{le="2.400000"} 2120
maxscale_events_queued_seconds_bucket{le="2.500000"} 2120
maxscale_events_queued_seconds_bucket{le="2.600000"} 2120
maxscale_events_queued_seconds_bucket{le="2.700000"} 2120
maxscale_events_queued_seconds_bucket{le="2.800000"} 2120
maxscale_events_queued_seconds_bucket{le="2.900000"} 2120
maxscale_events_queued_seconds_bucket{le="+Inf"} 2120
maxscale_events_queued_seconds_sum 212
maxscale_events_queued_seconds_count 2120
```
