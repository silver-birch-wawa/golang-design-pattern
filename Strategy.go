package main

import "fmt"

/*出行方式*/
type Itravel interface {
	Travel()
}

/*飞机*/
type Aircraft struct{}

/*火车*/
type Train struct{}

/*走路*/
type Walk struct{}

/*具体策略类 1：飞机出行*/
func (this Aircraft) Travel() {
	fmt.Println("飞机出行")
}

/*具体策略类 2：火车出行*/
func (this Train) Travel() {
	fmt.Println("火车出行")
}

/*具体策略类 3：走路出行*/
func (this Walk) Travel() {
	fmt.Println("走路出行")
}

/*环境类*/
type User struct {
	Name    string
	Itravel Itravel
}

func (this User) travel() {
	fmt.Printf("%s", this.Name)
	this.Itravel.Travel()
}

func main() {
	user := User{"张三", Aircraft{}}
	user.travel()

	user = User{"李四", new(Train)}
	user.travel()

	user = User{"王五", &Walk{}}
	user.travel()
}
