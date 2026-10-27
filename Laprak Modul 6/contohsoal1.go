package main

import "fmt"

func main3() {
	var a, b int
	var j int
	fmt.Scan(&a, &b)
	for j = a; j <= b; j += 1 {
		fmt.Print(j, " ")
	}
}
