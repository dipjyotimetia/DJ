package performance

import (
	"time"
)

// Config test config
type Config struct {
	Concurrency  uint
	TotalRequest uint
	Connections  uint
	InfluxDB     map[string]string
	Suite        []TestSuite
}

type TestSuite struct {
	API         string
	Host        string
	TestName    string
	RequestBody map[string]interface{}
	Latency     map[int]int64
	Expected    Conditions
}

type Conditions struct {
	TotalReqCount uint64
	AverageTime   time.Duration
}
