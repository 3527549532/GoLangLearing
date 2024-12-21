package main

import (
	"Demo/p1"
	"fmt"
)

func main() {
	f1()
	f2()
	p1.F3()
}

func f1() {
	fmt.Println("调用方法f1")
}
