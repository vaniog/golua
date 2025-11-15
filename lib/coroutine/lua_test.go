package coroutine_test

import (
	"testing"

	"github.com/vaniog/golua/lib"
	"github.com/vaniog/golua/luatesting"
)

func TestCoroutineLib(t *testing.T) {
	luatesting.RunLuaTestsInDir(t, "lua", lib.LoadAll)
}
