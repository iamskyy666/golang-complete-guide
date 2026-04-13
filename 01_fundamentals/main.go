package main

import "fmt"

// custom-logger (enums)
type LogLevel int

const (
	LevelTrace   = iota
	LevelDebug   = 1
	LevelWarning = 2
	LevelInfo    = 3
	LevelError   = 4
)

var levelNames = []string{"Trace 🟢", "Debug 🐞", "Warning ⚠️", "Info ☑️", "Err. 💀"}

func (l LogLevel) String() string {
	if l < LevelTrace || l > LevelError {
		return "UNKNOWN"
	}

	return levelNames[l]
}

func PrintLogLevel(level LogLevel) {
	fmt.Printf("Log_Level: %d %s\n", level, level.String())
}

func main() {
// custom logger
PrintLogLevel(LevelTrace)
PrintLogLevel(LevelWarning)
PrintLogLevel(LevelDebug)
PrintLogLevel(LevelInfo)
PrintLogLevel(LevelError)
}

// O/P:
// $ go run main.go
// Log Level: 0 Trace 🟢
// Log Level: 1 Debug 🐞
// Log Level: 4 Err. 💀
// Log Level: 3 Info ☑️
// Log Level: 2 Warning ⚠️