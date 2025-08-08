## TIPE DATA NUMBER
### INTEGER
* int8 = -(128) - 127
* int16 = -32768 - 32767
* int32 = -(2147483648) - 2147483648
* int64 = -(9223372036854775808) - 9223372036854775808
* uint8 = 0 - 255
* uint16 = 0 - 65535
* uint32 = 0 - 4294967295
* uint64 = 0- 18446744073709551615

### FLOATING POINt
* float32 = 1.18 X 10^-38 s/d 3,4 X10^38
* float64 = 2.23 X 10^-304 s/d 1.180 X 10^308
* complex64
* complex128

### Alias tipe data
* byte => unit8
* rune => int32
* int => minimal int32
* uint => minimal uin32

## TIPE DATA BOOLEAN
* true
* false
## TIPE DATA STRING
* String ada tipe data kumpulan karakter
* Tipe data String direpresentasikan dengan kata kunci string
* Nilai data string di go-lang selalu diawali dengan karakter " dan diakhiri "
### Function untuk String
* len("string)
* "string"[number] (mengambil 1 karakter dari string)

## Variable
* Variable adalah tempat menyimpan data
* Di Go-Lang Variable hanya bisa menyimpan tipe data yang sama, jika kita ingin menyimpan data yang berbeda-beda jenis, kita harus membuat beberapa variable
* Untuk membuat variable, kita bisa menggunakan kata kunci **var**, lalu diikuti dengan nama variable dan tipe datanya

## CONSTANT
* Constant adalah tempat menyimpan data yang tidak bisa diubah
* Di Go-Lang Constant hanya bisa menyimpan tipe data yang sama, jika kita ingin menyimpan data yang berbeda-beda jenis, kita harus membuat beberapa constant
* Untuk membuat constant, kita bisa menggunakan kata kunci **const**, lalu diikuti dengan nama constant dan tipe datanya

## KONVERSI TIPE DATA
* Di Go-Lang ada tipe data yang bisa diubah menjadi tipe data yang lain, contohnya seperti int ke string

## TYPE DECLARATION
* Type declaration digunakan untuk membuat ulang tipe data baru dari tipe data yang sudah ada
* Type Declaration biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada, dengan tujuan agar lebih mudah dimengerti

## OPERATOR MATEMATIKA
+, -, *, /, %

## Augment Assignment
* +=
* -=
* *=
* /=
* %=

## Unary Operator
* ++ = increment
* -- = decrement
* \+ = Positif
* \- = Negatif
* ! = digunakan untuk membalikkan nilai boolean

## Operator Perbandingan
* ==
* !=
* >
* <
* >=
* <=

## Operator Boolean
* &&
* ||
* !

## Index Program Array
* Index Program Array dimulai dari 0

## Function Array
* len(array)
* array[index]
* array[index] = value

## Slice
* Slice adalah potongan dari array
* Slice mirip dengan array, namun slice bisa berubah
* Slice dan array selalu terkoneksi, dimana length tidak boleh lebih dari capacity
### Membuat Slice dari Array
* array[low:high] = membuat slice dari array dimulai index low sampai index sebelum high
* array[low:] = membuat slice dari array dimulai index low sampai index di akhir array
* array[:high] = membuat slice dari array dimulai index 0 sampai index sebelum high
* array[:] = membuat slice dari array dimulai index 0 sampai index di akhir array

## Function Slice
* len(slice)
* cap(slice)
* append(slice, value)
* make([]Type, length, capacity)
* copy(dest, src)
* slice[index]
* slice[index] = value

## TIPE DATA MAP
* Map adalah tipe data lain yang berisikan kumpulan data yang sama, namun bisa menentukan jenis tipe data index yang akan digunakan.
* Map adalah tipe data kumpulan key-value dimana kata kuncinya bersifat unik
* Jumlah data yang dimasukkan ke dalam map tidak terbatas, dengan syarat kunci harus unik, jika menggunakan kunci yang sama, maka data yang lama akan diubah oleh data baru

## Function Map
* len(map)
* map[key]
* map[key] = value
* make(map[TypeKey]TypeValue)
* delete(map, key)

## DEFER, PANIC, RECOVER
* Defer digunakan untuk menjalankan kode setelah function selesai di eksekusi
* Panic digunakan untuk menghentikan eksekusi program, tapi defer akan tetap dijalankan
* Recover digunakan untuk menjalankan kode jika terjadi panic
* ini merupakan konsep try catch di bahasa pemrograman GOLANG

## STRUCT
* Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
* Struct biasanya representasi data dalam program aplikasi yang kita buat
* Data di struct disimpan dalam field
* Sederhananya, struct adalah kumpulan field

## INTERFACE
* Interface adalah tipe data Abstract, tidak memiliki implementasi langsung
* Sebuah Interface berisikan definisi-definisi method
* Biasanya interface digunakan sebagai kontrak

## INTERFACE KOSONG
* Biasanya dalam pemrograman berorientasi objek, ada satu data parent dipuncak yang bisa dianggap sebagai semua implementasi data yang ada dibahasa pemrograman tersebut
* Interface kosong adalah interface yang tidak memiliki deklarasi method satupun, hal ini membuat secara otomatis semua tipe data akan menjadi implementasinya
* Interface kosong juga memiliki type alias bernama any

## NIL
* Nil adalah nilai kosong
* Nil memiliki tipe data tertentu, yaitu interface{}, nil memiliki nilai default bernilai nil
* Nill sendiri hanya bisa digunakan dibeberapa tipe data, seperti interface, function, map, slice, pointer dan channel

## POINTER
* Pointer adalah kemampuan membuat reference ke alamat memori yang sama, tanpa menduplikasi data yang sudah ada

## OPERATOR *
* Semua variable yang mengacu ke alamat memori yang sama akan memiliki nilai yang sama

## OPERATOR new
* Sebelumnya kita menggunakan & untuk membuat pointer, namun ada cara lain untuk membuat pointer, yaitu dengan menggunakan keyword new