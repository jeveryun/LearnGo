package main

import "fmt"

func main() {
	m := map[string]string{
		"name": "aaa",
		"sex": "1",
	}
	m2 := make(map[string]int)

	var m3 map[string]int

	fmt.Println(m,m2,m3)

	for k, v := range m {
		fmt.Println(k, v)
	}

	println(len(m))
	// if name, ok := m["name1"]; ok {
	// 	println(name)
	// } else {
	// 	println("key does not exist")
	// }
	// delete(m, "sex")
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }
}