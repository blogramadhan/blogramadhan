+++
title = "Skala yang Berhenti di 500"
date = 2026-09-12T13:30:00+07:00
draft = false
description = "Pagi ini aplikasi di ponsel saya menunjukkan angka 1770 untuk kualitas udara Pontianak. Skala resminya cuma sampai 500. Saya coba hitung sendiri dari mana angka itu datang."
tags = ["kalbar", "pontianak", "karhutla", "kabut-asap", "kualitas-udara", "gambut", "lingkungan"]
categories = ["Jurnal"]
featured_image = "/images/asap-1770-hero.svg"
toc = true
+++

Jam sepuluh pagi tadi saya buka aplikasi kualitas udara, dan angkanya **1770**.

Di bawahnya tulisan kecil: US AQI⁺. Polutan utama PM2.5, 963,0 µg/m³. Suhu 35 derajat, angin
5 km/jam, kelembapan 41 persen. Ikonnya gambar orang pakai masker gas.

Saya sudah beberapa minggu ini terbiasa melihat angka tiga digit —
[sejak Agustus](/posts/kemarau-yang-datang-berlapis/), sebetulnya. Yang bikin saya berhenti
sebentar pagi ini bukan digit keempatnya. Tanda plus kecil di sebelah "US AQI" itu.

## Ke mana perginya angka setelah 500

Skala AQI Amerika berhenti di 500. Itu ujungnya, tidak ada lanjutannya. Kategori paling gawat,
*Hazardous*, mencakup rentang 301 sampai 500, dan sesudah itu skalanya memang tidak menyediakan
apa-apa lagi. Dirancang untuk kota yang sesekali buruk.

Tanda plus itu penanda bahwa aplikasinya sudah keluar dari skala resmi dan meneruskan garisnya
sendiri. Saya penasaran meneruskannya pakai apa, jadi saya coba hitung.

Sejak revisi EPA Mei 2024, segmen teratas skala PM2.5 berjalan dari 225,5 µg/m³ (AQI 301) ke
325,4 µg/m³ (AQI 500). Kemiringannya keluar di angka 1,992 poin AQI per mikrogram. Kalau garis
itu diteruskan lurus lewat 500:

```
(963,0 − 325,4) × 1,992 + 500 = 1770
```

Persis. Dan waktu saya coba dengan angka lain yang beredar pagi ini — media lokal di sini
mengutip PM2.5 862 µg/m³ — hasilnya 1569, yang juga persis angka yang mereka sebut. Dua
pembacaan, penyedia data yang sama, cuma beda beberapa menit.

Jadi 1770 itu bukan hasil pengukuran. Itu ekstrapolasi. Garis yang dipanjangkan lewat ujung
tabelnya sendiri, dan saya tidak tahu apakah hubungan konsentrasi-ke-indeks masih punya arti di
wilayah itu. Yang saya tahu cuma aritmetikanya cocok.

ISPU punya masalah yang mirip dari arah sebaliknya: kategori Berbahaya di Indonesia mulai dari
301 dan memang tidak dikasih atap sama sekali. Awal September lalu BNPB mencatat ISPU di
Pontianak Tenggara menyentuh 1.016.

## Angka yang lebih layak dipegang

Yang sebenarnya diukur adalah baris bawahnya: **963,0 µg/m³**. Itu konsentrasi, bukan indeks.

Dua pembandingnya:

- Pedoman WHO 2021 untuk PM2.5 rata-rata 24 jam: **15 µg/m³**
- Baku mutu nasional, PP 22/2021, rata-rata 24 jam: **55 µg/m³**

963 itu sekitar **64 kali** pedoman WHO dan **17,5 kali** baku mutu kita sendiri. Juga hampir
**tiga kali** konsentrasi yang sudah bikin skala AQI-nya habis.

Perbandingannya tidak sepenuhnya setara — yang saya lihat pembacaan sesaat, sementara kedua
ambang itu untuk rata-rata 24 jam. Selisih sebesar ini tidak akan hilang cuma karena
dirata-ratakan, tapi angkanya jelas akan turun. Saya tidak tahu turun sampai berapa.

![Perbandingan konsentrasi PM2.5 Pontianak terhadap pedoman WHO, baku mutu nasional, dan ujung skala US AQI](/images/pm25-ujung-skala.svg)

## Empat puluh satu persen

Dari semua angka di layar itu, yang paling ganjil justru yang paling kecil.

Pontianak duduk persis di garis khatulistiwa. Kelembapan di sini normalnya tujuh puluh sampai
sembilan puluh persen, dan itu sebagian dari kenapa kota ini terasa seperti kota ini — udara yang
selalu terasa ada isinya. Empat puluh satu persen jam sepuluh pagi bukan angka Pontianak.

Anginnya 5 km/jam. Udara kering, panas 35 derajat, dan hampir tidak bergerak ke mana-mana. Asap
yang naik hari ini tidak punya tempat tujuan.

## Yang terbakar di bawah tanah

Sebagian besar yang terbakar adalah gambut, dan itu yang bikin urusannya panjang. Api gambut
tidak berhenti di permukaan. Ia turun ke lapisan bawah dan membara di sana, di luar jangkauan air
yang disiram dari atas. Hujan sebentar bisa memadamkan yang kelihatan tanpa menyentuh yang di
bawah, lalu asapnya naik lagi dua hari kemudian.

Per awal September, 72.009 hektare lahan gambut tercatat terbakar di enam provinsi prioritas.
Kalimantan Barat menyumbang 28.680,47 hektare — porsi terbesarnya. BMKG Supadio sempat mencatat
2.434 titik panas di Kalbar dalam 24 jam; pada 8 September angkanya turun ke 1.057.

Pemprov Kalbar memperpanjang status tanggap darurat sampai 20 September, diumumkan Wakil Gubernur
Krisantus Kurniawan dua hari lalu. Prakiraan kemaraunya sendiri baru berakhir sekitar pertengahan
Oktober.

## Yang sudah kelihatan angkanya

Kemenkes mencatat 113.336 kasus ISPA di tujuh provinsi terdampak sampai 9 September, tersebar di
63 kabupaten/kota. Yang berumur nol sampai lima tahun ada 26.545, sekitar 23 persen. Kalbar
dilaporkan 1.337 kasus.

Kemenkes menambahkan catatan yang saya hargai: angka itu tidak selalu bisa dibandingkan langsung
antar periode karena cakupan pelaporannya berbeda-beda. Laporan 27 Agustus menyebut 13.449 kasus
kumulatif. Lompatan ke 113 ribu dalam dua minggu terdengar seperti kenaikan delapan kali lipat,
dan mungkin sebagian besar memang begitu — tapi sebagian juga soal berapa banyak fasilitas
kesehatan yang sempat melapor.

Sekolah sudah lama tidak normal. Belajar dari rumah diperpanjang lagi untuk 7–11 September; di
Kalbar saja sekitar 1,1 juta siswa terdampak. Di Supadio, pesawat yang membawa Menteri Kehutanan
gagal mendarat pada 3 September karena jarak pandang tinggal sekitar 100 meter.

## Yang belum selesai

Saya tidak tahu berapa lama ini berlangsung. Prakiraannya pertengahan Oktober, tapi prakiraan
kemarau pernah meleset ke dua arah.

Saya juga tidak tahu apa akibat jangka panjang menghirup angka segini selama berminggu-minggu.
Yang tercatat sekarang ISPA, karena itu yang muncul cepat dan bisa dihitung di puskesmas. Yang
tidak masuk tabel mana pun adalah apa yang terjadi pada paru anak umur tiga tahun yang
menghabiskan September 2026 di kota ini. Kalau ada penelitian yang menelusuri kohort karhutla
Indonesia sampai jauh ke depan, saya belum ketemu.

Dan daftar yang bisa saya lakukan pendek sekali: tutup jendela, pakai masker yang benar-benar
menyaring — N95 atau KN95, masker bedah tidak menahan partikel seukuran PM2.5 — jangan olahraga
di luar, jangan bakar apa pun. Itu daftar yang terasa tidak sebanding dengan angkanya, dan saya
tidak punya tambahan.

Sekarang jam dua siang. Saya cek lagi, turun sedikit. Masih empat digit.

## Catatan asumsi

- Angka 1770 dan 963,0 µg/m³ saya ambil dari tangkapan layar aplikasi kualitas udara di ponsel
  saya, Sabtu 12 September 2026 pukul 10.00 WIB, di Pontianak. Satu titik, satu penyedia data,
  satu waktu. Pembacaan di kecamatan lain bisa beda jauh.
- Rekonstruksi rumus AQI⁺ di atas hitungan saya sendiri, bukan dokumentasi resmi. Yang saya punya
  cuma bukti bahwa aritmetikanya mereproduksi dua angka yang beredar hari ini. Itu dugaan yang
  kuat, bukan konfirmasi.
- Perbandingan 64× dan 17,5× menyandingkan pembacaan sesaat dengan ambang rata-rata 24 jam. Itu
  bukan perbandingan setara, dan saya tulis begitu supaya jelas.
- Saya bukan dokter, bukan peneliti kualitas udara, dan tidak bekerja di instansi yang menangani
  ini. Saya penduduk kota ini yang kebetulan suka mengutak-atik angka.
- Jangan percaya angka di blog orang, termasuk blog ini. Tautannya ada di bawah — kalau ada yang
  menentukan keputusan Anda, buka sendiri sumbernya.

## Sumber

- [Kabut Asap Karhutla Makin Parah, Kualitas Udara Pontianak Masuk Kategori Berbahaya](https://pontianakinformasi.co.id/news/kabut-asap-karhutla-makin-parah-kualitas-udara-pontianak-masuk-kategori-berbahaya/) — Pontianak Informasi, 12 September 2026
- [Pemprov Kalbar Perpanjang Status Darurat Asap Karhutla hingga 20 September 2026](https://koran-jakarta.com/2026-09-10/pemprov-kalbar-perpanjang-status-darurat-asap-karhutla-hingga-20-september-2026) — Koran Jakarta
- [Kasus ISPA akibat Karhutla Tembus 113 Ribu, 26 Ribu Balita Terdampak](https://periskop.id/nasional/20260910/kasus-ispa-akibat-karhutla-tembus-113-ribu-26-ribu-balita-terdampak) — Periskop, 10 September 2026
- [Update Karhutla 2026: 72.009 Hektare Lahan Gambut di 6 Provinsi Telah Hangus Terbakar](https://mediaindonesia.com/humaniora/930786/update-karhutla-2026-72009-hektare-lahan-gambut-di-6-provinsi-telah-hangus-terbakar) — Media Indonesia
- [BMKG Supadio catat 2.434 titik panas di Kalbar](https://jatim.antaranews.com/berita/1095559/bmkg-supadio-catat-2434-titik-panas-di-kalbar) — ANTARA
- [West Kalimantan ISPU surges above 1,000 as pollutants spread regionally](https://www.thestar.com.my/aseanplus/aseanplus-news/2026/09/02/west-kalimantan-ispu-surges-above-1000-as-pollutants-spread-regionally---national-disaster-mitigation-agency) — The Star, mengutip BNPB
- [Bandara Supadio Pontianak Terganggu Kabut Asap Karhutla](https://www.viva.co.id/bisnis/1926535-bandara-supadio-pontianak-terganggu-kabut-asap-karhutla-sejumlah-penerbangan-tertunda) — Viva
- [Siswa di Pontianak Sekolah Daring Imbas Asap Karhutla, Kasus ISPA Naik](https://www.cnnindonesia.com/nasional/20260811080741-20-1390848/siswa-di-pontianak-sekolah-daring-imbas-asap-karhutla-kasus-ispa-naik) — CNN Indonesia
- [Technical Assistance Document for the Reporting of Daily Air Quality (AQI)](https://www.airnow.gov/sites/default/files/2024-05/aqi-technical-assistance-document-may2024.pdf) — US EPA, Mei 2024, untuk breakpoint PM2.5
- [WHO global air quality guidelines (2021)](https://www.who.int/publications/i/item/9789240034228) — ambang PM2.5 24 jam
