+++
title = "Membangun Portal PBJ dengan Grit Framework"
date = 2026-10-14T11:00:00+07:00
draft = false
description = "Kenalan dengan Grit, framework full-stack Go + React yang bisa membuat API dan halaman admin dari satu perintah, dan cerita saya memakainya untuk membangun Portal PBJ."
tags = ["grit", "go", "react", "portal-pbj", "pengadaan", "framework"]
categories = ["Jurnal"]
toc = true
+++

Setiap kali memulai aplikasi web baru, saya mengerjakan hal yang sama dari awal: login, hak akses,
tabel database, API, halaman admin. Pekerjaannya tidak sulit, tapi memakan minggu-minggu pertama
sebelum saya sempat menyentuh masalah yang sebenarnya ingin saya selesaikan.

Untuk [Portal PBJ](https://portal.pbj.my.id), saya mencoba cara lain: memakai
[Grit Framework](https://gritframework.dev/). Tulisan ini berisi apa itu Grit, dan bagaimana
saya memakainya.

## Apa itu Grit?

Grit adalah framework full-stack untuk **Go di sisi server** dan **React di sisi tampilan**.

Idenya sederhana. Anda mendeskripsikan sebuah *resource*, misalnya "Produk punya nama dan harga".
Grit lalu membuatkan:

- model dan tabel database,
- API-nya,
- tipe data TypeScript yang cocok dengan model tadi,
- React hooks untuk memanggil API itu,
- halaman admin lengkap dengan tabel dan formulir.

Yang tadinya perlu berjam-jam menyambung backend dan frontend, di sini keluar dari satu perintah:

```bash
grit generate resource Product --fields "name:string,price:money"
```

Situs Grit menyebut dirinya *opinionated*. Artinya, struktur folder dan polanya sudah ditentukan.
Anda kehilangan sedikit kebebasan, tapi tidak perlu lagi memutuskan hal-hal kecil berulang kali.

### Isinya apa saja?

Dari situs resminya, fitur yang sudah tersedia sejak awal antara lain:

| Kebutuhan | Yang disediakan Grit |
| --- | --- |
| Login | Register, JWT dengan refresh token, OAuth, 2FA, passkeys |
| Hak akses | Role dan permission per resource (RBAC) |
| Pekerjaan latar | Antrean job |
| Berkas | Penyimpanan file (S3/R2/MinIO) |
| Pemantauan | Pulse untuk tracing dan query lambat |
| Keamanan | Audit log, rate limiting, header keamanan |
| Deploy | Ke server sendiri dengan systemd dan Caddy (TLS otomatis) |

Teknologi di baliknya: Go dengan GORM, React (Next.js atau TanStack Router), dan database
PostgreSQL, SQLite, atau MySQL.

> Catatan: daftar fitur dan klaim performa di atas berasal dari situs Grit sendiri. Saya belum
> mengujinya satu per satu.

## Mencobanya

Pasang CLI-nya:

```bash
go install github.com/MUKE-coder/grit/v3/cmd/grit@latest
```

Buat proyek baru. Pilihan `--triple` membuat tiga bagian terpisah: API, web publik, dan admin.

```bash
grit new my-app --triple
cd my-app
docker compose up -d     # Postgres, Redis, MinIO, Mailhog
pnpm install
grit migrate
grit seed
grit start
```

Setelah itu API berjalan di `:8080`, web di `:3000`, dan admin di `:3001`.

Perintah lain yang sering dipakai:

```bash
grit generate resource ...   # buat resource baru
grit migrate                 # jalankan migrasi database
grit deploy --host deploy@server.com --domain myapp.com
grit upgrade                 # perbarui template tanpa menimpa file yang sudah Anda ubah
```

## Portal PBJ dan masalah yang ingin dijawabnya

Data pengadaan pemerintah sebenarnya sudah terbuka di SPSE dan INAPROC. Masalahnya, datanya
tersebar: satu paket di satu LPSE, satu rencana di satu halaman RUP, satu sanksi di satu daftar.
Untuk melihat pola, orang harus membuka puluhan tab dan menyalin angka satu per satu.

Portal PBJ menarik data itu dari sumber resmi, menyatukannya dalam satu basis, lalu menyajikannya
sebagai panel kerja. Isinya antara lain pencarian paket penyedia, rencana umum pengadaan, paket
aktif, realisasi, daftar hitam, sampai analisis dan pengawasan.

Aplikasi seperti ini punya banyak sekali "bagian membosankan": akun pengguna, persetujuan akun,
role, tabel dengan filter, halaman admin. Di situlah Grit terasa cocok.

## Kenapa saya memilih Grit

Alasan utamanya satu: **satu repo, satu basis kode, untuk web, desktop, dan mobile.**

Biasanya tiga platform berarti tiga proyek. Perubahan kecil, misalnya menambah satu kolom di
tabel, harus saya ulangi di tiap tempat. Di Grit, backend Go yang sama melayani ketiganya, dan
tipe datanya ikut tersinkron ke sisi tampilan.

Alasan kedua lebih personal: ini pengalaman baru. Saya ingin mencoba membangun aplikasi
berbasis **Go** dengan **TanStack** di sisi tampilan, dan Grit menyediakan keduanya sebagai
pilihan bawaan. Jadi saya tidak perlu merakit sendiri dari nol.

## Bagaimana saya memakainya

Saya mencoba **semua varian** yang ditawarkan Grit, mulai dari web sampai desktop. Versi desktop
sudah saya jalankan di **Windows dan macOS**. Grit memakai Wails untuk bagian desktop ini.

Untuk sisi web, saya memilih **TanStack Router**, bukan Next.js, sebagai pilihan frontend.

Hasilnya, versi desktop berjalan di dua sistem operasi dari kode yang sama. Saya tidak perlu
menulis ulang tampilan untuk Windows dan untuk Mac. Bagian ini yang paling terasa manfaat
"satu repo"-nya.

Versi mobile adalah alasan lain saya memilih Grit. Untuk saat ini saya baru menjalankan web dan
desktop, jadi saya tidak akan bercerita soal mobile sebelum benar-benar mencobanya.

### Kerja sehari-hari

Alurnya hampir selalu sama. Saat Portal PBJ butuh data baru, misalnya daftar paket atau
penyedia, saya mendeskripsikan resource-nya, menjalankan generator, lalu menyesuaikan hasilnya.
Model, API, tipe TypeScript, dan halaman admin sudah ada sebelum saya mulai menulis logika
khusus.

Sistem akun dan role juga tidak perlu saya bangun dari nol. Grit sudah menyediakan login dan
RBAC. Itu membuat kartu panel yang menyesuaikan role akun dan persetujuan akun di Manajemen
Pengguna jadi lebih ringan dikerjakan.

### Yang tetap saya kerjakan sendiri

Grit memberi kerangka, bukan isi. Bagian yang menjadi inti Portal PBJ tetap saya tulis sendiri:

- menarik data dari SPSE dan INAPROC,
- menyatukan beberapa dataset menjadi satu basis,
- panel **Analisis & Pengawasan** yang membandingkan serapan rencana dengan realisasi, menghitung
  efisiensi lelang, dan menandai risiko.

Bagian ini sebenarnya yang paling menyita waktu. Grit membuat bagian lainnya lebih cepat,
sehingga waktu saya bisa tercurah ke sini.

### Cara saya menayangkannya

Portal PBJ berjalan di **server lokal** milik saya sendiri, lalu dibuka ke internet lewat
**Cloudflare Tunnel (cloudflared) dengan Zero Trust**. Server tidak perlu membuka port ke
publik. Cloudflare yang menjadi pintu depannya, dan akses ke portal bisa diatur lewat Zero Trust.

Saya pernah menulis cara kerja tunnel ini di
[Cloudflare Tunnel: Server Lokal Berasa Publik](/tutorials/cloudflare-tunnel-server-lokal-berasa-publik/).

## Yang saya suka

- **Satu sumber kebenaran.** Perubahan di model Go ikut terlihat di sisi tampilan, jadi lebih
  jarang ada ketidaksesuaian antara backend dan frontend.
- **Kerangka yang sudah lengkap.** Auth, role, dan halaman admin tidak perlu dibuat dulu sebelum
  mulai mengerjakan fitur utama.
- **Satu basis kode untuk banyak platform.** Web dan desktop (Windows dan Mac) tidak perlu dua
  proyek terpisah.

## Yang kurang cocok atau perlu diwaspadai

Karena ini pengalaman pertama saya dengan kombinasi Go dan TanStack, ada kurva belajarnya. Saya
harus memahami cara Grit menata proyek sebelum bisa mengubahnya dengan nyaman.

Beberapa hal yang memang perlu dipertimbangkan siapa pun sebelum memakai Grit:

- Karena *opinionated*, Anda mengikuti pola Grit. Kalau butuh struktur yang sangat berbeda, akan
  terasa membatasi.
- Kode hasil generator tetap harus Anda pahami, karena Anda yang merawatnya.
- Fitur lengkap berarti banyak bagian yang ikut terpasang. Pastikan Anda benar-benar membutuhkannya.

## Cocok untuk siapa?

Menurut saya, Grit cocok untuk Anda yang:

1. ingin membangun aplikasi data atau panel admin dengan cepat,
2. nyaman dengan Go dan React,
3. ingin hasil akhirnya bisa dipasang di server sendiri.

Kalau hanya butuh situs statis seperti blog ini, Grit berlebihan. Hugo lebih tepat.

## Penutup

Buat saya, Grit berhasil menjawab alasan awalnya: satu kode untuk beberapa platform, dan
waktu yang lebih banyak untuk masalah yang sebenarnya, yaitu membuat data pengadaan mudah
dibaca. Berikutnya saya ingin mencoba versi mobile.

Harapan saya sederhana: semoga Grit Framework semakin banyak digunakan. Framework seperti ini
paling berguna kalau komunitasnya tumbuh, dengan lebih banyak contoh, dokumentasi, dan
pengalaman orang lain yang bisa dipelajari. Tulisan ini saya buat sebagai sumbangan kecil ke
arah itu.

Portal PBJ masih tumbuh. Kalau Anda ingin mencobanya, alamatnya
[portal.pbj.my.id](https://portal.pbj.my.id). Kalau penasaran dengan Grit,
mulailah dari [gritframework.dev](https://gritframework.dev/).

## Sumber

- [Grit Framework](https://gritframework.dev/) — dokumentasi dan situs resmi
- [Portal PBJ](https://portal.pbj.my.id)
- [Portal PBJ di portofolio saya](/projects/portal-pbj/)
