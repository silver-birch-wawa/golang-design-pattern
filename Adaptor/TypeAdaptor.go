package main

import (
	"fmt"
)

type Oldinterface interface {
	show()
}
type Oldclass struct {
	Name string
}

func (this *Oldclass) show() {
	fmt.Println(this.Name)
}

type Targetinterface interface {
	Newshow()
}
type Newclass struct {
	Oldinterface
}

func (this *Newclass) Newshow() {
	this.show()
	fmt.Println("OldInterface already adapt !")
	// 旧struct里有的变量只能通过旧interface继承的的旧函数方法搞定
	// 在新的struct中木有旧struct的变量(因为interface里木有变量只有方法)
}
func main() {
	var Target Targetinterface
	Target = &Newclass{&Oldclass{Name: "OldClass"}}
	Target.Newshow()
}
