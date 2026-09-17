+++
title = "Python Dasar 02: Menyiapkan Alat"
date = 2026-09-17T13:30:00+07:00
draft = false
description = "Bagian kedua seri Belajar Python Dasar. Dua cara menjalankan Python: lewat Google Colab di browser, atau dipasang di komputer sendiri. Lengkap dengan cara memilih dan apa yang harus dilakukan kalau gagal."
tags = ["python", "tutorial", "belajar-python-dasar", "pemula", "colab", "vscode"]
categories = ["Teknis"]
featured_image = "/images/python-02-hero.svg"
toc = true
+++

Ini bagian kedua dari [seri Belajar Python Dasar](/tutorials/belajar-python-dasar/). Kalau
belum baca [Bagian 01](/tutorials/python-dasar-01-kenapa-python/), sebaiknya mampir dulu ke
sana. Isinya soal kapan Python sepadan dipelajari dan kapan tidak.

Hari ini kita pasang alatnya, lalu menjalankan satu baris kode. Kalau lancar, sekitar lima
belas menit selesai.

## Pilih satu jalur dulu

Ada dua cara menjalankan Python. Lewat browser, atau dipasang di komputer Anda.

Pilih satu saja. Kalau bingung, jawab empat pertanyaan ini.

**Boleh memasang aplikasi di laptop yang Anda pakai?** Banyak laptop kantor dikunci dan cuma
bisa dipasangi aplikasi oleh bagian IT. Kalau begitu keadaannya, pakai Colab.

**Data yang akan Anda olah nanti sifatnya rahasia atau internal?** Kalau iya, jangan Colab.
Colab itu layanan Google, dan berkas yang Anda unggah ke sana keluar dari komputer kantor.
Untuk data pegawai, data pengadaan, atau apa pun yang tidak boleh keluar, pasang Python di
komputer sendiri. Ini bukan soal Google-nya nakal. Ini soal data itu memang tidak boleh
pindah tempat.

**Sering bekerja tanpa internet?** Colab butuh koneksi. Python di komputer tidak.

**Cuma ingin mencoba dulu hari ini?** Colab. Tidak ada yang dipasang, tidak ada yang perlu
dibersihkan kalau ternyata Anda tidak suka.

Saran saya: mulai dari Colab untuk belajar, lalu pasang Python di komputer waktu mulai
menyentuh data pekerjaan yang sebenarnya.

## Jalur A: Google Colab

Yang dibutuhkan cuma akun Google dan browser.

1. Buka **colab.research.google.com**
2. Masuk dengan akun Google Anda
3. Klik **File → New notebook**, atau tombol *Notebook baru*
4. Akan muncul kotak kosong. Itu namanya *cell*. Ketik di dalamnya:

```python
print("Halo")
```

5. Tekan **Shift + Enter**, atau klik tombol segitiga di kiri kotak

Tunggu sebentar. Ada titik berputar, itu Colab sedang menyiapkan komputer di sisi Google.
Setelah itu, di bawah kotak akan muncul tulisan `Halo`.

Selamat, itu program pertama Anda.

Tiga hal yang perlu Anda tahu soal Colab sebelum terlanjur nyaman:

Berkas yang Anda unggah ke Colab akan hilang waktu sesinya berakhir. Kalau besok Anda buka
lagi, berkasnya sudah tidak ada dan harus diunggah ulang. Kodenya tersimpan, berkasnya tidak.

Colab perlu internet. Mati koneksi, mati juga kerjaannya.

Dan sekali lagi, jangan unggah data yang tidak boleh keluar kantor.

## Jalur B: Di komputer sendiri

Ada dua yang dipasang. Python-nya sendiri, dan satu aplikasi untuk mengetik kode.

### Memasang Python

1. Buka **python.org/downloads**
2. Ambil versi stabil terbaru. Waktu tulisan ini dibuat, itu **Python 3.14**. Kalau ada
   versi yang di belakangnya ada tulisan `rc` atau `beta`, lewati saja. Itu versi uji coba,
   bukan untuk dipakai sehari-hari.
3. Jalankan berkas yang terunduh.

Di Windows, ada satu kotak centang di layar pertama yang bunyinya kira-kira **"Add python.exe
to PATH"**. **Centang kotak itu.** Letaknya di bawah, kecil, dan gampang terlewat.

Kalau kotak itu tidak dicentang, Python tetap terpasang, tapi komputer Anda tidak tahu di
mana mencarinya. Nanti muncul pesan error yang bikin bingung. Ini kesalahan paling umum
waktu memasang Python di Windows, dan saya juga pernah kena.

Setelah selesai, buka **Command Prompt** dan ketik:

```
python --version
```

Kalau muncul nomor versi, berarti berhasil.

### Memasang VS Code

VS Code itu aplikasi untuk mengetik kode. Gratis, buatan Microsoft.

1. Buka **code.visualstudio.com**, unduh, pasang seperti aplikasi biasa
2. Buka VS Code, klik ikon kotak-kotak di bilah kiri (*Extensions*)
3. Ketik `Python` di kotak pencarian
4. Pasang yang penerbitnya **Microsoft**

### Menjalankan kode pertama

1. Buat folder baru di komputer, misalnya `Documents\belajar-python`
2. Di VS Code, klik **File → Open Folder**, pilih folder tadi
3. Buat berkas baru, beri nama `halo.py`. Akhiran `.py` itu wajib
4. Ketik satu baris:

```python
print("Halo")
```

5. Simpan dengan **Ctrl + S**
6. Klik tombol segitiga di pojok kanan atas

Di bagian bawah layar akan muncul jendela hitam, dan di situ tertulis `Halo`.

## Sedikit soal kode tadi

`print` artinya tampilkan. Apa pun yang Anda taruh di dalam kurungnya akan muncul di layar.

Coba yang ini, ketik ketiganya lalu jalankan:

```python
print("Halo")
print(2 + 3)
print("2 + 3")
```

Hasilnya:

```
Halo
5
2 + 3
```

Perhatikan dua baris terakhir. Yang kedua tidak pakai tanda kutip, jadi Python
menghitungnya dan menampilkan `5`. Yang ketiga pakai tanda kutip, jadi Python
memperlakukannya sebagai teks biasa dan menampilkannya apa adanya.

Tanda kutip itu menentukan. Kita bahas lebih lanjut di Bagian 03 dan 04.

## Kalau tidak jalan

Empat masalah yang paling sering muncul.

**`'python' is not recognized as an internal or external command`**

Ini pesan dari Windows, artinya kotak centang PATH tadi terlewat. Jalankan lagi berkas
pemasang Python, pilih *Modify*, lalu pastikan opsi PATH-nya aktif. Atau pasang ulang dari
awal dan centang kotaknya.

**`SyntaxError: unterminated string literal`**

Tanda kutipnya kurang satu. Biasanya yang penutup.

```python
print("Halo)
```

**`SyntaxError: invalid character '“' (U+201C)`**

Ini yang sering kena kalau Anda menyalin kode dari Word, WhatsApp, atau PDF. Aplikasi itu
diam-diam mengubah tanda kutip lurus `"` menjadi tanda kutip melengkung `“` `”`. Bentuknya
mirip, tapi Python tidak menerimanya.

Hapus kutipnya, lalu ketik ulang langsung di VS Code atau Colab.

Ini alasan lain kenapa saya menyarankan mengetik ulang kode, bukan menyalin-tempel.

**`IndentationError: unexpected indent`**

Ada spasi di awal baris yang seharusnya tidak ada. Di Python, spasi di depan itu punya arti.
Hapus sampai kodenya menempel di tepi kiri.

## Latihan

Ketik dan jalankan ini, ganti isinya dengan data Anda sendiri:

```python
print("Nama saya Kurnia.")
print("Saya bekerja di Pontianak.")
print(2026 - 1990)
```

Kalau ketiganya muncul di layar, alat Anda sudah siap dan Bagian 03 tinggal dilanjutkan.

Coba juga sengaja bikin salah. Hapus satu tanda kutip, lalu jalankan. Baca pesan errornya
sampai habis. Nanti Anda akan sering ketemu pesan-pesan seperti ini, dan lebih enak kalau
sudah kenal duluan.

## Berikutnya

Di **Bagian 03** kita mulai menyimpan nilai ke dalam variabel, dan berkenalan dengan tipe
data. Di situ Python mulai terasa seperti alat kerja, bukan cuma layar yang menampilkan
tulisan.

Tautannya muncul di [halaman silabus](/tutorials/belajar-python-dasar/) begitu terbit.
