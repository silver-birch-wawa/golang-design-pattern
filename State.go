package main

import "fmt"

type Account struct {
	State       ActionState
	HealthValue int
}

type ActionState interface {
	View()
	Comment()
	Post()
}

type CloseState struct {
}

func (c *CloseState) View() {
	fmt.Println("账号被封，无法看帖")
}

func (c *CloseState) Comment() {
	fmt.Println("抱歉，你的健康值小于-10，不能评论")
}
func (c *CloseState) Post() {
	fmt.Println("抱歉，你的健康值小于0，不能发帖")
}

type RestrictedState struct {
}

func (r *RestrictedState) View() {
	fmt.Println("正常看帖")
}

func (r *RestrictedState) Comment() {
	fmt.Println("正常评论")
}
func (r *RestrictedState) Post() {
	fmt.Println("抱歉，你的健康值小于0，不能发帖")
}

func (a *Account) changeState() {
	if a.HealthValue <= -10 {
		a.State = &CloseState{}
	} else if a.HealthValue > -10 && a.HealthValue <= 0 {
		a.State = &RestrictedState{}
	} // else if a.HealthValue > 0 {
	// 	a.State = &NormalState{}
	// }
}

///给账户设定健康值
func (a *Account) SetHealth(value int) {
	a.HealthValue = value
	a.changeState()
}

func main() {
	acc := new(Account)
	acc.SetHealth(-20)
	acc.State.View()
}
