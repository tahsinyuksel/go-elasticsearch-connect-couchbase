package metric

import (
	"github.com/Trendyol/go-dcp-client/helpers"
	"github.com/Trendyol/go-elasticsearch-connect-couchbase/elasticsearch/bulk"
	"github.com/prometheus/client_golang/prometheus"
)

type Collector struct {
	bulk *bulk.Bulk

	processLatency            *prometheus.Desc
	bulkRequestProcessLatency *prometheus.Desc
}

func (s *Collector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(s, ch)
}

func (s *Collector) Collect(ch chan<- prometheus.Metric) {
	bulkMetric := s.bulk.GetMetric()

	ch <- prometheus.MustNewConstMetric(
		s.processLatency,
		prometheus.GaugeValue,
		float64(bulkMetric.ProcessLatencyMs),
		[]string{}...,
	)

	ch <- prometheus.MustNewConstMetric(
		s.bulkRequestProcessLatency,
		prometheus.GaugeValue,
		float64(bulkMetric.BulkRequestProcessLatencyMs),
		[]string{}...,
	)
}

func NewMetricCollector(bulk *bulk.Bulk) *Collector {
	return &Collector{
		bulk: bulk,

		processLatency: prometheus.NewDesc(
			prometheus.BuildFQName(helpers.Name, "elasticsearch_connector_latency_ms", "current"),
			"Elasticsearch connector latency ms",
			[]string{},
			nil,
		),

		bulkRequestProcessLatency: prometheus.NewDesc(
			prometheus.BuildFQName(helpers.Name, "elasticsearch_connector_bulk_request_process_latency_ms", "current"),
			"Elasticsearch connector bulk request process latency ms",
			[]string{},
			nil,
		),
	}
}
