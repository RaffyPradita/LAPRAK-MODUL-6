package main

import "fmt"

func main4() {
	var basis, pangkat int
	// Menerima masukan dua bilangan bulat positif
	fmt.Scan(&basis, &pangkat)

	hasil := 1

	// Perulangan untuk mengalikan basis sebanyak nilai pangkat
	for i := 0; i < pangkat; i++ {
		hasil *= basis
	}

	// Menampilkan hasil pemangkatan
	fmt.Println(hasil)
}

/*PROGRAM HitungPemangkatan

DEKLARASI:
    basis, pangkat, hasil, i : INTEGER

ALGORITMA:
    INPUT (basis, pangkat)

    hasil <- 1

    FOR i <- 1 TO pangkat DO
        hasil <- hasil * basis
    ENDFOR

    OUTPUT (hasil)
ENDPROGRAM*/
