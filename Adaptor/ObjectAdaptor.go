package main

import (
	"fmt"
)

type Oldinterface interface {
	show()
	SetName(string)
}
type Oldclass struct {
	Name string
}

func (this *Oldclass) show() {
	fmt.Println(this.Name)
}
func (this *Oldclass) SetName(name string) {
	this.Name = name
}

type Newclass struct {
	Old Oldinterface
}

func (this *Newclass) NewClass() {
	this.Old = new(Oldclass)
	// 接口木有成员变量只能写函数改值
	this.Old.SetName("Old")
}
func (this *Newclass) NewShow() {
	this.Old.show()
	fmt.Println("OldInterface already adapt !")
}
func main() {
	newinterface := new(Newclass)
	newinterface.NewClass()
	newinterface.NewShow()
}
