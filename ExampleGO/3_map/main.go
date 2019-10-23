package main

import "fmt"

func main() {
	var m1 map[int]bool
	m2 := map[int]bool{}
	m2[1] = true

	test1 := m2[1]
	test2, ok2 := m2[2]
	_, ok3 := m2[3]

	fmt.Println(m1, m2, test1, test2, ok2, ok3)
}
