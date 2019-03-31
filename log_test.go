package log_test

import (
	"github.com/funtoy/log"
	"testing"
)

func TestInit(t *testing.T) {
	log.Debug("debug message.")
	log.Debug("debug message.")
	log.Info("info message.")
	log.Info("info message.")
	log.Warn("warn message.")
	log.Warn("warn message.")
	log.Error("error message.")
	log.SetNoColor()
	log.SetJson()
	log.Error("error message.")

}
