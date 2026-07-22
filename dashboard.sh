#!/usr/bin/env bash
# Jalankan dashboard admin lokal untuk mengelola tulisan blog.
# Pakai: ./dashboard.sh
set -e
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "Membuka dashboard di http://127.0.0.1:1414 …"
echo "(Tekan Ctrl+C untuk berhenti. Jalankan 'hugo server -D' di terminal lain untuk pratinjau situs.)"
# Semua path diberikan absolut: dashboard membaca konten, folder gambar, dan
# hugo.toml dari root proyek — bukan dari tools/dashboard tempat kode berjalan.
cd "$ROOT/tools/dashboard"
exec go run . \
  -content "$ROOT/content" \
  -static "$ROOT/static" \
  -config "$ROOT/hugo.toml"
