package main

import(
	"fmt"
)


func main() {
	N := 1000
	a := []int32{}

	fmt.Println("start")
	for i:=0; i < N; i++{
		for j:=0; j < N; j++{
			a = append(a, int32(j + i * N))
		}
	}
	fmt.Println("end")
}
