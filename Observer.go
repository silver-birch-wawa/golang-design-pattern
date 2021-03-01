package main

import (
	"fmt"
)

type Subject interface {
	Attach(Observer)
	Detach(Observer)
	NotifyAll()
}

type ConcreteSubject struct {
	Observers []Observer
}

func (this *ConcreteSubject) Attach(observer Observer) {
	this.Observers = append(this.Observers, observer)
}
func (this *ConcreteSubject) Detach(observer Observer) {
	for index, obs := range this.Observers {
		if obs == observer {
			this.Observers = append(this.Observers[:index], this.Observers[index+1:]...)
		}
	}
}
func (this *ConcreteSubject) NotifyAll() {
	for _, value := range this.Observers {
		value.Update("close")
	}
}

type Observer interface {
	Update(string)
}

type ConcreteObserver struct {
}

func (this *ConcreteObserver) Update(str string) {
	fmt.Println("receive notify from subject:", str)
}

func main() {
	Newspaper := &ConcreteSubject{}
	Newspaper.Observers = make([]Observer, 0)
	Reader := &ConcreteObserver{}
	Newspaper.Attach(Reader)
	Newspaper.NotifyAll()
	Newspaper.Detach(Reader)
	Newspaper.NotifyAll()
}
