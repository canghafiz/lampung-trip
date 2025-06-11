# 🌴 Lampung Trip

**Lampung Trip** adalah aplikasi web dinamis untuk mengeksplorasi destinasi wisata dan villa terbaik di Provinsi Lampung. Aplikasi ini dibangun dengan **Vue.js** di sisi frontend dan **Golang** di sisi backend. Sistem menggunakan autentikasi JWT dan role-based access control untuk mengatur hak akses pengguna berdasarkan peran sebagai Admin atau Karyawan.

---

## 🔗 Contoh Hasil 

  - [Sisi Customer](https://lampungtrip.netlify.app/)
  - [Sisi Admin & Karyawan](https://admin-lampungtrip.netlify.app/)

---

## ✨ Fitur Utama

- 🔐 **Autentikasi JWT**
  - Login dan logout aman menggunakan token JWT.
  - Hak akses dibedakan berdasarkan peran pengguna (Admin dan Karyawan).

- 🔎 **Pencarian Destinasi**
  - Pencarian tempat wisata dan villa berdasarkan kategori.
  - Tampilan responsif dan dinamis untuk pengguna umum.

- 🧑‍💼 **Role Admin**
  - CRUD data destinasi wisata.
  - CRUD data villa.
  - Manajemen data karyawan (tambah, edit, hapus, lihat).

- 👨‍🔧 **Role Karyawan**
  - CRUD data destinasi wisata.
  - CRUD data villa.

---

## ⚙️ Teknologi yang Digunakan

### Frontend
- [Vue.js 3](https://vuejs.org/)
- Vue Router
- Axios

### Backend
- [Golang](https://golang.org/)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/) (ORM untuk MySQL)
- [JWT (JSON Web Token)](https://github.com/golang-jwt/jwt)

### Database
- MySQL

---

## 📁 Struktur Proyek

```text
lampung-trip/
├── backend/                 # Backend menggunakan Golang
│   ├── main.go
│   ├── controllers/         # Logika handler untuk setiap endpoint
│   ├── models/              # Model database dan relasi
│   ├── routes/              # Routing endpoint
│   └── middlewares/         # Middleware untuk JWT, Role, dll
├── admin-app/               # Frontend menggunakan Vue 3
│   ├── public/
│   ├── src/
│   │   ├── assets/
│   │   ├── components/      # Komponen UI (button, input, dll)
│   │   ├── pages/           # Halaman (Login, Dashboard, Wisata, Villa, dll)
│   │   ├── models/          # Model Request & Response Api
│   │   ├── services/        # Api Servis
│   │   ├── stores/          # State management pinia
│   │   ├── utils/           # Fungsi - fungsi yang sering digunakan
│   │   └── App.vue
├── customer-app/            # Frontend menggunakan Vue 3
│   ├── public/
│   ├── src/
│   │   ├── assets/
│   │   ├── components/      # Komponen UI (button, input, dll)
│   │   ├── pages/           # Halaman (Login, Dashboard, Wisata, Villa, dll)
│   │   ├── models/          # Model Request & Response Api
│   │   ├── services/        # Api Servis
│   │   ├── stores/          # State management pinia
│   │   ├── utils/           # Fungsi - fungsi yang sering digunakan
│   │   └── App.vue
└── README.md                # Dokumentasi proyek
```

## ⚙️ Konfigurasi Environment

Buat file `.env` di dalam folder `backend/` dan masukkan konfigurasi berikut:

```env
APP_PORT=:8080
APP_DOMAIN=https://lampungtrip.lonemanga.ink/

SECRET_KEY=Test123

DB_PORT=3306
DB_HOST=localhost
DB_USER=root
DB_PASS=
DB_NAME=db_lampung_trip

```

## 📧 Kontak Pengembang

Dikembangkan oleh: **Hafizarrahman**  
GitHub: [https://github.com/canghafiz](https://github.com/canghafiz)  
Email: [fizrahman47@gmail.com](fizrahman47@gmail.com) & [hfizrrhman@gmail.com](hfizrrhman@gmail.com) && [hfizzrhman@gmail.com](hfizzrhman@gmail.com)
