package lib

import (
	"github.com/vaniog/golua/lib/base"
	"github.com/vaniog/golua/lib/coroutine"
	"github.com/vaniog/golua/lib/debuglib"
	"github.com/vaniog/golua/lib/golib"
	"github.com/vaniog/golua/lib/iolib"
	"github.com/vaniog/golua/lib/mathlib"
	"github.com/vaniog/golua/lib/oslib"
	"github.com/vaniog/golua/lib/packagelib"
	"github.com/vaniog/golua/lib/runtimelib"
	"github.com/vaniog/golua/lib/stringlib"
	"github.com/vaniog/golua/lib/tablelib"
	"github.com/vaniog/golua/lib/utf8lib"
	rt "github.com/vaniog/golua/runtime"
)

func LoadLibs(r *rt.Runtime, loaders ...packagelib.Loader) func() {
	var cleanups []func()
	for _, loader := range loaders {
		cleanup := loader.Run(r)
		if cleanup != nil {
			cleanups = append(cleanups, cleanup)
		}
	}
	return func() {
		for i := len(cleanups) - 1; i >= 0; i-- {
			cleanups[i]()
		}
	}
}

func LoadAll(r *rt.Runtime) func() {
	return LoadLibs(
		r,
		base.LibLoader,
		packagelib.LibLoader,
		coroutine.LibLoader,
		stringlib.LibLoader,
		tablelib.LibLoader,
		mathlib.LibLoader,
		iolib.LibLoader,
		utf8lib.LibLoader,
		oslib.LibLoader,
		debuglib.LibLoader,
		golib.LibLoader,
		runtimelib.LibLoader,
	)
}
