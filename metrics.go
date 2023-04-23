package metrics

import (
    "errors"
    "time"

    "github.com/rcrowley/go-metrics"
)

type DBEngine string

const (
    InfluxDB DBEngine = "influxdb"
)

// NewMetrics - The connection status with influxdb is not guaranteed here.
func NewMetrics(
    submitPeriod time.Duration,
    dbEngine DBEngine,
    dbConfig interface{},
) error {
    switch dbEngine {
    case InfluxDB:
        config := dbConfig.(*InfluxDBConfig)
        newInfluxDB(submitPeriod, config)
    default:
        return errors.New("unsupported DB Engine")
    }
    return nil
}

type g struct {
    gauge metrics.Gauge
}

func Gauge(name string) *g {
    return &g{gauge: metrics.GetOrRegisterGauge(name, nil)}
}

func (gg *g) Update(value int64) {
    gg.gauge.Update(value)
}
func (gg *g) Value() int64 {
    return gg.gauge.Value()
}

type c struct {
    counter metrics.Counter
}

func Counter(name string) *c {
    return &c{counter: metrics.GetOrRegisterCounter(name, nil)}
}

func (cc *c) Inc(value int64) {
    cc.counter.Inc(value)
}

func (cc *c) Dec(value int64) {
    cc.counter.Dec(value)
}

func (cc *c) Clear() {
    cc.counter.Clear()
}

func (cc *c) Count() int64 {
    return cc.counter.Count()
}

type h struct {
    hist metrics.Histogram
}

func Hist(name string) *h {
    return &h{metrics.GetOrRegisterHistogram(name, nil, metrics.NewUniformSample(1024))}
}

func (hh *h) Update(value int64) {
    hh.hist.Update(value)
}

// TODO: Add other interface for query

type m struct {
    meter metrics.Meter
}

func Meter(name string) *m {
    return &m{metrics.GetOrRegisterMeter(name, nil)}
}

func (mm *m) Mark(value int64) {
    mm.meter.Mark(value)
}

func (mm *m) Stop() {
    mm.meter.Stop()
}

// TODO: Add Timer
