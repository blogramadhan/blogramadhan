+++
title = "Python Dasar 04: String, Merapikan Teks"
date = 2026-10-02T18:00:00+07:00
draft = false
description = "Bagian keempat seri Belajar Python Dasar. Menggabung, memotong, mengubah huruf, dan membuang spasi berlebih — lalu menyeragamkan nama unit kerja yang ditulis lima orang dengan lima gaya berbeda."
tags = ["python", "tutorial", "belajar-python-dasar", "pemula", "string", "teks"]
categories = ["Teknis"]
featured_image = "/images/python-04-hero.svg"
toc = true
+++

Ini bagian keempat dari [seri Belajar Python Dasar](/tutorials/belajar-python-dasar/). Di
[Bagian 03](/tutorials/python-dasar-03-variabel-dan-tipe-data/) kita berkenalan dengan `str`
sebagai salah satu dari empat tipe. Hari ini kita menggarapnya.

Kenapa teks dapat jatah satu bagian sendiri? Karena hampir semua data yang Anda temui di
pekerjaan datang sebagai teks, dan hampir semuanya berantakan. Nama yang ditulis dengan huruf
besar semua, spasi nyasar di ujung, nama unit yang berbeda di tiap berkas padahal maksudnya
sama.

Di akhir tulisan ini ada latihan yang mungkin pernah Anda kerjakan manual.

## Tanda kutip, dan masalah tanda petik satu

Di Python, teks boleh diapit kutip ganda atau kutip tunggal. Keduanya sama saja.

```python
print("Kantor Gubernur Kalbar")
print('Kantor Gubernur Kalbar')
```

Pilihannya jadi penting kalau teks Anda mengandung tanda kutip.

```python
print('Dia bilang "halo"')
```

```
Dia bilang "halo"
```

Kutip tunggal di luar, kutip ganda di dalam. Kalau dibalik, Python akan mengira teksnya
selesai di kutip kedua lalu bingung dengan sisanya.

Hal yang sama kena pada apostrof, dan ini sering muncul di nama-nama kita.

```python
print('Pondok Al-Qur'an')
```

```
SyntaxError: unterminated string literal
```

Apostrof di tengah kata dianggap kutip penutup. Pakai kutip ganda, beres:

```python
print("Pondok Al-Qur'an")
```

```
Pondok Al-Qur'an
```

Untuk teks beberapa baris, pakai tiga kutip.

```python
alamat = """Jalan Ahmad Yani
Pontianak"""
print(alamat)
```

```
Jalan Ahmad Yani
Pontianak
```

Kebiasaan saya: pakai kutip ganda untuk semuanya, kecuali kalau di dalamnya ada kutip ganda.
Yang penting konsisten.

## Menggabung teks dengan f-string

Di Bagian 03 kita menyambung teks dengan `+`, dan harus membungkus angka dengan `str()` dulu.
Ada cara yang jauh lebih enak.

```python
nama = "Dinas Pendidikan"
jumlah = 45

print(f"{nama} punya {jumlah} pegawai")
```

```
Dinas Pendidikan punya 45 pegawai
```

Huruf `f` sebelum tanda kutip itu kuncinya. Dengan `f` di depan, apa pun yang Anda tulis di
dalam kurung kurawal akan diganti dengan isinya. Angka tidak perlu diubah jadi teks lebih
dulu — Python mengurusnya sendiri.

Di dalam kurung kurawal boleh ada hitungan, dan boleh diatur jumlah angka di belakang koma.

```python
anggaran = 12.5
print(f"Per pegawai: {anggaran / jumlah:.3f} miliar")
```

```
Per pegawai: 0.278 miliar
```

Tanda `:.3f` artinya tampilkan sebagai desimal dengan tiga angka di belakang koma. Ini
mengatasi masalah `0.30000000000000004` yang kita temui di Bagian 03.

Mulai sekarang, pakai f-string. Cara `+` tetap jalan, tapi lebih panjang dan lebih mudah
salah.

## Mengubah besar-kecil huruf

Empat perintah yang paling sering dipakai.

```python
teks = "dinas pendidikan"

print(teks.upper())
print(teks.title())
print(teks.capitalize())
print("DINAS PENDIDIKAN".lower())
```

```
DINAS PENDIDIKAN
Dinas Pendidikan
Dinas pendidikan
dinas pendidikan
```

`title()` mengawali tiap kata dengan huruf besar. `capitalize()` cuma huruf pertama kalimat.

Sekarang bagian yang menjebak banyak pemula. Coba lihat isi `teks` setelah semua itu
dijalankan:

```python
print(teks)
```

```
dinas pendidikan
```

Masih huruf kecil semua. **Perintah-perintah ini tidak mengubah teks aslinya.** Mereka
membuat teks baru dan menyerahkannya kepada Anda. Kalau hasilnya ingin disimpan, simpan:

```python
teks = teks.upper()
```

Ini berlaku untuk semua perintah teks di tulisan ini. Teks di Python bersifat tetap; yang
bisa Anda lakukan adalah membuat versi barunya.

## Membuang spasi yang tidak kelihatan

Spasi di ujung teks adalah sumber bug yang paling menyebalkan, karena tidak terlihat di layar.

```python
kotor = "  Dinas Pendidikan  "
print(kotor.strip())
```

`strip()` membuang spasi di awal dan akhir. Ada juga `lstrip()` untuk kiri saja dan
`rstrip()` untuk kanan saja, tapi `strip()` yang paling sering dipakai.

Satu hal yang `strip()` **tidak** urus: spasi berlebih di tengah.

```python
ganda = "Dinas   Pendidikan"
print(ganda.strip())
```

```
Dinas   Pendidikan
```

Tetap tiga spasi di tengah. Untuk yang ini triknya memecah lalu menyatukan lagi — kita bahas
dua seksi lagi.

Kalau ingin melihat spasi yang tersembunyi, bungkus dengan `repr()`:

```python
print(repr("  Dinas Pendidikan  "))
```

```
'  Dinas Pendidikan  '
```

Sekarang spasinya terlihat. Perintah ini berguna waktu Anda yakin dua teks sama tapi Python
bersikeras berbeda.

## Mengambil sebagian teks

Tiap huruf punya nomor urut, dimulai dari **nol**.

```python
unit = "Dinas Pendidikan"

print(len(unit))
print(unit[0])
print(unit[-1])
```

```
16
D
n
```

`len()` menghitung jumlah karakter. `unit[0]` mengambil huruf pertama. Nomor minus menghitung
dari belakang, jadi `unit[-1]` huruf terakhir.

Untuk mengambil sepotong, pakai titik dua.

```python
print(unit[0:5])
print(unit[6:])
```

```
Dinas
Pendidikan
```

Bacanya: dari nomor 0 sampai **sebelum** nomor 5. Batas kanannya tidak ikut. Ini sering
bikin salah hitung di awal, dan lama-lama terbiasa.

Kalau batasnya dikosongkan, berarti dari awal atau sampai akhir.

## Mencari dan mengganti

```python
unit = "Dinas Pendidikan"

print(unit.replace("Dinas", "Din."))
print("Pendidikan" in unit)
print(unit.startswith("Dinas"))
```

```
Din. Pendidikan
True
True
```

`replace()` mengganti semua kemunculan. `in` memeriksa apakah sepotong teks ada di dalamnya,
dan hasilnya `bool` — tipe yang kita kenal di Bagian 03. Ada juga `endswith()` yang memeriksa
bagian akhir.

Kalau Anda butuh posisinya, bukan sekadar ada atau tidak:

```python
print(unit.find("Pendidikan"))
print(unit.find("Kesehatan"))
```

```
6
-1
```

`-1` artinya tidak ketemu.

## Memecah dan menyatukan

Dua perintah ini pasangan, dan keduanya akan sering Anda pakai nanti waktu membaca berkas
CSV di Bagian 14.

```python
baris = "Pontianak,Singkawang,Sintang"
print(baris.split(","))
```

```
['Pontianak', 'Singkawang', 'Sintang']
```

`split()` memotong teks di tiap tanda yang Anda sebut, hasilnya daftar. Kurung siku itu
*list*, dan kita bahas tuntas di Bagian 07.

Kebalikannya `join()`:

```python
kota = ["Pontianak", "Singkawang", "Sintang"]
print(" - ".join(kota))
```

```
Pontianak - Singkawang - Sintang
```

Perhatikan urutan penulisannya, karena ini sering tertukar: yang di depan adalah **pemisahnya**,
dan daftarnya ditaruh di dalam kurung.

Sekarang trik yang saya janjikan tadi. Kalau `split()` dipanggil **tanpa tanda apa pun di
dalamnya**, ia memotong di setiap kelompok spasi dan membuang yang kosong.

```python
ganda = "Dinas   Pendidikan"
print(" ".join(ganda.split()))
```

```
Dinas Pendidikan
```

Satu baris itu menyelesaikan spasi di awal, di akhir, dan berapa pun jumlahnya di tengah.

## Latihan: lima gaya, satu bentuk

Ini situasi yang dijanjikan di Bagian 03. Lima orang mengisi formulir, dan nama unit yang
sama ditulis dengan lima cara:

```python
daftar = [
    "  dinas pendidikan  ",
    "DINAS PENDIDIKAN",
    "Dinas  Pendidikan",
    "dinas Pendidikan ",
    "Dinas Pendidikan",
]
```

Lima-limanya maksudnya satu unit yang sama. Tapi bagi komputer, kelimanya berbeda — jadi
kalau Anda merekap jumlah per unit, hasilnya lima baris, bukan satu.

Perapiannya cukup satu baris:

```python
nama = "  dinas   PENDIDIKAN "
rapi = " ".join(nama.split()).title()
print(rapi)
```

```
Dinas Pendidikan
```

Bacanya dari dalam ke luar. `nama.split()` memecah di tiap kelompok spasi dan membuang yang
kosong. `" ".join(...)` menyatukannya kembali dengan satu spasi. `.title()` menyeragamkan
huruf besarnya.

**Tugas Anda:** jalankan baris itu untuk kelima isi `daftar` di atas, satu per satu, dan
pastikan kelimanya menghasilkan `Dinas Pendidikan` yang sama persis.

Untuk sekarang salin-tempel lima kali tidak apa-apa. Di Bagian 06 kita belajar perulangan,
dan pekerjaan ini selesai dalam tiga baris berapa pun banyaknya data.

**Satu peringatan soal `title()`.** Coba ini:

```python
print("dinas pendidikan dan kebudayaan".title())
```

```
Dinas Pendidikan Dan Kebudayaan
```

Kata *dan* ikut berhuruf besar, padahal menurut kaidah bahasa Indonesia seharusnya tidak.
Python tidak tahu kata sambung. Untuk nama unit yang mengandung *dan*, *di*, atau *untuk*,
Anda masih perlu memperbaikinya sendiri — dan setelah Bagian 08 kita punya cara yang lebih
rapi untuk itu.

## Kalau tidak jalan

**`AttributeError: 'int' object has no attribute 'upper'`**

Anda memakai perintah teks pada angka. Perintah seperti `upper()` dan `strip()` cuma ada di
`str`. Cek dulu dengan `type()`.

**`AttributeError: 'str' object has no attribute 'Upper'. Did you mean: 'upper'?`**

Huruf besar-kecilnya salah. Semua perintah teks ditulis huruf kecil. Python bahkan menebakkan
maksud Anda di pesan errornya.

**`TypeError: 'str' object does not support item assignment`**

Anda mencoba mengganti satu huruf dengan `teks[0] = "X"`. Tidak bisa — teks bersifat tetap.
Buat teks baru, misalnya dengan `replace()`.

**`IndexError: string index out of range`**

Nomor urutnya melewati panjang teks. Ingat penomorannya mulai dari nol, jadi huruf terakhir
dari teks 16 karakter ada di nomor 15, bukan 16.

**Teksnya sudah dirapikan tapi tidak berubah**

Hasilnya tidak disimpan. `teks.strip()` saja tidak cukup; tulis `teks = teks.strip()`.

## Berikutnya

Di **Bagian 05** kita mengajari program mengambil keputusan dengan `if`, `elif`, dan `else`.
Di situ `bool` yang kita temui di Bagian 03 dan perbandingan seperti `in` tadi mulai benar-benar
terpakai. Latihannya menandai baris yang nilainya melewati batas.

Tautannya muncul di [halaman silabus](/tutorials/belajar-python-dasar/) begitu terbit.
