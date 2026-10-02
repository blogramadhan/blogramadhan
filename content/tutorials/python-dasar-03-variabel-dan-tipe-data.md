+++
title = "Python Dasar 03: Variabel dan Tipe Data"
date = 2026-09-25T15:00:00+07:00
draft = false
description = "Bagian ketiga seri Belajar Python Dasar. Menyimpan nilai dan memberinya nama, empat tipe dasar, dan kenapa 10 + 5 jadi 15 sementara \"10\" + \"5\" jadi 105."
tags = ["python", "tutorial", "belajar-python-dasar", "pemula", "variabel", "tipe-data"]
categories = ["Teknis"]
featured_image = "/images/python-03-hero.svg"
toc = true
+++

Ini bagian ketiga dari [seri Belajar Python Dasar](/tutorials/belajar-python-dasar/). Alat
Anda semestinya sudah siap dari [Bagian 02](/tutorials/python-dasar-02-menyiapkan-alat/) —
entah lewat Colab atau Python di komputer sendiri. Keduanya sama saja untuk hari ini.

Di bagian sebelumnya saya menutup dengan satu janji: tanda kutip itu menentukan, dan kita
bahas kenapa. Hari ini janjinya saya tepati.

## Variabel itu nama untuk sebuah nilai

Sejauh ini kita cuma menampilkan sesuatu lalu melupakannya. Variabel membuat Python
mengingatnya.

```python
nama_unit = "Dinas Pendidikan"
jumlah_pegawai = 42

print(nama_unit)
print(jumlah_pegawai)
```

Hasilnya:

```
Dinas Pendidikan
42
```

Tanda `=` di situ bukan "sama dengan" seperti di pelajaran matematika. Bacanya "simpan ke".
Baris pertama artinya: simpan teks *Dinas Pendidikan* ke dalam nama `nama_unit`.

Karena cuma nama, isinya bisa diganti kapan saja.

```python
jumlah_pegawai = 42
jumlah_pegawai = 45
print(jumlah_pegawai)
```

Yang muncul `45`. Nilai lamanya hilang begitu ditimpa.

Satu hal kecil yang berguna: `print` bisa menerima beberapa nilai sekaligus kalau dipisah
koma, dan dia tidak rewel soal tipenya.

```python
print(nama_unit, jumlah_pegawai)
```

```
Dinas Pendidikan 45
```

## Aturan menamai variabel

Beberapa aturan wajib, sisanya kebiasaan.

Yang wajib:

- Boleh huruf, angka, dan garis bawah — tapi **tidak boleh diawali angka**. `nama2` boleh,
  `2nama` ditolak.
- **Tidak boleh ada spasi.** `jumlah pegawai` tidak bisa; pakai `jumlah_pegawai`.
- **Huruf besar dan kecil dibedakan.** `nama` dan `Nama` dianggap dua variabel berbeda.

Yang sebaiknya:

- Pakai huruf kecil semua, pisahkan kata dengan garis bawah. `jumlah_pegawai`, bukan
  `JumlahPegawai` atau `jmlpgw`.
- Pakai nama yang menjelaskan isinya. Anda akan membaca kode ini lagi bulan depan, dan
  `x` tidak akan memberi tahu apa-apa.
- Jangan pakai nama yang sudah dipakai Python sendiri, seperti `print`, `list`, atau `str`.
  Kodenya jalan, tapi perintah aslinya jadi rusak.

## Empat tipe dasar

Tiap nilai punya jenis. Python menyebutnya tipe, dan empat ini yang paling sering Anda temui.

| Jenis | Namanya di Python | Contoh |
|---|---|---|
| Teks | `str` | `"Pontianak"` |
| Bilangan bulat | `int` | `42` |
| Bilangan desimal | `float` | `12.5` |
| Benar atau salah | `bool` | `True` |

Kalau ragu, tanya langsung ke Python pakai `type()`.

```python
print(type("Pontianak"))
print(type(42))
print(type(12.5))
print(type(True))
```

```
<class 'str'>
<class 'int'>
<class 'float'>
<class 'bool'>
```

Tiga hal yang perlu dicatat sekarang, supaya tidak jadi masalah nanti.

**Desimal pakai titik, bukan koma.** Ini yang paling sering menjebak kita di Indonesia. Di
Python, dua belas setengah ditulis `12.5`. Kalau Anda menulis `12,5`, Python menganggapnya
dua nilai terpisah.

```python
print(3,5)
```

```
3 5
```

Bukan tiga koma lima, tapi tiga dan lima.

**`True` dan `False` diawali huruf besar**, dan tanpa tanda kutip. `"True"` dengan kutip itu
teks biasa, bukan nilai benar.

**Perbandingan menghasilkan `bool`.** Ini yang akan kita pakai banyak di Bagian 05.

```python
print(42 > 40)
print(42 == 40)
```

```
True
False
```

Perhatikan `==` dengan dua sama dengan. Satu `=` artinya menyimpan, dua `==` artinya
membandingkan. Tertukar sedikit, hasilnya beda jauh.

## Kenapa membedakannya penting

Sekarang bagian yang saya janjikan di Bagian 02. Jalankan dua baris ini:

```python
print(10 + 5)
print("10" + "5")
```

```
15
105
```

Baris pertama menjumlahkan dua angka. Baris kedua **menyambung dua teks**, persis seperti
menempelkan dua potong kertas. Tanda `+` yang sama, pekerjaan yang berbeda, karena tipenya
berbeda.

Inilah gunanya tanda kutip. Dengan kutip, `10` itu teks. Tanpa kutip, `10` itu angka.

Sekali Anda paham ini, satu kelompok besar error yang tadinya membingungkan akan langsung
masuk akal.

## Jebakan yang paling sering: angka yang sebenarnya teks

Python bisa meminta masukan dari pengguna dengan `input()`.

```python
umur = input("Umur: ")
print(umur + 5)
```

Ketik `30`, tekan Enter, dan yang muncul:

```
TypeError: can only concatenate str (not "int") to str
```

Padahal Anda jelas-jelas mengetik angka. Masalahnya, **`input()` selalu menghasilkan teks**,
tanpa peduli apa yang Anda ketik. Buktikan sendiri:

```python
umur = input("Umur: ")
print(type(umur))
```

```
<class 'str'>
```

Ini bukan kasus langka. Nanti waktu kita membaca berkas CSV di Bagian 14, semua isinya juga
datang sebagai teks — termasuk kolom yang jelas berisi angka. Pegawai yang mengeluh
"kenapa jumlahnya jadi aneh" hampir selalu sedang menjumlahkan teks tanpa sadar.

## Mengubah tipe

Solusinya mengubah tipenya dulu. Tiga perintah yang dipakai: `int()`, `float()`, dan `str()`.

```python
umur = int(input("Umur: "))
print(umur + 5)
```

Sekarang `30` jadi `35`.

Arah sebaliknya juga ada. Kalau ingin menyambung angka ke dalam teks, ubah angkanya jadi teks.

```python
jumlah_pegawai = 45
print("Jumlah: " + str(jumlah_pegawai))
```

Tapi untuk keperluan menampilkan saja, koma di `print` tadi lebih gampang:

```python
print("Jumlah:", jumlah_pegawai)
```

Dua hal yang gagal, dan sebaiknya Anda lihat sendiri sekarang:

```python
int("empat")
```

```
ValueError: invalid literal for int() with base 10: 'empat'
```

Wajar — *empat* memang bukan angka bagi Python.

```python
int("3.5")
```

```
ValueError: invalid literal for int() with base 10: '3.5'
```

Yang ini lebih mengejutkan. `int()` cuma mau menerima teks yang berisi bilangan bulat. Untuk
teks desimal, lewat `float()` dulu: `int(float("3.5"))` menghasilkan `3`.

## Satu keanehan soal desimal

Coba ini:

```python
print(0.1 + 0.2)
```

```
0.30000000000000004
```

Bukan salah ketik saya, dan bukan bug Python. Komputer menyimpan bilangan desimal dalam
bentuk biner, dan sebagian pecahan tidak bisa diwakili dengan tepat di sana — sama seperti
sepertiga tidak bisa ditulis tuntas sebagai 0,333... dalam desimal.

Untuk sekadar menampilkan, bulatkan dengan `round()`:

```python
print(round(0.1 + 0.2, 2))
```

```
0.3
```

Angka `2` di situ artinya dua angka di belakang koma.

Untuk urusan uang, saran saya: bulatkan tiap kali menampilkan, dan jangan pernah memakai
`==` untuk membandingkan dua nilai desimal. Kalau perlu presisi rupiah yang benar-benar
kaku, ada cara lain yang kita sentuh nanti setelah dasarnya kuat.

Satu lagi yang sering mengagetkan: **pembagian selalu menghasilkan desimal**, walaupun
hasilnya bulat.

```python
print(50 / 5)
print(type(50 / 5))
```

```
10.0
<class 'float'>
```

## Kalau tidak jalan

Empat pesan yang paling sering muncul di bagian ini.

**`NameError: name 'nama_unit' is not defined`**

Python tidak mengenal nama itu. Biasanya salah ketik, atau variabelnya memang belum pernah
diisi. Di Colab, penyebab lain: sel yang berisi variabelnya belum pernah dijalankan, atau
Anda menjalankan sel-selnya tidak berurutan.

**`TypeError: can only concatenate str (not "int") to str`**

Anda menyambung teks dengan angka. Bungkus angkanya dengan `str()`, atau pakai koma di
`print`.

**`TypeError: unsupported operand type(s) for +: 'int' and 'str'`**

Kebalikannya — angka di depan, teks di belakang. Obatnya sama.

**`SyntaxError: invalid decimal literal`**

Nama variabelnya diawali angka, misalnya `2nama`. Ganti jadi `nama2` atau nama lain yang
diawali huruf.

## Latihan

Ketik ini, jangan disalin, lalu jalankan:

```python
nama_unit = "Dinas Pendidikan"
jumlah_pegawai = 42
anggaran_miliar = 12.5
sudah_lapor = True

print(nama_unit, jumlah_pegawai, anggaran_miliar, sudah_lapor)
```

Lalu kerjakan empat hal ini:

1. Ubah `jumlah_pegawai` jadi `45`, jalankan lagi, pastikan angkanya ikut berubah.
2. Hitung anggaran per pegawai dengan `anggaran_miliar / jumlah_pegawai`, tampilkan hasilnya,
   lalu cek tipenya dengan `type()`.
3. Jalankan `print(jumlah_pegawai + " orang")`. Baca pesan errornya sampai habis, baru
   perbaiki.
4. Jalankan `print(0.1 + 0.2)` sekali, supaya Anda pernah melihatnya sendiri.

Nomor tiga sengaja saya minta gagal dulu. Membaca error waktu Anda sudah tahu penyebabnya
jauh lebih enak daripada bertemu pertama kali saat sedang dikejar laporan.

## Berikutnya

Di **[Bagian 04](/tutorials/python-dasar-04-string-merapikan-teks/)** kita menggarap `str`
lebih dalam: memotong, menggabung, mengubah huruf, dan membuang spasi berlebih. Latihannya
menyeragamkan nama unit kerja yang ditulis lima orang dengan lima gaya berbeda — pekerjaan
yang mungkin sudah pernah Anda lakukan manual.
