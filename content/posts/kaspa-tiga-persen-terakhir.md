+++
title = "Kaspa dan Tiga Persen Terakhir"
date = 2026-09-10T21:30:00+07:00
draft = false
description = "Saya beli KAS sekali, lalu saya diamkan. Setelah dihitung ulang, hal yang paling sering dipuji dari Kaspa — emisinya yang hampir habis — ternyata juga sumber masalah terbesarnya."
tags = ["kripto", "kaspa", "kas", "proof-of-work", "investasi", "keuangan", "blockdag"]
categories = ["Catatan"]
featured_image = "/images/kaspa-hero.svg"
toc = true
+++

Berbeda dengan [kebiasaan harian saya di ETF emas](/posts/dca-etf-emas/), KAS saya beli
**sekali saja**, lewat Bitget, lalu saya diamkan. Tidak ada jadwal, tidak ada cicilan, tidak
ada yang perlu saya buka tiap pagi.

Tulisan ini catatan sewaktu saya memeriksa apakah keputusan mendiamkan itu masih masuk akal
setelah hampir setahun. Yang saya temukan tidak seperti dugaan saya: bagian yang paling sering
dipuji dari Kaspa — emisinya yang tinggal sedikit — ternyata bersambung langsung ke risiko
terbesarnya, dan keduanya adalah angka yang sama.

## Apa yang sebenarnya saya pegang

Kaspa itu proof-of-work, seperti Bitcoin. Bedanya ada di apa yang dilakukan pada blok yang
lahir bersamaan.

Di Bitcoin, kalau dua penambang menemukan blok pada detik yang sama, satu menang dan satu
dibuang jadi *orphan* — kerja yang hangus. Itu sebabnya Bitcoin harus menjaga jarak sepuluh
menit antarblok: kalau bloknya terlalu rapat, terlalu banyak kerja terbuang dan keamanannya
justru melemah.

**GHOSTDAG** menolak membuang apa pun. Blok-blok paralel tetap dipakai, dirangkai jadi graf
berarah (*blockDAG*), lalu diurutkan belakangan berdasarkan kelompok blok mana yang secara
kolektif menghabiskan kerja komputasi terbanyak. Tidak ada blok yang hangus, jadi laju blok
boleh dinaikkan tanpa menggerus keamanannya.

Ini bukan klaim pemasaran. Protokolnya dirancang [Yonatan Sompolinsky](https://kaspa.org/lore),
orang yang menulis makalah GHOST — makalah yang dikutip di whitepaper Ethereum. Dan
peluncurannya, 7 November 2021, benar-benar bersih: **tanpa premine, tanpa presale, tanpa
jatah pendiri atau ventura.** Setiap KAS yang ada sekarang ditambang. Di ruang yang penuh
token dengan tabel alokasi tim empat puluh persen, ini langka dan pantas dihitung.

Pada 5 Mei 2025, hard-fork **Crescendo** menaikkan laju dari 1 menjadi **10 blok per detik** —
dan sejauh ini jaringannya jalan di angka itu. Peta jalannya menyebut 25, lalu 40, lalu 100
BPS pada 2027.

Jadi ya, saya masih menganggap rekayasanya rapi. Masalahnya bukan di sana.

## Tiga persen terakhir

Kaspa tidak punya halving mendadak seperti Bitcoin. Ia memakai **chromatic halving**: hadiah
blok menyusut tiap bulan dengan faktor 2<sup>−1/12</sup> ≈ 0,9439, sehingga dalam dua belas
bulan ia persis tinggal separuh. Kurvanya mulus, bukan tangga.

Karena rumusnya deterministik, hadiah blok hari ini bisa dihitung sendiri tanpa menunggu
siapa pun mengumumkannya:

| Sejak | Per blok | BPS | Per detik |
|---|---:|---:|---:|
| Mei 2022 | 440 | 1 | 440 |
| Mei 2023 | 220 | 1 | 220 |
| Mei 2024 | 110 | 1 | 110 |
| Mei 2025 | 55 → 5,5 | 1 → 10 | 55 |
| Mei 2026 | 2,75 | 10 | 27,5 |
| **Sep 2026** | **2,18** | 10 | **21,8** |

Baris Mei 2025 itu yang sering bikin orang salah baca. Crescendo membagi hadiah per blok
dengan sepuluh **dan** mengalikan jumlah bloknya dengan sepuluh, jadi laju emisi per detik
tidak berubah sama sekali. Kolom yang perlu diikuti adalah kolom paling kanan.

Sebelum melanjutkan saya uji dulu kurvanya. Bulan Juni 2026, titik akhir publik jaringan
melaporkan hadiah **2,59565436 KAS per blok**. Rumus di atas meramalkan 2,75 × 0,9439 =
**2,5957**. Cocok sampai empat angka di belakang koma, jadi sisa hitungan di bawah ini
berpijak pada dasar yang benar.

Sekarang bagian yang membuat orang bersemangat:

| | |
|---|---:|
| Pasokan beredar | 27,692 miliar KAS |
| Pasokan maksimum | 28,704 miliar KAS |
| **Sudah ditambang** | **96,5%** |
| Sisa yang belum ditambang | 1,012 miliar — 3,5% |
| Inflasi tahunan sekarang | **2,5%** |
| Inflasi tahun depan | 1,2% |

Sebagai uji silang: emisi setahun ke depan 688 juta KAS, dan karena ia meluruh dengan laju
halving tahunan, seluruh emisi yang tersisa sampai kapan pun menjumlah 688 juta ÷ ln 2 ≈
**993 juta KAS** — kira-kira sama dengan 1,01 miliar sisa pasokan yang tercatat. Dua jalan
hitung yang berbeda bertemu di angka yang sama.

Ini memang argumen yang bagus. Tekanan jual dari penambang hampir habis. Tidak ada jurang
pembukaan kunci token, tidak ada ventura yang menunggu giliran keluar. Dalam dua tahun,
inflasinya praktis nol.

Saya berhenti di sini dan merasa puas. Lalu saya sadar saya baru saja menghitung setengah
dari sebuah persamaan.

## Yang tidak saya duga

Inflasi yang menuju nol itu bukan cuma "berhentinya tekanan jual". Ia juga **satu-satunya
sumber pendapatan penambang.**

Proof-of-work tidak aman karena matematikanya indah. Ia aman karena ada orang yang membakar
listrik sungguhan untuk menjaganya, dan orang itu dibayar. Kalau bayarannya menghilang,
mesinnya pergi — dan biaya untuk menyerang jaringan pergi bersamanya.

Berapa besar bayaran itu sekarang? Emisi setahun dikali harga:

> 688 juta KAS × $0,037 = **sekitar $25,5 juta setahun**

Itulah seluruh anggaran keamanan Kaspa. Dan karena emisinya halving tiap dua belas bulan,
begini bentuknya ke depan bila harganya diam di tempat:

![Grafik batang: anggaran keamanan tahunan Kaspa menyusut dari 25,5 juta dolar pada 2026 menjadi 1,6 juta dolar pada 2030, separuh tiap tahun](/images/kaspa-anggaran-keamanan.svg)

| September | Emisi setahun | Nilai @ $0,037 |
|---|---:|---:|
| 2026 | 688 juta KAS | **$25,5 juta** |
| 2027 | 344 juta | $12,7 juta |
| 2028 | 172 juta | $6,4 juta |
| 2029 | 86 juta | $3,2 juta |
| 2030 | 43 juta | **$1,6 juta** |

Empat tahun dari sekarang, jaringan proof-of-work ini dijaga oleh anggaran seharga satu ruko
di Jakarta Selatan — kecuali harganya naik.

Dan perhatikan berapa besar kenaikan yang dibutuhkan cuma untuk **bertahan di tempat**: harga
KAS harus berlipat dua tiap tahun, terus-menerus, hanya supaya anggaran keamanannya tidak
menyusut. Bukan supaya saya untung. Supaya jaringannya sama amannya dengan hari ini.

### Bagaimana dibanding Bitcoin

Saya kira perbandingannya akan memalukan. Ternyata lebih menarik dari itu:

| | Kaspa | Bitcoin |
|---|---:|---:|
| Anggaran keamanan setahun | $25,5 juta | $12,98 miliar |
| Terhadap kapitalisasinya sendiri | **2,49%** | **0,98%** |
| Emisi berkurang separuh tiap | **12 bulan** | **4 tahun** |

Dalam angka mutlak Bitcoin membelanjakan sekitar 500 kali lipat. Tapi **relatif terhadap
ukurannya sendiri, Kaspa justru membayar dua setengah kali lebih banyak** untuk keamanannya
hari ini. Itu bukan kelemahan — itu kekuatan yang jarang disebut.

Yang membedakan bukan besarannya, melainkan **jam dindingnya.** Bitcoin punya waktu empat
tahun untuk menumbuhkan pendapatan biaya transaksi sebelum setiap pemotongan berikutnya;
persoalan "siapa yang membayar penambang setelah emisi habis" baru benar-benar menagih di
sekitar 2140. Kaspa menghadapi soal yang sama persis, dengan empat kali lebih sedikit waktu
di tiap putaran, dan tagihannya jatuh sekitar **2030** — bukan abad depan.

Satu-satunya jalan keluar dari aritmetika ini adalah biaya transaksi yang menggantikan emisi.
Dan di sinilah masalahnya menajam: biaya transaksi di Kaspa sengaja dibuat nyaris nol, dan
pada 10 BPS ada banyak sekali ruang blok kosong yang perlu diisi sebelum kelangkaan
menciptakan biaya yang berarti. **Justru throughput tinggi yang membuat pasar biaya sulit
tumbuh.** Keunggulan teknisnya dan masalah pendapatannya adalah sifat yang sama, dilihat dari
dua sisi.

Saya belum menemukan jawaban yang meyakinkan untuk ini — bukan di dokumen resminya, bukan di
forum komunitasnya. Bukan berarti jawabannya tidak ada. Tapi ini pertanyaan yang mestinya
dijawab lebih keras daripada pertanyaan berapa BPS berikutnya.

Tanda-tanda awalnya sudah terlihat. Hashrate jaringan bertahan di kisaran **416–421 PH/s**
pada pertengahan 2026, jauh di bawah puncak 700+ PH/s pada 2024. Analisis mining musim semi
lalu menemukan mayoritas ASIC Kaspa merugi pada tarif listrik $0,07/kWh. Mesin yang pergi
duluan itu bukan kebetulan.

## Yang sudah dikirim, dan yang masih dijanjikan

Karena harga sering bergerak mengikuti janji, saya pisahkan keduanya.

**Sudah jalan:**

- **Crescendo** (5 Mei 2025) — 10 BPS di mainnet, berjalan sejak itu.
- **Toccata** (30 Juni 2026) — hard-fork yang membawa aset native ke L1, *covenants* untuk
  keterprograman, verifikasi bukti zero-knowledge, dan compiler bernama Silverscript.
- **Igra L2** — rollup EVM, mainnet sejak Januari 2026, mendukung Solidity.

**Belum:**

- **DAGKnight**, konsensus adaptif dengan finalitas sub-detik — dijadwalkan setelah Toccata,
  tanggal pastinya belum ada.
- **25 → 40 → 100 BPS** — target 2027.

Satu koreksi terhadap ramai-ramai bulan Juni: Toccata **bukan** smart contract serba-guna.
Tidak ada mesin virtual global di L1 Kaspa. Yang datang adalah keterprograman terbatas dan
terarah — cukup untuk vault, dompet pintar, aset tokenisasi, kanal state, tapi bukan "Ethereum
di atas blockDAG". Kritik yang paling sering saya baca, dan saya kira adil, berbunyi kira-kira
begini: smart contract sudah dijanjikan bertahun-tahun, datangnya terlambat, dan bentuknya
lebih kecil daripada yang dibayangkan orang.

## Membeli dan menyimpannya dari Indonesia

Ini bagian yang tidak saya temukan di tulisan mana pun, dan justru paling praktis.

Dari sisi legal, KAS aman: ia tercatat sebagai aset kripto nomor **525** dalam daftar resmi
(PerBa 2/2024, 545 aset) — daftar yang disusun Bappebti sebelum pengawasan kripto berpindah ke
**OJK pada 10 Januari 2025** lewat POJK 27/2024. Jadi memperdagangkannya di Indonesia sah.

KAS juga tersedia di banyak tempat: **Bybit Indonesia, Pintu, Mobee, Pluang, Nanovest**, dan
tentu bursa luar seperti Bitget. Kelihatannya berlimpah.

Tapi waktu saya periksa satu per satu, muncul perbedaan yang jarang disebut:

> Dari daftar itu, **hanya Bitget dan Bybit Indonesia yang mengizinkan penarikan KAS ke dompet
> pribadi.** Sisanya bisa dibeli dan dijual, tapi koinnya tidak bisa dibawa keluar.

Kalau tujuan Anda menyimpannya sendiri — dan untuk aset proof-of-work yang niatnya dipegang
lama, itu tujuan yang wajar — perbedaan ini bukan detail kecil. Di platform tanpa penarikan,
yang Anda pegang secara praktis adalah **catatan saldo di pembukuan bursa**, bukan koin yang
kuncinya ada di tangan Anda. Untuk trading itu tidak apa-apa. Untuk disimpan bertahun-tahun,
itu risiko yang berbeda jenis, bukan sekadar berbeda ukuran.

Itu sebabnya saya lewat Bitget, dan itu satu-satunya alasannya.

Kebijakan penarikan bisa berubah kapan saja, ke arah mana saja. Kalau ini penting bagi Anda,
periksa sendiri di platformnya hari Anda membaca ini — jangan percaya daftar di blog orang,
termasuk blog ini.

## Harga, dan ukuran yang sebenarnya

Supaya ada kerangka:

| | |
|---|---:|
| Harga | $0,037 — sekitar Rp653 |
| Kapitalisasi pasar | $1,03 miliar — sekitar Rp18,1 triliun |
| Peringkat | #72 |
| Puncak tertinggi | $0,2074 — 31 Juli 2024 |
| Jarak dari puncak | **−82%** |
| Volume 24 jam | $30,4 juta — 3,0% dari kapitalisasi |

Dua baris terakhir yang paling banyak bicara. Turun 82% dari puncak dua tahun lalu berarti
siapa pun yang masuk di euforia 2024 masih jauh di bawah air. Dan volume harian yang cuma tiga
persen dari kapitalisasi berarti buku ordernya tipis — posisi besar tidak bisa keluar tanpa
menggerakkan harganya sendiri.

Perlu saya sebut juga bahwa saya menulis ini setelah KAS naik 29% dalam sepekan dan 45% dalam
sebulan. Itu waktu yang buruk untuk menilai apa pun secara jernih, dan saya berusaha menghitung
seolah angka-angka itu tidak ada.

## Jadi, saya jual atau tidak?

Tidak. Tapi alasannya berubah, dan itu bagian yang penting.

Saya membelinya karena rekayasanya. Setelah menghitung, saya tetap memegangnya dengan sadar
bahwa **soal yang belum selesai bukan soal teknis, melainkan soal siapa yang membayar
keamanannya setelah 2030.** Itu pertanyaan terbuka, bukan cacat yang sudah terbukti — dan
posisi saya berukuran sesuai itu: kecil, spekulatif, uang yang boleh hilang seluruhnya.

Yang berubah adalah apa yang saya pantau. Dulu saya menunggu kabar BPS berikutnya. Sekarang
tiga hal ini:

1. **Apakah pendapatan biaya transaksi mulai berarti** dibanding hadiah blok. Ini satu-satunya
   metrik yang benar-benar menjawab persoalan di atas. Selama biaya tetap nol koma sekian
   persen dari pendapatan penambang, hitungan 2030 tidak bergerak.
2. **Ke mana arah hashrate.** Kalau ia terus turun seiring emisi yang menyusut, teorinya sedang
   terbukti di depan mata.
3. **Apakah aset native dan L2 melahirkan sesuatu yang dipakai orang** — bukan sekadar ada.
   Ruang blok yang terpakai adalah asal-usul biaya transaksi, jadi butir ini dan butir pertama
   sebenarnya satu.

Kalau butir pertama tidak juga bergerak dalam dua tahun, saya akan menulis catatan lanjutan
dengan kesimpulan yang mungkin berbeda.

Mendiamkan sesuatu bukan berarti tidak memikirkannya. Bedanya, sekarang saya tahu **angka apa**
yang harus saya lihat — dan itu saja sudah membuat malam ini sepadan.

## Catatan asumsi

**Posisi saya:** saya memegang KAS, dibeli sekali lewat Bitget dan tidak ditambah sejak itu.
Nominalnya tidak saya sebut, tapi ada, dan Anda berhak tahu itu sebelum menilai tulisan ini.

**Angka pasar** diambil 10 September 2026: KAS $0,03705, kapitalisasi $1,026 miliar, beredar
27,692 miliar, maksimum 28,704 miliar, volume 24 jam $30,4 juta, puncak $0,2074 pada 31 Juli
2024. Bitcoin $79.016 dengan kapitalisasi $1,33 triliun. Kurs Rp17.630 per dolar. Semuanya
bergerak; perlakukan sebagai potret, bukan patokan.

**Hadiah blok tidak saya kutip, tapi saya turunkan** dari rumus chromatic halving
(2<sup>−1/12</sup> per bulan), lalu diuji terhadap satu pembacaan jaringan bulan Juni 2026 dan
cocok sampai empat desimal. Proyeksi anggaran keamanan 2027–2030 mengasumsikan **harga diam di
$0,037** — itu jelas tidak akan terjadi. Tabel itu bukan ramalan harga; ia menunjukkan seberapa
besar kenaikan harga yang dibutuhkan hanya agar anggaran keamanan tidak menyusut.

Perhitungan Bitcoin memakai 3,125 BTC per blok dan 52.560 blok setahun. Emisi Kaspa dihitung
dari 21,8 KAS per detik × 31.536.000 detik. Angka hashrate 416–421 PH/s dan temuan
profitabilitas ASIC berasal dari laporan pihak ketiga pertengahan 2026, bukan pengukuran saya
sendiri.

**Ketersediaan bursa dan kebijakan penarikan** adalah hasil pemeriksaan saya sendiri per
September 2026. Ini yang paling cepat basi di seluruh tulisan ini. Periksa ulang.

Ini catatan pribadi, bukan rekomendasi investasi. Kripto bisa turun 82% dari puncaknya — dan
seperti terlihat di atas, yang ini sudah.

---

**Sumber:** [Kaspa — Lore & fair launch](https://kaspa.org/kaspas-fair-launch/) ·
[Crescendo hard-fork roadmap](https://kaspa.org/crescendo-hard-fork-roadmap-10bps/) ·
[Kaspa roadmap 2026–2027](https://ourcryptotalk.com/blog/kaspa-roadmap-2026-2027) ·
[Toccata hard-fork outlook — Michael Sutton](https://medium.com/@michaelsuttonil/kaspa-covenants-toccata-hard-fork-outlook-a4d81a40900c) ·
[CoinGecko — KAS](https://www.coingecko.com/en/coins/kaspa) ·
[Kaspa mining 2026 — BT-Miners](https://bt-miners.com/kaspa-mining-in-2026-is-kas-still-profitable-complete-roi-analysis-for-every-miner/) ·
[Daftar aset kripto Bappebti](https://news.tokocrypto.com/daftar-lengkap-545-aset-kripto-legal-terdaftar-bappebti-di-indonesia/) ·
[Peralihan pengawasan kripto ke OJK](https://ojk.go.id/id/berita-dan-kegiatan/siaran-pers/Pages/Bappebti-Kemendag-Alihkan-Tugas-Aset-Keuangan-Digital-termasuk-Aset-Kripto-serta-Derivatif-Keuangan-kepada-OJK-dan-BI.aspx)
