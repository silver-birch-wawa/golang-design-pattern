package main

import "fmt"

//命令抽象接口
type Command interface {
	Execute()
}

type StartCommand struct {
}

func (c *StartCommand) Execute() {
	fmt.Println("Command call by StartCommand")
}

type CloseCommand struct {
}

func (c *CloseCommand) Execute() {
	fmt.Println("Command call by CloseCommand")
}

//请求结构体
type Box struct {
	button1 Command
	button2 Command
}

//抽象请求对象
func NewBox(button1, button2 Command) *Box {
	return &Box{
		button1: button1,
		button2: button2,
	}
}

//封装命令执行函数
func (b *Box) Pressbutton1() {
	b.button1.Execute()
}

//封装命令执行函数
func (b *Box) Pressbutton2() {
	b.button2.Execute()
}

func main() {
	button1 := new(StartCommand)
	button2 := new(CloseCommand)
	newbox := NewBox(button1, button2)

	newbox.Pressbutton1()
	newbox.Pressbutton2()
}
