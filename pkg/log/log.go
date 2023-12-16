package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log *logrus.Logger

func init() {
	if Log == nil {
		Log = logrus.New()
		Log.SetReportCaller(true)
		Log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, TimestampFormat: "2006-01-02 15:04:05"})
		Log.SetLevel(logrus.ErrorLevel)
		Log.SetOutput(os.Stdout)
	}
}
