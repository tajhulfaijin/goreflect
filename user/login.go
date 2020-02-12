package user

import (
	"fmt"
	"goreflect/core"
)

type Login struct {
	*core.Application
}

func (x *Login) Init() {
	fmt.Println("Login.Init().CALLED")
	x.Application = &core.Application{}
	x.Application.Name = "Login"
	fmt.Println("x.Application.Name", x.Application.Name)
}

func (x *Login) BeforeProcess() {
	fmt.Println("Login.BeforeProcess")
}

func (x *Login) Run() {
	fmt.Println("Login.RUn")
}

// func init() {
// 	fmt.Println("Login.INIT_CALLED")

// 	// core.Register("Login", Login{})
// }
