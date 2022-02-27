# go-routine
Basic go routine, channel, buffer

<p>Golang Konkurensi atau biasa dikenal sebagai goroutine lightweight thread yang dikelola oleh runtime Go. Sama halnya dengan thread, hanya saja sifatnya
lebih ringan hanya memerlukan kurang lebih 2 KB memory saja untuk satu buah goroutine. Eksekusi goroutine bersifat assinchronous yang artinya eksekusi program
tidak saling tunggu dengan goroutine lain dan juga eksekusinya dijalankan di multi core processor sehingga kita dapat menentukan berapa banyak core yang aktif sesuai dengan kebutuhan.</p>
<br>

Penggunaan goroutine sangat sederhana, hanya menambahkan keyword go sebelum memanggil fungsi yang diinginkan. seperti contoh:
```
go f(x,y,z)
```
