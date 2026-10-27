package main

import "fmt"

func main5() {
	var n int
	// Menerima input bilangan bulat non-negatif
	fmt.Scan(&n)

	hasil := 1

	// Perulangan untuk menghitung faktorial
	for i := 1; i <= n; i++ {
		hasil *= i
	}

	// Menampilkan hasil faktorial
	fmt.Println(hasil)
}

/*PROGRAM HitungFaktorial

DEKLARASI:
    n, hasil, i : INTEGER

ALGORITMA:
    input (n)

    hasil <- 1

    FOR i <- 1 TO n DO
        hasil <- hasil * i
    ENDFOR

    output (hasil)
endprogram*/
