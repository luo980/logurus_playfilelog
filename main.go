package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	logrus.SetLevel(logrus.TraceLevel)

	logrus.Trace("trace msg")
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	//logrus.Fatal("fatal msg")
	//logrus.Panic("panic msg")

	logrus.WithFields(logrus.Fields{
		"when": "here",
		"body": "go",
	}).Warn("A walrus appears")

	f, _ := os.OpenFile("/home/luo980/test.log", os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0755)
	logrus.SetOutput(f)

	for true {
		time.Sleep(1 * time.Second)
		println("eeeeee")
		logtest()
	}

	println("Helere")
}

func logtest() {
	logrus.Trace("trace msg")
	time.Sleep(100 * time.Millisecond)
	logrus.Debug("debug msg")
	time.Sleep(200 * time.Millisecond)
	logrus.Info("info msg")
	time.Sleep(100 * time.Millisecond)
	logrus.Warn("warn msg")
	time.Sleep(200 * time.Millisecond)
	logrus.Error("error msg")
}
