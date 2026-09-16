+++
title = "Belajar Python Dasar: Silabus"
date = 2026-09-15T07:30:00+07:00
draft = false
description = "Peta lengkap seri Belajar Python Dasar — 15 bagian, dari nol sampai bisa membaca CSV dan menulis Excel sendiri. Ditulis untuk yang pekerjaannya berurusan dengan data, bukan untuk calon software engineer."
tags = ["python", "tutorial", "belajar-python-dasar", "pemula", "data"]
categories = ["Teknis"]
featured_image = "/images/python-dasar-hero.svg"
toc = true
+++

Beberapa kali saya ditanya rekan kerja dengan kalimat yang kurang lebih sama: *"Python itu
susah nggak, ya?"* Yang sebenarnya mereka tanyakan biasanya bukan itu. Yang mereka tanyakan
adalah apakah pekerjaan yang tiap bulan memakan dua hari — menyalin data dari satu berkas ke
berkas lain, merapikan nama yang penulisannya berbeda-beda, menjumlahkan per unit kerja —
bisa dikerjakan lebih cepat.

Jawabannya bisa. Tapi jalan menuju ke sana sering tersesat di tutorial yang salah sasaran:
terlalu cepat masuk ke hal yang tidak dibutuhkan, atau terlalu lama berputar di teori yang
tidak pernah dipakai.

Seri ini percobaan saya menulis jalan yang lurus. Tulisan ini silabusnya — peta seluruh
perjalanan, supaya Anda tahu sedang di mana dan masih berapa jauh.

## Untuk siapa seri ini

Saya menulisnya dengan satu orang tertentu di kepala: **seseorang yang pekerjaannya berurusan
dengan data, tapi belum pernah menulis satu baris kode pun.**

Mungkin Anda mengelola rekap bulanan. Mungkin Anda menyalin data dari aplikasi ke Excel, lalu
dari Excel ke laporan. Mungkin Anda sudah cukup mahir dengan rumus Excel dan mulai merasa ada
batasnya.

Kalau begitu, seri ini untuk Anda. **Tidak ada prasyarat.** Anda tidak perlu pernah belajar
pemrograman, tidak perlu latar belakang teknik informatika, dan tidak perlu komputer khusus.

Kalau Anda sudah bisa bahasa pemrograman lain dan cuma ingin cepat tahu sintaks Python, seri
ini akan terasa lambat — dokumentasi resmi Python lebih cocok untuk Anda.

## Yang akan Anda bisa di akhir

Bukan janji muluk. Ini daftar yang saya anggap realistis setelah 15 bagian:

- **Membaca berkas CSV dan menulis Excel lewat kode** — tanpa membuka aplikasinya sama sekali.
- **Merapikan data berantakan** — menyeragamkan penulisan, membuang duplikat, memisahkan kolom
  yang tergabung.
- **Mengubah pekerjaan berulang jadi skrip** yang bisa dijalankan lagi bulan depan tanpa
  mengulang dari awal.
- **Membaca pesan error tanpa panik** dan memperbaikinya sendiri.
- **Membaca kode Python orang lain** dan mengerti garis besarnya — pintu masuk untuk belajar
  apa pun berikutnya.

Yang **tidak** akan Anda dapat: kemampuan membangun aplikasi web, melatih model, atau melamar
sebagai software engineer. Itu perjalanan lain, dan seri ini cuma bagian awalnya.

## Peta seri

![Peta seri Belajar Python Dasar: lima tahap berisi lima belas bagian](/images/python-dasar-peta.svg)

Lima belas bagian, dikelompokkan jadi lima tahap. Urutannya bukan hiasan — tiap bagian
memakai apa yang dibangun sebelumnya.

### Tahap I — Mulai

| Bagian | Isinya |
|---|---|
| **01** · [Kenapa Python, dan Kapan Bukan](/tutorials/python-dasar-01-kenapa-python/) ✅ | Apa yang Python kerjakan dengan baik, dan kapan Excel atau alat lain sebenarnya pilihan yang lebih waras. Supaya Anda tidak memakai obeng untuk memaku. |
| **02** · Menyiapkan Alat: Colab atau Komputer Sendiri | Dua jalur. **Google Colab** cukup browser, cocok kalau laptop kantor terkunci dan tidak bisa memasang aplikasi. **Python + VS Code** di komputer sendiri untuk yang bisa. Termasuk cara memilih di antara keduanya. |

### Tahap II — Dasar Bahasa

| Bagian | Isinya |
|---|---|
| **03** · Variabel dan Tipe Data | Menyimpan nilai dan memberinya nama. Empat tipe dasar: teks, bilangan bulat, desimal, benar/salah — dan kenapa membedakannya menyelamatkan Anda dari bug yang membingungkan. |
| **04** · String: Merapikan Teks | Memotong, menggabung, mengubah huruf, membuang spasi berlebih. Latihan: menyeragamkan nama unit kerja yang ditulis lima orang dengan lima gaya berbeda. |
| **05** · Percabangan: if, elif, else | Membuat program mengambil keputusan. Latihan: menandai baris yang nilainya melewati batas. |
| **06** · Perulangan: for dan while | Mengerjakan hal yang sama berkali-kali tanpa menyalin kode. Di sinilah Python mulai terasa menghemat waktu. |

### Tahap III — Menyimpan Banyak Data

| Bagian | Isinya |
|---|---|
| **07** · List dan Tuple | Menyimpan banyak nilai dalam satu wadah. Mengambil, menambah, mengurutkan — dan kapan sebaiknya memakai tuple. |
| **08** · Dictionary dan Set | Pasangan kunci–nilai, wadah paling berguna untuk data nyata. Latihan: merekap jumlah per kategori. |
| **09** · Comprehension | Menulis perulangan dalam satu baris. Ringkas, khas Python — dan batasnya, supaya tidak berubah jadi kode yang tak terbaca. |

### Tahap IV — Merapikan Kode

| Bagian | Isinya |
|---|---|
| **10** · Fungsi | Membungkus langkah berulang jadi satu nama. Parameter, nilai kembali, dan kapan sebaiknya sebuah kode dipecah. |
| **11** · Modul, Paket, dan pip | Memakai kode orang lain, memasang paket, dan memecah program jadi beberapa berkas. |
| **12** · Menangani Error | Membaca pesan error tanpa gentar, `try`/`except`, dan memeriksa data masukan yang tidak bisa dipercaya. Bagian yang paling sering dilewati tutorial lain, padahal paling sering dipakai. |

### Tahap V — Keluar dari Layar

| Bagian | Isinya |
|---|---|
| **13** · Membaca dan Menulis Berkas | Membuka berkas teks, menyimpan hasil, dan mengelola *path* supaya tidak kacau di Windows. |
| **14** · CSV dan Excel | Format yang paling sering Anda temui. Membaca CSV dengan modul bawaan, lalu menulis `.xlsx` dengan **openpyxl**. |
| **15** · Proyek Penutup: Rekap Otomatis | Menggabungkan semuanya: membaca beberapa berkas CSV, merapikan isinya, merekap per kategori, lalu menyimpannya jadi satu berkas Excel yang rapi. |

## Cara memakai seri ini

Empat hal yang menurut saya menentukan, jauh lebih daripada pilihan tutorialnya:

**Ketik ulang kodenya, jangan salin-tempel.** Ini terdengar sepele dan memang menyebalkan.
Tapi mengetik memaksa Anda membaca tiap karakter, dan salah ketik adalah cara tercepat belajar
membaca pesan error. Menyalin-tempel menghasilkan kode yang jalan tanpa pemahaman yang ikut
jalan.

**Satu bagian, satu duduk.** Tiap bagian saya rancang selesai dalam sekali duduk. Menumpuk
tiga bagian dalam semalam terasa produktif malam itu dan hilang seminggu kemudian.

**Kalau macet, tetap lanjut.** Ada konsep yang baru masuk akal setelah Anda melihatnya dipakai
di bagian berikutnya. Menunggu sampai paham 100% sebelum lanjut adalah cara paling umum
berhenti di tengah jalan.

**Pakai data Anda sendiri.** Tiap bagian ada latihannya, tapi latihan yang paling berguna
selalu yang memakai berkas yang benar-benar ada di komputer Anda. Ganti contoh saya dengan
data pekerjaan Anda sesegera mungkin.

## Yang sengaja tidak dibahas

Supaya jelas batasnya sejak awal:

- **pandas** — pustaka analisis data yang hampir pasti akan Anda pakai nanti. Saya tidak
  memasukkannya justru karena ia terlalu berguna: orang yang langsung belajar pandas sering
  bisa menyalin resep tanpa paham apa yang terjadi. Setelah seri ini, pandas akan terasa jauh
  lebih masuk akal.
- **OOP (class dan object)** — diperkenalkan sekilas di penutup sebagai arah lanjutan, tidak
  dibahas penuh. Untuk pekerjaan yang jadi sasaran seri ini, ia belum dibutuhkan.
- **Pengembangan web, machine learning, otomasi browser** — semuanya pintu yang terbuka setelah
  dasarnya kokoh, dan semuanya di luar cakupan.

## Kapan terbitnya

Saya menulisnya bertahap, sesuai prinsip yang saya pegang di blog ini: pelan, tapi tidak
berhenti. Bagian yang sudah terbit ditandai ✅ dan judulnya bisa diklik. **Tautan lain saya
isi di halaman ini setiap kali satu bagian selesai**, jadi halaman ini bisa Anda simpan
sebagai titik masuk.

Kalau ada bagian yang menurut Anda kurang, atau ada kasus di pekerjaan Anda yang cocok
dijadikan contoh, saya senang mendengarnya lewat kolom komentar di bawah.

Silakan mulai dari **[Bagian 01: Kenapa Python, dan Kapan Bukan](/tutorials/python-dasar-01-kenapa-python/)**.
