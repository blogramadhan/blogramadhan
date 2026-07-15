#!/usr/bin/env bash
# Jalankan dashboard admin lokal untuk mengelola tulisan blog.
# Pakai: ./dashboard.sh
set -e
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "Membuka dashboard di http://127.0.0.1:1414 …"
echo "(Tekan Ctrl+C untuk berhenti. Jalankan 'hugo server -D' di terminal lain untuk pratinjau situs.)"
cd "$ROOT/tools/dashboard"
exec go run . -content "$ROOT/content"
