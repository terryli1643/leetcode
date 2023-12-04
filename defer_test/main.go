package main

import "fmt"

func main() {
	fmt.Println(A())
}

func A() (a int) {
	defer func(i int) {
		fmt.Printf("defer a: %d\n", a)
		fmt.Printf("defer i: %d\n", i)
		a = 4
	}(a)
	a = 1
	return 2
}
