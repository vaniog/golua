package tablelib_test

import (
	"testing"

	"github.com/vaniog/golua/lib"
	"github.com/vaniog/golua/luatesting"
)

func TestTable(t *testing.T) {
	luatesting.RunLuaTestsInDir(t, "lua", lib.LoadAll)
}
