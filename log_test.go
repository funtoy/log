package zapLog_test

import (
	"github.com/funtoy/zapLog"
	"testing"
)

func TestNewLogger(t *testing.T) {
	zapLog.Debug("hello main Debug")
	zapLog.Infof("hello main Info %v", "hehehe")
	zapLog.Info("hello main Info2")
	zapLog.Warn("Hi Gateway Im Debug")

	//var a []int
	zapLog.Errorf("Hi Gateway  Im Info")
	zapLog.Info("Hi Gateway  Im Info")

}
