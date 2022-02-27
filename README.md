# go-routine
Basic go routine, channel, buffer

<p align=justify>Golang Konkurensi atau biasa dikenal sebagai goroutine lightweight thread yang dikelola oleh runtime Go. Sama halnya dengan thread, hanya saja sifatnya
lebih ringan hanya memerlukan kurang lebih 2 KB memory saja untuk satu buah goroutine. Eksekusi goroutine bersifat assinchronous yang artinya eksekusi program
tidak saling tunggu dengan goroutine lain dan juga eksekusinya dijalankan di multi core processor sehingga kita dapat menentukan berapa banyak core yang aktif sesuai dengan kebutuhan. Bahkan fungsi utama yaitu main() juga merupakan sebuah goroutine.</p><br>

Penggunaan goroutine sangat sederhana, hanya menambahkan keyword ```go``` sebelum memanggil fungsi yang diinginkan. Contoh:
```
go f(x,y,z)
```
Berikut potongan kode yang biasa digunakan:
```
# Untuk menentukan jumlah core yang diaktifkan
runtime.GOMAXPROCS(n)

# delay 100 ms
time.Sleep(100 * time.Millisecond)

# fungsi tunggu
fmt.Scanln(&x)
```

Terimakasih sudah membaca<br>
cheers,<br>
AlfianCikoa
