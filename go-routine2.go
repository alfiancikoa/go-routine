package main

import (
	"fmt"
	"time"
)

/*
 	Berikut adalah contoh di mana goroutine tidak dapat dieksekusi
	karena function utama yaitu main telah selesai mengeksekusi program yang ada di dalamnya.
	sehingga function Says belum sempat dieksekusi tapi function mainnya sudah selesai duluan
*/

func Says(str string) {
	for i := 0; i < 5; i++ {
		// Delay 100 ms
		time.Sleep(100 * time.Microsecond)
		fmt.Println(str)
	}
}

func main() {
	go Says("Hello World")
	/*
		Nothing Happen
	*/
}
