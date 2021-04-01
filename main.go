package main

import (
	"fmt"
	"github.com/p9c/parallelcoin/cmd/pod"
	"github.com/p9c/parallelcoin/pkg/interrupt"
	"github.com/p9c/parallelcoin/pkg/limits"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/debug"
	"runtime/trace"
)

func main() {
	// set some runtime parameters to fit the workload
	runtime.GOMAXPROCS(runtime.NumCPU() * 3)
	debug.SetGCPercent(10)
	var e error
	if runtime.GOOS != "darwin" {
		if e = limits.SetLimits(); E.Chk(e) { // todo: doesn't work on non-linux
			_, _ = fmt.Fprintf(os.Stderr, "failed to set limits: %v\n", e)
			os.Exit(1)
		}
	}
	// start up the tracer if it is enabled
	var f *os.File
	if os.Getenv("POD_TRACE") == "on" {
		D.Ln("starting trace")
		tracePath := fmt.Sprintf("%v.trace", fmt.Sprint(os.Args))
		if f, e = os.Create(tracePath); E.Chk(e) {
			E.Ln(
				"tracing env POD_TRACE=on but we can't write to trace file",
				fmt.Sprintf("'%s'", tracePath),
				e,
			)
		} else {
			if e = trace.Start(f); E.Chk(e) {
			} else {
				D.Ln("tracing started")
				defer trace.Stop()
				defer func() {
					if e := f.Close(); E.Chk(e) {
					}
				}()
				interrupt.AddHandler(
					func() {
						D.Ln("stopping trace")
						trace.Stop()
						e := f.Close()
						if e != nil {
						}
					},
				)
			}
		}
	}
	// start the main application
	res := pod.Init()
	D.Ln("returning value", res, os.Args)
	// shut down the tracer if it was started
	if os.Getenv("POD_TRACE") == "on" {
		D.Ln("stopping trace")
		trace.Stop()
		defer func() {
			if e := f.Close(); E.Chk(e) {
			}
		}()
	}
	if res != 0 {
		E.Ln("quitting with error")
		os.Exit(res)
	}
	
}
