<script setup>
import {computed, onMounted, ref, watch} from "vue";
import MainLayout from "../layouts/MainLayout.vue";
import {useRoute, useRouter} from "vue-router";
import {userStore} from "../store/user_store.js";
import {slugToText} from "../utils/slug.js";
import {villaStore} from "../store/villa_store.js";
import {UlasanVillaRequest} from "../model/request/ulasan_villa_request.js";
import {VillaService} from "../service/villa_service.js";

const router = useRouter()
const route = useRoute()

// Store
const useVillaStore = villaStore()
const useUserStore = userStore()

// State
const loading = ref(false);

watch(
    () => route.params,
    async () => {
      document.title = `Ulasan - ${slugToText(route.params.slug)}`;

      await useVillaStore.setSingleData(route.params.id, route.params.slug);
      if (useVillaStore.state.singleData === null) {
        document.title = "Lampung Trip";
        await router.push("/");
      }
    },
    { immediate: true }
);


const rating = ref(3.0);
const komentar = ref("");
const currentUrl = window.location.href;

function copyLink() {
  navigator.clipboard.writeText(currentUrl).then(() => {
    alert("Link berhasil disalin!");
  }).catch(() => {
    alert("Gagal menyalin link.");
  });
}

async function kirimUlasan(e) {
  e.preventDefault();
  loading.value = true;

  if (useUserStore.user.userId !== 0) {
    const request = new UlasanVillaRequest(parseInt(useVillaStore.state.singleData.id), parseInt(useUserStore.user.biodata.id), parseFloat(rating.value), komentar.value);

    await VillaService.CreateUlasan(request).then((response) => {
      // Success
      if (response === null) {
        window.location.reload();
      }
    }).finally(
        loading.value = false
    )
  } else {
    await router.push("/masuk")
  }
}

onMounted(async () => {
  await useVillaStore.setSingleData(route.params.id, route.params.slug);
  if (useVillaStore.state.singleData === null) {
    document.title = "Lampung Trip";
    await router.push("/");
  }
})

const listData = computed(() => {
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

const checkUlasan = computed(() => {
  let available = true;

  useVillaStore.state.singleData?.ulasan.filter(item => {
    if (item.userId === useUserStore.user.userId) {
      available = false;
    }
  });

  return available;
})
</script>

<template>
  <MainLayout>
    <div class="my-5">
      <!-- Kembali | Share -->
      <div class="d-flex justify-content-between align-items-start mb-4">
        <div class="d-flex flex-column align-items-start">
          <div class="d-flex align-items-center pointer" onclick="history.back()">
            <i class="material-icons">arrow_back</i> Kembali
          </div>
        </div>

        <!-- SHARE BUTTON -->
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
              <a class="dropdown-item d-flex align-items-center" href="#" @click.prevent="copyLink">
                <i class="material-icons me-2">link</i> Salin Link
              </a>
            </li>
            <li>
              <a class="dropdown-item d-flex align-items-center"
                 :href="`https://wa.me/?text=${encodeURIComponent(currentUrl)}`" target="_blank">
                <i class="material-icons me-2">chat</i> Bagikan ke WhatsApp
              </a>
            </li>
            <li>
              <a class="dropdown-item d-flex align-items-center"
                 :href="`https://www.facebook.com/sharer/sharer.php?u=${encodeURIComponent(currentUrl)}`"
                 target="_blank">
                <i class="material-icons me-2">facebook</i> Bagikan ke Facebook
              </a>
            </li>
          </ul>
        </div>
      </div>

      <!-- Comment Area -->
      <form v-if="checkUlasan" class="p-4 rounded mb-4" @submit="kirimUlasan">
        <h4 class="mb-4 text-center">Berikan Ulasan Anda</h4>

        <div class="row g-3">
          <div class="col-12 col-md-6">
            <label for="ratingRange" class="form-label d-block text-center mb-2">
              <strong class="font-color-primary">Rating Anda:</strong>
              <span class="fw-bold">{{ rating }}</span> dari 5
            </label>
            <input type="range" class="form-range" min="1" max="5" step="1" id="ratingRange" v-model="rating" required/>
            <div class="d-flex justify-content-between mt-1 text-muted small">
              <span>1 (Sangat Buruk)</span>
              <span>5 (Sangat Baik)</span>
            </div>
          </div>

          <div class="col-12 col-md-6">
            <label for="komentarInput" class="form-label">Tulis Komentar Anda:</label>
            <textarea class="form-control" id="komentarInput" rows="5"
                      v-model="komentar" placeholder="Bagikan pengalaman atau opini Anda di sini..." required></textarea>
          </div>
        </div>

        <div class="d-grid mt-4">
          <button
              type="submit"
              class="py-1 fs-5 bg-color-primary mt-4 w-100 rounded-5"
          >
          <span v-if="loading" class="spinner-border" role="status">
            <span class="sr-only"></span>
          </span>

            <span v-else>
            Kirim Ulasan
          </span>
          </button>
        </div>
      </form>

      <!-- List Ulasan -->
      <div class="row">
        <div class="col-sm-6 col-md-6 mb-4" v-for="ulasan of listData" :key="ulasan.userId">
          <div class="card-body border rounded p-3">
            <div class="card-title d-flex justify-content-between align-items-center w-full">
              <div class="ratio ratio-1x1 rounded-circle" style="width: 72px; height: 72px">
                <img
                    :src="ulasan.photo === '' ? 'https://placehold.co/400' : ulasan.photo"
                    class="img-fluid rounded-circle object-fit-cover"
                    alt="Foto Profil Pengguna"
                />
              </div>
              <div class="col-3 text-end">
                <span class="d-inline-flex align-items-center">
                  <i class="material-icons fs-1 ms-1 text-warning">star</i>
                  {{ulasan.rating}}
                </span>
              </div>
            </div>
            <div class="w-100">
              <span class="fw-bold">{{ulasan.nama}}</span><br />
              <span>
                {{ulasan.keterangan}}
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </MainLayout>
</template>

<style scoped>
.dropdown-menu {
  min-width: 220px;
}

.dropdown-item i.material-icons {
  font-size: 18px;
}

.card-body {
  background-color: #f9f9f9;
}
</style>
