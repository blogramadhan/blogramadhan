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

## Dashboard admin (menulis lewat antarmuka)

Tersedia dua cara mengelola tulisan tanpa menyunting Markdown secara manual.

### 1. Dashboard lokal (untuk dipakai sekarang)

Aplikasi kecil yang berjalan di komputer Anda dan menyunting berkas `content/` langsung —
daftar tulisan, form buat baru, editor dengan pratinjau Markdown instan.

```bash
./dashboard.sh
# Buka http://127.0.0.1:1414
```

Editor punya dua mode yang bisa diganti kapan saja:

- **Visual (WYSIWYG)** — menulis dengan format langsung terlihat, lengkap toolbar
  (tebal, miring, judul, kutipan, daftar, tautan, gambar, kode, garis). Otomatis
  dikonversi ke Markdown saat menyimpan.
- **Markdown** — menyunting sumber Markdown mentah dengan pratinjau instan di sampingnya.

Tips: jalankan `hugo server -D` di terminal lain agar tombol "↗ Situs" pada dashboard
membuka pratinjau tema aslinya. Pintasan **Ctrl/Cmd+S** untuk menyimpan.

### 2. Decap CMS di `/admin` (untuk setelah deploy)

Panel admin berbasis web yang menyatu dengan situs, menyimpan lewat Git. Editornya
sudah WYSIWYG bawaan (widget `markdown` Decap = rich text dengan toggle Markdown).
Berkas konfigurasinya sudah disiapkan di `static/admin/`. Untuk mengaktifkannya:

- **Uji lokal:** jalankan `npx decap-server` lalu buka `http://localhost:1313/admin/`
  saat `hugo server` aktif (memakai `local_backend: true`).
- **Di produksi:** deploy situs, lalu pilih backend Git di `static/admin/config.yml`:
  - *Netlify:* aktifkan **Identity** + **Git Gateway** (cocok dengan `name: git-gateway`).
  - *GitHub OAuth:* ganti backend menjadi `name: github` dengan `repo: user/repo`.

## Menulis tulisan baru

```bash
hugo new posts/judul-tulisan-anda.md
```

Lalu buka berkasnya, ubah `draft = true` menjadi `draft = false` bila sudah siap terbit.

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
│   └── projects/        # Portofolio
├── layouts/             # Template HTML (tema custom)
│   ├── _default/        # baseof, single, list, about
│   ├── partials/        # head, header, footer, kartu
│   ├── projects/        # Template khusus portofolio
│   └── index.html       # Beranda
├── assets/css/main.css  # Gaya elegan-klasik + mode gelap
├── static/admin/        # Decap CMS (panel /admin untuk setelah deploy)
├── tools/dashboard/     # Dashboard admin lokal (Go)
└── dashboard.sh         # Peluncur dashboard lokal
```

## Kustomisasi cepat

- **Identitas & sosial** → `hugo.toml`, bagian `[params]`.
- **Warna & tipografi** → token `:root` di bagian atas `assets/css/main.css`.
- **Menu navigasi** → `[menu]` di `hugo.toml`.
- **`baseURL`** → ganti di `hugo.toml` sebelum menerbitkan (mis. ke domain Anda).
