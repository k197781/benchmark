package main

import "fmt"

func main() {
	sum := 0
	for num := 1; num <= 500000000; num++ {
		sum += num
	}
	fmt.Println(sum)
}
