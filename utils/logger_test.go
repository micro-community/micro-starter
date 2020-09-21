package utils

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogrus(t *testing.T) {
	SetServiceName("test")
	Error(nil, "A group of walrus emerges from the ocean")
	UseJSONFormatter()
	Info(nil, "A group of walrus emerges from the ocean")
}

func TestSetting(t *testing.T) {
	SetServiceName("log-example, logrus")
	Info(nil, "hello logger in  logrus")
	Debug(nil, "debug logrus")
	SetLevel(logrus.DebugLevel)
	UseJSONFormatter()
	Infof(nil, "logrus: test format:%s", "hello logger")
}
