package main

import "fmt"

type Visitor interface {
	VisitConcreteElementA(*ConcreteElementA)
	VisitConcreteElementB(*ConcreteElementB)
}

type ConcreteVisitor1 struct {
}

func (v *ConcreteVisitor1) VisitConcreteElementA(ele *ConcreteElementA) {
	ele.Operation()
}

func (v *ConcreteVisitor1) VisitConcreteElementB(ele *ConcreteElementB) {
	ele.Operation()
}

type Element interface {
	accept(Visitor)
}

type ConcreteElementA struct {
}

func (e *ConcreteElementA) accept(visitor Visitor) {
	visitor.VisitConcreteElementA(e)
}

func (e *ConcreteElementA) Operation() {
	fmt.Println("eleA operation")
}

type ConcreteElementB struct {
}

func (e *ConcreteElementB) accept(visitor Visitor) {
	visitor.VisitConcreteElementB(e)
}

func (e *ConcreteElementB) Operation() {
	fmt.Println("eleB operation")
}

type ObjectStructure struct {
	col []Element
}

func (o *ObjectStructure) accept(visitor Visitor) {
	for _, e := range o.col {
		e.accept(visitor)
	}
}
func (o *ObjectStructure) addElement(ele ...Element) {
	for _, e := range ele {
		o.col = append(o.col, e)
	}
}

func main() {
	os := new(ObjectStructure)
	eleA := new(ConcreteElementA)
	eleB := new(ConcreteElementB)

	os.addElement(eleA, eleB)
	visitor := new(ConcreteVisitor1)
	os.accept(visitor)
}
