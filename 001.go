package main

import "fmt"

func main() {
	sum := 0
	n := 1000
	for i := 0; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
