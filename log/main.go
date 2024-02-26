package log

import (
	"fmt"
)

const (
	MaxLogSize      = 10
	VisibleLogLines = 4
)

var (
	LogLines []string
	LogReady bool = false
)

// Write new log line to display buffer
func Logf(format string, a ...any) {
	if !LogReady {
		return
	}

	formatted := fmt.Sprintf(format, a...)

	LogLines = append(LogLines, formatted)
	for len(LogLines) > MaxLogSize {
		LogLines = LogLines[1:]
	}
}

func InitLog() {
	LogLines = make([]string, 0)
	LogReady = true
}
