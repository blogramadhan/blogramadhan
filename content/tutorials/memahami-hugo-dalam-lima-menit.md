+++
title = "Memahami Hugo dalam Lima Menit"
date = 2026-07-05T13:30:00.000Z
draft = false
description = "Peta ringkas cara kerja Hugo: dari berkas Markdown menjadi situs statis yang cepat."
tags = ["hugo", "web", "tutorial"]
categories = ["Teknis"]
featured_image = "/images/hugo-lima-menit-hero.svg"
toc = true
aliases = ["/posts/memahami-hugo-dalam-lima-menit/"]
+++

Hugo itu *static site generator*, alat yang mengubah tumpukan berkas teks jadi situs web utuh
yang siap disajikan. Tidak ada basis data, tidak ada server yang perlu berpikir tiap kali
halaman dibuka. Semuanya sudah jadi sejak awal.

Blog ini jalan di atasnya. Waktu pertama memasangnya, yang lama saya pahami bukan perintahnya
(perintahnya cuma dua) melainkan bagaimana potongan-potongannya nyambung. Jadi itu yang saya
tulis di sini.

## Konten: tempat tulisan tinggal

Semua tulisan hidup di dalam folder `content/`. Setiap berkas Markdown diawali *front matter*,
sepotong metadata di antara tanda `+++` atau `---`:

```toml
+++
title = "Judul Tulisan"
date = 2026-07-05
tags = ["hugo"]
+++
```

Di bawahnya Anda menulis dengan Markdown biasa, dan Hugo yang mengubahnya jadi HTML.

## Layout: cara tulisan ditampilkan

Folder `layouts/` berisi cetakan HTML. Hugo mencocokkan tiap halaman dengan cetakan yang tepat
lewat aturan *lookup order*. Yang paling sering dipakai:

- `_default/baseof.html` — kerangka utama seluruh halaman.
- `_default/single.html` — satu tulisan.
- `_default/list.html` — daftar tulisan dalam sebuah bagian.

Bagian *lookup order* ini yang dulu bikin saya bingung agak lama. Intinya Hugo mencari dari
yang paling khusus ke yang paling umum, dan berhenti di yang pertama ketemu. Jadi kalau sebuah
halaman tampil dengan cetakan yang salah, biasanya bukan cetakannya yang keliru, tapi ada
cetakan lain yang lebih khusus dan menang duluan.

## Assets: gaya dan skrip

Berkas di `assets/` bisa diproses lewat *Hugo Pipes*, dikecilkan, disidik-jari untuk *cache
busting*, bahkan dikompilasi. Contohnya memuat CSS:

```go-html-template
{{ $css := resources.Get "css/main.css" | minify | fingerprint }}
<link rel="stylesheet" href="{{ $css.RelPermalink }}">
```

Bedanya dengan `static/`: apa pun yang ditaruh di `static/` disalin apa adanya tanpa diproses.
Gambar dan favicon biasanya cukup di sana.

## Menjalankannya

Cukup dua perintah:

```bash
hugo server -D   # pratinjau lokal, termasuk draft
hugo             # bangun situs final ke folder public/
```

`hugo server` memuat ulang browser tiap kali berkas disimpan, jadi biasanya saya biarkan jalan
di satu tab terminal sepanjang menulis.

Itu pondasinya. Sisanya menumpuk sendiri sambil jalan.
