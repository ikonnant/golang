package main

import "fmt"

func main() {
	v1 := 1
	v2 := &v1
	*v2 = 10

	fmt.Println(v1)

	s1 := []int{0, 0, 0}
	fmt.Println(s1)
	Sl(s1)
	fmt.Println(s1)

	Point(&v1)
	fmt.Println(v1)
}

func Sl(item []int) {
	item[1] = 2
}

func Point(a *int) {
	*a++
}
