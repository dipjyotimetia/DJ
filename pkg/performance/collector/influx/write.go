package influx

import (
	"fmt"
	"github.com/bojand/ghz/runner"
	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"time"
)

func WriteReportToInflux(host, token string, reportDetails *runner.Report) {
	client := asyncWriteClient(host, token)
	writeAPI := client.WriteAPI("my-org", "my-bucket")
	for _, details := range reportDetails.Details {
		p := influxdb2.NewPoint(
			"performance",
			map[string]string{
				"id":       fmt.Sprintf("rack_%v", uuid.New().ID()),
				"vendor":   "DJ",
				"hostname": fmt.Sprintf("host_%v", host),
			},
			map[string]interface{}{
				"TimeStamp": details.Timestamp,
				"Status":    details.Status,
				"Latency":   details.Latency,
				"Error":     details.Error,
			},
			time.Now())
		writeAPI.WritePoint(p)
	}
	for _, latency := range reportDetails.LatencyDistribution {
		p := influxdb2.NewPoint(
			"latency",
			map[string]string{
				"id":       fmt.Sprintf("rack_%v", uuid.New().ID()),
				"vendor":   "DJ",
				"hostname": fmt.Sprintf("host_%v", host),
			},
			map[string]interface{}{
				"Percentage": latency.Percentage,
				"Latency":    latency.Latency,
			},
			time.Now())
		writeAPI.WritePoint(p)
	}
	writeAPI.Flush()
	client.Close()
}
