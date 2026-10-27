package main

import "fmt"

func main9() {
	var a, b int
	fmt.Scan(&a, &b)
	for i := a; i <= b; i++ {
		fmt.Print(i)
		if i < b {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

/*program BarisanBilangan
kamus
    a, b, i : integer

algoritma
    read(a, b)

    for i ← a to b do
        write(i, " ")
    endfor
endprogram*/
