package performance

import (
	"encoding/json"
	"fmt"
	"github.com/bojand/ghz/runner"
	"github.com/goutils/pkg/performance/collector/influx"
	"os"
	"time"
)

// ResultDetails performance result details
type ResultDetails struct {
	TestName    string
	AverageTime int64
	Count       uint64
	TotalTime   int64
	Date        time.Time
	RPS         float64
	Latency     []runner.LatencyDistribution
}

// TestRunner core performance test runner
func TestRunner(config *Config) {
	for _, suite := range config.Suite {
		report, err := runner.Run(
			suite.API,
			suite.Host,
			runner.WithName(suite.TestName),
			//  runner.WithProtoFile("blog/blogPb/blog.proto", []string{}),
			runner.WithData(suite.RequestBody),
			runner.WithConcurrency(config.Concurrency),    // concurrency = 50
			runner.WithTotalRequests(config.TotalRequest), // totalReq = 300
			runner.WithConnections(config.Connections),    // connections = 20
			runner.WithInsecure(true),
		)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		ValidateLatency(suite.Latency, report)
		// ValidatePerformanceBenchmarks(suite.Expected.TotalReqCount, suite.Expected.AverageTime, report)
		influx.WriteReportToInflux(config.InfluxDB, report)

		obj, _ := json.Marshal(ResultDetails{
			TestName:    report.Name,
			AverageTime: report.Average.Milliseconds(),
			Count:       report.Count,
			TotalTime:   report.Total.Milliseconds(),
			Date:        report.Date,
			RPS:         report.Rps,
			Latency:     report.LatencyDistribution,
		})
		fmt.Println(string(obj))
	}
}
