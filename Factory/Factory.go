package main

import "fmt"

type KeyBoard interface {
	Input(string)
}

type HPKeyBoard struct {
}

func (board *HPKeyBoard) Input(s string) {
	fmt.Println("HP:", s)
}

type DELLKeyBoard struct {
}

func (board *DELLKeyBoard) Input(s string) {
	fmt.Println("DELL:", s)
}

type HPKeyBoardFactory struct {
}

func (board *HPKeyBoardFactory) GetInstance() KeyBoard {
	return new(HPKeyBoard)
}

type DELLKeyBoardFactory struct {
}

func (board *DELLKeyBoardFactory) GetInstance() KeyBoard {
	return new(DELLKeyBoard)
}

func main() {
	f := new(DELLKeyBoardFactory)
	k := f.GetInstance()
	k.Input("Enter")
}
