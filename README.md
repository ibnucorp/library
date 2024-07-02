# Projek UAS : API Perpustakaan

#### Pembuatan projek API perpustakaan dengan golang untuk menyelesaikan tugas mata kuliah Pemrosesan Data Terdistribusi

Pendahuluan

Dokumentasi ini menjelaskan tentang API Perpustakaan yang dibangun menggunakan Golang dan terhubung ke database MySQL. API ini menyediakan fungsionalitas CRUD (Create, Read, Update, Delete) untuk mengelola data buku dalam perpustakaan.

Teknologi yang Digunakan:

- Golang: Bahasa pemrograman backend
- MySQL: Basis data relasional
- Postman: Alat untuk menguji API

Versi:

- Golang: 1.22.1
- Database: 10.4.28-MariaDB (XAMPP)

Fitur:

- Menambahkan buku baru
- Mengambil daftar semua buku
- Mengambil detail buku berdasarkan ID
- Memperbarui informasi buku
- Menghapus buku

Panduan Penggunaan:

1. Instalasi:

   - Pastikan Anda memiliki Golang versi 1.22.1 terinstal di komputer Anda.
     Unduh dan instal XAMPP.
   - Buat database MySQL bernama "db_library".

2. Menjalankan API:

   - Buka terminal dan navigasikan ke direktori project API Anda.
   - Jalankan perintah berikut `go run main.go`

3. Menguji API:

   - Buka Postman dan buat request baru.
   - Atur URL API Anda (misalnya: http://localhost:8080/books).
   - Pilih metode HTTP yang sesuai (GET, POST, PUT, DELETE) untuk operasi yang ingin Anda lakukan.
   - Isi body request dengan data JSON yang sesuai untuk operasi CRUD.
   - Kirim request dan lihat responsnya di Postman.
