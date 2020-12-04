package main

import (
	"encoding/json"
)

type Response struct {
	Data json.RawMessage `json:"data"`
}

type ServerData struct {
	ID         string `json:"id"`
	Attributes struct {
		GtidBinlogPos  string `json:"gtid_binlog_pos"`
		GtidCurrentPos string `json:"gtid_current_pos"`
		LastEvent      string `json:"last_event"`
		MasterID       int    `json:"master_id"`
		Name           string `json:"name"`
		NodeID         int    `json:"node_id"`
		ReadOnly       bool   `json:"read_only"`
		ReplicationLag int    `json:"replication_lag"`
		ServerID       int    `json:"server_id"`
		State          string `json:"state"`
		Statistics     struct {
			ActiveOperations      int    `json:"active_operations"`
			AdaptiveAvgSelectTime string `json:"adaptive_avg_select_time"`
			ConnectionPoolEmpty   int    `json:"connection_pool_empty"`
			Connections           int    `json:"connections"`
			MaxConnections        int    `json:"max_connections"`
			MaxPoolSize           int    `json:"max_pool_size"`
			PersistentConnections int    `json:"persistent_connections"`
			ReusedConnections     int    `json:"reused_connections"`
			RoutedPackets         int    `json:"routed_packets"`
			TotalConnections      int    `json:"total_connections"`
		} `json:"statistics"`
		TriggeredAt   string `json:"triggered_at"`
		VersionString string `json:"version_string"`
	} `json:"attributes"`
}

type ServiceData struct {
	Attributes struct {
		Connections int `json:"connections"`
		Listeners   []struct {
			Attributes struct {
				State string `json:"state"`
			} `json:"attributes"`
			ID   string `json:"id"`
			Type string `json:"type"`
		} `json:"listeners"`
		Router            string `json:"router"`
		RouterDiagnostics struct {
			Queries               int                     `json:"queries"`
			ReplayedTransactions  int                     `json:"replayed_transactions"`
			RoTransactions        int                     `json:"ro_transactions"`
			RouteAll              int                     `json:"route_all"`
			RouteMaster           int                     `json:"route_master"`
			RouteSlave            int                     `json:"route_slave"`
			RwTransactions        int                     `json:"rw_transactions"`
			ServerQueryStatistics []ServerQueryStatistics `json:"server_query_statistics"`
		} `json:"router_diagnostics"`
		Started    string `json:"started"`
		State      string `json:"state"`
		Statistics struct {
			ActiveOperations int `json:"active_operations"`
			Connections      int `json:"connections"`
			MaxConnections   int `json:"max_connections"`
			RoutedPackets    int `json:"routed_packets"`
			TotalConnections int `json:"total_connections"`
		} `json:"statistics"`
		TotalConnections int `json:"total_connections"`
	} `json:"attributes"`
	ID string `json:"id"`
}

type ServerQueryStatistics struct {
	AvgSelectsPerSession int     `json:"avg_selects_per_session"`
	AvgSessActivePct     float64 `json:"avg_sess_active_pct"`
	AvgSessDuration      string  `json:"avg_sess_duration"`
	ID                   string  `json:"id"`
	Read                 int     `json:"read"`
	Total                int     `json:"total"`
	Write                int     `json:"write"`
}
