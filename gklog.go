package gklog

import (
	"flag"
	"io/ioutil"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"k8s.io/klog"
)

type Logger struct {
	l log.Logger
}

func Init(l log.Logger, verbosity int32) {
	fs := flag.NewFlagSet("gklog", flag.ContinueOnError)
	klog.InitFlags(fs)

	fs.Set("logtostderr", "false")
	fs.Set("alsologtostderr", "false")
	fs.Set("skip_headers", "true")
	v := klog.Level(verbosity)
	fs.Set("v", v.String())

	klog.SetOutput(ioutil.Discard)
	klog.SetOutputBySeverity("INFO", writer{level.Debug(l)})
	klog.SetOutputBySeverity("WARNING", writer{level.Warn(l)})
	klog.SetOutputBySeverity("ERROR", writer{level.Error(l)})
	klog.SetOutputBySeverity("FATAL", writer{level.Error(l)})
}

type writer struct {
	log.Logger
}

func (w writer) Write(p []byte) (n int, err error) {
	return len(p), w.Log("msg", string(p))
}
