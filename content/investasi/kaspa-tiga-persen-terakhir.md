+++
title = "Kaspa dan Tiga Persen Terakhir"
date = 2026-09-10T21:30:00+07:00
draft = false
description = "Saya beli KAS sekali, lalu saya diamkan. Waktu hitungannya saya buka lagi, ternyata hal yang paling sering dipuji dari Kaspa dan hal yang paling membuat saya waswas itu angka yang sama."
tags = ["kripto", "kaspa", "kas", "proof-of-work", "investasi", "keuangan", "blockdag"]
categories = ["Catatan"]
featured_image = "/images/kaspa-hero.svg"
toc = true
+++

KAS saya beli sekali, lewat Bitget, lalu saya diamkan. Tidak seperti
[ETF emas yang saya cicil tiap hari bursa](/posts/dca-etf-emas/), di sini tidak ada jadwal apa
pun. Beli, sudah, tutup aplikasi.

Malam ini saya buka lagi hitungannya, sebetulnya cuma karena penasaran apakah mendiamkannya
masih masuk akal setelah hampir setahun.

Ternyata ada yang luput. Angka yang paling sering dipakai orang untuk memuji Kaspa — pasokannya
yang hampir habis ditambang — adalah angka yang sama dengan yang sekarang membuat
saya agak waswas. Saya perlu beberapa hari untuk sadar keduanya sisi dari benda yang sama.

## Apa yang sebenarnya saya pegang

Kaspa itu proof-of-work, sama seperti Bitcoin. Yang beda cuma satu hal, tapi hal itu besar:
apa yang dilakukan terhadap blok yang lahir barengan.

Di Bitcoin, kalau dua penambang menemukan blok pada detik yang sama, satu menang dan satu
dibuang. Jadi *orphan* — kerja yang hangus. Itu sebabnya Bitcoin menjaga jarak sepuluh menit
antarblok. Kalau bloknya terlalu rapat, terlalu banyak kerja terbuang, dan keamanannya malah
melemah.

**GHOSTDAG** tidak membuang apa-apa. Blok-blok paralel tetap dipakai, disusun jadi graf
berarah (*blockDAG*), lalu diurutkan belakangan berdasarkan kelompok blok mana yang secara
kolektif paling banyak menghabiskan kerja komputasi. Karena tidak ada yang hangus, laju
bloknya boleh dinaikkan tinggi-tinggi.

Saya perlu membaca penjelasan itu beberapa kali sebelum benar-benar masuk.

Dan protokolnya bukan karangan anonim. Yang menulis [Yonatan Sompolinsky](https://kaspa.org/lore) —
orang di balik makalah GHOST yang dikutip di whitepaper Ethereum. Peluncurannya,
[7 November 2021](https://kaspa.org/kaspas-fair-launch/), juga bersih: tanpa premine, tanpa
presale, tanpa jatah pendiri. Semua KAS yang beredar sekarang ditambang. Di ruang yang isinya
token dengan tabel alokasi tim empat puluh persen, itu langka.

Lalu pada 5 Mei 2025, hard-fork
[Crescendo](https://kaspa.org/crescendo-hard-fork-roadmap-10bps/) menaikkan laju dari 1 jadi
**10 blok per detik**, dan sejak itu jaringannya jalan di angka segitu. Peta jalannya menyebut
25, lalu 40, lalu 100 BPS di 2027.

Sampai sini saya masih suka. Rekayasanya rapi, dan masalahnya memang bukan di situ.

## Tiga persen terakhir

Kaspa tidak punya halving mendadak. Ia pakai **chromatic halving**: hadiah blok menyusut
sedikit tiap bulan, faktornya 2<sup>−1/12</sup> ≈ 0,9439, jadi dalam dua belas bulan pas
tinggal separuh. Kurvanya mulus, bukan tangga.

Karena rumusnya pasti, hadiah blok hari ini bisa dihitung sendiri. Itu yang saya lakukan,
soalnya angka yang beredar di internet berantakan. Ada yang menulis 55 KAS per blok, ada yang
110. Dua-duanya angka lama dari sebelum Crescendo.

| Sejak | Per blok | BPS | Per detik |
|---|---:|---:|---:|
| Mei 2022 | 440 | 1 | 440 |
| Mei 2023 | 220 | 1 | 220 |
| Mei 2024 | 110 | 1 | 110 |
| Mei 2025 | 55 → 5,5 | 1 → 10 | 55 |
| Mei 2026 | 2,75 | 10 | 27,5 |
| **Sep 2026** | **2,18** | 10 | **21,8** |

Baris Mei 2025 itu yang bikin orang salah baca — saya juga waktu pertama. Crescendo membagi
hadiah per blok dengan sepuluh sekaligus mengalikan jumlah bloknya dengan sepuluh, jadi laju
emisi per detiknya tidak berubah sama sekali. Kolom yang perlu diikuti yang paling kanan.

Sebelum lanjut saya uji dulu kurvanya. Juni 2026, titik akhir publik jaringan melaporkan
hadiah **2,59565436 KAS per blok**. Rumus di atas meramalkan 2,75 × 0,9439 = **2,5957**. Cocok
sampai empat angka di belakang koma, jadi sisa hitungan di bawah berpijak di tempat yang
benar.

Sekarang bagian yang bikin orang bersemangat:

| | |
|---|---:|
| Pasokan beredar | 27,692 miliar KAS |
| Pasokan maksimum | 28,704 miliar KAS |
| **Sudah ditambang** | **96,5%** |
| Sisa yang belum ditambang | 1,012 miliar — 3,5% |
| Inflasi tahunan sekarang | **2,5%** |
| Inflasi tahun depan | 1,2% |

Saya uji silang sekali lagi, karena saya tidak begitu percaya pada hitungan saya sendiri.
Emisi setahun ke depan 688 juta KAS. Karena ia meluruh dengan laju halving tahunan, seluruh
emisi yang tersisa sampai kapan pun menjumlah 688 juta ÷ ln 2 ≈ **993 juta KAS**. Sisa pasokan
yang tercatat 1,01 miliar. Cukup dekat.

Ini memang argumen yang bagus — dan saya termasuk yang terbujuk olehnya waktu beli. Tekanan
jual dari penambang hampir habis. Tidak ada jurang pembukaan kunci token, tidak ada ventura
yang antre keluar. Dua tahun lagi inflasinya praktis nol.

Saya sempat berhenti di situ dan merasa cukup. Baru beberapa hari kemudian saya sadar saya
cuma menghitung separuh dari sesuatu.

## Yang tidak saya duga

Inflasi yang menuju nol itu bukan cuma berhentinya tekanan jual. Ia juga satu-satunya sumber
pendapatan penambang.

Proof-of-work aman bukan karena matematikanya indah. Ia aman karena ada orang yang membakar
listrik sungguhan untuk menjaganya, dan orang itu dibayar. Kalau bayarannya habis, mesinnya
pergi — dan ongkos untuk menyerang jaringan ikut pergi bersamanya.

Berapa bayarannya sekarang? Emisi setahun dikali harga.

> 688 juta KAS × $0,037 = **sekitar $25,5 juta setahun**

Itu seluruh anggaran keamanan Kaspa. Bukan sebagian — seluruhnya. Dan karena emisinya halving
tiap dua belas bulan, begini bentuknya ke depan kalau harganya diam di tempat:

![Grafik batang: anggaran keamanan tahunan Kaspa menyusut dari 25,5 juta dolar pada 2026 menjadi 1,6 juta dolar pada 2030, separuh tiap tahun](/images/kaspa-anggaran-keamanan.svg)

| September | Emisi setahun | Nilai @ $0,037 |
|---|---:|---:|
| 2026 | 688 juta KAS | **$25,5 juta** |
| 2027 | 344 juta | $12,7 juta |
| 2028 | 172 juta | $6,4 juta |
| 2029 | 86 juta | $3,2 juta |
| 2030 | 43 juta | **$1,6 juta** |

Empat tahun lagi, jaringan proof-of-work ini dijaga anggaran satu koma enam juta dolar
setahun — kira-kira seharga satu ruko.

Dan lihat berapa kenaikan harga yang dibutuhkan cuma supaya angkanya tidak turun. KAS harus
berlipat dua tiap tahun, terus-menerus. Bukan supaya saya untung. Supaya jaringannya sama
amannya dengan hari ini.

### Bagaimana dibanding Bitcoin

Saya kira perbandingannya bakal memalukan. Ternyata tidak.

| | Kaspa | Bitcoin |
|---|---:|---:|
| Anggaran keamanan setahun | $25,5 juta | $12,98 miliar |
| Terhadap kapitalisasinya sendiri | **2,49%** | **0,98%** |
| Emisi berkurang separuh tiap | **12 bulan** | **4 tahun** |

Dalam angka mutlak Bitcoin membelanjakan sekitar 500 kali lipat. Tapi relatif terhadap
ukurannya sendiri, Kaspa justru membayar dua setengah kali lebih banyak untuk keamanannya hari
ini. Ini jarang disebut orang, dan buat saya ini poin yang menguntungkan Kaspa.

Yang beda bukan besarannya, tapi kecepatan jamnya. Bitcoin punya empat tahun untuk menumbuhkan
pendapatan biaya transaksi sebelum tiap pemotongan berikutnya, dan pertanyaan "siapa yang
membayar penambang setelah emisi habis" baru benar-benar menagih di sekitar 2140. Kaspa
menghadapi pertanyaan yang sama persis dengan waktu empat kali lebih sedikit di tiap putaran,
dan tagihannya jatuh sekitar 2030 — bukan abad depan.

Jalan keluarnya cuma satu: biaya transaksi harus menggantikan emisi. Dan di sinilah saya
mentok. Biaya transaksi di Kaspa sengaja dibuat nyaris nol, lalu pada 10 BPS ruang bloknya
berlimpah. Ruang yang berlimpah tidak menciptakan kelangkaan, dan tanpa kelangkaan tidak ada
biaya yang berarti. Jadi throughput tinggi yang jadi daya tarik utamanya itu justru yang
membuat pasar biaya sulit tumbuh.

Saya tidak menemukan jawaban yang meyakinkan untuk ini. Sudah saya cari di dokumen resminya
dan di forum komunitasnya, dan yang saya temukan kebanyakan keyakinan, bukan hitungan. Mungkin
saya yang kurang jauh mencarinya. Tapi menurut saya ini pertanyaan yang mestinya lebih sering
ditanyakan daripada pertanyaan berapa BPS berikutnya.

Tanda awalnya sudah kelihatan. Hashrate jaringan bertahan di kisaran **416–421 PH/s** pada
pertengahan 2026, jauh di bawah puncak 700+ PH/s pada 2024.
[Analisis mining musim semi lalu](https://bt-miners.com/kaspa-mining-in-2026-is-kas-still-profitable-complete-roi-analysis-for-every-miner/)
menemukan mayoritas ASIC Kaspa merugi di tarif listrik $0,07/kWh. Mesin-mesin itu tidak pergi
tanpa sebab.

## Yang sudah dikirim, dan yang masih dijanjikan

Harga sering bergerak mengikuti janji, jadi saya pisahkan mana yang sudah ada dan mana yang
belum.

**Sudah jalan:**

- **Crescendo** (5 Mei 2025) — 10 BPS di mainnet, berjalan sejak itu.
- **Toccata** (30 Juni 2026) — hard-fork yang membawa aset native ke L1, *covenants* untuk
  keterprograman, verifikasi bukti zero-knowledge, dan compiler bernama Silverscript.
- **Igra L2** — rollup EVM, mainnet sejak Januari 2026, mendukung Solidity.

**Belum:**

- **DAGKnight** — konsensus adaptif dengan finalitas sub-detik. Dijadwalkan setelah Toccata,
  tanggal pastinya belum ada.
- **25 → 40 → 100 BPS** — target 2027.

Satu koreksi buat ramai-ramai bulan Juni kemarin:
[Toccata](https://medium.com/@michaelsuttonil/kaspa-covenants-toccata-hard-fork-outlook-a4d81a40900c)
bukan smart contract serba-guna. Tidak ada mesin virtual global di L1 Kaspa. Yang datang
keterprograman yang terbatas dan terarah — cukup untuk vault, dompet pintar, aset tokenisasi,
kanal state, tapi bukan "Ethereum di atas blockDAG".

Kritik yang paling sering saya baca kira-kira begini: smart contract sudah dijanjikan
bertahun-tahun, datangnya telat, dan bentuknya lebih kecil daripada yang dibayangkan orang.
Saya kira itu adil.

## Membeli dan menyimpannya dari Indonesia

Bagian ini yang paling praktis, dan anehnya paling jarang ditulis.

Dari sisi legal aman. KAS tercatat sebagai aset kripto nomor **525** di
[daftar resmi](https://news.tokocrypto.com/daftar-lengkap-545-aset-kripto-legal-terdaftar-bappebti-di-indonesia/)
(PerBa 2/2024, 545 aset) — daftar yang disusun Bappebti sebelum pengawasan kripto
[pindah ke OJK pada 10 Januari 2025](https://ojk.go.id/id/berita-dan-kegiatan/siaran-pers/Pages/Bappebti-Kemendag-Alihkan-Tugas-Aset-Keuangan-Digital-termasuk-Aset-Kripto-serta-Derivatif-Keuangan-kepada-OJK-dan-BI.aspx)
lewat POJK 27/2024. Jadi memperdagangkannya di sini sah.

Tempatnya juga banyak: **Bybit Indonesia, Pintu, Mobee, Pluang, Nanovest**, ditambah bursa luar
seperti Bitget. Kelihatannya berlimpah.

Tapi waktu saya periksa satu per satu, ada perbedaan yang tidak muncul di halaman depan mana
pun:

> Dari daftar itu, **hanya Bitget dan Bybit Indonesia yang mengizinkan penarikan KAS ke dompet
> pribadi.** Sisanya bisa dibeli dan dijual, tapi koinnya tidak bisa dibawa keluar.

Kalau niat Anda menyimpannya sendiri — dan untuk aset proof-of-work yang mau dipegang lama itu
niat yang wajar — perbedaan ini bukan detail kecil. Di platform yang tidak mengizinkan
penarikan, yang Anda pegang sebetulnya catatan saldo di pembukuan bursa. Buat trading tidak
masalah. Buat disimpan bertahun-tahun, itu jenis risiko yang lain sama sekali.

Itu alasan saya lewat Bitget. Cuma itu, tidak ada alasan lain.

Kebijakan penarikan bisa berubah kapan saja dan ke arah mana saja. Kalau ini penting buat
Anda, periksa sendiri di platformnya hari Anda membaca ini. Jangan percaya daftar di blog
orang, termasuk blog ini.

## Harga, dan ukuran yang sebenarnya

Supaya ada kerangkanya:

| | |
|---|---:|
| Harga | $0,037 — sekitar Rp653 |
| Kapitalisasi pasar | $1,03 miliar — sekitar Rp18,1 triliun |
| Peringkat | #72 |
| Puncak tertinggi | $0,2074 — 31 Juli 2024 |
| Jarak dari puncak | **−82%** |
| Volume 24 jam | $30,4 juta — 3,0% dari kapitalisasi |

Dua baris terakhir yang paling banyak bicara. Turun 82% dari puncak dua tahun lalu artinya
siapa pun yang masuk waktu euforia 2024 masih jauh di bawah air. Dan volume harian yang cuma
tiga persen dari kapitalisasi artinya buku ordernya tipis, jadi posisi besar tidak bisa keluar
tanpa menggerakkan harganya sendiri.

Saya juga perlu jujur bahwa saya menulis ini setelah KAS naik 29% dalam sepekan dan 45% dalam
sebulan. Itu waktu yang buruk untuk menilai apa pun dengan kepala dingin. Saya sudah berusaha
menghitung seolah-olah angka itu tidak ada, tapi silakan diskon sendiri.

## Jadi, saya jual atau tidak?

Tidak. Tapi alasannya sudah berubah, dan itu yang menurut saya penting.

Dulu saya pegang karena rekayasanya. Sekarang saya pegang sambil tahu bahwa soal yang belum
selesai itu bukan soal teknis, tapi soal siapa yang membayar keamanannya setelah 2030. Itu
pertanyaan terbuka, belum jadi cacat yang terbukti, dan posisi saya ukurannya menyesuaikan
itu: kecil, spekulatif, uang yang boleh hilang semua.

Yang berubah cuma apa yang saya pantau. Dulu saya menunggu kabar BPS berikutnya — yang
ternyata bukan angka yang penting-penting amat. Sekarang tiga ini:

1. **Apakah pendapatan biaya transaksi mulai berarti** dibanding hadiah blok. Ini satu-satunya
   metrik yang benar-benar menjawab persoalan di atas. Selama biaya masih nol koma sekian
   persen dari pendapatan penambang, hitungan 2030 tidak bergerak.
2. **Ke mana arah hashrate.** Kalau ia terus turun seiring emisi yang menyusut, teorinya sedang
   terbukti di depan mata.
3. **Apakah aset native dan L2 melahirkan sesuatu yang benar-benar dipakai orang**, bukan
   sekadar ada. Ruang blok yang terpakai itu asal-usul biaya transaksi, jadi butir ini dan
   butir pertama sebetulnya satu.

Kalau butir pertama tidak juga bergerak dalam dua tahun, saya akan menulis catatan lanjutan,
dan kesimpulannya mungkin lain.

Koinnya tetap saya diamkan seperti sebelumnya. Bedanya, sekarang saya tahu angka mana yang
harus dilihat kalau suatu saat mau berubah pikiran.

## Catatan asumsi

**Posisi saya:** saya memegang KAS, dibeli sekali lewat Bitget dan tidak ditambah sejak itu.
Nominalnya tidak saya sebut, tapi ada, dan Anda berhak tahu itu sebelum menilai tulisan ini.

**Angka pasar** diambil 10 September 2026 dari [CoinGecko](https://www.coingecko.com/en/coins/kaspa):
KAS $0,03705, kapitalisasi $1,026 miliar, beredar 27,692 miliar, maksimum 28,704 miliar, volume
24 jam $30,4 juta, puncak $0,2074 pada 31 Juli 2024. Bitcoin $79.016 dengan kapitalisasi $1,33
triliun. Kurs Rp17.630 per dolar. Semuanya bergerak, jadi perlakukan sebagai potret.

**Hadiah blok tidak saya kutip, saya turunkan** dari rumus chromatic halving
(2<sup>−1/12</sup> per bulan), lalu saya uji ke satu pembacaan jaringan bulan Juni 2026 dan
cocok sampai empat desimal. Proyeksi anggaran keamanan 2027–2030 mengasumsikan harga diam di
$0,037 — itu jelas tidak akan terjadi. Tabel itu bukan ramalan harga. Ia menunjukkan seberapa
besar kenaikan harga yang dibutuhkan hanya supaya anggaran keamanannya tidak menyusut.

Hitungan Bitcoin memakai 3,125 BTC per blok dan 52.560 blok setahun. Emisi Kaspa dihitung dari
21,8 KAS per detik dikali 31.536.000 detik. Angka hashrate 416–421 PH/s dan temuan
profitabilitas ASIC berasal dari laporan pihak ketiga pertengahan 2026, bukan pengukuran saya
sendiri.

**Ketersediaan bursa dan kebijakan penarikan** hasil pemeriksaan saya sendiri per September
2026. Ini yang paling cepat basi di seluruh tulisan ini. Periksa ulang.

Ini catatan pribadi, bukan rekomendasi investasi. Kripto bisa turun 82% dari puncaknya, dan
seperti terlihat di atas, yang ini sudah.

## Sumber

- [Kaspa's Fair Launch — kaspa.org](https://kaspa.org/kaspas-fair-launch/)
- [Unveiling the "Crescendo" Hard-Fork roadmap: 10BPS and more — kaspa.org](https://kaspa.org/crescendo-hard-fork-roadmap-10bps/)
- [Kaspa Covenants++ "Toccata" Hard-Fork Outlook — Michael Sutton](https://medium.com/@michaelsuttonil/kaspa-covenants-toccata-hard-fork-outlook-a4d81a40900c)
- [Kaspa Roadmap 2026–2027: Every Upgrade, and What It Means — Our Crypto Talk](https://ourcryptotalk.com/blog/kaspa-roadmap-2026-2027)
- [Kaspa (KAS) price, supply & market data — CoinGecko](https://www.coingecko.com/en/coins/kaspa)
- [Kaspa Mining in 2026: Is KAS Still Profitable? — BT-Miners](https://bt-miners.com/kaspa-mining-in-2026-is-kas-still-profitable-complete-roi-analysis-for-every-miner/)
- [Daftar Lengkap 545 Aset Kripto Legal Terdaftar Bappebti — Tokocrypto](https://news.tokocrypto.com/daftar-lengkap-545-aset-kripto-legal-terdaftar-bappebti-di-indonesia/)
- [Bappebti Alihkan Pengaturan dan Pengawasan Aset Kripto kepada OJK dan BI — OJK](https://ojk.go.id/id/berita-dan-kegiatan/siaran-pers/Pages/Bappebti-Kemendag-Alihkan-Tugas-Aset-Keuangan-Digital-termasuk-Aset-Kripto-serta-Derivatif-Keuangan-kepada-OJK-dan-BI.aspx)
