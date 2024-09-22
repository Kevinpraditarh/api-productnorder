# API Product Order

## Deskripsi
API untuk mengelola produk dan pesanan menggunakan Go dengan database MySQL. API ini menyediakan endpoint untuk melakukan operasi CRUD pada produk dan pesanan.

## Fitur
- Mendapatkan daftar produk
- Membuat produk baru
- Mendapatkan detail produk
- Memperbarui produk
- Menghapus produk
- Mendapatkan daftar pesanan
- Membuat pesanan baru
- Mendapatkan detail pesanan
- Menghapus pesanan

## Persyaratan
- Go 1.15 atau lebih baru
- MySQL 5.7 atau lebih baru
- Library yang diperlukan:
  - `github.com/go-sql-driver/mysql`

## Instalasi
1. **Clone repository:**
   ```bash
   - git clone <repository-url>
   - cd api-productnorder

2. **Install dependency**
   ```bash
   - go mod tidy

3. **Konfigurasi database**
    

4. **Jalankan Aplikasi**
    - go run main.go