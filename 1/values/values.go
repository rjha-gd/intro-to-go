package main

import "fmt"

func main() {
	// integers
	var a int = 10
	var b = -20
	c := 0
	c = a + b
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)

	var d int8 = 100
	var e int16 = -100
	var f int32 = 1025
	var g int64 = 1234567890123456789
	var h uint64 = 2 << 62
	fmt.Println(d, e, f, g, h)

	// floating point
	const i float32 = 3.1416
	var j float64 = 1234567890.123456
	var k = 2.71828
	fmt.Printf("%v, %v, %v\n", i, j, k)
	fmt.Printf("%T\n", k)

	// string
	var l = "Good morning!"
	m := "おはようございます！"
	fmt.Printf("'%s' in Japanese is '%s'.\n", l, m)

	// bool
	var n bool = true
	var o = false
	p := n && o
	fmt.Println(p)
}
