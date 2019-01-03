package main

import (
	"os"

	"github.com/go-kit/kit/log"
	"github.com/simonpasquier/gklog"
	"k8s.io/klog"
)

func main() {
	glogger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	glogger = log.With(glogger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	gklog.Init(glogger, 6)

	klog.Info("some information")
	klog.V(7).Info("should not be visible")
	klog.Warning("some warning")
	klog.Error("some error")
}
