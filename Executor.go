package main

import (
	"./silLib/silIO"
	"github.com/yuin/gopher-lua"
	"flag"
	"./service"
)

var Host string
var Port int
var Lua string

func init() {
	flag.StringVar(&Host, "host", "127.0.0.1", "host address of the APIGateway which this executor link up with")
	flag.IntVar(&Port, "port", 9090, "port for APIGateway which this executor link up with")
	flag.StringVar(&Lua, "lua", "/var/script.lua", "lua script's path")
	flag.Parse()

}

func main() {
	service.Host = Host
	service.Port = Port
	L := lua.NewState()
	defer L.Close()
	L.PreloadModule("SilIO", silIO.Loader)
	if err := L.DoFile(Lua); err != nil {
		panic(err)
	}
}
