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

type KeyBoardFactory struct {
}

func (board *KeyBoardFactory) GetInstance(Type int) KeyBoard {
	if Type == 0 {
		return new(HPKeyBoard)
	} else {
		return new(DELLKeyBoard)
	}
}
func main() {
	var k KeyBoard
	f := new(KeyBoardFactory)
	k = f.GetInstance(1)
	k.Input("Enter")
}
