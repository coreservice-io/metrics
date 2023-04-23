## Introduction
We can use metrics to measure the performance for code performance, api calling, and other statistics data. The metrics service decouple the metrics calling and persistent layer.

### Persistent service:
    - [x] InfluxDB

### InfluxDB Version: 2.X+

### Metrics kinds
- Gauge
  - A gauge is a metric that represents a single numerical value that can arbitrarily go up and down. Gauges are typically used for measured values like temperatures or current memory usage, but also "counts" that can go up and down, like the number of concurrent requests.
- Counter
  - A counter is a cumulative metric that represents a single monotonically increasing counter whose value can only increase or be reset to zero on restart. For example, you can use a counter to represent the number of requests served, tasks completed, or errors.
- Hist
  - A histogram samples observations (usually things like request durations or response sizes) and counts them in configurable buckets. It also provides a sum of all observed values.
- Meter
  - A Meter measures event occurrences count and rate:

## Examples
```go
    NewMetrics(
        10 * time.Minute(),
        metrics.InfluxDB,
        metrics.InfluxDBConfig{},
    )

    metrics.Gauge("api.query.buckets").Update(1)

    g := metrics.Gauge("api.query.buckets")
    g.Update(1)
    g.Update(2)

```