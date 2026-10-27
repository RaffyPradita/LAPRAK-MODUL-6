package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for k := 1; k <= n; k++ {
		i := n - k + 1
		fmt.Printf("Request %d diproses dengan prioritas %d\n", k, i)
	}
}
