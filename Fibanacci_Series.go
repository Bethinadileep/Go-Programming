package main

import "fmt"

func main() {
	a := 0
	b := 1
	c := b
	d := 10
	fmt.Println(a)
	fmt.Println(b)
	for d >= 1 {
		c = a + b
		fmt.Println(c)
		a = b
		b = c
		d = d - 1
	}

}
