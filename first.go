package main

import "fmt"

func main() {
	Tri1(10)
	Tri2(10)
}

func Tri1(num int) {

	for i := 1; i <= num; i++ {

		for x := num - i; x > 0; x-- {
			fmt.Print(" ")
		}

		for j := 1; j <= i; j++ {
			fmt.Print(" *")
		}
		fmt.Println()
	}
}

func Tri2(num int) {
	for i := num; i > 0; i-- {

		for x := num - i; x > 0; x-- {
			fmt.Print(" ")
		}

		for j := 1; j <= i; j++ {
			fmt.Print(" *")
		}
		fmt.Println()
	}
}
