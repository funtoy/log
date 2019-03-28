package log_test

import (
	"testing"
	"log"
)

func TestInit(t *testing.T) {
	r := log.NewWithFile("./logs/platform/debug")
	r.SetRotateByDay()
	r.Debug("debug message.")
	r.Debug("debug message.")
	r.Info("info message.")
	r.Info("info message.")
	r.Warn("warn message.")
	r.Warn("warn message.")
	r.Error("error message.")
	r.Error("error message.")

}
