package stats

type Stats struct {
	MaxConnections     int64 `json:"max_connections"`
	CurrentConnections int   `json:"current_connections"`
}
