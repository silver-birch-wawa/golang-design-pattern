package main

import (
	"fmt"
)

type IMediator interface {
	Communicate(string, ICountry)
}

type Mediator struct {
	cn ChinaCountry
	us AmericaCountry
}

func (m *Mediator) Communicate(msg string, country ICountry) {
	// 中国发msg米国收，否则相反
	if country.GetName() == "China" {
		m.us.ReceiveMessage(msg)
	} else {
		m.cn.ReceiveMessage(msg)
	}
}

type ICountry interface {
	Country(string, IMediator)
	GetName() string
	SendMessage(string)
	ReceiveMessage(string)
}

type ChinaCountry struct {
	ICountry
	CountryName string
	mediator    IMediator
}

func (i *ChinaCountry) Country(name string, mediator IMediator) {
	i.mediator = mediator
	i.CountryName = name
}

func (i *ChinaCountry) GetName() string {
	return i.CountryName
}

func (i *ChinaCountry) SendMessage(msg string) {
	i.mediator.Communicate(msg, i)
}
func (i *ChinaCountry) ReceiveMessage(msg string) {
	fmt.Println("cn received.")
}

type AmericaCountry struct {
	ICountry
	CountryName string
	mediator    IMediator
}

func (i *AmericaCountry) Country(name string, mediator IMediator) {
	i.mediator = mediator
	i.CountryName = name
}

func (i *AmericaCountry) GetName() string {
	return i.CountryName
}
func (i *AmericaCountry) SendMessage(msg string) {
	i.mediator.Communicate(msg, i)
}
func (i *AmericaCountry) ReceiveMessage(msg string) {
	fmt.Println("cn received.")
}
func main() {
	//定义的时候，接口实现的是指针接收，所以用指针初始化
	m := new(Mediator)
	cn := new(ChinaCountry)
	us := new(AmericaCountry)
	cn.Country("China", m)
	us.Country("America", m)
	cn.SendMessage("Hello")
}
