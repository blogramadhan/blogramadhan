# Rizko — Jurnal & Karya

Blog pribadi bergaya editorial klasik, dibangun dengan [Hugo](https://gohugo.io)
dan tema custom (tanpa tema pihak ketiga).

## Menjalankan secara lokal

```bash
# Pratinjau dengan live-reload (termasuk draft)
hugo server -D

# Buka http://localhost:1313
```

## Membangun untuk produksi

```bash
hugo --gc --minify
# Hasil ada di folder public/
```

## Dashboard lokal (menulis lewat antarmuka)

Aplikasi kecil yang berjalan di komputer Anda dan menyunting berkas `content/` langsung —
daftar tulisan, form buat baru, editor dengan pratinjau Markdown instan.

```bash
./dashboard.sh
# Buka http://127.0.0.1:1414
```

Editor punya dua mode yang bisa diganti kapan saja:

- **Visual (WYSIWYG)** — menulis dengan format langsung terlihat, lengkap toolbar
  (tebal, miring, judul, kutipan, daftar, tautan, gambar, video, kode, garis).
  Otomatis dikonversi ke Markdown saat menyimpan.
- **Markdown** — menyunting sumber Markdown mentah dengan pratinjau instan di sampingnya.

#### Identitas situs & teks beranda

Tiga baris di beranda — pil peran, judul besar, dan paragraf pengantar — **bukan
konten**, melainkan parameter di `hugo.toml`. Karena itu dulu tidak muncul di daftar
tulisan. Sekarang tersedia di sidebar: **Situs → Identitas & Beranda**.

| Teks di beranda | Kunci `hugo.toml` |
| --- | --- |
| Pil kecil di atas judul | `params.role` |
| Judul besar | `params.tagline` |
| Paragraf pengantar | `params.intro` |

Formulirnya juga mencakup judul situs, nama penulis, deskripsi SEO, dan tautan sosial,
lengkap dengan pratinjau hero yang ikut berubah saat diketik.

Penyuntingannya bersifat *bedah*: hanya baris kunci yang bersangkutan yang diganti,
sehingga komentar, urutan, indentasi, dan bagian lain (`[menu]`, `[markup]`) tetap utuh.
Berkas ditulis lewat berkas sementara lalu di-*rename*, jadi bila gagal di tengah jalan
`hugo.toml` lama tidak ikut rusak. Nilai berisi kutip, `\`, `#`, atau baris baru
di-*escape* otomatis agar TOML tetap sah.

Sengaja dibatasi pada teks yang tampil ke pembaca — `baseURL`, `[menu]`, dan setelan
markup tetap disunting lewat berkas, karena salah ketik di sana bisa menggagalkan build.

#### Gambar

Tiga cara menyisipkan, semuanya berujung sama:

- **Seret & lepas** berkas ke editor,
- **Tempel** (Ctrl+V) langsung dari papan klip — praktis untuk tangkapan layar,
- Tombol **🖼** untuk memilih berkas atau mengambil dari **pustaka media** (gambar
  yang pernah diunggah).

Sebelum dikirim, gambar diperkecil ke maksimal **1600px** sisi terpanjang dan
dikompres ke **WebP** (kualitas 82%) — dilakukan di peramban, jadi server Go tetap
tanpa dependensi. Foto 4MB dari HP biasanya turun ke ratusan KB. Berkas disimpan ke
`static/images/` dan disisipkan sebagai `![alt](/images/nama.webp)`.

> SVG dan GIF dilewatkan apa adanya — vektor tak perlu diraster, dan animasi GIF akan
> hilang bila dikonversi.

#### Video

Tombol **▶** menerima tautan YouTube dalam bentuk apa pun (`youtube.com/watch?v=…`,
`youtu.be/…`, `shorts/…`, atau ID-nya saja) dan menyisipkan shortcode:

```
{{< youtube dQw4w9WgXcQ >}}
{{< youtube id="dQw4w9WgXcQ" caption="Keterangan" >}}
```

Pemutarnya responsif dan memakai domain `youtube-nocookie.com`.

#### Kode

Tombol **{ }** membuka dialog dengan **pilihan bahasa**. Nama bahasa itu penting: ia
ikut ke pagar Markdown (```` ```go ````) dan dipakai Hugo/Chroma untuk mewarnai
sintaks saat terbit. Tanpa itu kode terbit hitam-putih.

Di situs, tiap blok kode otomatis mendapat kepala berisi label bahasa dan tombol
**Salin**.

Tips: jalankan `hugo server -D` di terminal lain agar tombol "↗ Situs" pada dashboard
membuka pratinjau tema aslinya. Pintasan **Ctrl/Cmd+S** untuk menyimpan.

## Menulis tulisan baru

```bash
hugo new posts/judul-tulisan-anda.md
```

Lalu buka berkasnya, ubah `draft = true` menjadi `draft = false` bila sudah siap terbit.

## Menambah catatan investasi

```bash
hugo new investasi/judul-catatan.md
```

Tulisan di `content/investasi/` punya halaman daftarnya sendiri di `/investasi/` —
sejajar dengan Tutorial dan Portofolio, bukan sub-kategori Blog.

## Menambah proyek portofolio

```bash
hugo new projects/nama-proyek.md
```

Isi front matter dengan `year`, `role`, `tech`, dan opsional `repo` / `demo`.

## Struktur

```
├── hugo.toml            # Konfigurasi situs
├── archetypes/          # Cetakan front matter untuk konten baru
├── content/
│   ├── about/           # Halaman "Tentang"
│   ├── posts/           # Tulisan blog
│   ├── investasi/       # Catatan investasi
│   ├── tutorials/       # Tutorial teknis
│   └── projects/        # Portofolio
├── layouts/             # Template HTML (tema custom)
│   ├── _default/
│   │   ├── _markup/     # Render hook: blok kode (label bahasa + tombol salin)
│   │   └── …            # baseof, single, list, about
│   ├── partials/        # head, header, footer, kartu, skrip
│   ├── shortcodes/      # youtube.html (embed responsif)
│   ├── projects/        # Template khusus portofolio
│   └── index.html       # Beranda
├── assets/css/
│   ├── main.css         # Tema modern-profesional + mode gelap
│   └── syntax.css       # Pewarnaan sintaks (dibangkitkan, jangan disunting)
├── static/
│   ├── images/          # Gambar unggahan
│   └── _headers         # Header HTTP untuk Cloudflare
├── tools/dashboard/     # Dashboard lokal (Go, pustaka standar saja)
├── wrangler.jsonc       # Konfigurasi Cloudflare Workers
└── dashboard.sh         # Peluncur dashboard lokal
```

## Deploy ke Cloudflare

Situs disajikan **Cloudflare Workers static assets**, dibangun otomatis dari repo ini
lewat Workers Builds.

Pengaturan build di dasbor Cloudflare:

| Kolom | Isi |
|---|---|
| Build command | `hugo --gc --minify` |
| Deploy command | `npx wrangler deploy` |
| Build output directory | `public` |
| Variabel `HUGO_VERSION` | `0.166.0` |

Konfigurasi Worker-nya ada di `wrangler.jsonc`. Tidak ada berkas skrip — Worker ini murni
menyajikan isi `public/`.

Header HTTP diatur lewat `static/_headers`, yang disalin Hugo ke `public/_headers` saat
build.

Deploy manual dari komputer, kalau sewaktu-waktu perlu:

```bash
hugo --gc --minify
npx wrangler deploy
```

## Kustomisasi cepat

- **Identitas & sosial** → `hugo.toml`, bagian `[params]`.
- **Warna & tipografi** → token `:root` di bagian atas `assets/css/main.css`.
  Palet dashboard di `tools/dashboard/ui.html` sengaja disamakan — ubah keduanya
  bila ingin tetap seragam.
- **Menu navigasi** → `[menu]` di `hugo.toml`.
- **`baseURL`** → ganti di `hugo.toml` sebelum menerbitkan (mis. ke domain Anda).
- **Gaya pewarnaan kode** → `assets/css/syntax.css` dibangkitkan, bukan ditulis tangan.
  Untuk ganti gaya, jalankan `hugo gen chromastyles --style=NAMA` (lihat komentar di
  berkas itu) dan sesuaikan `style` di `[markup.highlight]`.
