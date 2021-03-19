package farmer

import (
	// TODO
	//
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	OpenLog("TEST")
	Info("这是Info")
	Warn("这是Warn")
	Error("这是Error")
	state1, e := os.Stat("/opt/farmer-log/TEST/TEST-WARN.log")
	Assert(nil == e)
	Assert(!state1.IsDir())
	state2, e := os.Stat("/opt/farmer-log/TEST/TEST-ERROR.log")
	Assert(nil == e)
	Assert(!state2.IsDir())
	CloseLog()
	return
}
