package main

import "fmt"

func addInt(in1 int, in2 int) int {
	return in1 + in2
}

func calculation(in1 int, in2 int) (add int, mul int) {
	add = in1 + in2
	mul = in1 * in2
	return
}

func main() {
	var a, b, c, d, e int
	var arrayTest [10]int

	a = 5
	b = 3
	c = addInt(a, b)
	d, e = calculation(a, b)
	fmt.Println("C : ", c)
	fmt.Println("add : ", d)
	fmt.Println("mul : ", e)

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			arrayTest[i] = 10
		} else {
			arrayTest[i] = i
		}
		fmt.Println(arrayTest[i])
	}
}
