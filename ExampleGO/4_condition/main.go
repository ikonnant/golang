package main

import "fmt"

func main() {
	a := true
	if a {
		fmt.Println("ok a")
	}

	b := 1
	if b == 1 {
		fmt.Println("ok b")
	}

	mm := map[int]string{
		1: "test1",
		2: "test2",
	}
	num := 0

	if val, ok := mm[num]; ok {
		fmt.Println("key", num, "=", val)
	} else {
		fmt.Println("key", num, "не существует")
	}

	s1 := []int{10, 15, 16, 18, 33, 57, 90, 900, 112, 9}
	var val int
	var i int

	for i < 8 {
		if i < 3 {
			i++
			continue
		}

		val = s1[i]
		fmt.Println(i, ":", val)
		i++
	}

	for i := 0; i <= 3; i++ {
		fmt.Println(i)
	}

	for i, val := range s1 {
		fmt.Println(i, ":", val)
	}

	mm[1] = "test1"
	switch mm[1] {
	case "test1":
		fmt.Println("test1")
	case "test11":
		fmt.Println("test11")
		fallthrough
	default:
		fmt.Println("default")
	}

loop:
	for i, val := range s1 {
		fmt.Println(i, ":", val)
		switch i {
		case 3:
			fmt.Println("break loop")
			break loop
		}
	}
}
