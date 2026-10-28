+++
title = "Konversi Data"
date = 2026-08-28T09:00:00+07:00
draft = false
description = "Alat untuk mengubah JSON dan CSV jadi Excel, menarik dataset pengadaan langsung dari ISB Inaproc, menampilkan profil RUP dan tender, serta menyusun laporan sesuai template KPK."
year = "2026"
role = "Perancang & pengembang"
tech = ["Go", "Gothic", "WebAssembly", "Tailwind CSS"]
demo = "https://konversi-data.lpse.dev"
featured_image = "/images/konversi-data-cover.svg"
+++

**Konversi Data** awalnya alat kecil untuk satu pekerjaan: mengubah berkas JSON dan CSV jadi
Excel. Sekarang isinya lebih banyak. Ia bisa menarik data pengadaan langsung dari ISB Inaproc,
menampilkan profil RUP dan tender sebuah instansi, dan menyusun laporan dalam format yang
diminta KPK.

Alamatnya juga pindah, dari `konversi-data.pbj.my.id` ke
[konversi-data.lpse.dev](https://konversi-data.lpse.dev).

## Dari mana awalnya

Data publik sering disajikan sebagai JSON dari sebuah API atau CSV mentah, sementara orang
yang membutuhkannya bekerja di Excel. Jalan pintasnya biasanya menempelkan teks ke
spreadsheet lalu berkelahi dengan pemisah kolom, atau membuka situs konversi asing yang minta
unggahan dulu dan menampilkan iklan sesudahnya.

Bagian itu masih ada dan masih jadi pintu depannya:

- **Upload File**: pilih atau seret berkas `.json` atau `.csv`, maksimal 50 MB.
- **Dari URL**: tempel alamat publik yang langsung mengarah ke berkas JSON atau CSV, tanpa
  perlu mengunduhnya dulu.

Nama sheet bisa ditentukan sendiri. Kalau dikosongkan, namanya `Data`. Sekarang ada juga
tombol **Pratinjau** untuk melihat isinya sebelum dikonversi.

## Tarik data dari ISB Inaproc

Ini tambahan yang paling sering saya pakai. Daripada mengunduh dataset satu per satu dari
data.inaproc.id, cukup masukkan token ISB, kode KLPD, dan rentang tahun (maksimal lima tahun),
lalu centang dataset yang dibutuhkan.

Datasetnya dikelompokkan per modul:

- **RUP**: paket penyedia dan swakelola terumumkan, paket anggaran, lokasi paket, master
  satker, program, kegiatan, struktur anggaran, dan riwayat kaji ulang.
- **Tender & Non Tender**: pengumuman, jadwal tahapan, peserta, kontrak, eKontrak (SPPBJ,
  BAPBAST, SPMK/SPP), nilai, tender selesai, dan pencatatan non-tender serta swakelola.
- **E-Katalog**: instansi, satker, dan arsip paket e-purchasing, termasuk Katalog V6.
- **BELA (Toko Daring)**: realisasi.

Hasilnya satu berkas Excel dengan satu sheet per dataset, ditambah sheet Info yang mencatat
apa yang ditarik. Kalau butuh data mentah, keluarannya bisa juga JSON utuh.

Ada dua versi API yang bisa dipilih. **v1** memakai paginasi dan datanya paling lengkap, tapi
dataset besar bisa makan beberapa menit. **Legacy** cuma satu permintaan dan jauh lebih cepat,
dengan risiko sebagian dataset kurang lengkap. Saya biarkan dua-duanya karena memang ada
saatnya kecepatan lebih penting.

Tokennya hanya diteruskan ke ISB untuk permintaan itu dan tidak disimpan di server. Kalau mau,
token bisa diingat di browser sendiri supaya tidak perlu ditempel ulang.

## Analisa data

Dengan token, kode KLPD, dan tahun yang sama, ada dua tampilan ringkas:

- **Profil RUP**: persentase input RUP dari paket penyedia, paket swakelola terumumkan, dan
  struktur anggaran.
- **Profil Tender**: jumlah dan nilai paket, efisiensi, distribusi metode dan status, diambil
  dari pengumuman tender, eKontrak, dan status PDN/UKM paket RUP.

Ini bukan dasbor lengkap. Gunanya untuk menjawab pertanyaan cepat sebelum saya membuka
spreadsheet atau [Portal PBJ](/projects/portal-pbj/).

## Laporan KPK

Bagian ini dikunci dengan PIN karena memang dibuat untuk pekerjaan tertentu. Ia menyusun
berkas Excel sesuai template permintaan data KPK, langsung dari ISB Inaproc:

- Daftar RUP
- Daftar paket konsolidasi
- Daftar paket pengadaan langsung (2024 sampai 2026)
- Daftar kontrak
- Paket PBJ strategis (sepuluh pagu terbesar)
- Daftar pengadaan e-purchasing lama

Tidak semua kolom di template itu tersedia di ISB. Untuk e-purchasing lama misalnya, baru 19
dari 40 kolom yang bisa diisi; sisanya cuma ada di Kendali Inaproc. Kolom yang tidak tersedia
dibiarkan kosong dan dijelaskan di sheet Catatan. Saya sengaja tidak mengisi apa pun dengan
perkiraan, karena laporan ini akan dibaca orang yang tugasnya mencari kejanggalan.

Ada juga mode cepat yang memakai endpoint legacy untuk anggaran RUP, pengumuman, dan
non-tender. Hasilnya sudah saya cocokkan setara dengan v1, dengan jumlah permintaan yang jauh
lebih sedikit. Dataset lain tetap lewat v1.

## Excel yang langsung bisa dipakai

Semua jalur di atas berujung ke berkas Excel yang sama rapinya:

- Tipe data dikenali otomatis. Angka tetap angka, sementara NIK dan nomor HP tidak berubah jadi
  notasi ilmiah atau kehilangan nol di depan.
- Header ditebalkan, baris header dibekukan, lebar kolom disesuaikan.
- Berkas diproses di memori dan tidak disimpan.

## Cara kerjanya

Versi pertama dibangun dengan React, Chakra UI, dan Vite, dengan konversi di layanan terpisah
yang dipanggil lewat axios. Versi sekarang ditulis ulang dengan **Gothic**, framework Go yang
menjalankan sebagian logikanya di browser lewat **WebAssembly**, dengan tampilan memakai
**Tailwind CSS**.

## Catatan

Di versi lama halaman ini, saya menulis bahwa saya sengaja menahan godaan menambah fitur,
karena alat yang hanya melakukan satu hal lebih mudah dipercaya. Ternyata tidak bertahan lama.
Pratinjau yang dulu saya sebut sebagai godaan sekarang ada.

Alasannya praktis. Hampir semua data yang saya ubah ke Excel ternyata datang dari sumber yang
sama, ISB Inaproc, dan langkah mengunduh lalu mengonversinya selalu berulang. Jadi saya
memasukkan langkah itu ke alatnya sendiri. Pintu depannya tetap sama: unggah berkas, dapat
Excel, tanpa akun.
