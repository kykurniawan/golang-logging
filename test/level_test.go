package test

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()

	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("this is a trace")
	logger.Debug("this is a debug")
	logger.Info("this is an info")
	logger.Warn("this is a warn")
	logger.Error("this is an error")
}
