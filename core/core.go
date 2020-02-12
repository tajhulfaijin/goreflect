package core

import (
	"fmt"
	"reflect"
)

// Modules list all of available modules
var Modules = map[string]reflect.Type{}

type Application struct {
	Name string
}

// Register is used to register module in a module
func Register(name string, iface interface{}) {
	Modules[name] = reflect.TypeOf(iface)
}

func (x *Application) Init() {
	fmt.Println("Application.INIT")
}

func (x *Application) BeforeProcess() {
	fmt.Println("Application.BeforeProcess")
}

func (x *Application) AfterProcess() {
	fmt.Println("Application.BeforeProcess")
}

// IsMethodExists return a method is exists on an interface
func IsMethodExists(value interface{}, name string) bool {
	_, exists := reflect.TypeOf(value).MethodByName(name)
	return exists
}

// CallMethod is used to invoke method on an interface
// Ref : https://gitlab.com/iMil/derpina/-/blob/master/main.go
func CallMethod(value interface{}, name string, args ...interface{}) interface{} {
	var v = reflect.ValueOf(value).MethodByName(name)
	if v.IsValid() {
		param := make([]reflect.Value, len(args))
		for i, _ := range args {
			param[i] = reflect.ValueOf(args[i])
		}
		var r = v.Call(param)
		if len(r) > 0 && r[0].IsValid() {
			return r[0].Interface()
		}
	}
	return nil
}
