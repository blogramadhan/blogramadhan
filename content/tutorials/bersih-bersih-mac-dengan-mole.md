+++
title = "Bersih-bersih Mac dengan Mole"
date = 2026-09-25T14:00:00+07:00
draft = false
description = "Mole itu alat gratis yang dijalankan lewat Terminal untuk membersihkan sampah di Mac. Panduan memakainya dengan aman, dan angka jujur dari setengah tahun saya memakainya."
tags = ["tutorial", "mac", "macos", "mole", "terminal", "perawatan", "cli"]
categories = ["Teknis"]
featured_image = "/images/mole-hero.svg"
toc = true
+++

Beberapa bulan lalu Mac saya mulai mengeluh ruang penyimpanannya menipis. Saya periksa
folder-folder saya sendiri, dan tidak ada yang besar. Dokumen, foto, proyek — semuanya wajar.

Yang memenuhi disk bukan berkas saya. Itu tumpukan yang dibuat sendiri oleh aplikasi-aplikasi
yang saya pakai, di tempat yang tidak pernah saya buka.

**Mole** adalah alat gratis untuk menyapu tumpukan itu. Ia dijalankan lewat Terminal, bukan
lewat jendela aplikasi. Kedengarannya menakutkan kalau Anda belum pernah memakai Terminal,
tapi perintahnya cuma beberapa kata dan ada cara aman untuk mencobanya lebih dulu.

Tulisan ini panduan memakainya, ditutup dengan angka dari log Mac saya sendiri setelah
setengah tahun.

## Kenapa Mac bisa penuh sendiri

Hampir setiap aplikasi menyimpan salinan sementara supaya lain kali ia bekerja lebih cepat.
Peramban menyimpan gambar halaman yang pernah Anda buka. Aplikasi pengembang menyimpan hasil
unduhan dan hasil kompilasi supaya tidak mengulang pekerjaan. Sistem menyimpan catatan
kejadian dan laporan kerusakan.

Semua itu namanya *cache*, dan gunanya memang ada. Masalahnya, banyak aplikasi rajin membuat
dan malas membersihkan. Tumpukannya tinggal di `~/Library/Caches` dan beberapa tempat lain,
bertahun-tahun, tanpa ada yang menengok.

Ada satu sifat penting yang membuat ini aman dibersihkan: cache bisa dibuat ulang. Kalau
dihapus, aplikasinya tidak rusak — ia cuma perlu bekerja sedikit lebih lambat sekali, lalu
membuat cache barunya lagi.

## Memasang Mole

Kalau Anda sudah punya [Homebrew](https://brew.sh), satu baris:

```bash
brew install mole
```

Kalau belum, ada pemasang resminya:

```bash
curl -fsSL https://raw.githubusercontent.com/tw93/mole/main/install.sh | bash
```

Cara kedua mengunduh skrip dari internet lalu langsung menjalankannya. Itu praktik yang lazim,
tapi artinya Anda mempercayai sumbernya sepenuhnya. Saya sendiri pakai Homebrew.

Syaratnya macOS 12 ke atas, jalan di Mac Intel maupun Apple Silicon. Setelah terpasang,
perintahnya dipanggil dengan `mo`:

```bash
mo --version
```

Kalau keluar nomor versi, pemasangannya berhasil.

## Aturan pertama: lihat dulu, jangan hapus

Ini bagian yang paling ingin saya tekankan. Sebelum menghapus apa pun, jalankan dulu versi
pratinjaunya:

```bash
mo clean --dry-run
```

`--dry-run` artinya "jalankan pura-pura". Mole akan memindai dan menunjukkan daftar apa saja
yang akan dihapus beserta ukurannya, lalu berhenti. Tidak ada satu berkas pun yang hilang.

Baca daftarnya. Kalau ada yang Anda tidak kenal atau tidak rela, Anda tahu itu sebelum
terjadi, bukan sesudah.

Hampir semua perintah Mole punya `--dry-run`. Biasakan memakainya di kali pertama, untuk
perintah apa pun.

## Membersihkan

Kalau daftarnya sudah Anda lihat dan tidak ada yang mengganggu:

```bash
mo clean
```

Mole akan menyapu cache, log sistem, laporan kerusakan, dan sisa-sisa aplikasi. Prosesnya
beberapa menit, tergantung seberapa lama Mac Anda tidak pernah dibersihkan.

Kalau ada cache tertentu yang Anda ingin selalu dilewati — misalnya aplikasi yang lambat sekali
membangun ulang cache-nya — ada daftar lindung:

```bash
mo clean --whitelist
```

## Mencopot aplikasi sampai ke akarnya

Ini kegunaan Mole yang menurut saya paling terasa, di luar urusan cache.

Waktu Anda menyeret aplikasi ke Trash, yang terbuang cuma aplikasinya. Pengaturan, cache,
dan berkas pendukungnya tetap tinggal di `~/Library`, kadang sampai ratusan megabita, dan
tidak akan pernah terpakai lagi.

```bash
mo uninstall
```

Mole menampilkan daftar aplikasi terpasang, Anda pilih, dan ia mencari semua jejaknya sekalian.
Waktu saya mencopot sebuah editor kode bulan ini, yang ikut terangkat bukan cuma aplikasinya
tapi juga folder dukungannya, cache-nya, penyimpanan HTTP-nya, dan empat berkas preferensi
yang namanya tidak mirip sama sekali dengan nama aplikasinya. Saya tidak akan menemukan itu
sendiri.

Satu hal yang menenangkan: di log saya, semua yang diangkat `mo uninstall` tercatat dipindahkan
ke **Trash**, bukan dihapus langsung. Jadi kalau salah pilih, masih bisa dikembalikan selama
Trash belum dikosongkan.

## Perintah lain yang saya pakai

| Perintah | Gunanya |
|---|---|
| `mo analyze` | Menelusuri isi disk, melihat folder mana yang paling gemuk |
| `mo purge` | Membuang sisa proyek pemrograman lama seperti `node_modules` |
| `mo installer` | Mencari berkas pemasang lama yang menumpuk (.dmg, .pkg) |
| `mo status` | Layar pemantau CPU, memori, disk, dan jaringan |
| `mo history` | Melihat riwayat pembersihan yang pernah dijalankan |

Kalau Anda tidak pernah menulis kode, `mo purge` kemungkinan besar tidak menemukan apa-apa,
dan itu wajar. Sisanya berguna untuk siapa saja.

Mengetik `mo` saja tanpa tambahan apa-apa akan membuka menu, jadi Anda tidak perlu menghafal
daftar di atas.

## Yang saya dapat setelah setengah tahun

Sekarang bagian angkanya. Mole menyimpan catatan tiap kali dijalankan, di
`~/Library/Logs/mole/`, jadi saya bisa menghitung mundur apa saja yang sudah terjadi.

Sejak Maret 2026 saya menjalankannya sepuluh kali. Empat di antaranya `mo clean`:

![Empat sesi pembersihan: 22 GB pada Maret, 2,8 GB akhir Maret, 13,8 GB pada Juni, dan 12,4 GB pada September](/images/mole-sesi-clean.svg)

Totalnya **53,4 GB** dari 9.825 berkas.

Yang paling banyak memakan tempat, urut dari terbesar:

| Lokasi | Ruang |
|---|---:|
| Cache peramban | 8,0 GB |
| `/private/var` (berkas sementara sistem) | 6,8 GB |
| Cache Homebrew | 6,6 GB |
| Unduhan paket pemrograman | 4,9 GB |
| Cache alat pengembang lain | 3,2 GB |

Kalau Anda bukan pengembang, daftar Anda akan tampak berbeda — porsi peramban dan sistem
biasanya lebih dominan.

Tapi angka yang sebenarnya paling berguna bukan totalnya. Lihat dua batang terakhir di grafik
tadi. Juni membebaskan 13,8 GB, lalu September membebaskan 12,4 GB lagi.

Dalam 3,8 bulan, sampahnya tumbuh kembali hampir sebanyak semula. Sekitar 3,3 GB per bulan.

Jadi jangan berharap sekali bersih lalu selesai. Yang masuk akal adalah menjalankannya setiap
beberapa bulan, seperti membuang sampah rumah. Saya sendiri kira-kira tiga bulan sekali, dan
itu terasa cukup.

## Kalau ada sesuatu yang gagal

Dari 9.825 berkas, tujuh gagal dihapus dengan keterangan *permission denied*. Ini normal.

macOS punya pengaman bernama SIP yang melindungi sebagian folder sistem, dan beberapa folder
aplikasi memang terkunci. Mole akan melewatinya dan melanjutkan. Anda tidak perlu memaksa,
dan sebaiknya memang tidak.

Mole sengaja dirancang jalan tanpa `sudo`, dan hanya meminta izin admin kalau benar-benar
perlu. Kalau Anda bosan mengetik kata sandi, ada:

```bash
mo touchid
```

yang membuat permintaan izin bisa dijawab dengan sidik jari.

## Sebelum Anda menjalankannya

Beberapa hal yang sebaiknya Anda tahu, karena alat ini menghapus berkas:

- **Pakai `--dry-run` di kali pertama.** Untuk tiap perintah, bukan cuma `clean`.
- **Anggap `mo clean` permanen.** Saya memastikan `mo uninstall` memindahkan ke Trash, tapi
  saya belum memastikan hal yang sama untuk `clean`. Perlakukan sebagai hilang.
- **Aplikasi akan lambat sekali di pembukaan pertama** setelah cache-nya dibuang. Ini wajar
  dan cuma sekali.
- **Punya cadangan.** Time Machine atau apa pun. Ini berlaku untuk semua alat yang menghapus,
  bukan cuma Mole.

Mole yang dibahas di sini adalah versi Terminal, gratis dan berlisensi GPL-3.0. Ada juga
aplikasi Mac berjendela dengan nama sama yang dijual terpisah; saya belum mencobanya.

## Ringkasnya

Tiga perintah yang cukup untuk memulai:

```bash
mo clean --dry-run    # lihat apa yang akan dihapus
mo clean              # hapus
mo uninstall          # copot aplikasi sampai bersih
```

Selebihnya bisa Anda telusuri sendiri lewat menu `mo`.

Angka di tulisan ini dari mesin saya: Mole 1.55.0, macOS 26.6.2, Apple Silicon, dipasang lewat
Homebrew, log terhitung 25 September 2026. Nama aplikasi yang saya copot sengaja tidak saya
sebut. Hasil di Mac Anda hampir pasti berbeda — yang paling menentukan adalah sudah berapa lama
ia tidak pernah dibersihkan.

## Sumber

- [tw93/Mole — GitHub](https://github.com/tw93/Mole)
- [mole.fit — situs resmi](https://mole.fit/)
- [Mole combines CleanMyMac, AppCleaner, and DaisyDisk into one free, open-source tool — XDA](https://www.xda-developers.com/mole-mac-combines-paid-tools/)
