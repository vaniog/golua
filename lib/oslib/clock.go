//go:build !windows
// +build !windows

package oslib

import (
	// "syscall"

	"time"

	rt "github.com/arnodel/golua/runtime"
)

func clock(t *rt.Thread, c *rt.GoCont) (rt.Cont, error) {
	// var rusage syscall.Rusage
	// _ = syscall.Getrusage(syscall.RUSAGE_SELF, &rusage) // ignore errors
	// time := float64(rusage.Utime.Sec+rusage.Stime.Sec) + float64(rusage.Utime.Usec+rusage.Stime.Usec)/1000000.0
	return c.PushingNext1(t.Runtime, rt.FloatValue(float64(time.Now().UnixNano())/1000000.)), nil
}
