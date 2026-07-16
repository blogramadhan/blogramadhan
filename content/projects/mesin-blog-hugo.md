+++
title = "Tema Hugo Buatan Sendiri"
date = 2026-07-12
draft = false
description = "Tema Hugo custom tanpa kerangka pihak ketiga — tipografi bersih, mode gelap, dan penyajian kode yang enak dibaca."
year = "2026"
role = "Desain & Pengembangan"
tech = ["Hugo", "Go Templates", "CSS"]
# repo = "https://github.com/"
# demo = "https://example.com/"
+++

Tema ini lahir dari keinginan sederhana: sebuah blog yang tenang dibaca dan jujur pada
isinya. Saya membangunnya dari nol di atas Hugo, tanpa kerangka pihak ketiga.

Versi pertamanya bergaya majalah cetak — serif, *drop cap*, pemisah berornamen. Cantik,
tetapi tidak jujur: yang paling sering saya tulis adalah catatan teknis, dan ornamen itu
justru menghalangi hal yang mestinya paling menonjol, yaitu kode. Jadi temanya saya
rombak ke arah yang lebih lugas.

## Sorotan

- **Tipografi bersih** dengan *Inter* untuk teks dan *JetBrains Mono* untuk kode —
  dipilih agar padat, presisi, dan tidak melelahkan saat dibaca lama.
- **Penyajian kode sebagai warga kelas satu**: pewarnaan sintaks, label bahasa, dan
  tombol salin di setiap blok.
- **Mode terang & gelap** yang menyesuaikan preferensi sistem, dengan tombol manual dan
  tanpa kedip saat halaman dimuat. Palet kode punya versi tersendiri untuk keduanya,
  bukan satu palet yang dipaksakan.
- **Nihil JavaScript berat** — hanya dua skrip kecil: peralihan tema dan tombol salin,
  yang pun cuma dimuat pada halaman yang memang berisi kode. Sisanya HTML dan CSS murni
  yang dibangun saat *build*.

## Pelajaran

Membangun tema sendiri memaksa saya memahami *lookup order* Hugo dan Hugo Pipes secara
mendalam. Hasilnya bukan cuma tema, tetapi juga peta mental yang jauh lebih jelas tentang
cara kerja situs statis.

Pelajaran paling mahal justru datang dari pewarnaan kode. Konfigurasinya sudah menyebut
sebuah gaya, jadi saya pikir semuanya beres — padahal Hugo disetel memancarkan *class*
CSS, sementara berkas gayanya tidak pernah ada. Kode saya terbit hitam-putih tanpa saya
sadari, dan tidak ada satu pun pesan galat yang memberi tahu. Konfigurasi yang terlihat
benar tidak berarti keluarannya benar; yang menentukan adalah apa yang sampai ke pembaca.
