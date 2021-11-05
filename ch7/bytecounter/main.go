package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello1"))
	fmt.Println(c)

	c = 0
	var name = "Dolly111"
	fmt.Fprintf(&c, "Hello123, %s", name)
	fmt.Println(c)
}