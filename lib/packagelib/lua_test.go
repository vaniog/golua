package packagelib_test

import (
	"testing"

	"github.com/vaniog/golua/lib"
	"github.com/vaniog/golua/luatesting"
)

func TestPackageLib(t *testing.T) {
	luatesting.RunLuaTestsInDir(t, "lua", lib.LoadAll)
}
