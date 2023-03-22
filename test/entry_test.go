package test

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	entry := logrus.NewEntry(logger)

	entry.WithField("name", "Rizky Kurniawan")

	entry.Info("Hello")
}
