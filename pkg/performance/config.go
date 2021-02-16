package performance

import (
	"time"
)

// Config test config
type Config struct {
	API          string
	Host         string
	TestName     string
	RequestBody  interface{}
	Concurrency  uint
	TotalRequest uint
	Connections  uint
	Latency      map[int]int64
	Expected     Conditions
	InfluxDB     InfluxDBConnection
}

type Conditions struct {
	TotalReqCount uint64
	AverageTime   time.Duration
}

type InfluxDBConnection struct {
	Host  string
	Token string
}
