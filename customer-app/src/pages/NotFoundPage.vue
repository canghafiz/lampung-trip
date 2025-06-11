<script setup>
import { useRouter } from 'vue-router'; // Import useRouter untuk navigasi programatis

// Dapatkan instance router
const router = useRouter();

// Fungsi untuk kembali ke halaman utama
function goToHomePage() {
  router.push('/'); // Menggunakan push agar pengguna bisa kembali jika diperlukan
}
</script>

<template>
  <div class="d-flex align-items-center justify-content-center vh-100 bg-light">
    <div class="text-center">
      <h1 class="display-1 fw-bold font-color-primary">404</h1>
      <p class="fs-3"> <span class="font-color-primary">Oups!</span> Halaman tidak ditemukan.</p>
      <p class="lead">
        Maaf, halaman yang Anda cari tidak ada.
      </p>
      <button @click="goToHomePage" class="btn btn-primary mt-3">
        <i class="bi bi-house-door-fill me-2"></i> Kembali ke Beranda
      </button>
    </div>
  </div>
</template>

<style scoped>
.text-danger {
  /* Opsional: Kustomisasi warna merah jika diinginkan */
  color: #103d47;
}

.btn-primary {
  /* Opsional: Kustomisasi warna tombol utama jika diinginkan */
  background-color: #103d47;
  border-color: #103d47;
}

.btn-primary:hover {
  background-color: #103d47;
  border-color: #103d47;
}

/* Memastikan ikon Bootstrap dimuat jika Anda menggunakannya */
/* Pastikan Anda sudah mengimpor Bootstrap Icons di proyek Anda
   Contoh di index.html atau di main.js:
   <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.11.3/font/bootstrap-icons.min.css">
*/
</style>