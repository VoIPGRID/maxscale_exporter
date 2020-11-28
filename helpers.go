package main

import "github.com/prometheus/client_golang/prometheus"

type Metric struct {
	Desc      *prometheus.Desc
	ValueType prometheus.ValueType
}

func (m *Metric) new(value float64, labelValues ...string) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		m.Desc,
		m.ValueType,
		value,
		labelValues...,
	)
}

func newDesc(subsystem, name, help string, variableLabels []string, t prometheus.ValueType) Metric {
	return Metric{
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, name),
			help, variableLabels, nil,
		), t,
	}
}
