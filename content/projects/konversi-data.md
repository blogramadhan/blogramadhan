+++
title = "Konversi Data"
date = 2026-08-28T09:00:00+07:00
draft = false
description = "Alat konversi berkas JSON dan CSV menjadi Excel (.xlsx) — lewat unggahan berkas atau langsung dari URL."
year = "2026"
role = "Perancang & pengembang"
tech = ["React", "Chakra UI", "Vite", "axios"]
demo = "https://konversi-data.pbj.my.id"
featured_image = "/images/konversi-data-cover.svg"
+++

**Konversi Data** adalah alat kecil untuk mengubah berkas JSON dan CSV menjadi Excel — cukup
unggah berkasnya atau tempel URL-nya, dan `.xlsx` langsung terunduh.

Ia lahir dari kejengkelan yang berulang. Data publik sering disajikan sebagai JSON dari sebuah
API atau CSV mentah, sementara orang yang membutuhkannya bekerja di Excel. Jalan pintasnya
biasanya menempelkan teks ke spreadsheet lalu berkelahi dengan pemisah kolom, atau membuka situs
konversi asing yang meminta unggahan lebih dulu dan menampilkan iklan sesudahnya.

## Dua cara memasukkan data

- **Upload File** — pilih berkas `.json` atau `.csv` dari perangkat, atau seret ke area unggah.
- **Dari URL** — tempel alamat berkas JSON/CSV yang sudah ada di internet, tanpa perlu
  mengunduhnya dulu. Ini yang paling sering saya pakai untuk data yang disajikan lewat API.

Nama sheet di berkas Excel bisa ditentukan sendiri sebelum konversi, sehingga hasilnya tidak
selalu bernama `Sheet1`. Berkasnya terunduh otomatis begitu proses selesai.

## Cara kerjanya

Antarmukanya berupa aplikasi satu halaman: **React** dengan komponen **Chakra UI**, dibundel
memakai **Vite**. Konversinya sendiri tidak terjadi di peramban — berkas atau URL dikirim lewat
**axios** ke layanan terpisah, yang mengembalikan berkas Excel sebagai *blob* untuk langsung
diunduh.

Pemisahan itu disengaja. Menaruh konversi di sisi peladen berarti berkas besar tidak membebani
peramban pengguna, dan logika penguraiannya bisa diperbaiki tanpa menyentuh antarmuka sama
sekali. Halaman depannya juga menampilkan penghitung sederhana — total konversi, jumlah hari
ini, dan format yang paling sering dipakai.

## Catatan

Alat ini sengaja dibuat sesempit mungkin: satu pekerjaan, tanpa akun, tanpa langkah tambahan.
Godaan untuk menambahkan fitur selalu ada — pratinjau tabel, pemilihan kolom, konversi
berbalik dari Excel ke JSON — tapi sejauh ini saya menahannya. Alat yang hanya melakukan satu
hal jauh lebih mudah dipercaya, dan jauh lebih jarang rusak.
