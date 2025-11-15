package utf8lib_test

import (
	"testing"

	"github.com/vaniog/golua/lib"
	"github.com/vaniog/golua/luatesting"
)

func TestUtf8Lib(t *testing.T) {
	luatesting.RunLuaTestsInDir(t, "lua", lib.LoadAll)
}
