<script setup>
import { sideNavStore } from "../store/side_nav.js";
import { villaStore } from "../store/villa_store.js";
import { computed, onMounted, ref } from "vue";
import { VillaService } from "../services/villa_service.js";
import { VillaResponse } from "../models/response/villa_response.js";
import TableData from "../components/Loading/TableData.vue";
// Pastikan path ke function.js benar dan fungsi diekspor
import { deleteMultipleGambar } from "../utils/function.js";
import { searchStore } from "../store/search_store.js";

// Store
const useSideNavStore = sideNavStore();
const useVillaStore = villaStore();
const useSearchStore = searchStore();

// --- Fungsi untuk mengambil dan memproses data villa ---
// Ini akan mengambil data dari API, memproses URL gambar, lalu memperbarui store.
const fetchAndProcessVillaData = async () => {
  useVillaStore.setLoadData(true);
  try {
    const response = await VillaService.GetAll();
    const result = VillaResponse.fromArray(response.data);

    // Proses setiap item di 'result' untuk memuat URL gambar dengan ngrok-skip
    const processedResult = await Promise.all(
        result.map(async (villa) => {
          // Pastikan villa.gambar ada dan memiliki elemen pertama dengan URL
          if (villa.gambar && villa.gambar.length > 0 && villa.gambar[0].url) {
            const validUrl = villa.gambar[0].url
            return {
              ...villa,
              // Tambahkan properti baru 'displayImageUrl' ke objek villa
              // Gunakan validUrl jika ada, jika tidak, fallback ke originalUrl, atau placeholder
              displayImageUrl: validUrl,
            };
          }
          return {
            ...villa,
            displayImageUrl: 'https://via.placeholder.com/60?text=No+Image', // Placeholder jika tidak ada gambar
          };
        })
    );

    // Update Store dengan data yang sudah diproses
    useVillaStore.initListData(processedResult);
    console.log("Successfully getting list data villa");
  } catch (error) {
    console.error("Error fetching or processing villa data:", error);
    // TODO: Tambahkan notifikasi error ke user jika diperlukan (misal: toast)
  } finally {
    useVillaStore.setLoadData(false);
  }
};

onMounted(async () => {
  // Hanya ambil data jika store kosong (untuk menghindari fetch berulang saat komponen di-mount)
  if (useVillaStore.state.listData.length === 0) {
    await fetchAndProcessVillaData();
  }
});

// State untuk Modal Hapus
const tampilModalHapus = ref(false);
const hapusId = ref("");
const gambar = ref([]); // Menyimpan data gambar dari item yang akan dihapus

const tampilkanModalHapus = (villaItem) => { // Ganti 'response' jadi 'villaItem' biar lebih jelas
  tampilModalHapus.value = true;
  hapusId.value = villaItem.id;
  gambar.value = villaItem.gambar;
};

const sembunyikanModalHapus = () => {
  tampilModalHapus.value = false;
  hapusId.value = "";
  gambar.value = []; // Reset array gambar setelah modal disembunyikan
};

const hapusModal = async () => {
  try {
    // Panggil fungsi untuk menghapus gambar dari storage (jika ada)
    await deleteMultipleGambar(gambar.value);

    // Panggil service untuk menghapus data villa dari backend
    await VillaService.Delete(hapusId.value);

    // Setelah berhasil dihapus, refresh daftar data di UI
    await fetchAndProcessVillaData();

    console.log(`Data villa dengan ID ${hapusId.value} berhasil dihapus.`);
    // TODO: Tambahkan notifikasi sukses ke user (misal: toast)
  } catch (error) {
    console.error("Error deleting villa data:", error);
    // TODO: Tambahkan notifikasi error ke user jika penghapusan gagal
  } finally {
    sembunyikanModalHapus(); // Selalu tutup modal setelah proses selesai
  }
};

const editVilla = (villa) => { // Tidak perlu async jika hanya update store
  useVillaStore.setEditData(villa);
  useSideNavStore.setPage("Kelola Villa");
  console.log(villa);
};

// Computed property untuk data yang ditampilkan di tabel (filter berdasarkan pencarian)
const dataList = computed(() => {
  if (useSearchStore.state.input !== "") {
    return useSearchStore.state.result;
  }
  return useVillaStore.state.listData;
});
</script>

<template>
  <div v-if="tampilModalHapus" class="modal-overlay">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title">Hapus Data</h5>
        <button
            type="button"
            class="close-button"
            @click="sembunyikanModalHapus"
        >
          &times;
        </button>
      </div>
      <div class="modal-body">
        <p>Apakah Anda yakin ingin menghapus data ini?</p>
      </div>
      <div class="modal-footer">
        <button
            type="button"
            class="modal-button btn-batal"
            @click="sembunyikanModalHapus"
        >
          Batal
        </button>
        <button
            type="button"
            class="modal-button btn-yakin"
            @click="hapusModal"
        >
          Yakin
        </button>
      </div>
    </div>
  </div>

  <TableData v-if="useVillaStore.state.loadData"/>
  <div v-if="!useVillaStore.state.loadData" class="card mb-4">
    <div class="card-header pb-0">
      <h6>Data Villa</h6>
    </div>
    <div class="card-body px-0 pt-0 pb-2">
      <div class="table-responsive p-0">
        <table class="table align-items-center mb-0">
          <thead>
          <tr>
            <th
                class="text-uppercase text-secondary text-xxs font-weight-bolder opacity-7"
            >
              Cover
            </th>
            <th
                class="text-uppercase text-secondary text-xxs font-weight-bolder opacity-7 ps-2"
            >
              Judul
            </th>
            <th
                class="text-uppercase text-secondary text-xxs font-weight-bolder opacity-7 ps-2"
            >
              Deskripsi
            </th>
            <th
                class="text-center text-uppercase text-secondary text-xxs font-weight-bolder opacity-7"
            >
              Rating
            </th>
            <th class="text-secondary opacity-7"></th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="villa in dataList" :key="villa.id">
            <td>
              <div class="d-flex px-2 py-1">
                <div>
                  <img
                      :src="villa.displayImageUrl" class="avatar avatar-sm me-3 border-radius-lg"
                      :alt="`Gambar ${villa.judul}`"/>
                </div>
              </div>
            </td>
            <td>
              <p class="text-xs font-weight-bold mb-0">
                {{ villa.judul }}
              </p>
            </td>
            <td>
              <p class="text-xs text-secondary mb-0 description-ellipsis">
                {{ villa.deskripsi }}
              </p>
            </td>
            <td class="align-middle text-center">
                <span class="text-secondary text-xs font-weight-bold">{{
                    Number(villa.rating).toFixed(1)
                  }}</span>
            </td>
            <td class="align-middle">
              <a
                  @click="editVilla(villa)"
                  href="javascript:"
                  class="btn btn-link text-dark px-3 mb-0"
                  data-toggle="tooltip"
                  data-original-title="Edit Villa"
              >
                <i
                    class="fas fa-pencil-alt text-sm me-2"
                    aria-hidden="true"
                ></i
                >Edit
              </a>
              <a
                  href="javascript:"
                  class="btn btn-link text-danger px-3 mb-0"
                  data-toggle="tooltip"
                  data-original-title="Hapus Villa"
                  @click="tampilkanModalHapus(villa)"
              >
                <i
                    class="far fa-trash-alt text-sm me-2"
                    aria-hidden="true"
                ></i
                >Hapus
              </a>
            </td>
          </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* --- (Style CSS Anda yang sudah ada, tanpa perubahan signifikan) --- */
.description-ellipsis {
  width: 100px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

/* Modal Styles */
.modal-overlay {
  display: flex;
  justify-content: center;
  align-items: center;
  position: fixed;
  z-index: 9999; /* Lebih tinggi dari elemen lain */
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.4);
}

.modal-content {
  background-color: #fefefe;
  margin: auto;
  padding: 20px;
  border: 1px solid #888;
  width: 80%;
  max-width: 500px;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-bottom: 15px;
  border-bottom: 1px solid #eee;
}

.modal-title {
  font-size: 1.2em;
  font-weight: bold;
  margin: 0;
}

.close-button {
  color: #aaa;
  font-size: 28px;
  font-weight: bold;
  cursor: pointer;
  border: none;
  background: none;
  padding: 0;
}

.close-button:hover,
.close-button:focus {
  color: black;
  text-decoration: none;
}

.modal-body {
  padding: 20px 0;
}

.modal-footer {
  display: flex;
  justify-content: flex-end;
  padding-top: 15px;
  border-top: 1px solid #eee;
}

.modal-button {
  padding: 10px 15px;
  margin-left: 10px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 1em;
}

.btn-batal {
  background: none;
  border: 1px solid #344767;
  color: #344767;
}

.btn-yakin {
  background-color: #344767;
  color: white;
}

.btn-batal:hover,
.btn-batal:focus {
  background-color: #344767;
  color: #ffffff;
}

.btn-yakin:hover,
.btn-yakin:focus {
  background-color: #1b293a;
}

/* Responsiveness */
@media (max-width: 600px) {
  .modal-content {
    width: 95%;
  }
}
</style>