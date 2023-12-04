package main

import "fmt"

func main() {

	// m1 := struct {
	// 	a int64
	// 	b int64
	// }{a: 1, b: 2}

	// m2 := struct {
	// 	a int64
	// 	b int64
	// }{a: 1, b: 2}

	// if m1 == m2 {
	// 	fmt.Println(1)
	// }
	var a uint8 = 2
	var b uint8 = 3
	fmt.Print(a - b)
}

func A() (a int) {
	defer func(i int) {
		fmt.Println(a)
		fmt.Println(i)
	}(a)
	a = 1
	return 2
}
