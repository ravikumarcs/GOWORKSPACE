package main

import "fmt"

type A struct {
	year int
}

func (a A) greet() {
	fmt.Println("Hello GolangUK", a.year)
}

type B struct {
	A
}

func (b B) greet() {
	fmt.Println("Welcome to GolangUK", b.year)
}

func main() {
	var a A
	a.year = 2016
	a.greet()

	var b B
	b.year = 2016
	b.greet()
}
