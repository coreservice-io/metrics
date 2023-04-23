package metrics

import (
    "time"

    "github.com/rcrowley/go-metrics"
    reporter "github.com/zakhio/go-metrics-influxdb"
)

type InfluxDBConfig struct {
    ServerUrl      string
    OrganizationId string
    BucketId       string
    Measurement    string
    Token          string
    AlignTimestamp bool
}

func newInfluxDB(
    submitPeriod time.Duration,
    config *InfluxDBConfig,
) {
    go reporter.InfluxDB(
        metrics.DefaultRegistry,
        submitPeriod,
        config.ServerUrl,
        config.OrganizationId,
        config.BucketId,
        config.Measurement,
        config.Token,
        config.AlignTimestamp,
    )
}
