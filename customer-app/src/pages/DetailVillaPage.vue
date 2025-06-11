<script setup>
import { computed, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import { generateSlug, slugToText } from "../utils/slug.js";
import MainLayout from "../layouts/MainLayout.vue";
import { formatRupiah, parseRupiahToNumber } from "../utils/currency.js";
import {userStore} from "../store/user_store.js";
import {villaStore} from "../store/villa_store.js";

const route = useRoute();
const router = useRouter();

// Store
const useUserStore = userStore()
const useVillaStore = villaStore();

// Watch perubahan route (id/slug) untuk ambil data villa
watch(
    () => route.params,
    async () => {
      document.title = slugToText(route.params.slug);

      await useVillaStore.setSingleData(route.params.id, route.params.slug);
      if (useVillaStore.state.singleData === null) {
        document.title = "Lampung Trip";
        await router.push("/");
      }
    },
    { immediate: true }
);

// Format data wisata untuk ditampilkan
const villa = computed(() => {
  const data = useVillaStore.state.singleData;
  if (!data) return null;

  return {
    ...data
  };
});

// Villa lainnya (selain yang sedang dibuka)
const villaLainnya = computed(() => {
  return useVillaStore.state.listData.filter(item => {
    return item.id !== useVillaStore.state.singleData?.id;
  });
});

const currentUrl = window.location.href

function copyLink() {
  navigator.clipboard.writeText(currentUrl).then(() => {
    alert("Link berhasil disalin!")
  }).catch(() => {
    alert("Gagal menyalin link.")
  })
}

const listUlasan = computed(() => {
  const currentUserId = useUserStore.user.userId;
  const ulasan = useVillaStore.state.singleData?.ulasan;

  // Jika tidak ada ulasan atau singleData tidak ada, kembalikan array kosong
  if (!ulasan || ulasan.length === 0) {
    return [];
  }

  // Filter ulasan milik user yang sedang login
  const myUlasan = ulasan.filter(item => item.userId === currentUserId);

  // Filter ulasan milik user lain
  const otherUlasan = ulasan.filter(item => item.userId !== currentUserId);

  // Gabungkan: ulasan saya di depan, diikuti ulasan orang lain
  // Jika ada beberapa ulasan dari user yang sama, mereka semua akan ada di 'myUlasan'
  return [...myUlasan, ...otherUlasan];
});
</script>


<template>
  <MainLayout>
    <span v-if="useVillaStore.state.singleData === null" class="spinner-border" role="status">
      <span class="sr-only"></span>
    </span>

    <div v-if="useVillaStore.state.singleData !== null" class="my-5">
      <!-- Kembali | Judul | Share -->
      <div class="d-flex justify-content-between align-items-start">
        <div class="d-flex flex-column align-items-start">
          <div
              class="d-flex align-items-center pointer"
              onclick="history.back()"
          >
            <i class="material-icons">arrow_back</i> Kembali
          </div>
          <h1 class="my-4 fw-bold">{{ villa.judul }}</h1>
        </div>
        <!-- SHARE BUTTON + MENU -->
        <div class="dropdown">
          <button
              type="button"
              class="bg-white rounded-circle border-0 d-flex align-items-center justify-content-center"
              aria-label="Share Link"
              style="width: 48px; height: 48px"
              data-bs-toggle="dropdown"
              aria-expanded="false"
          >
            <i class="material-icons font-color-primary fs-5">share</i>
          </button>
          <ul class="dropdown-menu dropdown-menu-end">
            <li>
              <a
                  class="dropdown-item d-flex align-items-center"
                  href="#"
                  @click.prevent="copyLink"
              >
                <i class="material-icons me-2">link</i> Salin Link
              </a>
            </li>
            <li>
              <a
                  class="dropdown-item d-flex align-items-center"
                  :href="`https://wa.me/?text=${encodeURIComponent(currentUrl)}`"
                  target="_blank"
                  rel="noopener"
              >
                <i class="material-icons me-2">chat</i> Bagikan ke WhatsApp
              </a>
            </li>
            <li>
              <a
                  class="dropdown-item d-flex align-items-center"
                  :href="`https://www.facebook.com/sharer/sharer.php?u=${encodeURIComponent(currentUrl)}`"
                  target="_blank"
                  rel="noopener"
              >
                <i class="material-icons me-2">facebook</i> Bagikan ke Facebook
              </a>
            </li>
          </ul>
        </div>

      </div>

      <!-- Gambar -->
      <div id="carouselExampleIndicators" class="carousel slide">
        <div class="carousel-indicators">
          <button
              v-for="(_, index) of villa.gambar"
              :key="index"
              type="button"
              data-bs-target="#carouselExampleIndicators"
              :data-bs-slide-to="index"
              :class="{ active: index === 0 }"
              :aria-current="index === 0 ? 'true' : 'false'"
              :aria-label="'Slide ' + (index + 1)"
          ></button>
        </div>
        <div class="carousel-inner">
          <div
              v-for="(gambar, index) of villa.gambar"
              :key="index"
              :class="{
        'carousel-item': true,
        'active': index === 0,
        'h-50': true
    }"
          >
            <center>
              <img
                  id="hero"
                  :src="gambar.url"
                  class="d-block" :alt="gambar.alt || 'Gambar Villa'"
                  style="
                  height: 300px;
@media screen and (max-width: 767px) {
  #hero {
    width: 100%;
    display: none;
  }
}"
              />
            </center>
          </div>
        </div>
        <button
            class="carousel-control-prev"
            type="button"
            data-bs-target="#carouselExampleIndicators"
            data-bs-slide="prev"
        >
              <span
                  class="carousel-control-prev-icon"
                  aria-hidden="true"
              ></span>
          <span class="visually-hidden">Previous</span>
        </button>
        <button
            class="carousel-control-next"
            type="button"
            data-bs-target="#carouselExampleIndicators"
            data-bs-slide="next"
        >
              <span
                  class="carousel-control-next-icon"
                  aria-hidden="true"
              ></span>
          <span class="visually-hidden">Next</span>
        </button>
      </div>

      <!-- Deskripsi & Info -->
      <div class="row my-4 mx-2 mx-md-0">
        <div class="col-md-8">
          <h5 class="my-4">Deskripsi</h5>
          <p>
            {{villa.deskripsi}}
          </p>
        </div>
        <div
            class="col-md-4 px-2 py-4 rounded-4 d-flex flex-column justify-content-center"
            style="background-color: #ebebeb"
        >
          <div class="w-full mb-4">
            <div class="text-center">
              <span class="fw-bold text-secondary">Tiket Masuk</span> <br />
              <span class="text-black">{{formatRupiah(parseRupiahToNumber(villa.info.harga))}} / malam</span>
            </div>
          </div>
          <a href="https://wa.me/6287801756633" target="_blank" rel="noopener noreferrer">
            <button class="py-2 bg-dark text-white rounded-4 border-0 w-100">
              Kontak Customer Support
            </button>
          </a>
        </div>
      </div>

      <!-- Fasilitas & Lokasi -->
      <div class="row my-4 mx-2 mx-md-0">
        <div class="col-md-8">
          <h5>Fasilitas</h5>
          <ul class="row list-unstyled">
            <li v-for="fasilitas of villa.fasilitas" class="col-md-3 col-6 mb-2 d-flex align-items-center">
              {{fasilitas.url_icon}} {{fasilitas.judul}}
            </li>
          </ul>
        </div>
        <div class="col-md-4">
          <h5>Lokasi</h5>
          <a :href="villa.info.url_lokasi" class="text-decoration-underline text-black fw-light" target="_blank">
            Klik untuk melihat lokasi
          </a>
        </div>
      </div>

      <!-- Ulasan -->
      <div>
        <div class="d-flex justify-content-between">
          <h5 class="fs-5">Ulasan pengguna lainnya</h5>
          <RouterLink class="text-decoration-underline text-black fw-light" :to="`/villa/${useVillaStore.state.singleData?.id}/${route.params.slug}/ulasan`">
            Lihat semua ulasan
          </RouterLink>
        </div>
        <div class="scroll-container">
          <div
              v-for="ulasan of listUlasan" :key="ulasan.userId"
              class="card h-100 rounded-4 bg-transparent border-0 d-inline-block"
              style="width: 200px"
          >
            <div class="card-body">
              <div
                  class="card-title d-flex justify-content-between align-items-center w-full"
              >
                <div
                    class="ratio ratio-1x1 rounded-circle"
                    style="width: 72px; height: 72px"
                >
                  <img
                      :src="ulasan.photo === '' ? 'https://placehold.co/400' : ulasan.photo"
                      class="img-fluid rounded-circle object-fit-cover"
                      alt="Foto Profil Pengguna"
                  />
                </div>
                <div class="col-3 text-end">
                      <span class="d-inline-flex align-items-center">
                        <i class="material-icons fs-1 ms-1 text-warning"
                        >star</i
                        >
                        {{ulasan.rating}}
                      </span>
                </div>
              </div>
              <div class="w-100 text-truncate">
                <span class="fw-bold">{{ulasan.nama}}</span><br />
                <span
                >{{ulasan.keterangan}}</span
                >
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Lainnya -->
      <div class="mt-4">
        <h5 class="fs-5 text-start">Villa lainnya</h5>
        <div class="scroll-container">
          <RouterLink
              v-for="data of villaLainnya"
              :key="data.id" :to="`/villa/${data.id}/${generateSlug(data.judul)}`"
              class="router-link-item"
          >
            <div
                class="card h-100 rounded-4 bg-transparent border-0 d-inline-block"
                style="width: 300px"
            >
              <img
                  class="card-img-top rounded-4"
                  :src="data.gambar[0]?.url || 'https://via.placeholder.com/300x200?text=Gambar+Tidak+Tersedia'" :alt="data.judul || 'Gambar Villa'"
              />
              <div class="card-body">
                <div class="card-title row align-items-center">
                  <div class="col-9 text-truncate fs-5">
                    {{ data.judul }}
                  </div>
                  <div class="col-3 text-end">
          <span class="d-inline-flex align-items-center">
            <i class="material-icons fs-5 ms-1">star</i>
            {{ Number(data.rating).toFixed(1)}} </span>
                  </div>
                </div>
                <p class="w-100 text-truncate">
                  {{ data.deskripsi || 'Deskripsi singkat tidak tersedia.' }}
                </p>
              </div>
            </div>
          </RouterLink>
        </div>
      </div>
    </div>
  </MainLayout>
</template>

<style scoped>
.scroll-container {
  overflow-x: auto;
  white-space: nowrap; /* Penting untuk RouterLink sebagai inline-block */
  padding: 10px;
  -ms-overflow-style: none; /* IE and Edge */
  scrollbar-width: none; /* Firefox */
}

.scroll-container::-webkit-scrollbar {
  display: none; /* Webkit (Chrome, Safari, Opera) */
}

/* Mengatur RouterLink agar berperilaku seperti inline-block */
.scroll-container > .router-link-item { /* Beri kelas pada RouterLink */
  display: inline-block;
  vertical-align: top;
  margin-right: 24px; /* Jarak antar kartu */
}

.card-img-top {
  width: 100%;
  /* --- Bagian Penting untuk Tinggi Konsisten --- */
  height: 200px; /* Atur tinggi gambar secara eksplisit */
  object-fit: cover; /* Penting: Memastikan gambar mengisi area tanpa distorsi */
}

/* Gaya tambahan untuk card */
.card.h-100 {
  height: 100% !important; /* Pastikan card mengambil tinggi penuh kontainer flex jika ada */
}

.card-body {
  /* Sesuaikan padding atau tinggi minimum jika diperlukan */
  min-height: 120px; /* Contoh: Memberi tinggi minimum untuk body card */
}

/* Mengatur teks judul dan deskripsi agar tidak terlalu panjang */
.text-truncate {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>