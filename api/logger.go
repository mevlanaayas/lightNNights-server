package api

import (
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	// logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.SetOutput(os.Stdout)

	logrus.SetLevel(logrus.InfoLevel)
}
