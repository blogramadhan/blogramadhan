+++
title = "Portal PBJ"
date = 2026-08-28T10:00:00+07:00
draft = false
description = "Portal analisis pengadaan barang/jasa pemerintah — penyedia, paket aktif, RUP, realisasi, dan risiko dalam satu panel kerja."
year = "2026"
role = "Perancang & pengembang"
demo = "https://portal.pbj.my.id"
featured_image = "/images/portal-pbj-cover.svg"
# Isi daftar teknologinya bila ingin tampil sebagai tag di kartu portofolio:
# tech = ["…", "…"]
+++

**Portal PBJ** adalah portal analisis pengadaan barang/jasa pemerintah: menelusuri penyedia,
memantau paket yang sedang berjalan, membaca rencana anggaran, dan mengukur realisasinya —
semuanya dari satu panel kerja.

Data pengadaan sebenarnya sudah terbuka, tersebar di SPSE dan INAPROC. Masalahnya bukan
ketersediaan, melainkan bentuknya: satu paket di satu LPSE, satu rencana di satu halaman RUP,
satu sanksi di satu daftar. Untuk melihat pola, seseorang harus membuka puluhan tab dan menyalin
angka satu per satu. Portal ini menarik data itu dari sumber resmi, menyatukannya dalam satu
basis, lalu menyajikannya sebagai panel yang bisa ditelusuri — dari profil satu penyedia sampai
serapan anggaran lintas instansi.

## Panel kerja

Halaman awalnya sengaja bukan kotak pencarian, melainkan pilihan panel. Kartu yang muncul
menyesuaikan role akun, sehingga tiap pengguna langsung mendarat di alat yang relevan baginya.

**Menelusuri**

- **Pencarian Paket Penyedia** — profil penyedia, nilai kontrak, sebaran daerah, LPSE, dan
  daftar paket yang pernah dimenangkan.
- **Rencana Umum Pengadaan** — RUP pada tahap perencanaan: nilai anggaran, cara pengadaan
  (penyedia/swakelola), jenis, sumber dana, status PDN, dan sebarannya per instansi.
- **Paket Pengadaan Aktif** — tender dan non-tender yang sedang berjalan, lengkap dengan fase
  aktif, jenis pengadaan, dan instansi pemiliknya.

**Mengawasi**

- **Realisasi Pengadaan** — realisasi lintas kanal: E-Katalog, Tender, Non Tender, dan Toko
  Daring, dilihat dari nilai, status, kanal, instansi, maupun penyedia.
- **Analisis & Pengawasan** — bagian yang paling banyak menyita waktu saya. Ia menjahit
  beberapa dataset sekaligus untuk membandingkan serapan rencana dengan realisasi, menghitung
  efisiensi lelang, menandai risiko integritas, dan mengukur kepatuhan PDN serta UMK-K.
- **Daftar Hitam** — penyedia yang sedang terkena sanksi, status aktif atau berakhir, trennya,
  dan detail SK yang mendasarinya.

**Mengelola**

- **Vendor Management System** — katalog vendor dari sisi toko: profil, KBLI, produk, rasio
  produk yang sudah dan belum pernah dibeli, hingga rekap penyedia yang belum tersentuh
  pembelian di tiap kabupaten/kota.
- **Manajemen Pengguna** dan **Data Statistik** — panel administratif untuk approval akun, role,
  aktivitas login, moderasi testimoni, serta analisis dataset gabungan.

## Sumber data

Seluruh isinya berasal dari sistem resmi: **SPSE (spse.inaproc.id)** dan **INAPROC
(data.inaproc.id)**. Portal ini tidak menciptakan angka sendiri — ia hanya memindahkan,
menautkan, dan menyusun ulang apa yang sudah dipublikasikan, lalu menghitung turunannya.

## Catatan

Yang paling menarik dari mengerjakan ini bukan bagian teknisnya, melainkan menyadari betapa
banyak pertanyaan sederhana yang sebelumnya sulit dijawab. "Penyedia mana yang menang paling
banyak di daerah ini?" atau "berapa rencana yang benar-benar terealisasi tahun ini?" — keduanya
seharusnya cukup satu klik, dan sekarang memang begitu.

Portal ini masih tumbuh. Beberapa panel sudah mapan, beberapa lain masih saya rapikan.
