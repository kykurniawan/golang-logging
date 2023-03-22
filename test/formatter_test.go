package test

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestFormatter(t *testing.T) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Warn("Warning")
}
