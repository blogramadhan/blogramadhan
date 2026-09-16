+++
title = "Python Dasar 01: Kenapa Python, dan Kapan Bukan"
date = 2026-09-16T10:00:00+07:00
draft = false
description = "Bagian pertama seri Belajar Python Dasar. Apa yang Python kerjakan dengan baik, kapan Excel justru pilihan yang lebih waras, dan hitungan kasar apakah belajar ini sepadan untuk Anda."
tags = ["python", "tutorial", "belajar-python-dasar", "pemula", "excel"]
categories = ["Teknis"]
featured_image = "/images/python-01-hero.svg"
toc = true
+++

Ini bagian pertama dari [seri Belajar Python Dasar](/tutorials/belajar-python-dasar/).

Hari ini tidak ada yang perlu dipasang, dan tidak ada kode yang perlu diketik. Tugas Anda
cuma satu: memutuskan apakah seri ini layak menghabiskan waktu Anda.

Saya menaruh bagian ini di depan karena sudah beberapa kali melihat kejadian yang sama. Ada
teman kantor yang ikut kelas Python, rajin beberapa minggu, lalu pelan-pelan kembali
mengerjakan semuanya di Excel. Waktu saya tanya kenapa, jawabannya masuk akal. Pekerjaannya
memang tidak butuh Python.

## Python itu apa

Python itu bahasa pemrograman. Anda menulis perintah di sebuah berkas teks, lalu Python
membaca berkas itu dari atas ke bawah dan mengerjakan perintahnya satu per satu.

Tidak perlu dikompilasi dulu, tidak perlu membangun aplikasi. Tulis satu baris, jalankan,
langsung kelihatan hasilnya. Itu yang bikin Python enak dipakai untuk pekerjaan kecil yang
sifatnya sekali pakai.

Oh ya, namanya bukan dari ular. Guido van Rossum, yang membuatnya, kebetulan sedang senang
menonton Monty Python waktu itu.

## Yang bisa dikerjakan Python

Yang paling sering saya pakai: mengulang pekerjaan yang sama. Kalau tiap bulan Anda melakukan
langkah yang persis sama, Python bisa menyimpan langkah itu dan menjalankannya lagi kapan
saja.

Python juga tidak keberatan membuka banyak berkas sekaligus. Mau 24, 240, atau 2.400,
bedanya cuma di waktu tunggu.

Lalu soal ukuran data. Satu lembar Excel maksimal 1.048.576 baris. Kelihatannya banyak,
sampai suatu hari Anda mengekspor data transaksi setahun dan sadar angkanya terpotong di
baris terakhir tanpa peringatan apa pun.

Satu lagi yang baru saya hargai belakangan: kode meninggalkan jejak. Enam bulan lagi, waktu
ada yang bertanya kenapa angka di laporan itu begitu, Anda tinggal buka berkasnya dan baca.
Kalau prosesnya berupa klik-klik di Excel, biasanya saya sudah lupa.

## Contohnya begini

Misalkan ada 24 berkas CSV. Laporan bulanan dua tahun, satu berkas satu bulan, kolomnya sama
semua. Anda diminta menggabungkannya jadi satu.

Cara Excel: buka berkas pertama, blok, salin, tempel ke berkas gabungan. Buka berkas kedua,
salin, tempel. Begitu terus 22 kali lagi. Kalau ada satu berkas yang urutan kolomnya beda
sendiri, biasanya baru ketahuan setelah semuanya tergabung.

Cara Python:

```python
import csv
import glob

semua = []
berkas = glob.glob("laporan/*.csv")

for nama in berkas:
    with open(nama, newline="", encoding="utf-8") as f:
        semua.extend(csv.DictReader(f))

print(len(semua), "baris dari", len(berkas), "berkas")
```

Kalau kode itu belum ada artinya buat Anda, wajar. Anda memang belum belajar apa-apa. Saya
menaruhnya di sini cuma supaya Anda kenal bentuknya.

Dua hal yang menarik dari kode itu. Panjangnya delapan baris. Dan angka 24 tidak ditulis di
mana pun, jadi kalau bulan depan berkasnya bertambah jadi 25, kodenya tidak perlu disentuh.

Nanti di Bagian 14 Anda akan menulis sendiri kode seperti ini.

## Kapan Excel lebih masuk akal

Ini bagian yang jarang ada di tutorial lain, mungkin karena tidak enak ditulis.

Kalau pekerjaannya cuma sekali, pakai Excel. Menggabungkan tiga berkas hari ini dan tidak
akan diulang lagi? Sepuluh menit juga selesai. Menulis kodenya malah lebih lama.

Kalau Anda perlu melihat datanya sambil bekerja, Excel jauh lebih enak. Anda bisa menggulir,
menyorot, mengurutkan, dan tiba-tiba merasa ada yang janggal di satu baris. Python tidak
menunjukkan apa pun sampai Anda menyuruhnya.

Kalau hasilnya harus diserahkan ke orang lain, ujung-ujungnya tetap Excel juga. Atasan minta
lampiran, bukan skrip.

Kalau datanya dua ratus baris dan tiga kolom, ya sudah. Tidak ada yang perlu diotomasi di
situ.

Dan kalau tenggatnya besok pagi sementara Anda belum bisa Python, tolong jangan belajar
sekarang. Kerjakan dengan cara yang sudah Anda kuasai. Belajarnya setelah tenggat lewat.

Saya sendiri masih buka Excel hampir tiap hari.

## Kadang jawabannya bukan dua-duanya

Kalau datanya memang sudah ada di database, belajar SQL biasanya lebih langsung. Ambil yang
perlu saja, tidak usah unduh semuanya dulu.

Kalau Anda sudah nyaman di Excel dan yang dibutuhkan cuma transformasi rutin, Power Query
sudah ada di dalam Excel dan tidak menuntut Anda belajar bahasa baru.

Kalau yang diminta dasbor yang jalan sendiri, Power BI atau Looker Studio lebih cocok
daripada skrip apa pun.

Dan sebelum menulis apa-apa, tanya dulu ke tim yang mengurus aplikasinya. Saya pernah
menghabiskan dua hari menulis skrip untuk sesuatu yang ternyata sudah tersedia sebagai menu
ekspor. Tidak ada yang memberi tahu, karena saya juga tidak bertanya.

## Sepadan atau tidak

Angka di bawah ini karangan saya, cuma untuk menggambarkan polanya. Ganti dengan angka Anda
sendiri.

Katakanlah ada rekap bulanan yang makan waktu 2 jam tiap bulan. Setahun berarti 24 jam.

| | Tahun pertama | Tahun kedua |
|---|---:|---:|
| Tetap pakai Excel | 24 jam | 24 jam |
| Belajar Python + tulis skrip | 30 + 4 jam | — |
| Jalankan skrip (5 menit/bulan) | 1 jam | 1 jam |
| **Total** | **35 jam** | **1 jam** |

Lihat kolom tahun pertama. Python kalah. Anda menghabiskan 35 jam untuk pekerjaan yang kalau
dikerjakan manual cuma 24 jam.

Tahun kedua baru terbalik, dan sesudah itu tidak pernah balik lagi.

Jadi sebelum mulai, tanyakan satu hal ke diri sendiri: pekerjaan ini akan saya kerjakan lagi
tahun depan atau tidak? Kalau iya, lanjut. Kalau tidak, tutup halaman ini dan buka Excel,
tidak usah merasa bersalah.

Satu catatan. Waktu 30 jam untuk belajar itu dibayar sekali saja, bukan tiap ada pekerjaan
baru. Skrip kedua Anda jauh lebih cepat jadinya.

## Ringkasnya

Python untuk pekerjaan yang berulang, jumlahnya banyak, dan harus sama persis tiap kali.
Excel untuk yang sekali jalan, perlu dilihat langsung, dan harus diserahkan ke orang.

Kalau ternyata pekerjaan Anda memang berulang, lanjut ke Bagian 02. Di sana kita siapkan
alatnya: Google Colab kalau laptop kantor Anda terkunci, atau pasang Python langsung di
komputer kalau memang bisa.

Bagian itu belum terbit waktu tulisan ini naik. Saya menulisnya berurutan, dan tautannya
muncul di [halaman silabus](/tutorials/belajar-python-dasar/) begitu selesai.
