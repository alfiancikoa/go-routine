package main

import "fmt"

/*
	Berikut merupakan contoh penerapan goroutine yang dalam kasusnya untuk mencari jumlah
	dari satu deret array.
	proses pengerjaannya dibagi dua yaitu:
	- goroutine pertama menghitung dari angka pertama sampai jumlah angka bagi dua
	- goroutine kedua menghitung dari angka tengah sampai akhir
*/

func Sum(s []int, c chan int) {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum = sum + s[i]
	}
	c <- sum // kirim hasil jumlah ke channel c
}

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6}
	c := make(chan int) // channel c
	// Pembagian perhitungan dibagi dua kali
	go Sum(arr[:len(arr)/2], c)
	go Sum(arr[len(arr)/2:], c)
	// setelah perhitungan di fungsi goroutine, hasilnya akan dikembalikan ke channel c
	// kemudian ditampung di variable x dan y
	x, y := <-c, <-c
	// tampilan hasil perhitungn
	fmt.Println(x, y, x+y)
	/*
		x = 15
		y = 6
		total x+y = 21
	*/
}
