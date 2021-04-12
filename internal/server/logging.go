package server

import (
	"github.com/sirupsen/logrus"
)

var logger *logrus.Entry

func setLogger(log *logrus.Entry) {
	logger = log
}
