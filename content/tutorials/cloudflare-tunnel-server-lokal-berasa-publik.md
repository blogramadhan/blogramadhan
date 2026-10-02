+++
title = "Cloudflare Tunnel: Server Lokal yang Berasa Publik"
date = 2026-10-02T19:00:00+07:00
draft = false
description = "Memasang Cloudflare Tunnel di Ubuntu supaya server tanpa IP publik bisa diakses lewat domain sendiri — tanpa membuka port — lalu menguncinya dengan Zero Trust Access supaya tidak jadi milik seluruh internet."
tags = ["tutorial", "cloudflare", "zero-trust", "tunnel", "ubuntu", "jaringan", "keamanan", "server"]
categories = ["Teknis"]
featured_image = "/images/cf-tunnel-hero.svg"
toc = true
+++

Ada aplikasi yang jalan di server kantor Anda. Dasbor internal, Grafana, atau sesuatu yang
Anda bangun sendiri. Di jaringan lokal ia terbuka di `localhost:3000` dan semuanya baik-baik
saja.

Lalu Anda perlu membukanya dari luar.

Jalur yang biasa ditempuh semuanya punya masalah. Minta IP publik ke bagian jaringan — lama,
dan mungkin ditolak. Buka port di firewall — berarti mengundang seluruh internet mengetuk
pintu itu tiap hari. Pasang VPN — repot buat Anda, lebih repot buat orang yang mau
mengaksesnya.

**Cloudflare Tunnel** menyelesaikannya dengan membalik arah koneksi. Dan **Zero Trust Access**
memastikan yang terbuka itu tetap milik Anda, bukan milik semua orang.

## Cara kerjanya, singkat saja

Biasanya, supaya server bisa diakses dari luar, server itu harus *menunggu* dihubungi. Itu
sebabnya ada port yang dibuka, dan itu sebabnya ada yang perlu dijaga.

Tunnel membalik arahnya. Sebuah program kecil bernama `cloudflared` berjalan di server Anda,
lalu **menelepon keluar** ke jaringan Cloudflare dan membiarkan sambungan itu terbuka. Waktu
ada pengunjung, permintaannya lewat jalur yang sudah terbuka tadi.

![Perjalanan satu permintaan: pengunjung, Access memeriksa, terowongan, lalu cloudflared di server meneruskan ke aplikasi di localhost](/images/cf-tunnel-alur.svg)

Hasilnya: firewall masuk Anda tidak perlu disentuh sama sekali. Tidak ada port yang dibuka,
tidak ada IP publik yang dibutuhkan. Dari luar, orang mengakses `app.domain-anda.com` seperti
situs biasa, lengkap dengan HTTPS.

## Yang perlu disiapkan

- **Domain yang DNS-nya dikelola Cloudflare.** Ini syarat mutlak. Kalau nameserver domain
  Anda masih di tempat lain, pindahkan dulu.
- **Akun Cloudflare.** Paket gratis sudah cukup; Tunnel dan Access tersedia tanpa bayar untuk
  pemakaian wajar.
- **Server Ubuntu** dengan akses `sudo` dan sebuah aplikasi yang sudah jalan di salah satu
  port lokal.

Cara cepat memeriksa syarat pertama, dari terminal mana pun:

```bash
dig +short NS domain-anda.com
```

Kalau jawabannya berakhiran `ns.cloudflare.com`, Anda sudah siap.

## Memasang cloudflared

Pakai repositori resmi Cloudflare, bukan mengunduh `.deb` satuan — supaya ikut terbarui lewat
`apt upgrade`.

```bash
sudo mkdir -p --mode=0755 /usr/share/keyrings
curl -fsSL https://pkg.cloudflare.com/cloudflare-main.gpg \
  | sudo tee /usr/share/keyrings/cloudflare-main.gpg >/dev/null
echo 'deb [signed-by=/usr/share/keyrings/cloudflare-main.gpg] https://pkg.cloudflare.com/cloudflared any main' \
  | sudo tee /etc/apt/sources.list.d/cloudflared.list
sudo apt update && sudo apt install cloudflared
```

Periksa:

```bash
cloudflared --version
```

Kalau Anda pernah memasang ini sebelum akhir Oktober 2025, kunci GPG-nya sudah berganti.
Jalankan ulang perintah di atas supaya pembaruannya tetap masuk.

## Dua cara membuat tunnel

| Cara | Konfigurasinya di | Cocok untuk |
|---|---|---|
| **Dasbor** | Situs Cloudflare | Paling cepat, ubah rute tanpa menyentuh server |
| **Baris perintah** | Berkas di server | Bisa masuk Git, enak untuk beberapa lingkungan |

Cloudflare menganjurkan cara dasbor. Saya pakai cara baris perintah di panduan ini, karena ia
memperlihatkan apa yang sebenarnya terjadi — dan setelah paham, pindah ke dasbor jadi gampang.

Kalau Anda memilih dasbor: buka **Zero Trust → Networks → Tunnels → Create**, salin tokennya,
lalu jalankan satu baris ini di server dan lompat ke seksi Access.

```bash
sudo cloudflared service install <TOKEN>
```

## Membuat tunnel lewat baris perintah

Pertama, hubungkan server dengan akun Cloudflare Anda:

```bash
cloudflared tunnel login
```

Perintah ini menampilkan sebuah tautan. Buka di browser, pilih domain Anda, dan izinkan.
Sertifikatnya tersimpan di `~/.cloudflared/cert.pem`.

Buat tunnel, beri nama bebas:

```bash
cloudflared tunnel create kantor
```

Keluarannya memuat **UUID** tunnel dan lokasi berkas kredensialnya. Catat keduanya.

Arahkan sebuah nama domain ke tunnel itu:

```bash
cloudflared tunnel route dns kantor app.domain-anda.com
```

Satu perintah itu membuat catatan DNS-nya sekaligus. Anda tidak perlu menyentuh panel DNS.

## Menyambungkan ke aplikasi Anda

Buat berkas konfigurasi:

```bash
sudo mkdir -p /etc/cloudflared
sudo nano /etc/cloudflared/config.yml
```

Isi dengan ini, ganti UUID dan nama domainnya:

```yaml
tunnel: 6ff42ae2-765d-4adf-8112-31c55c1551ef
credentials-file: /root/.cloudflared/6ff42ae2-765d-4adf-8112-31c55c1551ef.json

ingress:
  - hostname: app.domain-anda.com
    service: http://localhost:3000
  - service: http_status:404
```

Bagian `ingress` itu daftar aturan, dibaca dari atas ke bawah, dan yang pertama cocok yang
dipakai. Baris terakhir wajib ada sebagai penampung sisa — tanpa itu `cloudflared` menolak
berjalan.

Mau beberapa aplikasi sekaligus? Tambah saja barisnya:

```yaml
ingress:
  - hostname: app.domain-anda.com
    service: http://localhost:3000
  - hostname: grafana.domain-anda.com
    service: http://localhost:9090
  - service: http_status:404
```

Periksa konfigurasinya sebelum dijalankan:

```bash
cloudflared tunnel ingress validate
```

Lalu coba jalankan di depan mata:

```bash
cloudflared tunnel run kantor
```

Buka `app.domain-anda.com` di browser. Kalau aplikasi Anda muncul, terowongannya hidup.

## Menjadikannya layanan

Selama masih dijalankan dari terminal, ia mati begitu Anda keluar. Pasang sebagai layanan:

```bash
sudo cloudflared service install
sudo systemctl enable --now cloudflared
sudo systemctl status cloudflared
```

Lihat lognya kalau perlu:

```bash
sudo journalctl -u cloudflared -f
```

## Berhenti di sini berarti Anda baru saja membuka server

Ini bagian yang paling ingin saya tekankan, dan yang paling sering dilewati tutorial lain.

Terowongan Anda sekarang jalan. Siapa pun yang tahu alamatnya — atau menemukannya dari
catatan sertifikat yang memang publik — bisa membuka aplikasi itu. Anda memang tidak membuka
port, tapi hasil akhirnya setara: aplikasi internal Anda sekarang ada di internet.

Memasang tunnel tanpa gerbang apa pun itu membuka pintu dengan cara yang lebih rapi, bukan
menguncinya. Untuk alat internal, gerbang itu Access — dan itu yang kita pasang sekarang.

## Mengunci dengan Zero Trust Access

Access berdiri di depan aplikasi Anda dan bertanya "siapa Anda" sebelum permintaannya
diteruskan. Sifat bawaannya **menolak**: tanpa kebijakan yang mengizinkan, semua orang
ditolak.

Di dasbor Cloudflare, buka **Zero Trust → Access → Applications → Add an application**, lalu
pilih **Self-hosted**.

Isi tiga hal: nama aplikasi, nama domain yang tadi (`app.domain-anda.com`), dan **session
duration** — berapa lama seseorang tetap dianggap sudah login sebelum diminta masuk lagi. Buat
dasbor internal, 24 jam biasanya masuk akal.

Lalu buat kebijakannya. Yang paling sederhana:

| Kolom | Isi |
|---|---|
| Action | **Allow** |
| Selector | **Emails ending in** |
| Value | `@instansi-anda.go.id` |

Itu artinya siapa pun dengan alamat surel di domain itu boleh masuk, dan selain itu tidak.
Kalau penggunanya cuma beberapa orang, pakai selector **Emails** dan daftarkan satu per satu.

Ada empat jenis tindakan, dan beda-beda gunanya:

- **Allow** — izinkan yang cocok. Ini yang paling sering dipakai.
- **Block** — tolak yang cocok, walau ada aturan Allow yang juga cocok.
- **Bypass** — matikan pemeriksaan sama sekali untuk lalu lintas yang cocok. Hati-hati:
  permintaannya juga tidak tercatat di log.
- **Service Auth** — untuk mesin, bukan orang. Dipakai kalau ada skrip atau layanan lain yang
  perlu memanggil aplikasi Anda.

Urutan pemeriksaannya: Service Auth dulu, lalu Bypass, lalu Allow, baru Block.

Untuk cara login, Cloudflare menyediakan **One-time PIN** tanpa perlu menyambungkan apa pun —
pengguna memasukkan surel, kodenya dikirim ke sana. Kalau instansi Anda memakai Google
Workspace atau Microsoft Entra, menyambungkannya sebagai penyedia identitas membuat
pengalamannya jauh lebih mulus.

## Mengujinya

Buka `app.domain-anda.com` di jendela penyamaran. Yang harus muncul sekarang halaman login
Cloudflare, bukan aplikasi Anda.

Masuk dengan surel yang cocok dengan kebijakan — aplikasi muncul. Coba lagi dengan surel yang
tidak cocok — Anda ditolak.

Kalau aplikasi Anda langsung muncul tanpa diminta login, berarti nama domain di Access tidak
sama persis dengan yang ada di `config.yml`. Periksa huruf per huruf.

## Kalau tidak jalan

**`Error 1016 (Origin DNS Error)`**

Tunnel-nya tidak berjalan atau tidak tersambung. Periksa dengan `sudo systemctl status
cloudflared` dan `cloudflared tunnel info kantor`.

**`502 Bad Gateway`**

Terowongannya hidup, tapi aplikasi di belakangnya tidak. Pastikan aplikasi Anda benar-benar
mendengarkan di port yang Anda tulis: `ss -tlnp | grep 3000`.

**Aplikasi Anda memakai HTTPS dengan sertifikat sendiri**

Tambahkan di bawah baris `service`:

```yaml
    originRequest:
      noTLSVerify: true
```

Ini untuk sertifikat yang Anda tanda tangani sendiri di jaringan lokal. Jangan dipakai untuk
tujuan di luar jaringan Anda.

**`Connection already registered`**

Ada proses `cloudflared` lama yang masih hidup. `sudo pkill cloudflared`, tunggu satu menit,
lalu jalankan lagi.

**Tunnel jalan tapi domainnya tidak ketemu**

Rute DNS-nya belum dibuat. Jalankan `cloudflared tunnel route dns kantor app.domain-anda.com`.

## Yang perlu diingat

**Access menjaga lalu lintas lewat Cloudflare, bukan server Anda.** Kalau aplikasi itu masih
bisa dijangkau langsung lewat IP lokal dari dalam jaringan, Access tidak menghalangi apa-apa
di jalur itu. Untuk benar-benar rapat, aplikasinya sebaiknya cuma mendengarkan di
`127.0.0.1`, bukan `0.0.0.0`.

**Sambungan yang panjang bisa terputus** waktu `cloudflared` memperbarui diri. SSH, WebSocket,
dan unggahan besar paling terasa. Di server produksi, pakai `--no-autoupdate` dan atur sendiri
kapan diperbarui.

**Satu tunnel boleh punya beberapa `cloudflared`.** Kalau layanannya penting, jalankan dua di
mesin berbeda; Cloudflare membagi lalu lintasnya sendiri.

**Access bukan satu-satunya gerbang yang benar.** Ia cocok untuk alat internal yang
penggunanya sudah Anda kenal. Kalau yang Anda terowongkan justru aplikasi yang memang
ditujukan untuk orang luar — punya halaman login sendiri, pendaftaran, dan peran pengguna —
maka login aplikasinya itulah gerbangnya, dan memasang Access di depan malah menutup pintu
bagi orang yang seharusnya masuk. Yang tidak boleh adalah tidak ada gerbang sama sekali.

**Gratisnya memang gratis**, tapi ini tetap layanan pihak ketiga. Seluruh lalu lintas Anda
lewat Cloudflare, dan mereka bisa melihatnya. Untuk data internal yang sensitif, itu
pertimbangan yang perlu dibicarakan dulu, bukan diputuskan sendiri oleh yang memasang.

## Ringkasnya

```bash
# pasang
sudo mkdir -p --mode=0755 /usr/share/keyrings
curl -fsSL https://pkg.cloudflare.com/cloudflare-main.gpg \
  | sudo tee /usr/share/keyrings/cloudflare-main.gpg >/dev/null
echo 'deb [signed-by=/usr/share/keyrings/cloudflare-main.gpg] https://pkg.cloudflare.com/cloudflared any main' \
  | sudo tee /etc/apt/sources.list.d/cloudflared.list
sudo apt update && sudo apt install cloudflared

# buat tunnel
cloudflared tunnel login
cloudflared tunnel create kantor
cloudflared tunnel route dns kantor app.domain-anda.com

# tulis /etc/cloudflared/config.yml, lalu
cloudflared tunnel ingress validate
sudo cloudflared service install
sudo systemctl enable --now cloudflared
```

Lalu pasang Access di dasbor, dan jangan berhenti sebelum itu.

## Catatan

Susunan ini saya pakai sendiri. [Portal PBJ](/projects/portal-pbj/) di
[portal.pbj.my.id](https://portal.pbj.my.id) berjalan di balik Cloudflare Tunnel — bukan di
server dengan IP publik, dan tanpa port yang dibuka. Jadi langkah-langkah di atas bukan
teori yang saya kutip dari dokumentasi lalu rapikan.

Yang perlu saya jelaskan: Portal PBJ **tidak** memakai Access di depannya, dan itu disengaja.
Lihat seksi sebelum ini soal kapan Access justru menghalangi. Untuk panduan ini saya tetap
menulis Access sebagai langkah wajib, karena kasus yang paling umum — dasbor internal, alat
kerja satu tim — memang begitu.

Perintah-perintahnya sendiri tidak saya jalankan ulang sambil menulis, jadi urutannya saya
cocokkan ke dokumentasi resmi Cloudflare per Oktober 2026. Kalau ada langkah yang berbeda di
versi `cloudflared` Anda, dokumentasinya yang benar.

Tampilan dasbor Cloudflare sering berubah. Kalau nama menunya bergeser, cari kata kuncinya —
alurnya tetap sama: buat aplikasi self-hosted, tentukan nama domainnya, tambahkan kebijakan
Allow.

## Sumber

- [Cloudflare Tunnel — dokumentasi resmi](https://developers.cloudflare.com/cloudflare-one/connections/connect-networks/)
- [Unduhan dan pemasangan cloudflared](https://developers.cloudflare.com/cloudflare-one/connections/connect-networks/downloads/)
- [Cloudflare Package Repository](https://pkg.cloudflare.com/)
- [Self-hosted application di Access](https://developers.cloudflare.com/cloudflare-one/applications/configure-apps/self-hosted-public-app/)
- [Access policies: action dan selector](https://developers.cloudflare.com/cloudflare-one/access-controls/policies/)
