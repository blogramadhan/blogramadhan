+++
title = "Python Dasar 07: List dan Tuple"
date = 2026-10-05T08:00:00+07:00
draft = false
description = "Bagian ketujuh seri Belajar Python Dasar. Menyimpan banyak nilai dalam satu wadah: mengambil isinya, memotong sebagian, menambah, membuang, mengurutkan, dan kapan sebaiknya memakai tuple."
tags = ["python", "tutorial", "belajar-python-dasar", "pemula", "list", "tuple"]
categories = ["Teknis"]
featured_image = "/images/python-07-hero.svg"
toc = true
+++

Ini bagian ketujuh dari [seri Belajar Python Dasar](/tutorials/belajar-python-dasar/).

Sejak [Bagian 06](/tutorials/python-dasar-06-perulangan/) kita sudah memakai kurung siku
tanpa penjelasan. Saya bilang waktu itu: nanti dibahas tuntas di Bagian 07.

Sekarang waktunya.

## Satu wadah untuk banyak nilai

Sampai Bagian 05, tiap nilai butuh variabel sendiri. Lima unit kerja berarti lima variabel,
dan itu melelahkan.

List menyimpan semuanya dalam satu nama:

```python
unit = ["Dinas Pendidikan", "Dinas Kesehatan", "Dinas PU"]
print(unit)
print(len(unit))
```

```
['Dinas Pendidikan', 'Dinas Kesehatan', 'Dinas PU']
3
```

Kurung siku untuk membuatnya, koma untuk memisahkan isinya. `len()` memberi tahu ada berapa
isinya.

Isinya tidak harus teks. Boleh angka, boleh campur, walaupun campur-campur biasanya tanda
datanya belum rapi.

## Mengambil isinya, dan hitungan yang mulai dari nol

Untuk mengambil satu isi, sebut nomornya di dalam kurung siku:

```python
print(unit[0])
print(unit[1])
print(unit[2])
```

```
Dinas Pendidikan
Dinas Kesehatan
Dinas PU
```

Perhatikan nomornya. Isi pertama bernomor **0**, bukan 1.

Ini bikin bingung hampir semua orang di awal, termasuk saya dulu. Anggap saja nomor itu bukan
"urutan ke berapa", melainkan "berapa langkah dari awal". Isi pertama berjarak nol langkah
dari awal.

Akibatnya, daftar berisi tiga nilai punya nomor 0, 1, dan 2. Nomor 3 tidak ada:

```python
print(unit[3])
```

```
IndexError: list index out of range
```

Nomor terakhir selalu `len()` dikurangi satu.

## Menghitung dari belakang

Kalau yang Anda butuhkan isi terakhir, tidak perlu menghitung dulu. Pakai angka negatif:

```python
print(unit[-1])
print(unit[-2])
```

```
Dinas PU
Dinas Kesehatan
```

`-1` artinya satu langkah dari ujung belakang. Ini sering lebih praktis daripada
`unit[len(unit) - 1]`.

## Mengambil sebagian

Dua nomor dipisah titik dua akan mengambil sepotong:

```python
print(unit[0:2])
print(unit[:2])
print(unit[1:])
```

```
['Dinas Pendidikan', 'Dinas Kesehatan']
['Dinas Pendidikan', 'Dinas Kesehatan']
['Dinas Kesehatan', 'Dinas PU']
```

Nomor yang pertama ikut diambil, nomor yang kedua tidak. Jadi `[0:2]` berarti nomor 0 dan 1
saja.

Kedengarannya aneh, tapi ada enaknya: `[0:2]` selalu menghasilkan tepat dua isi. Angka
belakang dikurangi angka depan sama dengan jumlah yang Anda dapat.

Kalau angka depannya dikosongkan, berarti dari awal. Kalau angka belakangnya dikosongkan,
berarti sampai habis.

## Mengubah, menambah, membuang

Isi list bisa diganti dengan menyebut nomornya:

```python
unit[1] = "Dinas Kesehatan dan KB"
print(unit)
```

```
['Dinas Pendidikan', 'Dinas Kesehatan dan KB', 'Dinas PU']
```

Menambah di belakang pakai `append`. Menambah di posisi tertentu pakai `insert`:

```python
unit.append("Dinas Sosial")
unit.insert(0, "Sekretariat Daerah")
print(unit)
```

```
['Sekretariat Daerah', 'Dinas Pendidikan', 'Dinas Kesehatan dan KB', 'Dinas PU', 'Dinas Sosial']
```

Membuang bisa dua cara. `remove` membuang berdasarkan isinya, `pop` membuang yang terakhir
sambil menyerahkannya ke Anda:

```python
unit.remove("Dinas PU")
terakhir = unit.pop()
print(terakhir)
print(unit)
```

```
Dinas Sosial
['Sekretariat Daerah', 'Dinas Pendidikan', 'Dinas Kesehatan dan KB']
```

Kalau yang dibuang tidak ada di dalam daftar, `remove` protes:

```
ValueError: list.remove(x): x not in list
```

## Memeriksa dan menghitung

`in` memeriksa apakah sesuatu ada di dalam daftar. Hasilnya `True` atau `False`, seperti
syarat di Bagian 05:

```python
print("Sekretariat Daerah" in unit)
print("Dinas Perhubungan" in unit)
```

```
True
False
```

Untuk daftar berisi angka, ada beberapa perintah bawaan yang langsung berguna:

```python
nilai = [820, 150, 640, 310]
print(sum(nilai))
print(max(nilai))
print(min(nilai))
print(sum(nilai) / len(nilai))
```

```
1920
820
150
480.0
```

Empat baris itu menggantikan beberapa kolom rumus di Excel.

## Mengurutkan

```python
nilai = [820, 150, 640, 310]
nilai.sort()
print(nilai)

nilai.sort(reverse=True)
print(nilai)
```

```
[150, 310, 640, 820]
[820, 640, 310, 150]
```

`sort()` mengubah daftar aslinya. Kalau Anda masih butuh urutan aslinya, pakai `sorted()`
yang membuat daftar baru dan membiarkan yang lama:

```python
asli = [820, 150, 640, 310]
urut = sorted(asli)
print(asli)
print(urut)
```

```
[820, 150, 640, 310]
[150, 310, 640, 820]
```

Bedanya halus tapi penting. `sort()` merapikan di tempat. `sorted()` membuat salinan yang
rapi.

### Mengurutkan teks tidak sesederhana kelihatannya

Coba urutkan nama unit kerja:

```python
unit = ["Dinas PU", "Dinas Pendidikan", "Dinas Kesehatan"]
print(sorted(unit))
```

```
['Dinas Kesehatan', 'Dinas PU', 'Dinas Pendidikan']
```

Lihat urutannya. `Dinas PU` mendahului `Dinas Pendidikan`. Secara alfabet itu jelas salah.

Penyebabnya, Python membandingkan teks lewat kode angka tiap hurufnya, dan semua huruf besar
punya kode lebih kecil daripada huruf kecil. Huruf `U` bernilai 85, huruf `e` bernilai 101.
Jadi di mata Python, `PU` memang lebih kecil daripada `Pendidikan`.

Perbaikannya, suruh Python membandingkan versi huruf kecilnya saja:

```python
print(sorted(unit, key=str.lower))
```

```
['Dinas Kesehatan', 'Dinas Pendidikan', 'Dinas PU']
```

`key=str.lower` artinya: khusus untuk membandingkan, anggap semuanya huruf kecil. Isi aslinya
tidak ikut diubah.

Kalau daftar Anda berisi nama orang atau nama instansi, hampir selalu pakai ini. Saya pernah
menyerahkan daftar yang urutannya kacau gara-gara lupa, dan yang menegur bukan Python,
melainkan atasan saya.

## Kesalahan yang tidak memberi pesan error

Ini jebakan yang hampir pasti akan Anda temui.

```python
nilai = [820, 150, 640]
nilai = nilai.sort()
print(nilai)
```

```
None
```

Daftarnya hilang. Diganti `None`.

Penyebabnya, `sort()` merapikan daftar di tempat dan tidak menyerahkan apa-apa kembali. Waktu
Anda menulis `nilai = nilai.sort()`, yang disimpan ke `nilai` adalah "tidak ada apa-apa" itu.

Aturannya: `sort()` ditulis sendirian, tanpa tanda sama dengan.

```python
nilai.sort()          # benar
nilai = sorted(nilai) # juga benar
nilai = nilai.sort()  # salah, dan tidak ada pesan error
```

Python tidak protes karena secara aturan bahasa tidak ada yang salah. Errornya baru muncul
beberapa baris kemudian, waktu Anda mencoba memakai `nilai` yang sudah telanjur kosong.

## tuple: daftar yang dikunci

Tuple itu seperti list, tapi memakai kurung biasa dan isinya tidak bisa diubah setelah dibuat.

```python
titik = (-0.0263, 109.3425)
print(titik[0])
print(titik[1])
```

```
-0.0263
109.3425
```

Mengambil isinya sama persis dengan list. Yang beda, mencoba mengubahnya akan ditolak:

```python
titik[0] = 1
```

```
TypeError: 'tuple' object does not support item assignment
```

Kenapa ada wadah yang sengaja dibikin tidak bisa diubah? Karena sebagian data memang tidak
seharusnya berubah. Koordinat sebuah titik, pasangan tahun dan bulan, satu baris data yang
sudah final.

Memakai tuple untuk data seperti itu bikin program Anda gagal cepat kalau ada kode lain yang
tidak sengaja mengubahnya. Lebih baik ditolak sekarang daripada ketahuan di laporan bulan
depan.

Untuk sehari-hari, Anda akan jauh lebih sering memakai list.

## Dua daftar sejajar, dan kenapa itu rapuh

Gabungan list dan perulangan dari Bagian 06 sudah cukup untuk banyak pekerjaan:

```python
unit = ["Dinas Pendidikan", "Dinas Kesehatan", "Dinas PU"]
realisasi = [82, 45, 91]

for i in range(len(unit)):
    print(unit[i], "->", realisasi[i], "%")
```

```
Dinas Pendidikan -> 82 %
Dinas Kesehatan -> 45 %
Dinas PU -> 91 %
```

Ini jalan, tapi perhatikan betapa rapuhnya. Kedua daftar harus sama panjang dan urutannya
harus sama persis. Kalau suatu hari ada yang menyisipkan satu unit di tengah tapi lupa
menyisipkan realisasinya juga, hasilnya kacau tanpa ada pesan error apa pun.

Pasangan nama dan nilai seperti ini punya wadah yang lebih cocok. Itu isi Bagian 08.

## Latihan 1: merapikan daftar unit kerja

```python
unit = ["Dinas PU", "Dinas Pendidikan", "Dinas Kesehatan"]

unit.append("Dinas Sosial")
unit.sort(key=str.lower)

print("Jumlah unit:", len(unit))
print("Urutan pertama:", unit[0])
print("Urutan terakhir:", unit[-1])
print("Ada Dinas Pendidikan?", "Dinas Pendidikan" in unit)

for nama in unit:
    print("-", nama)
```

```
Jumlah unit: 4
Urutan pertama: Dinas Kesehatan
Urutan terakhir: Dinas Sosial
Ada Dinas Pendidikan? True
- Dinas Kesehatan
- Dinas Pendidikan
- Dinas PU
- Dinas Sosial
```

Jalankan, lalu ubah daftarnya dengan unit kerja di tempat Anda sendiri. Coba juga hapus
`key=str.lower` dan lihat urutannya berubah.

## Latihan 2: tiga paket terbesar

```python
paket = [820, 150, 640, 310, 975, 420]

print("Jumlah paket :", len(paket))
print("Total nilai  :", sum(paket))
print("Tertinggi    :", max(paket))
print("Terendah     :", min(paket))
print("Rata-rata    :", sum(paket) / len(paket))

tiga = sorted(paket, reverse=True)[:3]
print("Tiga terbesar:", tiga)
```

```
Jumlah paket : 6
Total nilai  : 3315
Tertinggi    : 975
Terendah     : 150
Rata-rata    : 552.5
Tiga terbesar: [975, 820, 640]
```

Baris `tiga` itu menggabungkan dua hal yang baru dipelajari: urutkan dari besar ke kecil,
lalu ambil tiga dari depan. Dibaca dari kiri ke kanan seperti kalimat biasa.

## Kalau tidak jalan

**`IndexError: list index out of range`**

Nomor yang Anda minta melebihi isi daftar. Ingat, daftar berisi tiga nilai nomornya cuma 0, 1,
dan 2. Cek dengan `print(len(daftar))`.

**`ValueError: list.remove(x): x not in list`**

Yang mau dibuang tidak ada di dalam daftar. Sering karena beda spasi atau beda huruf besar
kecil. Teknik `repr()` dari Bagian 04 berguna di sini.

**`TypeError: 'tuple' object does not support item assignment`**

Anda mencoba mengubah isi tuple. Kalau datanya memang perlu berubah, pakai list, ganti kurung
biasa jadi kurung siku.

**`TypeError: '<' not supported between instances of 'str' and 'int'`**

Daftarnya berisi campuran teks dan angka, lalu Anda menyuruhnya mengurutkan. Python tidak tahu
cara membandingkan `"820"` dengan `640`. Samakan dulu tipenya.

**Urutannya aneh, yang berhuruf besar didahulukan**

Bukan error, tapi hasilnya salah. Pakai `sort(key=str.lower)` atau `sorted(daftar, key=str.lower)`.

**Daftarnya tiba-tiba jadi `None`**

Anda menulis `daftar = daftar.sort()`. Hapus bagian `daftar = ` di depannya.

## Berikutnya

Di **Bagian 08** kita masuk ke *dictionary*, wadah yang menyimpan pasangan nama dan nilai
sekaligus. Itu jawaban untuk masalah dua daftar sejajar tadi, dan wadah yang paling sering
saya pakai untuk data pekerjaan sungguhan.

Tautannya muncul di [halaman silabus](/tutorials/belajar-python-dasar/) begitu terbit.
