package test

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "Rizky").Info("Info saja")

	logger.WithFields(logrus.Fields{
		"username": "kykurniawan",
		"name":     "Rizky kurniawan",
	}).Info("Info user")
}
