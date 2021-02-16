package performance

import (
	"fmt"
	"github.com/bojand/ghz/runner"
	"github.com/goutils/pkg/performance/collector/influx"
	"os"
)

// PerformanceTestRunner core performance test runner
func PerformanceTestRunner(config Config) {
	report, err := runner.Run(
		config.API,
		config.Host,
		runner.WithName(config.TestName),
		// runner.WithProtoFile("blog/blogPb/blog.proto", []string{}), //TODO: gRPC Reflection api is configured
		runner.WithData(config.RequestBody),
		runner.WithConcurrency(config.Concurrency),    // concurrency = 50
		runner.WithTotalRequests(config.TotalRequest), // totalReq = 300
		runner.WithConnections(config.Connections),    // connections = 20
		runner.WithInsecure(true),
	)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	ValidateLatency(config.Latency, report)
	ValidatePerformanceBenchmarks(config.Expected.TotalReqCount, config.Expected.AverageTime, report)
	influx.WriteReportToInflux(config.InfluxDB.Host, config.InfluxDB.Token, report)
}
