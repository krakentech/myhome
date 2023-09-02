package logit

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

var (
	IsDebug              = false
	OutFormat            = "{{time}} {{prefix}}  {{type}} - {{message}}"
	TimeFormat           = "06.02.01-15:04:05"
	Prefix               = ""
	w          io.Writer = os.Stdout
	printLf              = fmt.Fprintln
	isTesting            = false
)

func SetWriter(newWriter io.Writer) {
	w = newWriter
}

func Debug(format string, a ...interface{}) {
	if IsDebug {
		printLine(Prefix, logTypeDebug, fmt.Sprintf(format, a...))
	}
}

func Info(format string, a ...interface{}) {
	printLine(Prefix, logTypeInfo, fmt.Sprintf(format, a...))
}

func Warn(format string, a ...interface{}) {
	printLine(Prefix, logTypeWarn, fmt.Sprintf(format, a...))
}

func Err(format string, a ...interface{}) {
	printLine(Prefix, logTypeError, fmt.Sprintf(format, a...))
}

func printLine(prefix string, lType logType, msg string) {
	out := OutFormat

	logTime := time.Now().Format(TimeFormat)
	if isTesting {
		logTime = "00.00.00-00:00:00"
	}

	out = strings.Replace(out, "{{time}}", logTime, 1)
	out = strings.Replace(out, "{{prefix}}", prefix, 1)
	out = strings.Replace(out, "{{type}}", string(lType), 1)
	out = strings.Replace(out, "{{message}}", msg, 1)

	_, err := printLf(w, out)
	if err != nil {
		fmt.Printf("failed to print line: %s", err.Error())
	}
}
