package main

import (
	"fmt"
	"time"
)

// Contoh Penggunaan Goroutine Sederhana

// Fungsing Says untuk menampilkan string sebanyak 5 kali
func Says(str string) {
	for i := 0; i < 5; i++ {
		// Delay 100 ms
		time.Sleep(100 * time.Millisecond)
		fmt.Println(str)
	}
}

func main() {
	fmt.Println("Hello World")
	go Says("world")
	Says("Hello")
	/*
		Hello World
		world
		Hello
		Hello
		world
		world
		Hello
		Hello
		world
		Hello
	*/
}
