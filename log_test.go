package log_test

import (
	"github.com/funtoy/log"
	"testing"
)

func TestNewLogger(t *testing.T) {
	log.Debug("this is a debug log")
	log.Infof("this is a info log %v", "hehehe")
	log.Info("this is a info log too")
	log.Warn("this is a warn log")
	log.Errorf("this is a error log")

}
