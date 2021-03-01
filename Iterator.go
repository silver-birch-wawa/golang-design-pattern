package main

import "fmt"

type Ints []int

func (i Ints) First() Iterator {
	return Iterator{
		data:  i,
		index: 0,
	}
}

type Iterator struct {
	data  Ints
	index int
}

func (i Iterator) Get() int {
	return i.data[i.index]
}

func (i Iterator) HasNext() bool {
	return i.index < len(i.data)
}

func (i Iterator) Next() Iterator {
	return Iterator{
		data:  i.data,
		index: i.index + 1,
	}
}

func main() {
	ints := Ints{1, 2, 3}
	for it := ints.First(); it.HasNext(); it = it.Next() {
		fmt.Println(it.Get())
	}
}
