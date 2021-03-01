package main

import (
	"fmt"
)

type config struct {
	Name string
}

var cfg *config

func initConfig() *config {
	if cfg == nil {
		cfg = &config{"config"}
		return cfg
	} else {
		return cfg
	}
}

func main() {
	f := initConfig()
	fmt.Println(f.Name)
}
