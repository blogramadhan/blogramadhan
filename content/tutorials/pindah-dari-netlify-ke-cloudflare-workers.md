+++
title = "Pindah dari Netlify ke Cloudflare Workers"
date = 2026-10-02T21:00:00+07:00
draft = false
description = "Catatan memindahkan blog Hugo ini dari Netlify ke Cloudflare Workers: langkah-langkahnya, satu build yang gagal beserta sebabnya, dan dua hal yang di Netlify otomatis tapi di Cloudflare harus disetel sendiri."
tags = ["tutorial", "cloudflare", "netlify", "hugo", "workers", "hosting", "migrasi"]
categories = ["Teknis"]
featured_image = "/images/netlify-ke-cloudflare-hero.svg"
toc = true
+++

Blog yang sedang Anda baca ini tinggal di Netlify sejak awal. Hari ini saya memindahkannya ke
Cloudflare Workers.

Pemindahannya sendiri sebentar. Yang memakan waktu justru hal-hal kecil yang di Netlify
berjalan diam-diam tanpa pernah saya pikirkan, dan ternyata di Cloudflare harus dinyalakan
satu per satu.

Tulisan ini catatan apa adanya, termasuk satu build yang gagal karena kesalahan saya sendiri.

## Kenapa pindah

Jujur saja: alasannya bukan Netlify bermasalah. Netlify bekerja baik selama ini.

Alasannya DNS saya sudah di Cloudflare. Selama ini permintaan pengunjung lewat Cloudflare
dulu, baru diteruskan ke Netlify. Memindahkan situsnya ke Cloudflare memangkas satu lompatan,
dan menyatukan semuanya di satu tempat.

Kalau DNS Anda tidak di Cloudflare, sebagian panduan ini tetap berlaku, tapi Anda perlu
memindahkan nameserver lebih dulu dan itu butuh waktu propagasi.

## Workers atau Pages

Cloudflare punya dua tempat untuk situs statis, dan keduanya ada di menu yang sama,
**Workers & Pages**.

**Pages** lebih sederhana. Tidak perlu berkas konfigurasi apa pun di repo; semua diatur dari
dasbor. Cara ini paling mirip Netlify.

**Workers dengan static assets** butuh satu berkas `wrangler.jsonc` di repo, tapi ke sanalah
Cloudflare mengarahkan pengembangannya. Dokumentasinya sendiri menyebut Workers punya
rangkaian fitur yang jauh lebih luas.

Perlu saya luruskan karena banyak yang salah kaprah: **Pages tidak sedang dihentikan.** Tidak
ada pengumuman penghentian, dan fiturnya masih dirawat. Kalau Anda ingin perubahan seminimal
mungkin, Pages pilihan yang sah.

Saya memilih Workers. Sisa tulisan ini memakai jalur itu.

## Dua berkas yang perlu disiapkan

**Pertama, `wrangler.jsonc` di akar repo.**

```jsonc
{
  "name": "blogramadhan",
  "compatibility_date": "2026-10-02",

  "assets": {
    "directory": "./public",
    "not_found_handling": "404-page"
  }
}
```

Tidak ada `"main"` di situ, dan itu disengaja. Tanpa berkas skrip, Worker ini murni menyajikan
isi folder. `not_found_handling` disetel `404-page` supaya alamat yang tidak ada disajikan
`public/404.html` bila tersedia.

Satu hal yang tidak perlu diisi: `html_handling`. Bawaannya `auto-trailing-slash`, dan itu
memang yang cocok dengan URL Hugo berbentuk `/posts/judul/`.

**Kedua, `static/_headers`** untuk menggantikan blok `[[headers]]` di `netlify.toml`.

```
/*
  X-Content-Type-Options: nosniff
  Referrer-Policy: strict-origin-when-cross-origin
```

Taruh di `static/`, bukan di akar. Hugo menyalinnya ke `public/` saat build, dan Cloudflare
membacanya dari sana. Berkasnya sendiri tidak ikut tersaji — saya sudah mengeceknya, `/_headers`
membalas 404.

## Menyambungkan ke repo

Di dasbor: **Workers & Pages → Create → Import a repository**, pilih reponya.

Pengaturan build-nya:

| Kolom | Isi |
|---|---|
| Build command | `hugo --gc --minify` |
| Deploy command | `npx wrangler deploy` |
| Build output directory | `public` |
| Variabel `HUGO_VERSION` | versi Hugo Anda |

Variabel terakhir itu penting dan mudah terlewat. Build image Cloudflare membawa Hugo versi
lama. Di percobaan pertama saya, lognya menunjukkan ia memakai `hugo v0.147.7` padahal mesin
saya 0.166.0. Situsnya tetap terbangun, tapi perilaku template dan Markdown bisa bergeser
antarversi, dan Anda tidak mau menemukannya lewat halaman yang tiba-tiba kosong.

## Build pertama saya gagal

Ini bagian yang paling berguna, jadi saya tulis lengkap.

Build Hugo-nya sukses — 209 halaman, 95 alias, selesai dalam 159 milidetik. Yang gagal
perintah deploy-nya:

```
[build] Running: npx hugo
[build] npm error could not determine executable to run
✘ [ERROR] Running custom build `npx hugo` failed.
```

Saya tidak pernah menyuruh siapa pun menjalankan `npx hugo`. Hugo bukan paket npm, jadi
wajar npm kebingungan. Pertanyaannya: dari mana perintah itu datang?

Jawabannya ada beberapa baris di atasnya:

```
📄 Create wrangler.jsonc:
  {
    "name": "blogramadhan",
    ...
  }
```

Wrangler sedang **membuat** `wrangler.jsonc`. Artinya ia tidak menemukan satu pun di repo.
Waktu tidak ada konfigurasi, wrangler menjalankan auto-configuration: ia menebak jenis
proyeknya, menyimpulkan sendiri perintah build-nya, lalu menjalankannya. Tebakannya `npx hugo`,
dan tebakan itu salah.

Penyebabnya memalukan sekaligus sederhana. **Saya belum meng-commit `wrangler.jsonc`.** Berkasnya
ada di komputer saya, tidak di repo. Cloudflare meng-clone repo, dan di sana tidak ada apa-apa.

```bash
git cat-file -e HEAD:wrangler.jsonc || echo "tidak ada di HEAD"
```

Satu baris itu akan menghemat waktu Anda kalau mengalami hal yang sama.

Saya sempat mencari opsi untuk mematikan auto-config lewat `wrangler deploy --help`, dan tidak
ada. Memang satu-satunya cara mencegahnya dengan menyediakan konfigurasinya di repo.

## Dua hal yang di Netlify otomatis

Setelah build berhasil dan situsnya tampil di URL `*.workers.dev`, saya pindahkan domainnya
lewat **Settings → Domains & Routes → Add custom domain**. Situsnya langsung hidup.

Lalu saya periksa satu per satu, dan menemukan dua hal yang ternyata tidak ikut pindah.

### 1. HTTP tidak dialihkan ke HTTPS

```bash
curl -sS -o /dev/null -D - "http://domain-anda.com/"
```

Yang saya harapkan 301. Yang saya dapat:

```
HTTP/1.1 200 OK
```

Pengunjung yang mengetik alamat tanpa `https://` dilayani lewat koneksi tanpa enkripsi. Di
Netlify ini otomatis; di Cloudflare ada sakelarnya sendiri di **SSL/TLS → Edge Certificates →
Always Use HTTPS**.

Menurut saya ini yang paling mendesak dari keduanya. Juga paling gampang terlewat, karena di
browser Anda sendiri tidak akan terasa — browser modern sudah mencoba HTTPS lebih dulu.

### 2. Alamat www tidak ada

```
curl: (6) Could not resolve host: www.domain-anda.com
```

Bukan salah setel, memang tidak ada catatan DNS-nya. Dua langkah.

**DNS → Records → Add record:** tipe `CNAME`, nama `www`, target domain apex Anda, status
**Proxied**. Awan oranyenya wajib — kalau DNS-only, permintaannya tidak lewat jaringan
Cloudflare dan aturan berikutnya tidak akan pernah jalan.

**Rules → Redirect Rules → Create rule**, mode wildcard:

| Kolom | Isi |
|---|---|
| Request URL | `https://www.*` |
| Target URL | `https://${1}` |
| Status code | `301` |
| Preserve query string | dicentang |

`${1}` itu bagian yang tertangkap wildcard. Hasilnya path ikut terbawa, jadi tautan lama ke
halaman dalam tidak jatuh ke beranda:

```
https://www.domain-anda.com/investasi/  ->  301  ->  https://domain-anda.com/investasi/
```

Setelah ini saya uji keempat pintu masuknya, dan semuanya bermuara ke alamat yang sama:

| Yang diketik | Berakhir di |
|---|---|
| `http://domain-anda.com/posts/` | `https://domain-anda.com/posts/` |
| `https://domain-anda.com/posts/` | `https://domain-anda.com/posts/` |
| `http://www.domain-anda.com/posts/` | `https://domain-anda.com/posts/` |
| `https://www.domain-anda.com/posts/` | `https://domain-anda.com/posts/` |

Yang dari `http://www` menempuh dua lompatan — dinaikkan ke HTTPS dulu, baru `www`-nya dibuang.
Itu wajar.

## Memastikan pindahnya benar

Jangan percaya pada tampilan situs yang kelihatan normal. Situs lama yang masih di-cache juga
kelihatan normal.

Cara paling meyakinkan: cari jejak host lama di header respons.

```bash
curl -sSI "https://domain-anda.com/?nocache=$RANDOM" | grep -ci netlify
```

Query acak di ujungnya untuk menembus cache. Jawaban `0` berarti tidak ada satu pun header
Netlify yang tersisa.

Lalu pastikan yang melayani memang Cloudflare:

```bash
curl -sSI "https://domain-anda.com/" | grep -iE "^server|^cf-ray"
```

Satu pemeriksaan lagi yang saya suka, karena sulit dipalsukan cache — tanggal sertifikatnya:

```bash
echo | openssl s_client -connect domain-anda.com:443 -servername domain-anda.com 2>/dev/null \
  | openssl x509 -noout -issuer -dates
```

Kalau `notBefore`-nya beberapa jam lalu, itu sertifikat yang baru diterbitkan Cloudflare, bukan
warisan host lama.

Terakhir, pastikan isinya memang hasil build terbaru. Cara termudah: buka alamat halaman yang
baru Anda tulis. Kalau halaman yang belum pernah ada di host lama sudah bisa dibuka, build-nya
memang berasal dari commit terbaru.

## Yang tidak perlu dikhawatirkan

**Alias Hugo tetap jalan.** Saya punya 95 alias dari tulisan yang pernah pindah folder. Alias
itu bukan fitur Netlify — Hugo membuatnya sendiri sebagai halaman HTML berisi `canonical` dan
`meta refresh`. Jadi ia ikut ke mana pun situsnya pindah.

Catatan kecil: alias membalas 200, bukan 301. Memang begitu bentuknya, dan sama persis seperti
waktu di Netlify.

**Tidak ada waktu mati.** Selama domainnya belum dialihkan, situs lama tetap melayani. Pindahkan
domainnya hanya setelah Anda melihat situsnya benar di URL `*.workers.dev`.

**Situs Netlify lamanya tidak hilang sendiri.** Ia masih hidup di akun Anda dan masih ikut
rebuild tiap kali Anda push. Hapus dari dasbor Netlify kalau sudah yakin, dan buang
`netlify.toml` dari repo supaya tidak membingungkan Anda sendiri setahun lagi.

## Soal observability

Cloudflare menawarkan Workers Logs, dan menambahkannya ke `wrangler.jsonc` memang gampang:

```jsonc
"observability": {
  "logs": {
    "enabled": true,
    "invocation_logs": true,
    "persist": true
  }
}
```

Tapi jangan berharap banyak kalau Worker Anda tanpa `main` seperti punya saya. Permintaan yang
dilayani langsung dari static assets **gratis dan tidak dihitung sebagai invocation**. Worker
baru benar-benar dipanggil kalau permintaannya tidak cocok dengan berkas mana pun.

Artinya hampir semua kunjungan tidak menyentuh kode Worker, jadi tidak ada yang tercatat.
Untuk angka pengunjung, tempatnya di **Analytics & Logs** tingkat domain — yang sudah tersedia
sejak DNS Anda di Cloudflare, tanpa setelan apa pun.

## Ringkasnya

1. Siapkan `wrangler.jsonc` dan `static/_headers`, lalu **commit dan push**.
2. Buat proyek Workers dari repo, isi build command dan `HUGO_VERSION`.
3. Periksa situsnya di URL `*.workers.dev`.
4. Baru pindahkan domainnya.
5. Nyalakan Always Use HTTPS.
6. Tambahkan DNS `www` dan aturan pengalihannya.
7. Buktikan pindahnya dengan `curl`, bukan dengan melihat situsnya.
8. Matikan situs lama dan buang `netlify.toml`.

Langkah pertama yang paling sering bikin celaka, dan itu karena kata "commit"-nya, bukan
berkasnya.

## Catatan

Semua angka dan pesan error di tulisan ini dari pemindahan situs ini sendiri pada 2 Oktober
2026, bukan contoh karangan. Log build yang gagal, keluaran `curl`, sampai tanggal sertifikat —
semuanya saya salin dari layar.

Satu hal yang saya tidak uji: jalur Cloudflare Pages. Saya memilih Workers sejak awal, jadi
perbandingan di seksi kedua berasal dari dokumentasi, bukan dari mencoba dua-duanya.

Tampilan dasbor Cloudflare sering berubah. Kalau nama menunya bergeser, cari kata kuncinya.

## Sumber

- [Workers static assets](https://developers.cloudflare.com/workers/static-assets/) — Cloudflare
- [Migrasi dari Pages ke Workers](https://developers.cloudflare.com/workers/static-assets/migration-guides/migrate-from-pages/) — Cloudflare
- [Berkas `_headers`](https://developers.cloudflare.com/workers/static-assets/headers/) — Cloudflare
- [Redirect www ke root](https://developers.cloudflare.com/rules/url-forwarding/examples/redirect-www-to-root/) — Cloudflare
- [Billing dan batasan static assets](https://developers.cloudflare.com/workers/static-assets/billing-and-limitations/) — Cloudflare
