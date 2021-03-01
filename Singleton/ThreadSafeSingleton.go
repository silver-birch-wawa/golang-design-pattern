package main

import (
	"fmt"
	"sync"
	"time"
)

type config struct {
	Name string
}

var cfg *config

var oSingle sync.Once

func initConfig() *config {
	oSingle.Do(
		func() {
			cfg = &config{"config"}

		})
	// cfg = &config{"config"}
	//比较用
	return cfg
}
func main() {
	//fmt.Println(&cfg.Name)
	for i := 0; i < 3; i++ {
		go func() {
			initConfig()
			fmt.Println(&cfg.Name)
			// fmt.Println(&cfg)
			//访问指针的地址木有用，因为指针地址没变，变的是struct对象
		}()
	}
	time.Sleep(time.Second)
}
