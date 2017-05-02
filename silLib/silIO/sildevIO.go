package silIO

import (
	"github.com/yuin/gopher-lua"
	"../../service"
)

func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), exports)
	L.SetField(mod, "name", lua.LString("SildevTeam's standard IO library"))
	L.Push(mod)
	return 1
}

var exports = map[string]lua.LGFunction{
	"testFunc": testFunc,
	"read":     read,
	"write":    write,
	"reg":      reg,
}

func testFunc(L *lua.LState) int {
	return 0
}
func read(L *lua.LState) int {
	ret, _ := service.GetIns().GetDate()
	L.Push(lua.LString(ret))
	return 1
}

func write(L *lua.LState) int {
	lv := L.Get(-1)
	if str, ok := lv.(lua.LString); ok {
		service.GetIns().PostData(string(str))
	}
	return 1
}

func reg(L *lua.LState) int {
	service.GetIns().RegExecutor()
	return 1
}
