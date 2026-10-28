+++
title = "Python Dasar 08: Dictionary dan Set"
date = 2026-10-28T08:00:00+07:00
draft = false
description = "Bagian kedelapan seri Belajar Python Dasar. Menyimpan pasangan nama dan nilai dengan dictionary, merekap jumlah per kategori, lalu memakai set untuk membuang duplikat dan mencari siapa yang belum lapor."
tags = ["python", "tutorial", "belajar-python-dasar", "pemula", "dictionary", "set"]
categories = ["Teknis"]
featured_image = "/images/python-08-hero.svg"
toc = true
+++

Ini bagian kedelapan dari [seri Belajar Python Dasar](/tutorials/belajar-python-dasar/).

[Bagian 07](/tutorials/python-dasar-07-list-dan-tuple/) ditutup dengan masalah dua daftar
sejajar. Nama unit di satu list, angka realisasinya di list lain, dan keduanya harus tetap
sama urutan. Sekali ada yang menyisipkan satu baris di tempat yang salah, semuanya bergeser
tanpa pesan error.

Dictionary menyelesaikan itu. Dan setelah Anda terbiasa, kemungkinan besar ini wadah yang
paling sering Anda pakai.

## Pasangan kunci dan nilai

```python
realisasi = {"Dinas Pendidikan": 82, "Dinas Kesehatan": 45, "Dinas PU": 91}
print(realisasi)
print(len(realisasi))
```

```
{'Dinas Pendidikan': 82, 'Dinas Kesehatan': 45, 'Dinas PU': 91}
3
```

Kurung kurawal, lalu pasangan-pasangan yang dipisah koma. Tiap pasangan ditulis
`kunci: nilai`. Di sini kuncinya nama unit, nilainya persen realisasi.

Nama dan angkanya sekarang menempel. Tidak ada lagi urutan yang harus dijaga.

Untuk mengambil nilainya, sebut kuncinya di kurung siku:

```python
print(realisasi["Dinas PU"])
```

```
91
```

Bandingkan dengan list. Di list Anda bertanya "isi nomor 2 apa?", di dictionary Anda bertanya
"realisasi Dinas PU berapa?". Pertanyaan kedua jauh lebih dekat dengan cara kita berpikir
soal data kantor.

## Kunci yang tidak ada

```python
print(realisasi["Dinas Sosial"])
```

```
KeyError: 'Dinas Sosial'
```

Dinas Sosial tidak ada di dalamnya, jadi Python menolak. Ini mirip `IndexError` di list.

Kalau Anda tidak yakin kuncinya ada, pakai `get`. Ia tidak protes, cuma menyerahkan `None`:

```python
print(realisasi.get("Dinas Sosial"))
print(realisasi.get("Dinas Sosial", 0))
```

```
None
0
```

Angka kedua di `get` adalah nilai cadangan, yang dipakai kalau kuncinya tidak ketemu. Ini
akan sangat berguna sebentar lagi, waktu kita merekap.

## Menambah, mengubah, membuang

Menambah dan mengubah memakai cara yang sama persis:

```python
realisasi["Dinas Sosial"] = 60
realisasi["Dinas Kesehatan"] = 52
print(realisasi)
```

```
{'Dinas Pendidikan': 82, 'Dinas Kesehatan': 52, 'Dinas PU': 91, 'Dinas Sosial': 60}
```

Kalau kuncinya belum ada, ia ditambahkan. Kalau sudah ada, nilainya ditimpa. Tidak ada
`append` di sini.

Akibatnya, satu kunci cuma bisa muncul sekali. Ini sifat yang berguna, tapi juga bisa
menjebak:

```python
d = {"Dinas PU": 91, "Dinas PU": 40}
print(d)
```

```
{'Dinas PU': 40}
```

Angka 91 hilang tanpa peringatan. Yang ditulis belakangan menang. Kalau Anda mengetik
dictionary panjang dengan tangan dan tidak sengaja menulis satu kunci dua kali, Python diam
saja.

Untuk membuang, ada `del` dan `pop`. Sama seperti di list, `pop` menyerahkan nilai yang
dibuang:

```python
del realisasi["Dinas Sosial"]
lama = realisasi.pop("Dinas PU")
print(lama)
print(realisasi)
```

```
91
{'Dinas Pendidikan': 82, 'Dinas Kesehatan': 52}
```

## `in` memeriksa kunci, bukan nilai

```python
realisasi = {"Dinas Pendidikan": 82, "Dinas Kesehatan": 45, "Dinas PU": 91}

print("Dinas PU" in realisasi)
print(91 in realisasi)
print(91 in realisasi.values())
```

```
True
False
True
```

Baris kedua sering bikin heran. Angka 91 jelas ada di sana. Tapi `in` pada dictionary hanya
melihat kuncinya. Kalau yang ingin diperiksa nilainya, sebut `.values()` secara terang.

## Perulangan di dictionary

Kalau dictionary dimasukkan ke `for` begitu saja, yang Anda dapat kuncinya:

```python
for nama in realisasi:
    print(nama)
```

```
Dinas Pendidikan
Dinas Kesehatan
Dinas PU
```

Biasanya Anda butuh keduanya sekaligus. Untuk itu ada `.items()`:

```python
for nama, persen in realisasi.items():
    print(nama, "->", persen, "%")
```

```
Dinas Pendidikan -> 82 %
Dinas Kesehatan -> 45 %
Dinas PU -> 91 %
```

Bandingkan dengan versi dua daftar sejajar di Bagian 07. Hasilnya sama, tapi tidak ada lagi
`range(len(...))` dan tidak ada lagi dua daftar yang harus dijaga sama panjang.

Perhatikan `nama, persen` di baris `for`. Tiap putaran, `.items()` menyerahkan satu pasangan,
dan Python langsung membaginya ke dua nama. Ini disebut *unpacking*, dan ia jalan juga untuk
tuple biasa:

```python
paket = ("Tender", 975)
metode, nilai = paket
print(metode)
print(nilai)
```

```
Tender
975
```

Jumlah nama di kiri harus sama dengan jumlah isi di kanan. Kita akan memakainya lagi di
latihan.

`.values()` juga bisa langsung dipakai dengan perintah bawaan yang sudah Anda kenal:

```python
print(sum(realisasi.values()) / len(realisasi))
```

```
72.66666666666667
```

Urutannya, omong-omong, selalu sesuai urutan waktu Anda memasukkannya. Ini berlaku sejak
Python 3.7. Di tutorial lama Anda mungkin membaca bahwa dictionary tidak punya urutan; itu
benar untuk versi yang sudah lama tidak dipakai.

## Mengurutkan berdasarkan nilai

`sorted()` pada dictionary mengurutkan kuncinya:

```python
print(sorted(realisasi))
```

```
['Dinas Kesehatan', 'Dinas PU', 'Dinas Pendidikan']
```

Lagi-lagi `Dinas PU` mendahului `Dinas Pendidikan`, karena alasan huruf besar yang sama di
Bagian 07.

Yang lebih sering dibutuhkan sebenarnya urutan berdasarkan nilainya. Siapa realisasinya
tertinggi, siapa terendah. Caranya pakai `key` lagi, kali ini diisi `realisasi.get`:

```python
print(sorted(realisasi, key=realisasi.get, reverse=True))
print(max(realisasi, key=realisasi.get))
```

```
['Dinas PU', 'Dinas Pendidikan', 'Dinas Kesehatan']
Dinas PU
```

Bacanya begini: urutkan kunci-kuncinya, tapi untuk membandingkan, pakai nilai yang didapat
dari `realisasi.get(kunci)`. Mirip `key=str.lower` di bagian sebelumnya, cuma alat
pembandingnya beda.

Kalau ini terasa agak gelap, tidak apa-apa. Setelah Bagian 10 soal fungsi, `key` akan jauh
lebih masuk akal. Untuk sekarang, cukup hafal polanya.

## Apa yang boleh jadi kunci

Teks dan angka boleh. Tuple juga boleh, dan di sinilah tuple dari Bagian 07 mulai terasa
gunanya:

```python
per_bulan = {(2026, 9): 14, (2026, 10): 21}
print(per_bulan[(2026, 10)])
```

```
21
```

Kuncinya pasangan tahun dan bulan. Jumlah paket per bulan, tanpa perlu menggabungkan tahun dan
bulan jadi satu teks.

List tidak boleh:

```python
per_bulan = {[2026, 10]: 21}
```

```
TypeError: unhashable type: 'list'
```

Pesan errornya memakai istilah teknis, tapi maksudnya sederhana. Kunci harus sesuatu yang
tidak bisa berubah. Kalau kunci bisa diubah setelah dimasukkan, dictionary tidak akan bisa
menemukannya lagi. List bisa diubah, jadi ditolak. Tuple dikunci, jadi diterima.

## Set: wadah tanpa duplikat

Set itu wadah yang isinya tidak pernah kembar. Paling sering dipakai untuk satu pekerjaan:
membuang duplikat.

```python
metode = ["Tender", "E-purchasing", "Tender", "Pengadaan Langsung", "E-purchasing", "Tender"]
unik = set(metode)

print(len(metode))
print(len(unik))
print(sorted(unik))
```

```
6
3
['E-purchasing', 'Pengadaan Langsung', 'Tender']
```

Enam baris data, tiga metode berbeda. Di Excel ini setara dengan *Remove Duplicates*, cuma
tidak mengubah data aslinya.

Kenapa saya memakai `sorted()` untuk menampilkannya? Karena set tidak menjaga urutan. Coba
`print(unik)` langsung, lalu jalankan programnya sekali lagi. Di komputer saya, dua kali
jalan menghasilkan dua urutan berbeda:

```
{'E-purchasing', 'Tender', 'Pengadaan Langsung'}
{'Pengadaan Langsung', 'Tender', 'E-purchasing'}
```

Isinya sama, urutannya acak. Karena itu set juga tidak bisa diambil isinya pakai nomor:

```python
print(unik[0])
```

```
TypeError: 'set' object is not subscriptable
```

Kalau butuh urutan, ubah dulu jadi list dengan `sorted()`.

Satu lagi yang perlu diingat. Set ditulis dengan kurung kurawal, sama seperti dictionary,
dan itu menimbulkan satu kejanggalan:

```python
print(type({}))
print(type(set()))
```

```
<class 'dict'>
<class 'set'>
```

Kurung kurawal kosong adalah dictionary kosong. Untuk set kosong, tulis `set()`, lalu isi
dengan `add`:

```python
s = set()
s.add("Tender")
s.add("Tender")
print(s, len(s))
```

```
{'Tender'} 1
```

Ditambahkan dua kali, tetap satu.

## Siapa yang belum lapor

Ini bagian set yang menurut saya paling sering menghemat waktu di kantor.

Bayangkan Anda punya daftar semua unit kerja, dan daftar unit yang sudah mengirim laporan.
Pertanyaannya: siapa yang belum?

```python
semua = {"Dinas Pendidikan", "Dinas Kesehatan", "Dinas PU", "Dinas Sosial", "Bappeda"}
sudah = {"Dinas PU", "Bappeda", "Dinas Pendidikan"}

print(sorted(semua - sudah))
print(sorted(semua & sudah))
```

```
['Dinas Kesehatan', 'Dinas Sosial']
['Bappeda', 'Dinas PU', 'Dinas Pendidikan']
```

Tanda minus artinya "yang ada di kiri tapi tidak ada di kanan". Tanda `&` artinya "yang ada
di keduanya".

Dulu saya mengerjakan ini dengan dua kolom di Excel dan `VLOOKUP`, lalu memfilter yang
hasilnya `#N/A`. Jalan juga, tapi tiap bulan harus disusun ulang.

Peringatannya sama dengan di Bagian 04: `"Dinas PU"` dan `"Dinas PU "` dengan spasi di
belakang dianggap dua unit berbeda. Kalau hasilnya menyebut unit yang Anda yakin sudah lapor,
periksa penulisannya dengan `repr()`, lalu rapikan dulu dengan `strip()` sebelum dimasukkan ke
set.

## Latihan 1: merekap jumlah per metode

Ini pola yang akan Anda tulis berkali-kali seumur hidup.

```python
metode = ["Tender", "E-purchasing", "Tender", "Pengadaan Langsung",
          "E-purchasing", "Tender", "Pengadaan Langsung", "E-purchasing",
          "E-purchasing", "Tender", "E-purchasing"]

jumlah = {}
for m in metode:
    if m in jumlah:
        jumlah[m] = jumlah[m] + 1
    else:
        jumlah[m] = 1

print(jumlah)
```

```
{'Tender': 4, 'E-purchasing': 5, 'Pengadaan Langsung': 2}
```

Mulai dari dictionary kosong. Tiap metode yang lewat, kalau sudah pernah dicatat, tambah
satu. Kalau belum, catat sebagai satu.

Empat baris `if`/`else` itu bisa dipadatkan jadi satu dengan `get` dan nilai cadangannya:

```python
jumlah = {}
for m in metode:
    jumlah[m] = jumlah.get(m, 0) + 1

print("Jumlah paket :", len(metode))
print("Jenis metode :", len(jumlah))
for m in sorted(jumlah, key=jumlah.get, reverse=True):
    print(m, "->", jumlah[m])
```

```
Jumlah paket : 11
Jenis metode : 3
E-purchasing -> 5
Tender -> 4
Pengadaan Langsung -> 2
```

`jumlah.get(m, 0)` artinya: ambil hitungan sejauh ini, atau nol kalau belum ada. Lalu tambah
satu. Hasilnya sama persis dengan versi `if`/`else`.

Versi mana yang dipakai terserah Anda. Saya sendiri butuh beberapa bulan sebelum versi `get`
terasa wajar dibaca. Kalau versi panjang lebih jelas di kepala Anda, pakai yang panjang.

## Latihan 2: total nilai per unit kerja

Sekarang datanya pasangan unit dan nilai paket, ditulis sebagai list berisi tuple:

```python
paket = [
    ("Dinas PU", 820),
    ("Dinas Pendidikan", 150),
    ("Dinas PU", 640),
    ("Dinas Kesehatan", 310),
    ("Dinas Pendidikan", 975),
    ("Dinas PU", 420),
]

total = {}
banyak = {}
for unit, nilai in paket:
    total[unit] = total.get(unit, 0) + nilai
    banyak[unit] = banyak.get(unit, 0) + 1

for unit in sorted(total, key=total.get, reverse=True):
    print(f"{unit}: {banyak[unit]} paket, total {total[unit]}")

print("Total semua:", sum(total.values()))
```

```
Dinas PU: 3 paket, total 1880
Dinas Pendidikan: 2 paket, total 1125
Dinas Kesehatan: 1 paket, total 310
Total semua: 3315
```

Angka-angka paketnya sama dengan Latihan 2 di Bagian 07, dan totalnya juga sama, 3315.
Bedanya, sekarang Anda tahu dari mana saja angka itu datang.

Kalau Anda pernah memakai PivotTable, ini pada dasarnya PivotTable dengan dua kolom nilai:
jumlah baris dan total. Bentuk data seperti `paket` itu juga kurang lebih yang akan Anda
dapat waktu membaca berkas CSV di Bagian 14, jadi pola ini akan muncul lagi.

## Kalau tidak jalan

**`KeyError: 'nama kunci'`**

Kunci yang Anda minta tidak ada. Paling sering karena beda huruf besar kecil atau ada spasi
tersembunyi: `"dinas pu"` tidak sama dengan `"Dinas PU"`. Cetak `list(d.keys())` untuk melihat
kunci yang sebenarnya ada, atau pakai `get` kalau memang wajar kuncinya kadang tidak ada.

**`TypeError: unhashable type: 'list'`**

Anda memakai list sebagai kunci dictionary atau sebagai isi set. Ganti dengan tuple.

**`TypeError: 'set' object is not subscriptable`**

Anda mencoba mengambil isi set dengan nomor, seperti `unik[0]`. Set tidak punya urutan. Ubah
dulu dengan `sorted(unik)`.

**`AttributeError: 'dict' object has no attribute 'append'`**

Kebiasaan dari list. Di dictionary, menambah isi cukup dengan `d[kunci] = nilai`. Di set,
pakai `add`.

**`AttributeError: 'dict' object has no attribute 'add'`**

Anda menulis `{}` dan mengira itu set kosong. Itu dictionary kosong. Tulis `set()`.

**Ada nilai yang hilang tanpa pesan apa pun**

Kemungkinan ada kunci yang ditulis dua kali, dan yang belakangan menimpa yang duluan. Hitung
dengan `len()` dan bandingkan dengan jumlah yang Anda harapkan.

**`ValueError: too many values to unpack`**

Jumlah nama di baris `for a, b in ...` tidak cocok dengan isi tiap pasangannya. Kalau
tuple-nya berisi tiga nilai, siapkan tiga nama.

## Berikutnya

Di **Bagian 09** kita membahas *comprehension*, cara menulis perulangan pendek dalam satu
baris untuk membuat list, dictionary, atau set baru. Ringkas sekali kalau dipakai dengan
benar, dan susah dibaca kalau dipaksakan.

Tautannya muncul di [halaman silabus](/tutorials/belajar-python-dasar/) begitu terbit.
