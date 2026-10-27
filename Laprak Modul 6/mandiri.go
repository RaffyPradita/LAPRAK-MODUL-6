package main

import "fmt"

func main7() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Println(2*i - 1)
	}
}
