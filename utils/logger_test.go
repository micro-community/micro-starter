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
	SetServiceName("log-example2")
	Info(nil, "hello logger!")
	Debug(nil, "debug")
	SetLevel(logrus.DebugLevel)
	Debug(nil, "debug")
	UseJSONFormatter()
	Info(nil, "hello logger!")
	Infof(nil, "test format:%s", "hello logger")
}
