package log

import (
	"testing"

	"github.com/op/go-logging"
)

func TestLogInfo(t *testing.T) {

	GetLog(logging.INFO).Infof("This is an info message called from %s", "TestLogInfo")
}
