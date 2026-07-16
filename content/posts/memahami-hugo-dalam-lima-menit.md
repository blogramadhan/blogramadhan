+++
title = "Memahami Hugo dalam Lima Menits"
date = 2026-07-05T13:30:00.000Z
draft = false
description = "Peta ringkas cara kerja Hugo: dari berkas Markdown menjadi situs statis yang cepat."
tags = ["hugo", "web", "tutorial"]
categories = ["Teknis"]
toc = true
+++

Hugo adalah *static site generator* — sebuah alat yang mengubah tumpukan berkas teks
menjadi situs web utuh yang siap disajikan. Tidak ada basis data, tidak ada server
yang perlu berpikir setiap kali halaman dibuka. Semuanya sudah jadi sejak awal.

Mari kita bedah bagian-bagian pentingnya.

## Konten: tempat tulisan tinggal

Semua tulisan hidup di dalam folder `content/`. Setiap berkas Markdown diawali
*front matter* — sepotong metadata di antara tanda `+++` atau `---`:

```toml
+++
title = "Judul Tulisan"
date = 2026-07-05
tags = ["hugo"]
+++
```

Di bawahnya, Anda menulis dengan Markdown biasa. Hugo yang akan mengubahnya menjadi HTML.

## Layout: cara tulisan ditampilkan

Folder `layouts/` berisi cetakan HTML. Hugo mencocokkan setiap halaman dengan cetakan
yang tepat lewat aturan *lookup order*. Yang paling sering dipakai:

- `_default/baseof.html` — kerangka utama seluruh halaman.
- `_default/single.html` — satu tulisan.
- `_default/list.html` — daftar tulisan dalam sebuah bagian.

## Assets: gaya dan skrip

Berkas di `assets/` bisa diproses lewat *Hugo Pipes* — dikecilkan, disidik-jari untuk
*cache busting*, bahkan dikompilasi. Contohnya memuat CSS:

```go-html-template
{{ $css := resources.Get "css/main.css" | minify | fingerprint }}
<link rel="stylesheet" href="{{ $css.RelPermalink }}">
```

## Menjalankannya

Cukup dua perintah yang perlu Anda ingat:

```bash
hugo server -D   # pratinjau lokal, termasuk draft
hugo             # bangun situs final ke folder public/
```

Itu saja pondasinya. Sisanya hanyalah menulis — dan Hugo akan mengurus perubahannya
menjadi halaman yang rapi dan cepat.
