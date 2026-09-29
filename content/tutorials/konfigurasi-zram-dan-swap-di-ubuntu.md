+++
title = "Konfigurasi ZRAM dan Swap di Ubuntu"
date = 2026-09-29T15:00:00+07:00
draft = false
description = "Panduan langkah demi langkah memasang swap file dan zram di Ubuntu, menyusun prioritas keduanya, dan menyetel swappiness supaya mesin tidak tercekik waktu RAM menipis."
tags = ["tutorial", "ubuntu", "linux", "zram", "swap", "memori", "server"]
categories = ["Teknis"]
featured_image = "/images/zram-swap-hero.svg"
toc = true
+++

Mesin Linux yang kehabisan RAM tidak memberi peringatan sopan. Ia melambat, lalu satu proses
mati tiba-tiba, dan di log tertulis sesuatu tentang *out of memory*.

Ada dua alat untuk mengurangi kemungkinan itu. **Swap** menitipkan sebagian isi memori ke
disk. **Zram** memadatkan isi memori dan menyimpannya tetap di RAM. Keduanya bisa dipakai
bersamaan, dan itu justru susunan yang paling masuk akal.

Panduan ini memakai Ubuntu 22.04 ke atas. Kira-kira dua puluh menit, sudah termasuk membaca.

## Bedanya apa

Bayangkan meja kerja yang penuh berkas.

**Swap** itu seperti memindahkan tumpukan berkas yang jarang dibuka ke lemari di ruang
sebelah. Meja jadi lega. Tapi kalau sewaktu-waktu berkas itu dibutuhkan, Anda harus berdiri,
jalan ke ruang sebelah, dan mencarinya. Lambat.

**Zram** itu seperti melipat berkasnya rapat-rapat lalu menaruhnya kembali di sudut meja.
Tetap di meja, jadi ambilnya cepat. Ongkosnya: Anda perlu waktu untuk melipat dan membuka
lipatan. Waktu itu diambil dari prosesor.

Karena dilipat, tempat yang dihemat lumayan. Data biasa umumnya menyusut dua sampai tiga kali
lipat. Jadi zram seukuran 2 GB sering bisa menampung 4–6 GB data.

Satu hal yang perlu diingat: zram **memakai RAM itu sendiri**. Ia bukan memori tambahan dari
langit. Anda menyerahkan sedikit tenaga prosesor untuk ditukar dengan ruang.

## Lihat dulu keadaan mesin Anda

Jangan memasang apa-apa sebelum tahu apa yang sudah ada.

```bash
free -h
swapon --show
```

`free -h` menampilkan total RAM dan swap. Kalau baris `Swap` isinya nol semua, mesin Anda
belum punya swap. Ini lumrah di VPS dan image cloud.

`swapon --show` menampilkan daftar swap yang aktif. Kalau tidak keluar apa-apa, memang belum
ada.

Cek juga apakah zram sudah jalan:

```bash
zramctl
```

Kosong berarti belum ada. Kalau perintahnya tidak ditemukan, pasang dulu `util-linux` —
tapi di Ubuntu biasanya sudah ada.

## Pilih jalurnya

Tidak semua mesin butuh dua-duanya.

| Keadaan | Saran |
|---|---|
| VPS kecil, RAM 1–2 GB, tanpa swap | Pasang keduanya. Ini yang paling terbantu. |
| Desktop RAM 8 GB ke atas, pemakaian biasa | Zram saja sudah cukup. |
| Server dengan beban yang tidak selalu terduga | Keduanya. Swap disk jadi jaring pengaman. |
| Beban kerja besar yang wajar menumpahkan memori | Swap file besar, zram kecil atau tanpa zram. |
| Database dengan latensi ketat | Hati-hati. Baca seksi terakhir dulu. |
| Disk SSD dan Anda ingin menghemat siklus tulisnya | Utamakan zram. |

Kalau ragu, susunan yang aman untuk hampir semua mesin: **zram seukuran setengah RAM, plus
swap file di disk sebagai cadangan terakhir.** Itu yang akan kita pasang di bawah.

## Membuat swap file

Empat perintah. Ganti `2G` sesuai kebutuhan; patokan ukuran ada di akhir seksi ini.

```bash
sudo fallocate -l 2G /swapfile
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile
```

Baris kedua penting. Tanpa `chmod 600`, berkas swap bisa dibaca pengguna lain — dan isinya
potongan memori aplikasi Anda. Perintah `swapon` akan menolak dan memberi peringatan kalau
izinnya terlalu longgar.

Periksa hasilnya:

```bash
swapon --show
```

Sekarang supaya tetap aktif setelah mesin dinyalakan ulang, tambahkan satu baris ke
`/etc/fstab`:

```bash
echo '/swapfile none swap sw,pri=-2 0 0' | sudo tee -a /etc/fstab
```

Bagian `pri=-2` itu angka prioritas. Kita beri nilai rendah dengan sengaja, supaya zram yang
dipakai lebih dulu. Penjelasannya sebentar lagi.

**Berapa besar?** Jawabannya berbeda tergantung Anda memasang zram atau tidak. Kalau zram
ikut dipasang, sebagian beban sudah ditampung di sana, jadi swap disknya boleh lebih kecil.

| RAM | Tanpa zram | Dengan zram setengah RAM |
|---|---|---|
| 1–2 GB | 2× RAM | Sama dengan RAM |
| 4–8 GB | Sama dengan RAM | Setengah RAM |
| 16 GB ke atas | 4–8 GB | 4 GB sudah cukup |

Contohnya server yang saya pakai menyusun panduan ini: RAM 8 GB, swap file 4 GB, zram
setengah RAM. Totalnya 8 GB ruang cadangan di atas 8 GB RAM, dan separuhnya yang cepat.

Satu pengecualian: kalau Anda ingin memakai fitur *hibernate*, swap disk harus setidaknya
sebesar RAM, berapa pun zram yang terpasang. Zram hilang waktu mesin mati, jadi ia tidak bisa
dipakai menyimpan isi memori. Di server, urusan ini biasanya tidak relevan.

## Memasang zram

Ubuntu punya dua paket untuk ini, dan **jangan pasang keduanya** karena akan berebut.

- `zram-tools` — cara lama, konfigurasinya di `/etc/default/zramswap`
- `systemd-zram-generator` — cara yang sekarang dianjurkan

Kita pakai yang kedua. Tersedia sejak Ubuntu 22.04.

```bash
sudo apt update
sudo apt install systemd-zram-generator
```

Lalu buat berkas konfigurasinya:

```bash
sudo nano /etc/systemd/zram-generator.conf
```

Isi dengan ini:

```ini
[zram0]
zram-size = ram / 2
compression-algorithm = zstd
swap-priority = 100
```

Tiga baris itu artinya:

- **`zram-size = ram / 2`** — besarnya setengah RAM. Karena isinya dipadatkan, daya tampung
  efektifnya lebih besar dari itu. Nilainya boleh ditulis sebagai ekspresi, misalnya
  `min(ram / 2, 4096)` untuk membatasi maksimum 4 GB.
- **`compression-algorithm = zstd`** — cara memadatkannya. `zstd` memberi rasio paling baik
  dengan ongkos prosesor yang masih wajar. Alternatifnya `lzo-rle`, lebih ringan tapi kurang
  padat.
- **`swap-priority = 100`** — angka prioritas, dibahas di seksi berikutnya.

Aktifkan tanpa perlu menyalakan ulang:

```bash
sudo systemctl daemon-reload
sudo systemctl start systemd-zram-setup@zram0.service
```

Periksa:

```bash
zramctl
swapon --show
```

`zramctl` menampilkan perangkat `/dev/zram0` beserta algoritma dan ukurannya. Setelah dipakai
sebentar, ia juga menunjukkan berapa banyak data yang tersimpan dibanding tempat yang
benar-benar terpakai. Dari situ terlihat rasio pemadatannya.

## Mengatur siapa dipakai duluan

Ini bagian yang paling sering terlewat, dan tanpa itu zram Anda mungkin jarang tersentuh.

Linux memakai swap berdasarkan angka prioritas. **Makin besar angkanya, makin dulu dipakai.**
Kalau semua swap punya prioritas sama, kernel membagi rata — termasuk melempar data ke disk
padahal zram masih lowong.

![Urutan pemakaian memori: RAM lebih dulu, lalu zram berprioritas 100, lalu swap file di disk berprioritas minus dua, dan terakhir OOM killer](/images/zram-swap-urutan.svg)

Susunan yang kita pasang tadi:

| Lapisan | Prioritas | Kapan dipakai |
|---|---:|---|
| zram | 100 | Begitu RAM mulai sesak |
| swap file | −2 | Kalau zram pun sudah penuh |

Pastikan dengan:

```bash
swapon --show
```

Kolom `PRIO` harus menunjukkan `100` untuk `/dev/zram0` dan `-2` untuk `/swapfile`. Kalau
angkanya belum sesuai, berarti `/etc/fstab` atau berkas konfigurasi zram belum tersimpan.

## Menyetel swappiness

`vm.swappiness` mengatur seberapa gampang kernel memindahkan data ke swap. Nilai bawaannya
`60`, dan itu disetel dengan asumsi swap ada di disk yang lambat.

Begitu swap tercepat Anda ada di dalam RAM, asumsinya berubah. Memindahkan data ke zram jauh
lebih murah daripada ke disk, jadi tidak ada alasan menahan diri.

```bash
sudo nano /etc/sysctl.d/99-zram.conf
```

Isi:

```ini
vm.swappiness = 180
vm.page-cluster = 0
```

Terapkan:

```bash
sudo sysctl --system
```

Angka `180` memang terlihat ekstrem kalau Anda terbiasa dengan saran lama "turunkan
swappiness ke 10". Saran itu benar untuk swap di *hard disk*, dan tidak lagi cocok untuk
zram. Sejak kernel 5.8 nilainya boleh sampai 200.

`vm.page-cluster = 0` mematikan pembacaan borongan. Di disk, membaca beberapa halaman
sekaligus menghemat gerakan kepala baca. Di RAM tidak ada kepala baca yang perlu digerakkan,
jadi borongan itu cuma pekerjaan tambahan.

## Memastikan semuanya jalan

Nyalakan ulang mesinnya, lalu jalankan tiga perintah ini:

```bash
free -h
swapon --show
zramctl
```

Yang Anda harapkan:

- `free -h` — baris `Swap` menunjukkan total zram dan swap file digabung.
- `swapon --show` — ada dua baris, prioritasnya 100 dan −2.
- `zramctl` — `/dev/zram0` muncul dan aktif.

Kalau ingin melihat apakah pernah ada proses yang mati kehabisan memori:

```bash
journalctl -k | grep -i "out of memory"
```

Kosong itu kabar baik.

## Kalau tidak jalan

**`fallocate: fallocate failed: Text file busy`**

Sudah ada swap file dengan nama itu dan sedang aktif. Matikan dulu: `sudo swapoff /swapfile`.

**Swap file di partisi Btrfs gagal atau bikin sistem rusak**

`fallocate` tidak cocok untuk Btrfs. Di sana swap file harus dibuat dengan cara khusus, dan
salah langkah bisa merusak berkasnya. Periksa dulu dengan `findmnt -no FSTYPE /`. Kalau
jawabannya `btrfs`, cari panduan khusus Btrfs — jangan ikuti perintah di atas.

**`swapon: /swapfile: insecure permissions 0644, 0600 suggested`**

`chmod 600` terlewat. Jalankan, lalu `swapon` lagi.

**`zramctl` kosong padahal paket sudah dipasang**

Konfigurasinya belum terbaca, atau layanannya belum dijalankan. Periksa:

```bash
systemctl status systemd-zram-setup@zram0.service
```

Pesan errornya biasanya menyebut baris mana di berkas konfigurasi yang bermasalah.

**Zram terpasang tapi tidak pernah terpakai**

Hampir selalu soal prioritas. Lihat kolom `PRIO` di `swapon --show`. Kalau zram tidak lebih
tinggi dari swap file, kernel tidak punya alasan memilihnya.

**Setelah memasang `zram-tools` dan `systemd-zram-generator` bersamaan**

Copot salah satu. `sudo apt remove zram-tools`, lalu nyalakan ulang.

## Kapan zram tidak menolong

Jujur soal batasnya, karena zram sering dijual terlalu manis.

**Kalau datanya memang tidak bisa dipadatkan.** Berkas video, gambar terkompresi, dan data
terenkripsi sudah padat dari sananya. Memadatkannya lagi cuma menghabiskan prosesor tanpa
menghemat tempat.

**Kalau prosesornya yang jadi leher botol.** Zram menukar CPU dengan RAM. Di mesin yang
prosesornya sudah tersengal, pertukaran itu merugikan.

**Kalau memang kurang RAM, bukan kurang swap.** Swap dan zram memperlambat kematian, bukan
mencegahnya. Mesin yang butuh 8 GB tidak akan nyaman di 2 GB, sebanyak apa pun swapnya.
Kalau OOM tetap sering muncul setelah semua ini, jawabannya menambah RAM atau memperkecil
bebannya.

**Untuk satu proses yang rakus dan sudah dikenal**, membatasi jatah memorinya lewat systemd
(`MemoryMax=`) biasanya lebih tepat sasaran daripada memperbesar swap.

## Ringkasnya

Seluruh isi panduan ini dalam sembilan perintah:

```bash
# swap file
sudo fallocate -l 2G /swapfile
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile
echo '/swapfile none swap sw,pri=-2 0 0' | sudo tee -a /etc/fstab

# zram
sudo apt install systemd-zram-generator
printf '[zram0]\nzram-size = ram / 2\ncompression-algorithm = zstd\nswap-priority = 100\n' \
  | sudo tee /etc/systemd/zram-generator.conf
sudo systemctl daemon-reload
sudo systemctl start systemd-zram-setup@zram0.service
```

Lalu swappiness di `/etc/sysctl.d/99-zram.conf`, dan periksa dengan `swapon --show`.

## Catatan

Susunan di panduan ini saya terapkan di sebuah server kantor: Ubuntu 24.04, RAM 8 GB, swap
file 4 GB, dan zram setengah RAM. Jadi langkah-langkahnya bukan cuma teori.

Yang belum saya ukur sendiri: rasio pemadatannya. Angka "menyusut dua sampai tiga kali lipat"
di awal tulisan berasal dari literatur, bukan dari `zramctl` di mesin itu setelah bebannya
berjalan beberapa hari. Kalau nanti sempat saya catat, seksi ini saya perbarui.

Selebihnya saya susun dari dokumentasi resmi zram-generator, manpage Ubuntu, dan wiki Debian.
Jalankan di mesin uji dulu kalau yang Anda pegang mesin produksi.

Nama paket dan ketersediaannya saya cocokkan ke daftar paket resmi Ubuntu per September 2026:
`systemd-zram-generator` ada di 22.04, 24.04, 25.10, dan 26.04. Nilai `vm.swappiness = 180`
dan `vm.page-cluster = 0` mengikuti anjuran yang lazim dipakai distribusi yang mengaktifkan
zram sejak awal; itu titik awal yang wajar, bukan angka keramat. Kalau mesin Anda terasa
aneh setelahnya, turunkan pelan-pelan dan amati.

## Sumber

- [systemd/zram-generator — repositori resmi](https://github.com/systemd/zram-generator)
- [zram-generator(8) — manpage Ubuntu](https://manpages.ubuntu.com/manpages/noble/man8/zram-generator.8.html)
- [ZRam — Debian Wiki](https://wiki.debian.org/ZRam)
- [zram — ArchWiki](https://wiki.archlinux.org/title/Zram)
- [Daftar paket zram di Ubuntu](https://packages.ubuntu.com/search?keywords=zram&searchon=names&suite=all&section=all)
