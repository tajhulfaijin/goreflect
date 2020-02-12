package main

import (
	"goreflect/core"
	"goreflect/user"
	"reflect"
	"strings"
)

func _registerModule() {
	core.Register("Login", user.Login{})
}

func main() {
	_registerModule()
	module := "Login"
	action := "run"
	mod := reflect.New(core.Modules[module]).Interface()
	if core.IsMethodExists(mod, strings.Title(action)) {
		core.CallMethod(mod, "Init")
		core.CallMethod(mod, "BeforeProcess")
		core.CallMethod(mod, "AfterProcess")
	}
}
