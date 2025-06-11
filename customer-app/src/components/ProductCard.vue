<script setup>
import {computed, ref, watchEffect} from "vue";
import {WisataResponse} from "../model/response/wisata_response.js";
import {VillaResponse} from "../model/response/villa_response.js";
import {generateSlug} from "../utils/slug.js";
import {tabStore} from "../store/tab_store.js";

// --- Store ---
const useTabStore = tabStore();

// --- Props ---
const props = defineProps({
  // Definisikan 'wisata' sebagai opsional dan default ke null
  wisata: {
    type: Object, // Lebih baik menggunakan Object daripada langsung WisataResponse karena bisa null
    default: null
  },
  // Definisikan 'villa' sebagai opsional dan default ke null
  villa: {
    type: Object, // Lebih baik menggunakan Object daripada langsung VillaResponse karena bisa null
    default: null
  },
  style: {
    type: String,
    default: "default"
  }
});

// --- State Reaktif untuk URL Gambar ---
// Inisialisasi dengan placeholder agar ada gambar default saat dimuat
const currentLoadedImageUrl = ref('https://placehold.co/400');
// Atau, jika Anda memiliki placeholder khusus:
// const currentLoadedImageUrl = ref('/placeholder-image.jpg');

// --- Computed Property untuk Data yang Ditampilkan ---
const displayedData = computed(() => {
  // Prioritaskan wisata jika ada dan merupakan instance dari WisataResponse
  if (props.wisata instanceof WisataResponse) {
    return props.wisata;
  }
  // Kemudian villa jika ada dan merupakan instance dari VillaResponse
  if (props.villa instanceof VillaResponse) {
    return props.villa;
  }
  // Jika keduanya null atau bukan instance yang benar, kembalikan null
  return null;
});

watchEffect(async () => {
  if (displayedData.value && displayedData.value.gambar && displayedData.value.gambar.length > 0) {
    currentLoadedImageUrl.value = displayedData.value.gambar[0].url;
  } else {
    // Jika tidak ada data atau tidak ada gambar yang valid
    currentLoadedImageUrl.value = 'https://placehold.co/400?text=No+Image';
  }
});

// --- Computed Property untuk Deskripsi ---
const descriptionText = computed(() => {
  if (displayedData.value && displayedData.value.deskripsi) {
    return displayedData.value.deskripsi;
  }
  return 'Tidak ada deskripsi.'; // Default teks jika tidak ada data atau deskripsi
});

// --- Computed Property untuk Judul ---
// Tambahkan juga untuk judul agar lebih aman
const titleText = computed(() => {
  if (displayedData.value && displayedData.value.judul) {
    return displayedData.value.judul;
  }
  return 'Tidak Ada Judul';
});

// --- Computed Property untuk Rating ---
// Tambahkan juga untuk rating agar lebih aman
const ratingText = computed(() => {
  if (displayedData.value && typeof displayedData.value.rating === 'number') {
    return Number(displayedData.value.rating).toFixed(1);
  }
  return 'N/A'; // Not Available
});

</script>

<template>
  <RouterLink
      v-if="displayedData"
      :to="`/${useTabStore.state.type.toLowerCase()}/${displayedData.id}/${generateSlug(displayedData.judul)}`"
  >
    <div :style="props.style" class="card h-100 rounded-4 bg-transparent border-0">
      <img
          class="card-img-top rounded-4"
          :src="currentLoadedImageUrl"
          :alt="`Gambar ${titleText}`"
      />
      <div class="card-body">
        <div class="card-title row align-items-center">
          <div class="col-9 text-truncate fs-5">
            {{ titleText }}
          </div>
          <div class="col-3 text-end">
          <span class="d-inline-flex align-items-center">
            <i class="material-icons fs-5 ms-1">star</i>
            {{ ratingText }}
          </span>
          </div>
        </div>
        <p class="w-100 text-truncate">
          {{ descriptionText }}
        </p>
      </div>
    </div>
  </RouterLink>
  <div v-else class="card-placeholder">
    Memuat...
  </div>
</template>

<style scoped>
.card {
  transition: transform 0.2s ease-in-out;
}

.card:hover {
  transform: translateY(-5px);
}

.card-img-top {
  width: 100%;
  height: 200px; /* Atur tinggi yang tetap atau gunakan object-fit */
  object-fit: cover; /* Pastikan gambar mengisi area tanpa distorsi */
}

.card-title {
  margin-bottom: 0.5rem;
}

.text-truncate {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.card-placeholder {
  /* Gaya untuk placeholder jika tidak ada data */
  width: 100%;
  height: 350px; /* Sesuaikan dengan tinggi kartu */
  display: flex;
  justify-content: center;
  align-items: center;
  border: 1px dashed #ccc;
  border-radius: 0.5rem;
  color: #888;
}
</style>