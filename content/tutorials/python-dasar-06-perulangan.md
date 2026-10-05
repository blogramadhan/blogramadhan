+++
title = "Python Dasar 06: Perulangan dengan for dan while"
date = 2026-10-04T12:00:00+07:00
draft = false
description = "Bagian keenam seri Belajar Python Dasar. Mengerjakan hal yang sama berkali-kali tanpa menyalin kode, bedanya for dan while, serta perulangan yang tidak pernah berhenti dan cara menghentikannya."
tags = ["python", "tutorial", "belajar-python-dasar", "pemula", "perulangan", "for", "while"]
categories = ["Teknis"]
featured_image = "/images/python-06-hero.svg"
toc = true
+++

Ini bagian keenam dari [seri Belajar Python Dasar](/tutorials/belajar-python-dasar/).

Dua kali saya menyuruh Anda menyalin-tempel kode lalu mengganti angkanya satu per satu.
[Lima nama unit di Bagian 04](/tutorials/python-dasar-04-string-merapikan-teks/), dan
[empat angka realisasi di Bagian 05](/tutorials/python-dasar-05-percabangan/). Keduanya saya
tutup dengan janji yang sama: nanti di Bagian 06 tidak perlu begitu lagi.

Hari ini janjinya saya tepati.

## for: untuk tiap isi

```python
for kota in ["Pontianak", "Singkawang", "Sintang"]:
    print(kota)
```

```
Pontianak
Singkawang
Sintang
```

Bacanya: untuk tiap kota di dalam daftar itu, tampilkan kotanya.

Yang terjadi, Python mengambil isi pertama, menaruhnya di variabel `kota`, menjalankan blok di
bawahnya, lalu kembali ke atas dan mengambil isi berikutnya. Begitu terus sampai habis.

Nama `kota` itu bebas. Bisa `k`, bisa `nama_kota`. Yang penting Anda memakai nama yang sama di
dalam bloknya.

Aturan bentuknya sama persis dengan `if` di Bagian 05: **titik dua** di ujung, dan **empat
spasi** di depan baris isinya.

Kurung siku itu *list*. Kita pakai sebentar di sini dan membahasnya tuntas di Bagian 07.

Teks juga bisa diulang, per huruf:

```python
for huruf in "Hugo":
    print(huruf)
```

```
H
u
g
o
```

## range: mengulang sekian kali

Kalau yang Anda butuhkan sekadar "ulangi lima kali", tidak perlu menulis daftarnya.

```python
for i in range(3):
    print(i)
```

```
0
1
2
```

Perhatikan: mulai dari **nol**, dan berhenti **sebelum** angka yang Anda sebut. Sama seperti
penomoran huruf di Bagian 04.

`range` menerima sampai tiga angka:

| Ditulis | Hasilnya |
|---|---|
| `range(3)` | 0, 1, 2 |
| `range(1, 4)` | 1, 2, 3 |
| `range(0, 10, 3)` | 0, 3, 6, 9 |

Angka ketiga itu lompatannya.

## Digabung dengan percabangan

Di sinilah dua bagian terakhir bertemu. Blok di dalam `for` boleh berisi `if`, asal menjoroknya
ditambah lagi.

```python
for nilai in [92, 81, 70, 45]:
    if nilai >= 90:
        print(nilai, "A")
    elif nilai >= 80:
        print(nilai, "B")
    else:
        print(nilai, "lainnya")
```

```
92 A
81 B
70 lainnya
45 lainnya
```

Hitung spasinya. Baris `if` menjorok empat spasi karena ia isi dari `for`. Baris `print`
menjorok delapan karena ia isi dari `if`, yang isi dari `for`.

## Menumpuk hasil

Sering kali Anda tidak cuma menampilkan, tapi menjumlahkan. Caranya menyiapkan wadah sebelum
perulangan dimulai.

```python
total = 0

for n in [10, 20, 30]:
    total = total + n

print(total)
```

```
60
```

`total` harus dibuat **sebelum** `for`. Kalau dibuat di dalamnya, ia akan direset ke nol tiap
putaran.

## while: selama syaratnya masih benar

`for` dipakai kalau Anda tahu apa yang mau diulang. `while` dipakai kalau yang Anda tahu cuma
kapan harus berhenti.

```python
sisa = 3

while sisa > 0:
    print("sisa", sisa)
    sisa = sisa - 1

print("habis")
```

```
sisa 3
sisa 2
sisa 1
habis
```

Syarat di `while` diperiksa ulang tiap putaran. Begitu `sisa` jadi nol, syaratnya salah, dan
perulangannya berhenti.

Dalam pekerjaan sehari-hari, `for` jauh lebih sering dipakai. `while` berguna waktu jumlah
putarannya memang belum diketahui — misalnya menunggu sesuatu selesai.

## Perulangan yang tidak pernah berhenti

Ini kesalahan khas `while`, dan akibatnya langsung terasa.

```python
sisa = 3

while sisa > 0:
    print("sisa", sisa)
```

Perhatikan apa yang hilang: tidak ada yang mengurangi `sisa`. Syaratnya selamanya benar, dan
program itu akan mencetak tanpa henti sampai Anda memaksanya berhenti.

Cara menghentikannya: **Ctrl + C** di terminal, atau tombol berhenti di Colab.

Tiap kali menulis `while`, periksa satu hal: adakah baris di dalam bloknya yang membuat
syaratnya suatu saat jadi salah? Kalau tidak ada, Anda baru saja membuat perulangan abadi.

## Keluar lebih awal, atau melewati satu putaran

Dua kata yang berguna di dalam perulangan mana pun.

`break` menghentikan perulangan sama sekali:

```python
for n in [10, 20, 30, 40]:
    if n > 25:
        print("ketemu", n)
        break
```

```
ketemu 30
```

Angka 40 tidak pernah diperiksa.

`continue` melewati sisa putaran ini dan lanjut ke isi berikutnya:

```python
for n in [10, 0, 30]:
    if n == 0:
        continue
    print(100 / n)
```

```
10.0
3.3333333333333335
```

Tanpa `continue` itu, pembagian dengan nol akan menghentikan program. Ini pola yang sering
dipakai untuk melewati data kosong atau rusak.

## Latihan 1: empat paket sekaligus

Ini kode dari Bagian 05, sekarang tanpa mengganti angka satu per satu.

```python
pagu = 800_000_000
daftar_realisasi = [820_000_000, 742_000_000, 600_000_000, 350_000_000]

for realisasi in daftar_realisasi:
    persen = realisasi / pagu * 100

    if persen > 100:
        status = "MELEWATI PAGU"
    elif persen >= 90:
        status = "aman"
    elif persen >= 70:
        status = "perlu dikejar"
    else:
        status = "rendah"

    rp = f"Rp{realisasi:,}".replace(",", ".")
    print(f"{rp} - {persen:.1f}% - {status}")
```

```
Rp820.000.000 - 102.5% - MELEWATI PAGU
Rp742.000.000 - 92.8% - aman
Rp600.000.000 - 75.0% - perlu dikejar
Rp350.000.000 - 43.8% - rendah
```

Satu hal yang berubah dari Bagian 05: hasil percabangannya tidak langsung dicetak, tapi
disimpan dulu ke variabel `status`, baru ditampilkan di baris terakhir. Pola ini membuat
bagian "memutuskan" dan bagian "menampilkan" tidak saling campur.

**Tugas Anda:** tambahkan satu angka lagi ke `daftar_realisasi`, jalankan, dan perhatikan
bahwa Anda tidak menyentuh satu baris pun di dalam perulangannya.

## Latihan 2: lima nama unit

Yang ini janji dari Bagian 04.

```python
daftar = [
    "  dinas pendidikan  ",
    "DINAS PENDIDIKAN",
    "Dinas  Pendidikan",
    "dinas Pendidikan ",
    "Dinas Pendidikan",
]

for nama in daftar:
    rapi = " ".join(nama.split()).title()
    print(repr(nama), "->", repr(rapi))
```

```
'  dinas pendidikan  ' -> 'Dinas Pendidikan'
'DINAS PENDIDIKAN' -> 'Dinas Pendidikan'
'Dinas  Pendidikan' -> 'Dinas Pendidikan'
'dinas Pendidikan ' -> 'Dinas Pendidikan'
'Dinas Pendidikan' -> 'Dinas Pendidikan'
```

`repr()` itu dari Bagian 04, supaya spasi yang tersembunyi kelihatan.

Dan inilah yang saya maksud dengan menghemat waktu. Daftarnya boleh berisi lima nama atau lima
ribu — kodenya tetap sepanjang ini.

## Kalau tidak jalan

**`SyntaxError: expected ':'`**

Titik dua di ujung baris `for` atau `while` terlupa.

**`IndentationError: expected an indented block after 'for' statement on line 1`**

Baris setelah `for` tidak menjorok.

**`TypeError: 'int' object is not iterable`**

Anda menulis `for x in 42:`. Angka tunggal tidak bisa diulang satu per satu. Yang bisa diulang
itu daftar, teks, atau `range()`. Kalau maksudnya mengulang 42 kali, tulis `range(42)`.

**Program jalan terus tanpa berhenti**

Perulangan `while` tanpa jalan keluar. Tekan **Ctrl + C**, lalu periksa apakah ada baris yang
mengubah nilai yang diuji di syaratnya.

**Hasilnya cuma satu baris, padahal datanya banyak**

Baris `print`-nya tidak menjorok, jadi ia berada di luar perulangan dan cuma jalan sekali
setelah semuanya selesai.

## Berikutnya

Di **[Bagian 07](/tutorials/python-dasar-07-list-dan-tuple/)** kita membahas *list* dengan
serius — wadah yang sejak tadi kita pakai sambil lalu. Menambah isi, mengambil sebagian,
mengurutkan, dan kapan sebaiknya memakai *tuple*.
