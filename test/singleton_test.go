package test

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSingleton(t *testing.T) {
	logrus.Info("Info")
	logrus.Error("Error")
	logrus.Warn("Warn")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("Hello")
}
