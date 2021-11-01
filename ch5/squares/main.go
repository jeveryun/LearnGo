package main

import "fmt"

func suqares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := suqares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}