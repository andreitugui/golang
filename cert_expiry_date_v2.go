package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"ioutil"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	certExpiryDateMetrics = []string{"certExpiryDateLabel", "certExpiryDate"}
)

// CertExpiryDateCollector asd
type CertExpiryDateCollector struct {
	CertExpiryDateMetrics map[string]*prometheus.GaugeVec
	logFileParseFailures  prometheus.Counter
}

// CertMetricValues asd
type CertMetricValues struct {
	CertExpiryDateMetric float64
	Label                string
}

// Collect asd
func (certExpiryDateCollector *CertExpiryDateCollector) Collect(metrics chan<- prometheus.Metric) {
	logger.Infoln("Start collecting tablespace availability metrics")
	startTime := time.Now()
	certExpiryDateCollector.reset()
	certExpiryDateCollector.scrape()
	PrometheusCollectorDuration.WithLabelValues("CertExpiryDate Collector", "scrape").Observe(float64(time.Since(startTime)) / float64(time.Millisecond))
	startTime = time.Now()
	metrics <- certExpiryDateCollector.logFileParseFailures
	for _, certExpiryDateMetric := range certExpiryDateCollector.CertExpiryDateMetrics {
		certExpiryDateMetric.Collect(metrics)
	}
	PrometheusCollectorDuration.WithLabelValues("certExpiryDate Collector", "collect_metrics").Observe(float64(time.Since(startTime)) / float64(time.Millisecond))
	logger.Infoln("End collecting tablespace availability metrics")
}

// Describe asd
func (certExpiryDateCollector *CertExpiryDateCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- certExpiryDateCollector.logFileParseFailures.Desc()
	for _, certExpiryDateMetric := range certExpiryDateCollector.CertExpiryDateMetrics {
		certExpiryDateMetric.Describe(ch)
	}
}

func (certExpiryDateCollector *CertExpiryDateCollector) reset() {
	for _, certExpiryDateMetric := range certExpiryDateCollector.CertExpiryDateMetrics {
		certExpiryDateMetric.Reset()
	}
}

func (certExpiryDateCollector *CertExpiryDateCollector) scrape() {
	var certificateFilePath = "/u01/app/oracle/product/12.1.0/dbhome_12102/network/admin/dbwallet/$(hostname -f)"
	info, err := os.Stat(certificateFilePath)
	if os.IsNotExist(err) {
		logger.Errorf("File %s does not exist: %s", certificateFilePath, err)
		certExpiryDateCollector.logFileParseFailures.Inc()
	}
	if info != nil {
		content, err := ioutil.ReadFile(certificateFilePath)
		if err != nil {
			logger.Errorf("File %s cannot be opened: %s", certificateFilePath, err)
		}
		bytes := []byte(string(content))

		block, _ := pem.Decode(bytes)
		var cert *x509.Certificate
		cert, _ = x509.ParseCertificate(block.Bytes)
		expiryDate := cert.NotAfter
		certExpiryDateMetric := float64(expiryDate.Unix())
		metrics := make(map[float64]CertMetricValues)
		metrics[certExpiryDateMetric] = CertMetricValues{
			CertExpiryDateMetric: certExpiryDateMetric,
			Label:                "ca",
		}

		for _, metric := range metrics {
			certExpiryDateCollector.CertExpiryDateMetrics[metric.CertExpiryDateMetric].WithLabelValues(metric.CertExpiryDateMetric).Set(metric.Label)
		}
	}
}

// NewCertExpiryDateCollector sda
func NewCertExpiryDateCollector() *CertExpiryDateCollector {
	certExpiryDateCollector := &CertExpiryDateCollector{
		logFileParseFailures: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: "oracle_certificate",
			Name:      "type",
			Help:      "Number of errors while parsing Log file.",
		}),
	}
	certExpiryDateCollector.CertExpiryDateMetrics = make(map[string]*prometheus.GaugeVec)
	for _, certExpiryDateMetric := range certExpiryDateMetrics {
		certExpiryDateCollector.CertExpiryDateMetrics[certExpiryDateMetric] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: fmt.Sprintf("certExpiryDate_metrics_%s, certExpiryDateMetric"),
			Name:      fmt.Sprintf("certExpiryDate_metrics_%s", certExpiryDateMetric),
			Help:      fmt.Sprintf("CertExpiryDate Metrics %s", certExpiryDateMetric),
		}, []string{"Schema_Name"})
	}

	return certExpiryDateCollector
}
