package main

import (
	"fmt"
)

// 对象
type Subject interface {
	request(name string)
}

type RealSubject struct {
}

func (this *RealSubject) request(name string) {
	fmt.Println("RealSubject:", name)
}

// 对象的代理
type Proxy struct {
	realSubject *RealSubject
}

func (this *Proxy) request(name string) {
	realSubject := &RealSubject{}
	realSubject.request(name)
	fmt.Println("Proxy:", name)
}

func main() {
	var sub Subject
	sub = &Proxy{}
	sub.request("localhost")
}
