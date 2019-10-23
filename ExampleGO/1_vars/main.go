package main

var v0 int = 100

const c0 uint8 = 200
const (
	c1 = iota
	c2
	_
	c4
)

func main() {
	const c5 = true
	var v1 float32 = 10
	var v2 = 10
	var v3, v4 int = 1, 2
	var v5 int
	var v6 string
	v6 = "х"
	var (
		v7 = 1.1
		v8 bool
	)
	v9 := "строка"
	v10, v5 := "1\n", 3
	v11 := '\xFF'
	v12 := `многострочная
строка\n`

	println(v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v9+v10, v11, v12)
	println(c0, c1, c2, c4, c5)
}
