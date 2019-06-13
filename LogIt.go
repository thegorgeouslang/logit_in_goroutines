// Author: James Mallon <jamesmallondev@gmail.com>
// logit package - lib created to print and write file logs
package logit

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

// Struct type syslog -
type syslog struct {
	file       *os.File
	Filepath   string
	log        *log.Logger
	categories map[string][]string
}

// to be used as an external pointer to the syslog struct type
var Syslog *syslog

// init function - initialize values and processes of the file
func init() {
	lg := syslog{}
	lg.Filepath = fmt.Sprintf("%s%s%s", "logs/", time.Now().Format("2006_01_02"), ".log")
	lg.loadCategories() // loads all categories
	Syslog = &lg
}

// startLog method - changes the default filepath
func (lg *syslog) startLog() {
	lg.file, _ = os.OpenFile(lg.Filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 1444)
	lg.log = log.New(lg.file, "", log.Ldate|log.Ltime)
}

// loadCategories method - loads all categories
func (lg *syslog) loadCategories() {
	lg.categories = map[string][]string{
		"emergency": {"Emergency:", "an emergency"},
		"alert":     {"Alert:", "an alert"},
		"critical":  {"Critical:", "a critical"},
		"error":     {"Error:", "an error"},
		"warning":   {"Warning:", "a warning"},
		"notice":    {"Notice:", "a notice"},
		"info":      {"Info:", "an info"},
		"debug":     {"Debug:", "a debug"},
	}
}

// AppendCategories method - it allow the user to append new categories
func (lg *syslog) AppendCategories(newCategories map[string][]string) {
	for k, v := range newCategories {
		lg.categories[k] = v
	}
}

// WriteLog method - writes the message to the log file
func (lg *syslog) WriteLog(category string, msg string, trace string) {
	lg.startLog()
	val, res := lg.categories[category]
	if !res {
		lg.log.Printf("The log category does not exists: %s", lg.GetTraceMsg())
	} else {
		lg.log.Printf("%s %s on %s", val[0], msg, trace)
	}
	defer lg.file.Close()
}

// GetTraceMsg method - get the full error stack trace
func (lg *syslog) GetTraceMsg() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fmt.Sprintf("%s:%d PID: %d", frame.File, frame.Line, os.Getpid())
}
