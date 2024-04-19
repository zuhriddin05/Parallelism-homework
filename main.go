package main

import (
	"fmt"
)

func faktorial(n int, ch chan int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	ch <- result
}
func main() {
	ch := make(chan int)
	var son int
	fmt.Printf("Son kiriting: ")
	fmt.Scanf("%d", &son)
	go faktorial(son, ch)
	faktorialResult := <-ch
	fmt.Println("Faktorial:", faktorialResult)
}
