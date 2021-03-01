package main

import "fmt"

// 包含获取产品以及配置产品的方法
type Builder interface {
	SetName(string)
	GetProduct() Product
}

type Director struct {
	// 负责指导产品的具体的生成过程
}

func (director *Director) Construct(builder Builder) {
	builder.SetName("aaa")
}

type ConcreteBuilder1 struct {
	Builder
	product Product
}

func (this *ConcreteBuilder1) SetName(name string) {
	this.product.name = name
}
func (this *ConcreteBuilder1) GetProduct() Product {
	return this.product
}

type Product struct {
	name string
}

func main() {
	director := new(Director)
	builder1 := new(ConcreteBuilder1)

	director.Construct(builder1)
	product := builder1.GetProduct()
	fmt.Println(product)
}
