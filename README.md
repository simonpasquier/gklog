# gklog

The goal is similar to [klog-gokit](https://github.com/simonpasquier/klog-gokit) but with a different approach. Rather than replacing the klog's public methods (eg `Info()`, `Error`, ...), this library leverages `klog.SetOutputBySeverity()` which passes log data to an `io.Writer` instance.

Check the [example program](https://github.com/simonpasquier/gklog/blob/master/examples/main.go) for usage.

Note that it isn't fully functional as messages can be logged several times (see https://github.com/kubernetes/klog/issues/23 for more details).
