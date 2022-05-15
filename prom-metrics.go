package main

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	DNSLookupMaxMS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_dns_lookup_max_ms",
		Help: "DNS LookUp Max time in ms",
	},
		[]string{"cluster", "url", "application"})

	DNSLookupMinMS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_dns_lookup_min_ms",
		Help: "DNS LookUp Min time in ms",
	},
		[]string{"cluster", "url", "application"})

	DNSLookupMS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_dns_lookup_ms",
		Help: "DNS LookUp time in ms",
	},
		[]string{"cluster", "url", "application"})

	TCPConnectionMaxMS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_tcp_connection_max_ms",
		Help: "TCP Connection Max time in ms",
	},
		[]string{"cluster", "url", "application"})

	TCPConnectionMinMS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_tcp_connection_min_ms",
		Help: "TCP Connection Min time in ms",
	},
		[]string{"cluster", "url", "application"})

	TCPConnectionMS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_tcp_connection_ms",
		Help: "TCP Connection time in ms",
	},
		[]string{"cluster", "url", "application"})

	TLSHandshakeMaxMS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_tls_handshake_max_ms",
		Help: "TLS Handshake Max time in ms",
	},
		[]string{"cluster", "url", "application"})

	TLSHandshakeMinMS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_tls_handshake_min_ms",
		Help: "TLS Handshake Min time in ms",
	},
		[]string{"cluster", "url", "application"})

	TLSHandshakeMS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_tls_handshake_ms",
		Help: "TLS Handshake time in ms",
	},
		[]string{"cluster", "url", "application"})

	serverProcessingMaxMS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_server_processing_max_ms",
		Help: "Server Processing Max time in ms",
	},
		[]string{"cluster", "url", "application"})

	serverProcessingMinMS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_server_processing_min_ms",
		Help: "Server Processing Min time in ms",
	},
		[]string{"cluster", "url", "application"})

	serverProcessingMS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_server_processing_ms",
		Help: "Server Processing time in ms",
	},
		[]string{"cluster", "url", "application"})

	contentTransferMaxMS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_content_transfer_max_ms",
		Help: "Content Transfer Max time in ms",
	},
		[]string{"cluster", "url", "application"})

	contentTransferMinMS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_content_transfer_min_ms",
		Help: "Content Transfer Min time in ms",
	},
		[]string{"cluster", "url", "application"})

	contentTransferMS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_content_transfer_ms",
		Help: "Content Transfer time in ms",
	},
		[]string{"cluster", "url", "application"})

	connectMaxMS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_connect_max_ms",
		Help: "Connect Max time in ms",
	},
		[]string{"cluster", "url", "application"})

	connectMinMS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_connect_min_ms",
		Help: "Connect Min time in ms",
	},
		[]string{"cluster", "url", "application"})

	connectMS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_connect_ms",
		Help: "Connect time in ms",
	},
		[]string{"cluster", "url", "application"})

	nameLookupMaxMS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_name_lookup_max_ms",
		Help: "Name Lookup Max time in ms",
	},
		[]string{"cluster", "url", "application"})

	nameLookupMinMS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_name_lookup_min_ms",
		Help: "Name Lookup Min time in ms",
	},
		[]string{"cluster", "url", "application"})

	nameLookupMS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_name_lookup_ms",
		Help: "Name Lookup time in ms",
	},
		[]string{"cluster", "url", "application"})

	preTransferMaxMS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_pre_transfer_max_ms",
		Help: "Pre Transfer Max time in ms",
	},
		[]string{"cluster", "url", "application"})

	preTransferMinMS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_pre_transfer_min_ms",
		Help: "Pre Transfer Min time in ms",
	},
		[]string{"cluster", "url", "application"})

	preTransferMS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_pre_transfer_ms",
		Help: "Pre Transfer time in ms",
	},
		[]string{"cluster", "url", "application"})

	startTransferMaxMS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_start_transfer_max_ms",
		Help: "Start Transfer Max time in ms",
	},
		[]string{"cluster", "url", "application"})

	startTransferMinMS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_start_transfer_min_ms",
		Help: "Start Transfer Min time in ms",
	},
		[]string{"cluster", "url", "application"})

	startTransferMS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_start_transfer_ms",
		Help: "Start Transfer time in ms",
	},
		[]string{"cluster", "url", "application"})

	totalMaxMS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_total_max_ms",
		Help: "Total Max time in ms",
	},
		[]string{"cluster", "url", "application"})

	totalMinMS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_total_min_ms",
		Help: "Total Min time in ms",
	},
		[]string{"cluster", "url", "application"})

	totalMS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_total_ms",
		Help: "Total time in ms",
	},
		[]string{"cluster", "url", "application"})

	httpResponseStats = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "http_tracer_response_stats",
		Help: "HTTP Response stats",
	},
		[]string{"cluster", "url", "application", "code"})
)

func setPromMetrics(cluster, url, application string, httpStats map[string]map[string]map[string]int64, responseStats map[string]map[int]int64) {
	DNSLookupMaxMS.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application}).Set(float64(httpStats[url]["max"][DNSLookup]))
	DNSLookupMinMS.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application}).Set(float64(httpStats[url]["min"][DNSLookup]))
	DNSLookupMS.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application}).Set(float64(httpStats[url]["cumulative"][DNSLookup]))
	TCPConnectionMaxMS.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application}).Set(float64(httpStats[url]["max"][TCPConnection]))
	TCPConnectionMinMS.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application}).Set(float64(httpStats[url]["min"][TCPConnection]))
	TCPConnectionMS.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application}).Set(float64(httpStats[url]["cumulative"][TCPConnection]))
	TLSHandshakeMaxMS.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application}).Set(float64(httpStats[url]["max"][TLSHandshake]))
	TLSHandshakeMinMS.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application}).Set(float64(httpStats[url]["min"][TLSHandshake]))
	TLSHandshakeMS.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application}).Set(float64(httpStats[url]["cumulative"][TLSHandshake]))
	serverProcessingMaxMS.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application}).Set(float64(httpStats[url]["max"][serverProcessing]))
	serverProcessingMinMS.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application}).Set(float64(httpStats[url]["min"][serverProcessing]))
	serverProcessingMS.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application}).Set(float64(httpStats[url]["cumulative"][serverProcessing]))
	contentTransferMaxMS.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application}).Set(float64(httpStats[url]["max"][contentTransfer]))
	contentTransferMinMS.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application}).Set(float64(httpStats[url]["min"][contentTransfer]))
	contentTransferMS.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application}).Set(float64(httpStats[url]["cumulative"][contentTransfer]))
	connectMaxMS.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application}).Set(float64(httpStats[url]["max"][connect]))
	connectMinMS.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application}).Set(float64(httpStats[url]["min"][connect]))
	connectMS.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application}).Set(float64(httpStats[url]["cumulative"][connect]))
	nameLookupMaxMS.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application}).Set(float64(httpStats[url]["max"][nameLookup]))
	nameLookupMinMS.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application}).Set(float64(httpStats[url]["min"][nameLookup]))
	nameLookupMS.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application}).Set(float64(httpStats[url]["cumulative"][nameLookup]))
	preTransferMaxMS.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application}).Set(float64(httpStats[url]["max"][preTransfer]))
	preTransferMinMS.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application}).Set(float64(httpStats[url]["min"][preTransfer]))
	preTransferMS.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application}).Set(float64(httpStats[url]["cumulative"][preTransfer]))
	startTransferMaxMS.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application}).Set(float64(httpStats[url]["max"][startTransfer]))
	startTransferMinMS.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application}).Set(float64(httpStats[url]["min"][startTransfer]))
	startTransferMS.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application}).Set(float64(httpStats[url]["cumulative"][startTransfer]))
	totalMaxMS.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application}).Set(float64(httpStats[url]["max"][total]))
	totalMinMS.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application}).Set(float64(httpStats[url]["min"][total]))
	totalMS.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application}).Set(float64(httpStats[url]["cumulative"][total]))
	for code := range responseStats[url] {
		httpResponseStats.With(prometheus.Labels{"cluster": cluster, "url": url, "application": application, "code": strconv.Itoa(code)}).Set(float64(responseStats[url][code]))
	}
}
