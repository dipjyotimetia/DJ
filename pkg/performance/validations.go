package performance

import (
	"github.com/bojand/ghz/runner"
	"log"
	"time"
)

// ValidatePerformanceBenchmarks validate total req count and average time
func ValidatePerformanceBenchmarks(totalReqCount uint64, averageTime time.Duration, report *runner.Report) {
	if totalReqCount > report.Count {
		log.Fatalf("Total request count more than the expected")
	}
	if averageTime > report.Average {
		log.Fatalf("Average request count more than the expected")
	}
}

// ValidateLatency validate latency
func ValidateLatency(latencyTime map[int]int64, report *runner.Report) {
	for _, details := range report.LatencyDistribution {
		switch details.Percentage {
		case 50:
			if details.Latency.Milliseconds() > latencyTime[50] {
				log.Fatalf("P50: Expected latency: %d received :%d", latencyTime[50], details.Latency.Milliseconds())
			}
		case 75:
			if details.Latency.Milliseconds() > latencyTime[75] {
				log.Fatalf("P75: Expected latency: %d received :%d", latencyTime[75], details.Latency.Milliseconds())
			}
		case 90:
			if details.Latency.Milliseconds() > latencyTime[90] {
				log.Fatalf("P90: Expected latency: %d received :%d", latencyTime[90], details.Latency.Milliseconds())
			}
		case 95:
			if details.Latency.Milliseconds() > latencyTime[95] {
				log.Fatalf("P95: Expected latency: %d received :%d", latencyTime[95], details.Latency.Milliseconds())
			}
		case 99:
			if details.Latency.Milliseconds() > latencyTime[99] {
				log.Fatalf("P99: Expected latency: %d received :%d", latencyTime[99], details.Latency.Milliseconds())
			}
		}
	}
}
