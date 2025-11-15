//go:build !noquotas
// +build !noquotas

package runtimelib_test

import (
	"testing"

	"github.com/vaniog/golua/lib"
	"github.com/vaniog/golua/luatesting"
)

func TestRuntimeLib(t *testing.T) {
	luatesting.RunLuaTestsInDir(t, "lua", lib.LoadAll)
}
