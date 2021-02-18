package influx

import (
	"github.com/influxdata/influxdb-client-go/v2"
)

func asyncWriteClient(host string, token string) influxdb2.Client {
	return influxdb2.NewClientWithOptions(host, token,
		influxdb2.DefaultOptions().SetBatchSize(20))
}

func blockWriteClient(host string, token string) influxdb2.Client {
	return influxdb2.NewClient(host, token)
}
