package config

import (
	"github.com/sirupsen/logrus"
)

func SetupLogrus() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
}
