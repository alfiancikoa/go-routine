# go-routine
Basic go routine, channel, buffer

<p align=justify>Golang konkurensi atau biasa dikenal sebagai goroutine merupakan <i>lightweight</i> thread yang dikelola oleh runtime Go. Sama halnya dengan thread, hanya saja sifatnya
lebih ringan hanya memerlukan kurang lebih <b>2KB</b> memory saja untuk satu buah goroutine. Eksekusi goroutine bersifat <i>assynchronous</i> yang artinya eksekusi program
tidak saling tunggu dengan goroutine lain dan juga eksekusinya dijalankan di multi core processor sehingga kita dapat menentukan berapa banyak core yang aktif sesuai dengan kebutuhan. Bahkan fungsi utama yaitu <b>main()</b> juga merupakan sebuah goroutine.</p><br>

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

## Channel
<p align=justify>Channels adalah sebuah kanal yang digunakan oleh konkuren goroutine untuk berkomunikasi. Mekanisme yang digunakan adalah serah terima data lewat channel yang sudah disiapkan. Channel dapat difungsikan sebagai pengirim data di sebuah goroutine dan bisa juga difungsikan sebagai penerima di goroutine lainnya. Proses pengiriman dan penerimaan data pada channel bersifat <b><i>blocking</i></b> atau <b><i>Assynchronous</i></b>.</p><br>

Cara Deklarasi dan penggunaan sebuah channel:
```
# Deklarasi
var nama_channel = make(chan type_data)
# contoh
var c = make(chan string)
atau
c := make(chan string)

# Penggunaan
ch <- v // kirim v ke channel ch
v := <- ch // menerima data dari ch dan set value ke v
```
