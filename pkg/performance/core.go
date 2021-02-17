package performance

import (
	"fmt"
	"github.com/bojand/ghz/runner"
	"github.com/goutils/pkg/performance/collector/influx"
	"os"
)

// TestRunner core performance test runner
func TestRunner(config *Config) {
	for _, suite := range config.Suite {
		report, err := runner.Run(
			suite.API,
			suite.Host,
			runner.WithName(suite.TestName),
			//  runner.WithProtoFile("blog/blogPb/blog.proto", []string{}), //TODO: gRPC Reflection api is configured
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
	}
}
