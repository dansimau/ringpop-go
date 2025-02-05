package mocks

import "github.com/uber-common/bark"
import "github.com/stretchr/testify/mock"

import "time"

type StatsReporter struct {
	mock.Mock
}

func (_m *StatsReporter) IncCounter(name string, tags bark.Tags, value int64) {
	_m.Called(name, tags, value)
}
func (_m *StatsReporter) UpdateGauge(name string, tags bark.Tags, value int64) {
	_m.Called(name, tags, value)
}
func (_m *StatsReporter) RecordTimer(name string, tags bark.Tags, d time.Duration) {
	_m.Called(name, tags, d)
}
