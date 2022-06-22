package main

import (
	"log"
	"time"

	"github.com/docktermj/go-xyzzy-helpers/logger"
)

// Values updated via "go install -ldflags" parameters.

var programName string = "unknown"
var buildVersion string = "0.0.0"
var buildIteration string = "0"

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const MessageIdFormat = "senzing-9999%04d"

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func complexProcess() string {
	time.Sleep(1000 * time.Second)
	return "slept"
}

// ----------------------------------------------------------------------------
// Main
// ----------------------------------------------------------------------------

func main() {

	// Configure the "log" standard library.

	log.SetFlags(log.Llongfile | log.Ldate | log.Lmicroseconds | log.LUTC)

	// Configure the logger. If not configured, no functions will print.

	logger.SetLevel(logger.LevelInfo)

	// Log a message.

	logger.LogMessage(MessageIdFormat, 2000, "Test message", "Variable1", "Variable2")

	// The following demonstrates the low-level calls for
	// Trace, Debug, Info, Warn, and Error.

	log.Println("Test Trace")
	logger.Trace("trace prints")
	logger.Tracef("trace A: %s B: %s C: %d", "aaa", "bbb", 35)

	log.Println("Test Debug")
	logger.Debug("debug prints")
	logger.Debugf("debug A: %s B: %s C: %d", "aaa", "bbb", 35)

	log.Println("Test Info")
	logger.Info("info prints")
	logger.Infof("info A: %s B: %s C: %d", "aaa", "bbb", 35)

	log.Println("Test Warn")
	logger.Warn("warn prints")
	logger.Warnf("warn A: %s B: %s C: %d", "aaa", "bbb", 35)

	log.Println("Test Error")
	logger.Error("error prints")
	logger.Errorf("error A: %s B: %s C: %d", "aaa", "bbb", 35)

	// Avoid long running logging when appropriate.

	if logger.IsDebug() {
		logger.Debugf("%s", complexProcess())
	}

	// Note:  the first Fatal or Panic issued will exit the program.

	log.Println("Test Fatal")
	//	logger.Fatal("fatal prints")
	//	logger.Fatalf("fatal A: %s B: %s C: %d", "aaa", "bbb", 35)

	log.Println("Test Panic")
	//		logger.Fatal("fatal prints")
	//	logger.Fatalf("fatal A: %s B: %s C: %d", "aaa", "bbb", 35)

	log.Println("Test varadic")
	_, err := time.LoadLocation("bob")
	logger.Info("Should be error: ", err)

	log.Println("End")

}
