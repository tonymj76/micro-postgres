package log

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetReportCaller(true)
	formatter := &logrus.TextFormatter{
		TimestampFormat:        time.StampMilli, // the "time" field configuratiom
		FullTimestamp:          true,
		DisableLevelTruncation: true, // log level field configuration
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			// this function is required when you want to introduce your custom format.
			// In my case I wanted file and line to look like this `file="engine.go:141`
			// but f.File provides a full path along with the file name.
			// So in `formatFilePath()` function I just trimmet everything before the file name
			// and added a line number in the end
			return "", fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		},
	}
	logrus.SetFormatter(formatter)
}
func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

// //CustomLog _
// func CustomLog() *logrus.Logger {
// 	log := logrus.New()
// 	log.Formatter = &logrus.TextFormatter{
// 		TimestampFormat:,
// 		FullTimestamp:   true,
// 		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
// 			// this function is required when you want to introduce your custom format.
// 			// In my case I wanted file and line to look like this `file="engine.go:141`
// 			// but f.File provides a full path along with the file name.
// 			// So in `formatFilePath()` function I just trimmet everything before the file name
// 			// and added a line number in the end
// 			return "", fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
// 		},
// 	}
// 	return log
// }

// func formatFilePath(path string) string {
// 	arr := strings.Split(path, "/")
// 	return arr[len(arr)-1]
// }

// log := logrus.New()
// 	log.Formatter = &logrus.TextFormatter{
// 		TimestampFormat: time.StampMilli,
// 		FullTimestamp:   true,
// 	}
