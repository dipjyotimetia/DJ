package influx

import (
	"fmt"
	"github.com/bojand/ghz/runner"
	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"time"
)

// WriteReportToInflux send report details to influxdb
func WriteReportToInflux(influxCon map[string]string, reportDetails *runner.Report) {
	if influxCon["enabled"] == "Yes" {
		client := asyncWriteClient(influxCon["host"], influxCon["token"])
		writeAPI := client.WriteAPI(influxCon["org"], influxCon["bucket"])

		point := influxdb2.NewPoint(
			"grpc",
			map[string]string{
				"id":       fmt.Sprintf("rack_%v", uuid.New().ID()),
				"vendor":   "DJ",
				"hostname": fmt.Sprintf(influxCon["host"]),
			},
			map[string]interface{}{
				"TestName":            reportDetails.Name,
				"EndReason":           reportDetails.EndReason,
				"Date":                reportDetails.Date,
				"TotalCount":          int(reportDetails.Count),
				"TotalDuration":       reportDetails.Total,
				"Average":             reportDetails.Average,
				"Fastest":             reportDetails.Fastest,
				"Slowest":             reportDetails.Slowest,
				"RPS":                 reportDetails.Rps,
				"ErrorDist":           reportDetails.ErrorDist,
				"StatusCodeDist":      reportDetails.StatusCodeDist,
				"LatencyDistribution": reportDetails.LatencyDistribution,
				"Details":             reportDetails.Details,
			},
			time.Now())
		writeAPI.WritePoint(point)
		writeAPI.Flush()
		client.Close()
	}
}
