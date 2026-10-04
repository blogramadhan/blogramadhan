+++
title = "Python Dasar 05: Percabangan dengan if, elif, else"
date = 2026-10-03T19:00:00+07:00
draft = false
description = "Bagian kelima seri Belajar Python Dasar. Mengajari program mengambil keputusan, kenapa spasi di depan menentukan segalanya, dan satu kesalahan urutan yang membuat seluruh percabangan salah tanpa pesan error."
tags = ["python", "tutorial", "belajar-python-dasar", "pemula", "percabangan", "if"]
categories = ["Teknis"]
featured_image = "/images/python-05-hero.svg"
toc = true
+++

Ini bagian kelima dari [seri Belajar Python Dasar](/tutorials/belajar-python-dasar/). Sampai
sekarang program kita cuma menjalankan semua baris dari atas ke bawah, tanpa pernah memilih.

Hari ini kita memberinya kemampuan memutuskan.

Di sini `bool` yang kita kenal di [Bagian 03](/tutorials/python-dasar-03-variabel-dan-tipe-data/)
dan `in` dari [Bagian 04](/tutorials/python-dasar-04-string-merapikan-teks/) baru benar-benar
terpakai.

## Bentuk paling sederhana

```python
nilai = 85

if nilai > 75:
    print("Lulus")
```

```
Lulus
```

Bacanya persis seperti kalimat biasa: kalau nilai lebih dari 75, tampilkan "Lulus".

Dua hal yang wajib ada dan sering terlewat. **Titik dua** di ujung baris `if`, dan **spasi di
depan** baris berikutnya.

## Spasi di depan itu bukan hiasan

Di banyak bahasa pemrograman lain, isi sebuah blok ditandai kurung kurawal. Python tidak
punya itu. Yang menentukan mana isi blok dan mana bukan adalah **seberapa menjorok baris itu
ke kanan**.

```python
nilai = 60

if nilai > 75:
    print("Lulus")
    print("Selamat")

print("Program selesai")
```

Dua baris yang menjorok itu isi dari `if`, jadi keduanya dilewati kalau syaratnya tidak
terpenuhi. Baris terakhir menempel di tepi kiri, jadi ia selalu jalan.

```
Program selesai
```

Pakai **empat spasi**. Itu kebiasaan yang disepakati hampir semua orang Python, dan VS Code
maupun Colab sudah otomatis melakukannya begitu Anda menekan Enter setelah titik dua.

Jangan mencampur spasi dan tab dalam satu berkas. Keduanya terlihat sama di layar, dan
Python akan menolak dengan pesan yang membingungkan.

## Kalau tidak, kerjakan yang ini

```python
nilai = 60

if nilai > 75:
    print("Lulus")
else:
    print("Belum lulus")
```

```
Belum lulus
```

`else` tidak punya syarat. Ia bagian yang dikerjakan kalau semua di atasnya tidak cocok.

## Lebih dari dua kemungkinan

Untuk pilihan bertingkat, pakai `elif` — singkatan dari *else if*.

```python
nilai = 81

if nilai >= 90:
    print("A")
elif nilai >= 80:
    print("B")
elif nilai >= 70:
    print("C")
else:
    print("D")
```

```
B
```

Cara kerjanya: Python memeriksa dari atas, satu per satu, lalu **berhenti di yang pertama
cocok**. Sisanya tidak diperiksa sama sekali.

Nilai 81 memang lebih besar dari 70, tapi barisnya tidak pernah sampai ke sana.

## Syarat yang bisa dipakai

Semua perbandingan ini menghasilkan `bool` — `True` atau `False` — seperti yang kita lihat di
Bagian 03.

| Ditulis | Artinya |
|---|---|
| `==` | sama dengan |
| `!=` | tidak sama dengan |
| `>` `<` | lebih besar, lebih kecil |
| `>=` `<=` | lebih besar atau sama, lebih kecil atau sama |
| `in` | ada di dalamnya |

`in` itu yang kita pakai di Bagian 04 untuk memeriksa isi teks:

```python
unit = "Dinas Pendidikan"

if "Dinas" in unit:
    print("Ini unit dinas")
```

Satu pengingat yang penting. **Satu `=` menyimpan, dua `==` membandingkan.** Tertukar sedikit
dan Python langsung menegur, tapi lebih enak tahu duluan.

## Menggabungkan beberapa syarat

Tiga kata penghubung: `and`, `or`, dan `not`.

```python
anggaran = 100
terpakai = 95

if terpakai > 90 and anggaran <= 100:
    print("Hampir habis, dan pagunya kecil")
```

`and` butuh dua-duanya benar. `or` cukup salah satu. `not` membalik hasilnya.

Kalau syaratnya berupa rentang, Python punya cara yang lebih enak dibaca daripada
kebanyakan bahasa lain:

```python
nilai = 85

if 70 <= nilai < 90:
    print("Di rentang menengah")
```

Itu boleh ditulis berantai seperti itu, dan artinya persis seperti yang terbaca.

## Kosong itu dianggap salah

Python punya kebiasaan yang awalnya mengagetkan: beberapa nilai dianggap `False` tanpa perlu
dibandingkan dengan apa pun.

```python
nama = ""

if nama:
    print("Ada isinya")
else:
    print("Kosong")
```

```
Kosong
```

Teks kosong, angka nol, dan `None` semuanya dianggap salah. Teks yang ada isinya dan angka
selain nol dianggap benar.

Ini berguna waktu merapikan data. Alih-alih menulis `if nama != "":`, cukup `if nama:`.

## Kesalahan yang tidak memberi pesan error

Sekarang bagian yang paling ingin saya tekankan, karena akibatnya diam-diam.

```python
nilai = 92

if nilai >= 70:
    print("C")
elif nilai >= 90:
    print("A")
```

```
C
```

Nilai 92 seharusnya A. Yang keluar C.

Tidak ada yang salah secara tata bahasa, jadi Python tidak mengeluh sama sekali. Masalahnya
pada **urutan**. Karena 92 memang lebih besar dari 70, cabang pertama cocok, Python berhenti,
dan cabang `>= 90` tidak pernah diperiksa.

Aturannya: untuk syarat bertingkat, **susun dari yang paling ketat ke yang paling longgar.**

Kesalahan ini tidak akan ketahuan dari pesan error. Satu-satunya cara menemukannya adalah
mencoba beberapa angka dan memeriksa hasilnya — termasuk angka di dekat tiap batas.

## Latihan: menandai serapan anggaran

Satu paket pengadaan, dan kita tandai status serapannya.

```python
nama_paket = "Pengadaan Meja Sekolah"
pagu = 800_000_000
realisasi = 742_000_000

persen = realisasi / pagu * 100

print(nama_paket)
print(f"Pagu      : Rp{pagu:,}".replace(",", "."))
print(f"Realisasi : Rp{realisasi:,}".replace(",", "."))
print(f"Serapan   : {persen:.1f}%")

if persen > 100:
    print("Status    : MELEWATI PAGU")
elif persen >= 90:
    print("Status    : aman")
elif persen >= 70:
    print("Status    : perlu dikejar")
else:
    print("Status    : rendah")
```

```
Pengadaan Meja Sekolah
Pagu      : Rp800.000.000
Realisasi : Rp742.000.000
Serapan   : 92.8%
Status    : aman
```

Dua hal kecil di kode itu. Garis bawah di `800_000_000` boleh dipakai untuk memisahkan angka
besar supaya terbaca; bagi Python nilainya sama saja dengan `800000000`. Dan `:,` memasang
pemisah ribuan gaya Inggris, jadi saya ganti komanya dengan titik memakai `replace()` dari
Bagian 04.

**Tugas Anda:** jalankan ulang dengan empat angka realisasi ini, dan pastikan statusnya sesuai
dugaan Anda.

| Realisasi | Serapan | Yang seharusnya muncul |
|---|---:|---|
| `820_000_000` | 102,5% | MELEWATI PAGU |
| `742_000_000` | 92,8% | aman |
| `600_000_000` | 75,0% | perlu dikejar |
| `350_000_000` | 43,8% | rendah |

Lalu satu tugas tambahan yang lebih penting: **bolak-balik urutan `elif`-nya** sehingga
`persen >= 70` naik ke atas. Jalankan lagi dengan `820_000_000`. Anda akan melihat sendiri
bagaimana sebuah program bisa salah tanpa mengeluh sedikit pun.

Untuk sekarang masih diganti manual satu per satu. Di Bagian 06 kita belajar perulangan, dan
keempatnya bisa diperiksa sekaligus.

## Kalau tidak jalan

**`IndentationError: expected an indented block after 'if' statement on line 1`**

Baris setelah `if` tidak menjorok. Tambahkan empat spasi di depannya.

**`SyntaxError: expected ':'`**

Titik dua di ujung baris `if`, `elif`, atau `else` terlupa.

**`SyntaxError: invalid syntax. Maybe you meant '==' or ':=' instead of '='?`**

Anda menulis `if x = 1:` dengan satu sama dengan. Untuk membandingkan, pakai `==`.

**`IndentationError: unexpected indent`**

Ada spasi di depan baris yang seharusnya menempel di tepi kiri.

**Program jalan tapi hasilnya salah**

Tidak ada pesan error, dan itu justru petunjuknya. Hampir selalu soal urutan `elif`. Periksa
apakah syarat yang longgar tidak diletakkan di atas yang ketat.

## Berikutnya

Di **[Bagian 06](/tutorials/python-dasar-06-perulangan/)** kita belajar perulangan dengan
`for` dan `while`. Di situlah empat angka di latihan tadi tidak perlu lagi diganti satu per
satu, dan Python mulai terasa benar-benar menghemat waktu.
