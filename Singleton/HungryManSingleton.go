package main

import (
	"fmt"
)

type config struct {
	Name string
}

var cfg *config = &config{"config"}

func initConfig() *config {
	return cfg
}

func main() {
	f := initConfig()
	fmt.Println(f.Name)
}
