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

### 2. Decap CMS di `/admin` (untuk setelah deploy)

Panel admin berbasis web yang menyatu dengan situs, menyimpan lewat Git. Editornya
sudah WYSIWYG bawaan (widget `markdown` Decap = rich text dengan toggle Markdown).
Berkas konfigurasinya sudah disiapkan di `static/admin/`.

Situs di-deploy di **Netlify**, dan login `/admin` memakai **GitHub OAuth** yang
dijembatani layanan OAuth bawaan Netlify — jadi tidak perlu server tambahan.
Sekali atur, lewat tiga langkah berikut:

**1. Buat GitHub OAuth App**

Buka <https://github.com/settings/developers> → **OAuth Apps** → **New OAuth App**:

- **Application name:** bebas (mis. "Blog CMS")
- **Homepage URL:** URL situs Netlify Anda (mis. `https://ramadhan.me`)
- **Authorization callback URL:** `https://api.netlify.com/auth/done` ← wajib persis ini

Klik **Register**, catat **Client ID**, lalu **Generate a new client secret** dan catat
**Client Secret** (hanya tampil sekali).

**2. Daftarkan ke Netlify**

Di dashboard Netlify, klik proyek/situs Anda lebih dulu (bukan menu tim/akun), lalu:

**Project configuration** → **Access & security** → **OAuth** → bagian
**Authentication Providers** → **Install Provider**:

- **Provider:** GitHub
- **Client ID** & **Client Secret:** isi dari langkah 1 → simpan

> Dashboard versi lama menamai menu ini **Site configuration** → **Access control**.
> Isinya sama, hanya beda penamaan.

**3. Pakai**

Buka `https://<situs-anda>/admin/` → **Login with GitHub** → izinkan akses. Setiap kali
menyimpan di CMS, Decap membuat commit ke repo `blogramadhan/blogramadhan`, dan Netlify
otomatis rebuild situs.

> Akun GitHub yang dipakai login harus punya akses tulis ke repo tersebut.

**Uji lokal (opsional):** jalankan `npx decap-server` di terminal terpisah lalu buka
`http://localhost:1313/admin/` saat `hugo server` aktif (memakai `local_backend: true`,
tanpa perlu login GitHub).

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
│   ├── admin/           # Decap CMS (panel /admin setelah deploy)
│   └── images/          # Gambar unggahan (dashboard & Decap)
├── tools/dashboard/     # Dashboard admin lokal (Go, pustaka standar saja)
└── dashboard.sh         # Peluncur dashboard lokal
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
