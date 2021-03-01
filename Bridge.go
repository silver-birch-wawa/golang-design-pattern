package main

import "fmt"

type ICoffee interface {
	OrderCoffee(count int)
}

type LargeCoffee struct {
	additives ICoffeeAdditives
}

func (l *LargeCoffee) OrderCoffee(count int) {
	fmt.Print("Large Coffee ")
	l.additives.AddAdditives()
	fmt.Println(count, " cup")
}

type SmallCoffee struct {
	additives ICoffeeAdditives
}

func (s *SmallCoffee) OrderCoffee(count int) {
	fmt.Print("Small Coffee ")
	s.additives.AddAdditives()
	fmt.Println(count, "cup")
}

type ICoffeeAdditives interface {
	AddAdditives()
}

type Milk struct{}

func (m *Milk) AddAdditives() {
	fmt.Print("With Milk ")
}

type Sugar struct{}

func (s *Sugar) AddAdditives() {
	fmt.Print("With Sugar ")
}

func main() {
	largeCoffeeMilk := LargeCoffee{
		additives: &Milk{},
	}
	largeCoffeeMilk.OrderCoffee(1)
	smallCoffeeMilk := SmallCoffee{
		additives: &Milk{},
	}
	smallCoffeeMilk.OrderCoffee(1)
	largeCoffeeSugar := LargeCoffee{
		additives: &Sugar{},
	}
	largeCoffeeSugar.OrderCoffee(1)
	smallCoffeeMSugar := SmallCoffee{
		additives: &Sugar{},
	}
	smallCoffeeMSugar.OrderCoffee(1)
}
