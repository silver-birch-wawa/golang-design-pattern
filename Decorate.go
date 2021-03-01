package main

import "fmt"

type Component interface {
	Operation()
}

type ConcreteComponent struct {
}

func (this *ConcreteComponent) Operation() {
	fmt.Println("ConcreteComponent operation")
}

type Decorator struct {
	C Component
}

func (this *Decorator) Operation() {
	this.C.Operation()
}

type ConcreteDecoratorA struct {
	Decorator
}

func (this *ConcreteDecoratorA) Operation() {
	this.Decorator.C.Operation()
	fmt.Println("---ConcreteDecorator---")
}
func main() {
	D := Decorator{&ConcreteComponent{}}
	c1 := &ConcreteDecoratorA{D}
	c1.Operation()
}
