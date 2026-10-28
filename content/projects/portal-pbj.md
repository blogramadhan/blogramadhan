+++
title = "Portal PBJ"
date = 2026-08-28T10:00:00+07:00
draft = false
description = "Portal analisis pengadaan barang/jasa pemerintah, dari rencana hingga realisasi: penyedia, paket aktif, RUP, realisasi, daftar hitam, dan analisis pengawasan lintas dataset dalam satu panel kerja."
year = "2026"
role = "Perancang & pengembang"
tech = ["Go", "Grit", "React", "TanStack Router", "Tailwind CSS"]
demo = "https://portal.pbj.my.id"
featured_image = "/images/portal-pbj-cover.svg"
+++

**Portal PBJ** adalah portal analisis pengadaan barang/jasa pemerintah. Isinya penelusuran
penyedia, paket yang sedang berjalan, rencana anggaran, realisasi, dan analisis pengawasan,
semuanya dari satu panel kerja.

Data pengadaan sebenarnya sudah terbuka di SPSE dan INAPROC. Yang menyulitkan bentuknya: satu
paket di satu LPSE, satu rencana di satu halaman RUP, satu sanksi di satu daftar. Untuk melihat
pola, orang harus membuka puluhan tab dan menyalin angka satu per satu. Portal ini menarik data
itu dari sumber resmi secara terjadwal, menyatukannya dalam satu basis, lalu menyajikannya
sebagai panel yang bisa ditelusuri.

## Panel kerja

Setelah masuk, halaman awalnya menampilkan ringkasan (jumlah paket, LPSE terdaftar, dan
instansi), lalu kartu-kartu panel. Kartu yang muncul menyesuaikan akun.

**Menelusuri**

- **Pencarian Paket Penyedia**: profil penyedia, nilai kontrak, sebaran daerah, LPSE, dan daftar
  paket yang pernah dimenangkan.
- **Rencana Umum Pengadaan**: RUP pada tahap perencanaan, dengan nilai anggaran, cara pengadaan
  (penyedia atau swakelola), jenis, sumber dana, status PDN, dan sebarannya per instansi.
- **Paket Pengadaan Aktif**: tender dan non-tender yang sedang berjalan, lengkap dengan fase
  aktif, jenis pengadaan, dan instansi pemiliknya.

**Mengawasi**

- **Realisasi Pengadaan**: realisasi E-Katalog, Tender, Non Tender, dan Toko Daring, dilihat
  dari nilai, status, kanal, instansi, maupun penyedia.
- **Daftar Hitam**: penyedia yang terkena sanksi, status aktif atau berakhir, trennya, LPSE, dan
  detail SK yang mendasarinya.
- **Analisis & Pengawasan**: panel yang paling lama saya kerjakan, sekarang dipecah jadi empat
  bagian.

**Mengelola**

- **Vendor Management System**: katalog vendor dari sisi toko, termasuk profil, KBLI, produk,
  rasio produk yang sudah dan belum pernah dibeli, sampai rekap penyedia yang belum pernah
  dibeli di tiap kabupaten/kota.
- **Manajemen Pengguna**: khusus admin. Approval akun, role, aktivitas login, dan moderasi
  komentar serta testimoni.

## Analisis & Pengawasan

Panel ini menjahit beberapa dataset sekaligus. Isinya empat bagian:

- **Serapan**: membandingkan rencana di RUP dengan realisasi lintas kanal, lalu menelusuri
  selisihnya dan apakah kode RUP-nya cocok.
- **Efisiensi**: penghematan tender yang sudah berkontrak dibandingkan pagu anggaran. Ukurannya
  pagu karena HPS tidak tersedia di sumber datanya.
- **Mitigasi Risiko**: menandai peserta tunggal, kontrak yang nilainya mendekati pagu, dominasi
  satu penyedia, dan penyedia yang cocok dengan daftar hitam.
- **PDN & UMK-K**: porsi nilai produk dalam negeri pada realisasi, dan populasi penyedia UMK-K di
  katalog nasional.

Tanda di Mitigasi Risiko bukan tuduhan. Peserta tunggal atau kontrak mendekati pagu bisa punya
penjelasan yang wajar. Gunanya menunjukkan paket mana yang layak dilihat lebih dulu.

## Akun dan akses

Portal ini memakai akun, dan akun baru perlu disetujui admin dulu. Ada dua jenis akun,
**Pemerintah** dan **Penyedia**. Akun penyedia hanya melihat Paket Pengadaan Aktif dan Daftar
Hitam, karena dua panel itu yang relevan untuk mereka. Selain pengguna biasa, ada role
**Auditor** dan **Admin**.

PIN masuk dikirim lewat email, dan bisa juga lewat bot Telegram setelah nomor HP dihubungkan.
Sesi yang dibiarkan diam selama 30 menit akan keluar sendiri, dengan peringatan dua menit
sebelumnya.

Di halaman depan sekarang ada **Berita & Update** untuk pengumuman, fitur baru, dan pembaruan
data, serta **Suara Pengguna** untuk testimoni. Komentar baru tampil setelah disetujui admin.

## Sumber data

Seluruh isinya berasal dari sistem resmi:

- **SPSE/LPSE (spse.inaproc.id)**: pemenang tender dan non-tender, paket aktif, nilai HPS,
  penawaran, hasil negosiasi, peserta, dan instansi penyelenggara.
- **E-Katalog/E-Purchasing (data.inaproc.id)**: realisasi transaksi E-Katalog 6.0, penyedia,
  instansi pembeli, nilai, status, metode, sumber dana, dan nilai PDN.

Datanya diambil otomatis secara terjadwal. Tidak semua paket di semua LPSE bisa terjangkau,
karena sebagian LPSE kadang punya keterbatasan akses pada saat pengambilan. Jadi angka di portal
bisa berbeda dengan portal resmi, entah karena waktu pengambilannya beda atau sumbernya belum
diperbarui.

Portal ini alat bantu analisis yang independen. Ia tidak terafiliasi dengan LKPP atau instansi
pemerintah mana pun, dan tidak mengarang angka di luar yang tersedia di sumber. Untuk laporan
resmi, pengaduan, atau audit formal, tetap konfirmasi ke sumber resminya.

## Cara kerjanya

Portal PBJ dibangun ulang dengan [Grit](https://gritframework.dev/): Go di sisi server, React
dengan TanStack Router di sisi tampilan. Login, role, dan halaman admin datang dari kerangka
Grit, sehingga waktu saya bisa dipakai untuk bagian intinya: menarik data, menyatukan dataset,
dan panel analisisnya.

Servernya berjalan di mesin lokal saya sendiri dan dibuka ke internet lewat Cloudflare Tunnel.
Ceritanya lebih panjang di
[Membangun Portal PBJ dengan Grit Framework](/posts/membangun-portal-pbj-dengan-grit-framework/).

## Catatan

Yang paling menarik dari mengerjakan ini adalah menyadari betapa banyak pertanyaan sederhana
yang sebelumnya sulit dijawab. "Penyedia mana yang menang paling banyak di daerah ini?" atau
"berapa rencana yang benar-benar terealisasi tahun ini?" Dua-duanya seharusnya cukup sekali
klik, dan sekarang memang begitu.

Portal ini masih tumbuh. Beberapa panel sudah mapan, beberapa lain masih saya rapikan.
