package main

import "fmt"

// KeyBoard
type KeyBoard interface {
	Input(string)
}

type HPKeyBoard struct {
}

func (board *HPKeyBoard) Input(s string) {
	fmt.Println("HP keyboard:", s)
}

type DELLKeyBoard struct {
}

func (board *DELLKeyBoard) Input(s string) {
	fmt.Println("DELL KeyBoard:", s)
}

// Monitor
type Monitor interface {
	Input(string)
}

type DELLMonitor struct {
}

func (m *DELLMonitor) Input(s string) {
	fmt.Println("DELL monitor:", s)
}

type HPMonitor struct {
}

func (m *HPMonitor) Input(s string) {
	fmt.Println("HP monitor:", s)
}

type IFactory interface {
	createMonitor()
	createKeyBoard()
}

type DELLFactory struct {
}

func (f *DELLFactory) createMonitor() Monitor {
	return new(DELLMonitor)
}

func (f *DELLFactory) createKeyBoard() KeyBoard {
	return new(DELLKeyBoard)
}

type HPFactory struct {
}

func (f *HPFactory) createMonitor() Monitor {
	return new(HPMonitor)
}

func (f *HPFactory) createKeyBoard() KeyBoard {
	return new(HPKeyBoard)
}

func main() {
	f := new(DELLFactory)
	k := f.createKeyBoard()
	m := f.createMonitor()
	k.Input("Enter")
	m.Input("Enter")

	ff := new(HPFactory)
	kk := ff.createKeyBoard()
	mm := ff.createMonitor()
	kk.Input("Enter")
	mm.Input("Enter")
}
