+++
title = "DCA ETF Emas"
date = 2026-09-10T07:45:00+07:00
draft = false
description = "Saya membeli satu lot XGLD dan satu lot XTRA tiap hari bursa. Setelah dihitung, ternyata yang saya lakukan bukan DCA — dan selisihnya lima kali lebih besar daripada soal frekuensi."
tags = ["investasi", "emas", "etf", "dca", "keuangan", "xgld", "xtra", "stockbit", "ipot", "power-fund-series"]
categories = ["Investasi"]
featured_image = "/images/dca-emas-hero.svg"
toc = true
aliases = ["/posts/dca-etf-emas/"]
+++

Sejak menulis [catatan tentang ETF emas](/investasi/etf-emas-indonesia/) bulan lalu, saya
menjalankan kebiasaan kecil: **tiap hari bursa, beli satu lot XGLD dan satu lot XTRA.**

Dua produk itu saya pilih karena denominasi unitnya paling kecil di antara lima ETF emas
yang tercatat — sekitar Rp256 dan Rp248 per unit, jadi satu lot masing-masing menebus
Rp25 ribuan. Nominal yang keluar tiap hari tidak saya patok; ia mengikuti harga, dan
rata-ratanya jatuh di sekitar **Rp50.400**.

Tulisan ini catatan sewaktu saya memeriksa apakah kebiasaan itu masuk akal. Jawabannya: dua
dari tiga bagiannya sudah benar, dan satu bagian ternyata keliru — bagian yang paling tidak
saya duga.

*Sembilan hari setelah tulisan ini terbit, kompromi yang saya pilih di bawah jadi tidak
perlu. Lihat [pembaruannya](#pembaruan-jalan-yang-lebih-lurus).*

## Berapa sebenarnya ongkosnya

Pertanyaan pertama yang wajar: dengan 490 order setahun, apakah saya digerus biaya
transaksi?

Saya pakai **Stockbit Sekuritas**. Ini yang tertulis di halaman biayanya:

| | Beli | Jual |
|---|---:|---:|
| Saham | 0,15% | 0,25% |
| **ETF**, right, warrant | **0,15%** | **0,15%** |

Komisinya **murni persentase, tanpa minimum per transaksi**. Ini bagian yang menentukan
segalanya: kalau biayanya proporsional, memecah pembelian jadi 490 potong kecil tidak lebih
mahal sepeser pun daripada memborong sekaligus. Order Rp25.000 kena komisi sekitar Rp38, dan
490 order kecil menjumlah persis sama dengan 12 order besar.

Setahun, dengan asumsi harga bergerak di volatilitas 15%:

| | Setahun |
|---|---:|
| Order | 490 |
| Dana tertanam | Rp12.350.205 |
| Rata-rata per hari | Rp50.409 |
| Fee broker 0,15% | Rp18.525 |
| **Spread** (asumsi setengah-spread 0,25%) | **Rp30.876** |
| Total ongkos | Rp49.401 — 0,40% |

Perhatikan baris yang saya tebalkan. **Spread memakan hampir dua kali lipat komisi broker.**
Selisih harga bid dan ask tidak muncul di tagihan mana pun, tapi di ETF yang baru tercatat
dan volumenya belum tebal, ia ongkos terbesar yang saya bayar. Membeli dua produk sekaligus
juga berarti saya menyeberangi dua buku order, bukan satu.

Konsekuensi praktisnya cuma satu, dan itu gratis: **pakai limit order, jangan market order.**
Menawar di harga yang wajar dan menunggu sebentar lebih murah daripada mengejar.

> Catatan untuk yang brokernya berbeda: ada segelintir sekuritas yang memberlakukan
> **minimum fee per transaksi**, dan kalau Anda kena itu, seluruh hitungan di atas berubah
> total — 490 order kecil jadi jauh lebih mahal daripada 12 order besar. Praktik ini
> [sudah jarang di Indonesia](https://rankia.id/sekuritas-fee-termurah-indonesia/), tapi
> lima menit membuka lembar biaya Anda sendiri tetap sepadan.

## Apakah harian lebih baik daripada bulanan?

Pertanyaan kedua, dan ini yang biasanya paling ramai diperdebatkan.

Saya simulasikan gerak harga acak berulang kali, membandingkan harga rata-rata per unit yang
didapat kalau menyetor harian, mingguan, atau bulanan, dengan **semua ongkos dinolkan** supaya
efek frekuensinya berdiri sendiri:

| Frekuensi | Harga rata-rata | Persentil 5 | Persentil 95 |
|---|---:|---:|---:|
| Harian | Rp254,53 | Rp220,30 | Rp291,54 |
| Mingguan | Rp254,60 | Rp220,21 | Rp291,84 |
| Bulanan | Rp254,72 | Rp219,73 | Rp292,91 |

Harian menang — dengan selisih **0,07%** terhadap bulanan.

Sekarang lihat dua kolom kanan. Jarak antara skenario baik dan skenario buruk membentang
**sekitar 28%**. Keunggulan frekuensi itu ada, tapi ia empat ratus kali lebih kecil daripada
kebisingan yang menaunginya. Dalam praktik, ia tidak terasa.

Jadi kenapa saya tetap harian? Bukan karena unggul, tapi karena **di Stockbit ia tidak
berongkos lebih**, dan dari dua pilihan berbiaya sama saya memilih yang paling tidak menuntut
saya berpikir. Nilai DCA harian ada di perilaku, bukan di matematika: tidak ada hari di mana
saya perlu bertanya "sekarang atau tunggu turun dulu?".

## Yang ternyata keliru

Sampai di sini saya sudah puas. Lalu saya sadari ada satu hal yang luput, dan efeknya lebih
besar daripada seluruh perdebatan frekuensi di atas.

Saya membeli **satu lot**, tiap hari. Bukan **sejumlah rupiah**, tiap hari.

Kedengarannya perbedaan sepele. Ternyata ia menghapus satu-satunya keunggulan matematis yang
dimiliki DCA.

Inti DCA bukan "beli rutin". Intinya adalah **nominal rupiahnya dikunci, jumlah unitnya
dibiarkan bergerak.** Waktu harga turun, uang yang sama memborong lebih banyak unit; waktu
harga naik, ia kebagian lebih sedikit. Efeknya otomatis: Anda selalu membeli lebih banyak di
harga murah, tanpa pernah memutuskan apa pun.

Kalau yang dikunci justru jumlah lotnya, mekanisme itu mati. Saya membeli 100 unit hari ini,
100 unit besok, sama banyak entah harganya sedang jatuh atau sedang mahal. Harga rata-rata
yang saya dapat jadi rata-rata biasa dari seluruh harga harian — bukan rata-rata yang condong
ke sisi murah.

Ini bisa dibuktikan tanpa simulasi apa pun. Ambil dua hari saja, harga Rp200 lalu Rp300:

| Cara membeli | Unit didapat | Total bayar | Harga rata-rata |
|---|---:|---:|---:|
| 100 unit tiap hari | 200 | Rp50.000 | **Rp250,00** |
| Rp30.000 tiap hari | 250 | Rp60.000 | **Rp240,00** |

Seluruh mekanismenya ada di kolom pertama. Dengan rupiah tetap, hari murah kebagian 150 unit
sementara hari mahal cuma 100 unit — tanpa saya memutuskan apa pun.

Dan dua angka di kolom terakhir itu punya nama. Rp250 adalah **rata-rata aritmetik** dari
harga harian: (200+300)/2. Rp240 adalah **rata-rata harmonik**: 2/(1/200+1/300). Untuk deret
harga mana pun, rata-rata aritmetik tidak pernah lebih kecil daripada rata-rata harmonik, dan
keduanya sama persis hanya bila harganya tidak bergerak sama sekali.

Jadi ini bukan strategi yang kadang menang kadang kalah. **Mengunci rupiah tidak bisa kalah
dari mengunci lot** — itu sifat bilangan, bukan ramalan pasar. Yang perlu disimulasikan bukan
*apakah* ia menang, melainkan *seberapa besar*, karena itulah yang bergantung pada seberapa
liar harga bergerak.

![Grafik batang: memilih frekuensi harian bernilai 0,07 persen, sementara mengunci rupiah alih-alih lot bernilai 0,37 sampai 4,03 persen tergantung volatilitas](/images/dca-emas-efek.svg)

Hasil simulasi di volatilitas 15%, memakai harga XGLD sebagai titik awal:

| Cara membeli | Harga rata-rata per unit |
|---|---:|
| **1 lot tetap tiap hari** — yang saya lakukan | Rp255,77 |
| **Rp-tetap tiap hari** — DCA sebenarnya | Rp254,82 |
| Selisih | **0,37%** |

Di jalur harga yang paling tidak menguntungkan sekalipun, selisih itu menipis sampai sekitar
0,03% — tapi tidak pernah hilang dan tidak pernah berbalik arah. Besarnya tumbuh bersama
gejolak harga:

| Volatilitas tahunan | Keunggulan rupiah-tetap |
|---|---:|
| 15% — emas dalam keadaan normal | 0,37% |
| 30% | 1,46% |
| 50% | 4,03% |

Bandingkan dengan keunggulan harian-atas-bulanan tadi: 0,07%. **Cara menghitung nominalnya
bernilai lima kali lipat dari frekuensinya** — dan tidak seperti frekuensi, ia tidak
tenggelam di kebisingan, karena arahnya tidak bergantung pada bagaimana pasar bergerak.

Saya menghabiskan berminggu-minggu memikirkan seberapa sering harus membeli, padahal soal
yang lebih besar ada di kata pertama singkatannya. *Dollar*-cost averaging. Yang
dirata-ratakan itu rupiahnya, bukan lotnya.

## Jadi, saya ubah atau tidak?

Sebelum buru-buru mengubah apa pun, saya lihat dulu besarannya. Nol koma tiga puluh tujuh
persen dari Rp12,35 juta adalah **sekitar Rp45.500 setahun**. Nyata, pasti, tapi tidak
mengubah hidup.

Dan cara saya sekarang punya kelebihannya sendiri, yang tidak muncul di simulasi mana pun:

- **Tidak ada yang perlu dihitung.** Buka aplikasi, beli satu lot, satu lot, selesai. Tidak
  ada pembagian, tidak ada sisa kas yang harus diingat.
- **Tidak ada uang menganggur.** Metode rupiah-tetap menyisakan receh tiap hari yang harus
  menunggu sampai cukup untuk satu lot penuh. Metode satu-lot selalu terpakai penuh.
- **Nominalnya justru naik saat harga naik**, yang bagi sebagian orang terasa lebih tenang
  daripada memborong deras saat harga jatuh.

Cuma, kelebihan yang terakhir itu sebenarnya kelemahan yang menyamar. Membeli lebih banyak
waktu harga murah justru bagian yang bikin DCA bekerja.

Yang saya pilih: **kompromi.** Frekuensinya tetap harian, dua produk tetap jalan, tapi jumlah
lotnya tidak lagi selalu satu. Saya patok anggaran harian di kepala — sekitar Rp50.000 — dan
kalau harga sedang turun cukup jauh, satu lot saya jadikan dua. Bukan formula, cuma condong
ke arah yang benar.

Itu tidak menangkap seluruh 0,37%, dan saya tahu itu. Tapi ia menangkap sebagian besar tanpa
membuat kebiasaannya jadi pekerjaan — dan kebiasaan yang terasa seperti pekerjaan biasanya
berhenti dalam tiga bulan.

## Pembaruan: jalan yang lebih lurus

Sembilan hari setelah tulisan ini terbit, kompromi di atas jadi tidak perlu.

XGLD ternyata juga bisa dibeli lewat **IPOT Fund**, di fitur bernama Power Fund Series. Saya
sudah mencobanya. Satu hal yang berbeda di sana persis hal yang saya keluhkan di seksi
sebelumnya: pesanannya berdenominasi rupiah.

Berita-berita yang mengabarkannya memakai angka "mulai 0,1 miligram" atau "mulai Rp250".
Itu lantai minimumnya, dan bukan bagian yang penting buat saya. Yang penting ada di tabel
perbandingan di dalam aplikasinya sendiri, di baris **Simple value in Rp** — bertanda centang
untuk Power Fund Series, bertanda silang untuk ETF biasa.

Kenapa satu baris itu membereskan seluruh seksi "Yang ternyata keliru" di atas: soalnya ada
di ukuran langkah.

Satu lot XGLD menebus sekitar Rp25.400. Anggaran harian saya untuk XGLD, dari Rp50.000 yang
saya bagi dua produk, juga sekitar Rp25.400. Jadi waktu saya menulis "kalau harga sedang
turun cukup jauh, satu lot saya jadikan dua", yang sebenarnya tersedia buat saya cuma dua
angka: Rp25.400 atau Rp50.800. Tidak ada apa pun di antaranya.

Instrumennya punya resolusi selebar seluruh belanja harian saya. Itu sebabnya kompromi tadi
terasa kasar waktu saya menuliskannya — tidak ada angka di antara dua pilihan itu untuk
dihitung.

Lewat jalur rupiah, sisa yang tidak terpakai paling banyak seharga satu unit — sekitar Rp254.
Sebelumnya paling banyak seharga satu lot. Seratus kali lebih rapat, persis sebesar ukuran
lotnya sendiri.

### Yang belum saya pastikan

Saya belum bisa bilang cara ini lebih murah.

Di hitungan di atas, yang paling mahal bukan komisi broker. Spread memakan Rp30.876, hampir
dua kali lipat komisi Rp18.525. Jadi pertanyaan yang menentukan bukan berapa biaya belinya,
tapi apakah saya masih menyeberangi bid-ask.

Tabel yang sama menyebut Power Fund Series menarik likuiditas dari empat sumber sekaligus —
manajer investasi, saham di portofolio, pasar sekunder BEI, dan dealer partisipan — sementara
ETF biasa cuma dua. Barisnya untuk harga juga mencentang Live NAV, bukan cuma harga
orderbook.

Kalau itu berarti transaksinya benar-benar terjadi di NAV, komponen biaya terbesar saya
berubah. Kalau ternyata tetap menyeberangi buku order, ia tidak berubah sama sekali.

Saya belum tahu yang mana. Tabel itu materi pemasaran, dan materi pemasaran bukan nota
transaksi. Tarif belinya pun belum saya cocokkan dengan yang benar-benar tercatat — sumber
publik yang saya temukan malah saling bertabrakan soal angkanya. Sampai saya punya beberapa
bulan nota, seksi ini menggantung.

Mekanismenya saya uraikan lebih panjang di
[catatan tersendiri](/investasi/beli-xgld-tanpa-menghitung-lot/).

Satu hal lagi dari tabel yang sama, yang tidak ada hubungannya dengan DCA tapi sebaiknya
disebut: Power Fund Series mengizinkan posisi dengan leverage. Membeli emas dengan utang
persoalan yang lain sama sekali, dan tidak ada di tulisan ini yang berlaku untuk itu.

## Kenapa dua produk

Satu pertanyaan yang belum saya jawab: kenapa XGLD **dan** XTRA, bukan salah satu saja?

Alasannya bukan diversifikasi harga — keduanya melacak emas yang sama, jadi tidak ada risiko
yang benar-benar tersebar. Yang tersebar adalah **risiko pengelolanya**: dua manajer
investasi, dua bank kustodian, dua kebijakan yang bisa berubah.

Ongkosnya: saya menyeberangi dua spread, bukan satu. Dan setelah melihat bahwa spread adalah
komponen biaya terbesar saya, itu bukan harga yang gratis.

Setelah beberapa bulan berjalan, saya berencana membandingkan keduanya pada hal yang benar-
benar bisa diukur — seberapa rapat harga pasarnya menempel NAB, dan seberapa lebar
spread-nya di jam sepi. Kalau salah satu ternyata konsisten lebih baik, tidak ada alasan
sentimental untuk mempertahankan dua-duanya.

## Catatan asumsi

Angka di atas **simulasi, bukan data historis**. ETF emas di BEI baru tercatat 10 Agustus
2026, jadi riwayat harganya belum cukup panjang untuk diuji serius.

Cara menghitungnya: saya bangkitkan jalur harga acak (gerak Brown geometrik) sepanjang satu
tahun bursa, menjalankan tiap strategi pada jalur yang sama persis, lalu mengulangnya ribuan
kali dan merata-ratakan hasilnya.

Yang saya asumsikan: harga awal Rp256 (XGLD) dan Rp248 (XTRA) per unit, fee beli 0,15%,
setengah-spread 0,25%, volatilitas tahunan 15%, 245 hari bursa, tanpa arah naik atau turun
yang diasumsikan. Semuanya bisa berbeda di kasus Anda. Perhitungan ini juga belum memasukkan
**biaya pengelolaan tahunan ETF** (sekitar 0,3–0,5%) maupun perlakuan pajaknya, yang
[masih abu-abu](/investasi/etf-emas-indonesia/).

Tarif Stockbit yang saya kutip adalah yang tercantum di
[halaman biaya resminya](https://help.stockbit.com/id/article/transaksi-saham-berapa-biaya-trading-di-stockbit-sekuritas-1lbkyq9/)
per September 2026. Lembar biaya bisa berubah kapan saja — jangan percaya angka di blog
orang, termasuk blog ini.

Seksi pembaruan ditambahkan 19 September 2026. Keterangan mekanisme Power Fund Series di
sana saya ambil dari tabel perbandingan di dalam aplikasi IPOT dan dari
[keterangan Indo Premier](https://www.indopremier.com/ipotnews/newsDetail.php?jdl=Bagaimana_Cara_Membeli_Unit_Penyertaan_Reksa_Dana_Exchange_Traded_Fund_&news_id=92050&group_news=IPOTNEWS&taging_subtype=MUTUALFUNDEDUCATION);
angka biaya dan perilaku spread-nya belum saya verifikasi dari nota transaksi sendiri.
Hitungan di badan tulisan tidak saya ubah — ia rekaman apa yang saya tahu pada 10 September.

Ini catatan pribadi, bukan rekomendasi investasi. Emas bisa turun, dan sudah pernah turun
lama.
