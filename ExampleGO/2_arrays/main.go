package main

import "fmt"

func main() {
	var a1 [3]int
	a2 := [...]int{1, 2, 3}
	var a3 [3][3]int
	a3[0][0] = 10

	fmt.Println(a1, a2, a3)

	var s1 []int
	s2 := append(s1, 1)
	s3 := append(s1, 1)
	s3 = append(s3, 2)
	s3 = append(s3, 3)

	s4 := append(s2, s3...)

	s5 := make([]int, 10, 20)
	s5 = append(s5, []int{1, 2, 3, 4, 5, 6}...)

	s6 := s5
	s6[1] = 10

	fmt.Println(s1, s2, s3, s4, s5, s6, s6[10:])
	fmt.Println(len(s1), len(s2), len(s3), len(s4), len(s5), len(s6))
	fmt.Println(cap(s1), cap(s2), cap(s3), cap(s4), cap(s5), cap(s6))
}
